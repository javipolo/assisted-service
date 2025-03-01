// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListClusterManifestsHandlerFunc turns a function with the right signature into a list cluster manifests handler
type ListClusterManifestsHandlerFunc func(ListClusterManifestsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListClusterManifestsHandlerFunc) Handle(params ListClusterManifestsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListClusterManifestsHandler interface for that can handle valid list cluster manifests params
type ListClusterManifestsHandler interface {
	Handle(ListClusterManifestsParams, interface{}) middleware.Responder
}

// NewListClusterManifests creates a new http.Handler for the list cluster manifests operation
func NewListClusterManifests(ctx *middleware.Context, handler ListClusterManifestsHandler) *ListClusterManifests {
	return &ListClusterManifests{Context: ctx, Handler: handler}
}

/* ListClusterManifests swagger:route GET /v1/clusters/{cluster_id}/manifests manifests listClusterManifests

Lists manifests for customizing cluster installation.

*/
type ListClusterManifests struct {
	Context *middleware.Context
	Handler ListClusterManifestsHandler
}

func (o *ListClusterManifests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListClusterManifestsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
