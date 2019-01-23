<p align="center"><a href="https://www.shyft.network/">
	<img src="https://www.shyft.network/images/shyft-logo-horizontal-tm.svg" alt="Shyft" height="100px" align="center">
</a></p>

# Shyft Block Explorer API

This repository contains the Shyft Block Explorer API which is primarily used to serve the Shyft Block Explorer UI by communicating with a Postgres instance attached to Shyft's go-empyrean blockchain.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

It is required to have Shyft's go-empyrean blockchain postgres instance running in order to utilize the API. The API is specifically looking for the postgres instances connection string in order to communicate
the data to the Shyft Block Explorer UI. Here you will find the [go-empyrean repo](https://github.com/ShyftNetwork/go-empyrean) and instructions on how to start the blockchain.

In addition you will need `go 1.11` or greater as the API makes use of [go modules](https://github.com/golang/go/wiki/Modules). You can install go versions [here](https://golang.org/dl/).

```
1. go-empyrean postgres instance
2. > go 1.11
```

### Installing

To get a development environment running locally please follow the below steps:

```
1. Create the following directory $GOPATH/src/github.com/ShyftNetwork/
2. git clone git@github.com:ShyftNetwork/blockexplorer_api.git
3. make install (this installs dependencies from go.mod)
4. make run
```

**You should see the server starting and listening on port 8080!**

### Running API through Docker

To run the API through a docker container please make sure your have a docker daemon running. Then run the following command:

```
> make docker
```

You should see the binary executable and a docker image being built once that is completed, you should see the server starting and listening on port 8080!**

## Running the tests

To run the API tests run the following command:

```
> make test
```

## Running the linter

To run the API linter, which uses *gometalinter* run the following command:

```
> make lint
```

## Deployment

For the most optimized deployment, you will want to create an executable binary. You can do so by running the following command:

```
> make build
```

This will create a binary in the root directory which than can be referenced in the `Dockerfile` as we have done.

