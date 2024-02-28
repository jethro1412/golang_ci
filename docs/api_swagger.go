package docs

import "github.com/jethro1412/golang_ci/api"

// swagger:route POST /api api idOfAPIEndpoint
// This api does some amazing stuff.
// responses:
//   200: ApiResponse

// This text will appear as description of your response body.
// swagger:response ApiResponse
type ApiResponseWrapper struct {
	// in:body
	Body api.ApiResponse
}

// swagger:parameters idOfApiEndpoint
type foobarParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body api.ApiRequests
}