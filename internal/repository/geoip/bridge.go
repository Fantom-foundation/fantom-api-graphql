package geoip

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	"github.com/oschwald/geoip2-golang"
	"net"
)

const defaultLocationNameFormat = "en"

// Bridge represents the bridge to GeoIP database services.
type Bridge struct {
	db  *geoip2.Reader
	log logger.Logger
}

// New creates a new Mongo Db connection bridge.
func New(cfg *config.PeerNetworking, log logger.Logger) (*Bridge, error) {
	db, err := geoip2.Open(cfg.GeoIPPath)
	if err != nil {
		log.Criticalf("can not open GeoIP database at %s; %s", cfg.GeoIPPath, err.Error())
		return nil, err
	}

	return &Bridge{
		db:  db,
		log: log,
	}, nil
}

// Close will terminate or finish all operations and close the connection to GeoIP database.
func (gib *Bridge) Close() {
	err := gib.db.Close()
	if err != nil {
		gib.log.Errorf("could not close GeoIP database; %s", err.Error())
	}
}

// Location provides location details for the given IP address.
func (gib *Bridge) Location(ip net.IP) (types.GeoLocation, error) {
	var loc types.GeoLocation

	record, err := gib.db.City(ip)
	if err != nil {
		gib.log.Errorf("could not lookup IP address %s, %s", ip.String(), err.Error())
		return loc, err
	}

	loc.Continent = record.Continent.Names[defaultLocationNameFormat]
	loc.Country = record.Country.Names[defaultLocationNameFormat]
	loc.City = record.City.Names[defaultLocationNameFormat]
	loc.TimeZone = record.Location.TimeZone
	loc.Latitude = record.Location.Latitude
	loc.Longitude = record.Location.Longitude
	loc.Accuracy = record.Location.AccuracyRadius

	if len(record.Subdivisions) > 0 {
		loc.Region = record.Subdivisions[0].Names[defaultLocationNameFormat]
	} else {
		loc.Region = loc.Country
	}
	return loc, nil
}
