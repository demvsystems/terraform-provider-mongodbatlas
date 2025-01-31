# terraform-provider-mongodbatlas
[![Build Status](https://travis-ci.org/akshaykarle/terraform-provider-mongodbatlas.svg?branch=master)](https://travis-ci.org/akshaykarle/terraform-provider-mongodbatlas)
[![GitHub release](https://img.shields.io/github/release/akshaykarle/terraform-provider-mongodbatlas.svg)](https://github.com/akshaykarle/terraform-provider-mongodbatlas/releases)
[![codecov](https://codecov.io/gh/akshaykarle/terraform-provider-mongodbatlas/branch/master/graph/badge.svg)](https://codecov.io/gh/akshaykarle/terraform-provider-mongodbatlas)
[![GitHub downloads](https://img.shields.io/github/downloads/akshaykarle/terraform-provider-mongodbatlas/total.svg)]()

Terraform provider for [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

## Requirements
-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.12 (to build the provider plugin)

## Installing the Provider
Follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory, run `terraform init` to initialize it.

## Usage
```
# Configure the MongoDB Atlas Provider
provider "mongodbatlas" {
  username = "${var.mongodb_atlas_username}"
  api_key = "${var.mongodb_atlas_api_key}"
}

# Create a Cluster
resource "mongodbatlas_cluster" "test" {
  # ...
}
```
Also look at the example under [/examples](/examples).

## Building the Provider
Clone and build the repository

```sh
go get github.com/demvsystems/terraform-provider-mongodbatlas
make build
```

## Updating the Provider

```sh
go get -u github.com/demvsystems/terraform-provider-mongodbatlas
make build
```

## Contributing
* Install project dependencies: `go get github.com/kardianos/govendor`
* Run tests: `make test`
* Build the binary: `make build`
