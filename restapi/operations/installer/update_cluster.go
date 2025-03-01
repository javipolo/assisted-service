// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateClusterHandlerFunc turns a function with the right signature into a update cluster handler
type UpdateClusterHandlerFunc func(UpdateClusterParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateClusterHandlerFunc) Handle(params UpdateClusterParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateClusterHandler interface for that can handle valid update cluster params
type UpdateClusterHandler interface {
	Handle(UpdateClusterParams, interface{}) middleware.Responder
}

// NewUpdateCluster creates a new http.Handler for the update cluster operation
func NewUpdateCluster(ctx *middleware.Context, handler UpdateClusterHandler) *UpdateCluster {
	return &UpdateCluster{Context: ctx, Handler: handler}
}

/* UpdateCluster swagger:route PATCH /v1/clusters/{cluster_id} installer updateCluster

Updates an OpenShift cluster definition.

*/
type UpdateCluster struct {
	Context *middleware.Context
	Handler UpdateClusterHandler
}

func (o *UpdateCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateClusterParams()
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
