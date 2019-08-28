# Optimizely Sidedoor
Exploratory project for developing a service version of the Optimizely SDK.

## Installing, Building and Running
The following `make` targets can be used to build and run the application:
* build - builds sidedoor and installs binary in bin/sidedoor
* test - recursively tests all .go files
* run - builds and executes the sidedoor binary
* clean - runs `go clean` and removes the bin/ dir
* install - runs `go get` to install all dependencies
* generate-api - generates APIs from the swagger spec


## Client Generation

### Prerequisites
This repo currently depends heavily on [OpenAPI](https://swagger.io/specification/) and [OpenAPI Generator](https://github.com/openapitools/openapi-generator) (a [fork](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/migration-from-swagger-codegen.md) of swagger-codegen).

To install the OpenAPI Generator on OSX:
```
brew install openapi-generator
```

To determine which generators are available you can execute `openapi-generator` without any arguments or refer to the generator source [docs](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/README.md):

Types of generators are either CLIENT, SERVER, DOCUMENTATION, SCHEMA and CONFIG.

### Generating
You can use the helper script `generate.sh` to experiment with the various generated assets.
```
scripts/generate.sh <GENERATOR_NAME>
```
We also provide a Make task `generate-api`:
```
make generate-api ARG-<GENERATOR_NAME>
```
