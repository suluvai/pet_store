# Pet Store Rest Api

Based on the given requirementts as per the https://petstore.swagger.io/ this rest api application is developed.

- It implements Pet endpoints only.
- Uses a simple api_key authentication.
- Added a simple test case for testng PetController with mocked PetRepository. But have not implemented tests for all test case scenarios. Although demostrated the ability to add tests.

## Prerequisites

- Docker
- Linux (Presumption). But can be built on Windows if required.

## Build

Builds the app in a docker container and copies the build binary to new container used to run the application.

```./run.sh build```

## Execute

Runs the docker container that runs the application.

```./run.sh```

