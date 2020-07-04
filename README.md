Axon
----

Core API and distributed focus group domain logic in golang.

Primary design principles are loosely pulled from:

  * Domain Driven Design (see book by Eric Evans)
  * Onion/Hexagonal Architecture. _Layers of implementation where inner layers
are ignorant of those stacked on top of them._

## Quickstart

There are two entrypoints:

	* `/cmd/script/main.go` for command line access.
	* `/cmd/server/main.go` for a server routing API endpoints.

`go run cmd/script/main.go` will run the script entrypoint.

## Testing
