# Fantom GraphQL API Server
[![Go Report Card](https://goreportcard.com/badge/github.com/Fantom-foundation/fantom-api-graphql)](https://goreportcard.com/report/github.com/Fantom-foundation/fantom-api-graphql)

GraphQL API server for Fantom powered blockchain network.

## Releases
Please check the [release tags](https://github.com/Fantom-foundation/fantom-api-graphql/tags) to get more details and to download previous releases.

#### Version 0.2.0, pending
This version connects with the Lachesis v.0.7.0-rc1. The SFC contract ABI bundled with the API is version 2.0.2-rc1.

The release brings new fluid delegations and rewards system. Each address is be able to delegate to multiple stakers. Delegation can be locked to certain time, at least 14 days and up to 1 year, to get higher rewards. Please check our website [Fantom.Foundation](https://fantom.foundation) and the [Special Fee Contract repository](https://github.com/Fantom-foundation/fantom-sfc) for more details.

#### Version 0.1.0, released on 8/2020
This is the version you want to be able to connect with Lachesis v.0.6.0-rc2. The SFC contract ABI bundled with this API release is the version 1.1.0-rc1. The release uses Lachesis API v0.6.0 which recognizes single delegation per address and no delegation locking.

## Building the source

Building `apiserver` requires a Go (version 1.13 or later). You can install
it using your favourite package manager. Once the dependencies are installed, run

```shell
make
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

Persistent data are stored in a MongoDB database. Going through the installation and
configuration process of MongoDB is out of scope here, please consult
[MongoDB manual](https://docs.mongodb.com/manual/) to install and configure appropriate
MongoDB environment for your deployment of the API server.
