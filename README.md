# Fantom GraphQL API Server
![Go Report Card](https://goreportcard.com/badge/github.com/Fantom-foundation/fantom-api-graphql)

High performance API server for Fantom powered blockchain network.

## Building the source

Building `apiserver` requires a Go (version 1.13 or later). You can install
it using your favourite package manager. Once the dependencies are installed, run

```shell
go build -o ./build/apiserver ./cmd/apiserver
```

The build output is ```build/apiserver``` executable.

You don't need to clone the project into $GOPATH, due to use of Go Modules you can 
use any location.

## Running the API server

To run the API Server you need access to a RPC interface of a full Lachesis node. Please
follow [Lachesis](https://github.com/Fantom-foundation/go-lachesis) instructions to build
and run the node. Alternatively you can obtain access to a remotely running instance
of Lachesis. 

We recommend using local IPC channel for communication between a Lachesis node and the 
API Server for performance and security reasons. Please consider security implications 
of opening Lachesis RPC to outside access, especially if you enable "personal" commands 
on your node while keeping your account keys in the Lachesis key store.

Persistent data are stored in a MongoDB database. Going through the instalation and 
configuration process of MongoDB is out of scope here, please consult 
[MongoDB manual](https://docs.mongodb.com/manual/) to install and configure appropriate 
MongoDB environment for your deployment of the API server.
