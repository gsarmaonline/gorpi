# Gorpi REST API integration

## Approach
The purpose of this module is to provide the fastest
way to declare RESTful APIs. Motivation is to provide
something similar to Ruby on Rails scaffolding.

There are 2 approaches that we can take here.
1. Provide a scaffolding based approach similar to Rails
2. Provide a framework which pre-defines the functions common
to the APIs.
3. Provide common tools which autofills the common tasks and easy way to
write new APIs.


### Observations
1. Scaffolding results in a higher amount of code being generated
regardless.
2. Since the code is being generated, it also
provides a higher level of control once the code is generated.
3. It is difficult to control/modify the behaviour of the scaffolded
code once generated.

### Flow for the integration
User will define the model and the RESTful routes they want to generate.
It can be of the format:
```go
userApiConfig := RestApiConfig{
    model: &User{},
    ignoreApis: [api.CreateApi],
    parent: &Person{}
}
restapi.GenerateAPI(userApiConfig)
```

The above block should generate the following:
- Routes for all the RESTful routes
- API layer for the RESTful routes
- DB action for the RESTful routes

## Common tools

- Extract request's unique ID
- Pagination 
- Database Operations
