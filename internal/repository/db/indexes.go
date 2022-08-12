// Package db implements bridge to persistent storage represented by Mongo database.
package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

// indexListQueueCapacity represents the capacity of index lists queue.
const indexListQueueCapacity = 25

// IndexListProvider represents a function providing list of indexes.
type indexListProvider func() []mongo.IndexModel

// IndexList represents a prescribed list of indexes for the given collection
type IndexList struct {
	Collection *mongo.Collection
	Indexes    []mongo.IndexModel
}

// updateDatabaseIndexes checks for indexes existence; if an expected index is not found, it creates it.
func (db *MongoDbBridge) updateDatabaseIndexes() {
	// define index list loaders
	var ixLoaders = map[string]indexListProvider{
		colNetworkNodes:      operaNodeCollectionIndexes,
		colLockedDelegations: lockedDelegationsIndexes,
	}

	// the DB bridge needs a way to terminate this thread
	sig := make(chan bool, 1)
	db.sig = append(db.sig, sig)

	// prep queue and start the updater
	iq := make(chan *IndexList, indexListQueueCapacity)
	db.wg.Add(1)
	go db.indexUpdater(iq, sig)

	// check indexes
	for cn, ld := range ixLoaders {
		iq <- &IndexList{
			Collection: db.client.Database(db.dbName).Collection(cn),
			Indexes:    ld(),
		}
	}

	// close the channel as no more updates will be sent
	close(iq)
}

// indexUpdater processes queue of index updates.
func (db *MongoDbBridge) indexUpdater(iq chan *IndexList, sig chan bool) {
	defer func() {
		db.wg.Done()
		db.log.Noticef("db index updater closed")
	}()

	for {
		select {
		case <-sig:
			return
		case ix, more := <-iq:
			// is this the end?
			if !more {
				db.log.Noticef("all indexes processed")
				return
			}

			// do the update
			err := db.updateIndexes(ix.Collection, ix.Indexes, sig)
			if err != nil {
				db.log.Errorf("%s index list sync failed; %s", ix.Collection.Name(), err.Error())
			}
		}
	}
}

// updateIndexes synchronizes indexes of the given DB collection with the given prescription.
func (db *MongoDbBridge) updateIndexes(col *mongo.Collection, list []mongo.IndexModel, sig chan bool) error {
	view := col.Indexes()

	// load list of all indexes known
	known, err := view.ListSpecifications(context.Background())
	if err != nil {
		return err
	}

	// loop prescriptions and make sure each index exists by name
	for _, ix := range list {
		// respect possible terminate signal
		select {
		case <-sig:
			sig <- true
			return nil
		default:
		}

		// create missing index
		err = db.createIndexIfNotExists(col, &view, ix, known)
		if err != nil {
			db.log.Errorf(err.Error())
		}
	}

	// loop indexes and remove indexes missing in the prescriptions
	for _, spec := range known {
		err = db.removeIndexIfShouldNotExists(col, &view, spec, list)
		if err != nil {
			db.log.Errorf(err.Error())
		}
	}

	return nil
}

// createIndexIfNotExists checks specific index model and creates the index on the DB collection if needed.
func (db *MongoDbBridge) createIndexIfNotExists(col *mongo.Collection, view *mongo.IndexView, ix mongo.IndexModel, known []*mongo.IndexSpecification) error {
	// throw if index is not explicitly named
	if ix.Options.Name == nil {
		return fmt.Errorf("index name not defined on %s", col.Name())
	}

	// do we know the index?
	for _, spec := range known {
		if spec.Name == *ix.Options.Name {
			return nil
		}
	}

	createdName, err := view.CreateOne(context.Background(), ix)
	if err != nil {
		return fmt.Errorf("failed to create index %s on %s", *ix.Options.Name, col.Name())
	}
	db.log.Noticef("created index %s on %s", createdName, col.Name())
	return nil
}

// removeIndexIfShouldNotExists checks specific index model and creates the index on the DB collection if needed.
func (db *MongoDbBridge) removeIndexIfShouldNotExists(col *mongo.Collection, view *mongo.IndexView, spec *mongo.IndexSpecification, list []mongo.IndexModel) error {
	// skip for _id_ index
	if spec.Name == "_id_" {
		return nil
	}

	// do we know the index?
	for _, ix := range list {
		if spec.Name == *ix.Options.Name {
			return nil
		}
	}

	_, err := view.DropOne(context.Background(), spec.Name)
	if err != nil {
		return fmt.Errorf("failed to drop index %s on %s", spec.Name, col.Name())
	}
	db.log.Noticef("dropped index %s on %s", spec.Name, col.Name())
	return nil
}
