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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/genesis/models"
)

// NewGenesisPeersPingParams creates a new GenesisPeersPingParams object
// with the default values initialized.
func NewGenesisPeersPingParams() *GenesisPeersPingParams {
	var ()
	return &GenesisPeersPingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGenesisPeersPingParamsWithTimeout creates a new GenesisPeersPingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGenesisPeersPingParamsWithTimeout(timeout time.Duration) *GenesisPeersPingParams {
	var ()
	return &GenesisPeersPingParams{

		timeout: timeout,
	}
}

// NewGenesisPeersPingParamsWithContext creates a new GenesisPeersPingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGenesisPeersPingParamsWithContext(ctx context.Context) *GenesisPeersPingParams {
	var ()
	return &GenesisPeersPingParams{

		Context: ctx,
	}
}

// NewGenesisPeersPingParamsWithHTTPClient creates a new GenesisPeersPingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGenesisPeersPingParamsWithHTTPClient(client *http.Client) *GenesisPeersPingParams {
	var ()
	return &GenesisPeersPingParams{
		HTTPClient: client,
	}
}

/*GenesisPeersPingParams contains all the parameters to send to the API endpoint
for the genesis peers ping operation typically these are written to a http.Request
*/
type GenesisPeersPingParams struct {

	/*Body
	  Request Body

	*/
	Body *models.PeerPing
	/*PeerID
	  Name of the Weaviate peer

	*/
	PeerID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the genesis peers ping params
func (o *GenesisPeersPingParams) WithTimeout(timeout time.Duration) *GenesisPeersPingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the genesis peers ping params
func (o *GenesisPeersPingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the genesis peers ping params
func (o *GenesisPeersPingParams) WithContext(ctx context.Context) *GenesisPeersPingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the genesis peers ping params
func (o *GenesisPeersPingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the genesis peers ping params
func (o *GenesisPeersPingParams) WithHTTPClient(client *http.Client) *GenesisPeersPingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the genesis peers ping params
func (o *GenesisPeersPingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the genesis peers ping params
func (o *GenesisPeersPingParams) WithBody(body *models.PeerPing) *GenesisPeersPingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the genesis peers ping params
func (o *GenesisPeersPingParams) SetBody(body *models.PeerPing) {
	o.Body = body
}

// WithPeerID adds the peerID to the genesis peers ping params
func (o *GenesisPeersPingParams) WithPeerID(peerID strfmt.UUID) *GenesisPeersPingParams {
	o.SetPeerID(peerID)
	return o
}

// SetPeerID adds the peerId to the genesis peers ping params
func (o *GenesisPeersPingParams) SetPeerID(peerID strfmt.UUID) {
	o.PeerID = peerID
}

// WriteToRequest writes these params to a swagger request
func (o *GenesisPeersPingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param peerId
	if err := r.SetPathParam("peerId", o.PeerID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
