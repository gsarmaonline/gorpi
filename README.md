# Go Rest APIs - Gorpi

Gorpi is a highly opinionated REST API framework powered by Gin.
The framework decides to provide few helpers and conventions which
make it very straightforward to establish simple CRUD APIs.

## Implementation flow

Let's assume that we have an object called `User`.

```go
type (
    User struct {
        AttrOne string
        AttrTwo int
    }
)
```

In order to expose the object in a REST API, the following steps are required.

1. Create a model and the underlying table in the configured database
2. Register routes
3. Support authentication
4. Access to database objects

This is the normal flow of an API. Given the Gorpi is going to implement the MVC
methodology, there are other standard places where overriding the original object
should be allowed.

Similar to how Rails provides hooks or filters for models, controllers, etc like
`BeforeSave`, `AfterSave` for models and `BeforeAction`, `AfterAction` for controllers.

For every API request, the flow is quite similar:

- Receive the request
- Check whether the request is proper with the correct information and is authorised
- Fetch or write something to the DB
- Return the response

## Related models

In common scenarios, models are connected to each other via some or the other way.
In this section, we will cover how parent child models are implemented and are expected to behave.

Example:

```go
type (
    User struct {
        Name string
        Age int
        Addresses []Address
    }

    Address struct {
        Line string
        State string
        Country string
    }
)
```

In this case, the child model, which is `Address`, cannot exist in the system without the existence of
the `User` object.

## Gorpi Elements

### Resources

In Rails, using `resources :model_one` creates the RESTful routes for the `model_one` resource and the
required actions are then implemented with custom logic.
Gorpi also uses the concept of `Resource` to leverage standard methods in the app.

### Routes

Routes are the building blocks of the API system.
Below is the object structure.

```go
type (
	Route struct {
		RequestURI    string
		RequestMethod string
		Handler       api.ApiHandlerFunc
		ResourceModel models.ResourceModel
	}
)
```

## Integrating with Gorpi

### Defining routes without a Resource

TBD

### Defining Resource

TBD
