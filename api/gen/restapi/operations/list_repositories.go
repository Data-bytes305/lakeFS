// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/treeverse/lakefs/api/gen/models"
)

// ListRepositoriesHandlerFunc turns a function with the right signature into a list repositories handler
type ListRepositoriesHandlerFunc func(ListRepositoriesParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRepositoriesHandlerFunc) Handle(params ListRepositoriesParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// ListRepositoriesHandler interface for that can handle valid list repositories params
type ListRepositoriesHandler interface {
	Handle(ListRepositoriesParams, *models.User) middleware.Responder
}

// NewListRepositories creates a new http.Handler for the list repositories operation
func NewListRepositories(ctx *middleware.Context, handler ListRepositoriesHandler) *ListRepositories {
	return &ListRepositories{Context: ctx, Handler: handler}
}

/*ListRepositories swagger:route GET /repositories listRepositories

list repositories

*/
type ListRepositories struct {
	Context *middleware.Context
	Handler ListRepositoriesHandler
}

func (o *ListRepositories) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListRepositoriesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
