# Go Rest APIs - Gorpi

Gorpi is a highly opinionated REST API framework powered by Gin.
The framework decides to provide few helpers and conventions which
make it very straightforward to establish simple CRUD APIs.

## Routing
Rails, the Ruby framework provides a method `resources` which defines
the RESTful routes by default for the mentioned resource.
It allows proper nesting of resources depending on where they are configured.
Though slightly laborious to get to that level of efficiency in Go, the API
layer will try to take the Rails resources as the guiding factor.

## Database connections
Gorpi provides a database library on top of gorm to provide easy access
to an ORM via the server object.

## TODO
- Authenticator
- Database connection
- Statistics
- Background workers
- Cache

