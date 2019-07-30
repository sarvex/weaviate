//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
// 
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewWeaviateActionsPatchParams creates a new WeaviateActionsPatchParams object
// with the default values initialized.
func NewWeaviateActionsPatchParams() *WeaviateActionsPatchParams {
	var ()
	return &WeaviateActionsPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionsPatchParamsWithTimeout creates a new WeaviateActionsPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionsPatchParamsWithTimeout(timeout time.Duration) *WeaviateActionsPatchParams {
	var ()
	return &WeaviateActionsPatchParams{

		timeout: timeout,
	}
}

// NewWeaviateActionsPatchParamsWithContext creates a new WeaviateActionsPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionsPatchParamsWithContext(ctx context.Context) *WeaviateActionsPatchParams {
	var ()
	return &WeaviateActionsPatchParams{

		Context: ctx,
	}
}

// NewWeaviateActionsPatchParamsWithHTTPClient creates a new WeaviateActionsPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionsPatchParamsWithHTTPClient(client *http.Client) *WeaviateActionsPatchParams {
	var ()
	return &WeaviateActionsPatchParams{
		HTTPClient: client,
	}
}

/*WeaviateActionsPatchParams contains all the parameters to send to the API endpoint
for the weaviate actions patch operation typically these are written to a http.Request
*/
type WeaviateActionsPatchParams struct {

	/*Body
	  JSONPatch document as defined by RFC 6902.

	*/
	Body []*models.PatchDocument
	/*ID
	  Unique ID of the Action.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) WithTimeout(timeout time.Duration) *WeaviateActionsPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) WithContext(ctx context.Context) *WeaviateActionsPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) WithHTTPClient(client *http.Client) *WeaviateActionsPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) WithBody(body []*models.PatchDocument) *WeaviateActionsPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) SetBody(body []*models.PatchDocument) {
	o.Body = body
}

// WithID adds the id to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) WithID(id strfmt.UUID) *WeaviateActionsPatchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the weaviate actions patch params
func (o *WeaviateActionsPatchParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionsPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
