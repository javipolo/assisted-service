// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2ListSupportedOpenshiftVersionsHandlerFunc turns a function with the right signature into a v2 list supported openshift versions handler
type V2ListSupportedOpenshiftVersionsHandlerFunc func(V2ListSupportedOpenshiftVersionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2ListSupportedOpenshiftVersionsHandlerFunc) Handle(params V2ListSupportedOpenshiftVersionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2ListSupportedOpenshiftVersionsHandler interface for that can handle valid v2 list supported openshift versions params
type V2ListSupportedOpenshiftVersionsHandler interface {
	Handle(V2ListSupportedOpenshiftVersionsParams, interface{}) middleware.Responder
}

// NewV2ListSupportedOpenshiftVersions creates a new http.Handler for the v2 list supported openshift versions operation
func NewV2ListSupportedOpenshiftVersions(ctx *middleware.Context, handler V2ListSupportedOpenshiftVersionsHandler) *V2ListSupportedOpenshiftVersions {
	return &V2ListSupportedOpenshiftVersions{Context: ctx, Handler: handler}
}

/* V2ListSupportedOpenshiftVersions swagger:route GET /v2/openshift-versions versions v2ListSupportedOpenshiftVersions

Retrieves the list of OpenShift supported versions.

*/
type V2ListSupportedOpenshiftVersions struct {
	Context *middleware.Context
	Handler V2ListSupportedOpenshiftVersionsHandler
}

func (o *V2ListSupportedOpenshiftVersions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2ListSupportedOpenshiftVersionsParams()
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
