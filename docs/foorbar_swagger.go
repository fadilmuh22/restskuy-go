package docs

import "github.com/pdrum/swagger-automation/api"

// swagger:route POST /foobar foobar-tag idOfFoobarEndpoint
// Foobar does some amazing stuff.
// responses:
//   200: foobarResponse

// This text will appear as description of your response body.
// swagger:response foobarResponse
type FoobarResponseWrapper struct {
	// in:body
	Body api.FooBarResponse
}

// swagger:parameters idOfFoobarEndpoint
type FoobarParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body api.FooBarRequest
}
