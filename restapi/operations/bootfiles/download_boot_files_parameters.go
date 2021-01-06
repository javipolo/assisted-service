// Code generated by go-swagger; DO NOT EDIT.

package bootfiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDownloadBootFilesParams creates a new DownloadBootFilesParams object
// no default values defined in spec.
func NewDownloadBootFilesParams() DownloadBootFilesParams {

	return DownloadBootFilesParams{}
}

// DownloadBootFilesParams contains all the bound params for the download boot files operation
// typically these are obtained from a http.Request
//
// swagger:parameters DownloadBootFiles
type DownloadBootFilesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The file type to download.
	  Required: true
	  In: query
	*/
	FileType string
	/*The corresponding OpenShift version for the boot file.
	  Required: true
	  In: query
	*/
	OpenshiftVersion string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDownloadBootFilesParams() beforehand.
func (o *DownloadBootFilesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFileType, qhkFileType, _ := qs.GetOK("file_type")
	if err := o.bindFileType(qFileType, qhkFileType, route.Formats); err != nil {
		res = append(res, err)
	}

	qOpenshiftVersion, qhkOpenshiftVersion, _ := qs.GetOK("openshift_version")
	if err := o.bindOpenshiftVersion(qOpenshiftVersion, qhkOpenshiftVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFileType binds and validates parameter FileType from query.
func (o *DownloadBootFilesParams) bindFileType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("file_type", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("file_type", "query", raw); err != nil {
		return err
	}

	o.FileType = raw

	if err := o.validateFileType(formats); err != nil {
		return err
	}

	return nil
}

// validateFileType carries on validations for parameter FileType
func (o *DownloadBootFilesParams) validateFileType(formats strfmt.Registry) error {

	if err := validate.EnumCase("file_type", "query", o.FileType, []interface{}{"initrd.img", "rootfs.img", "vmlinuz"}, true); err != nil {
		return err
	}

	return nil
}

// bindOpenshiftVersion binds and validates parameter OpenshiftVersion from query.
func (o *DownloadBootFilesParams) bindOpenshiftVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("openshift_version", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("openshift_version", "query", raw); err != nil {
		return err
	}

	o.OpenshiftVersion = raw

	return nil
}
