// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/henrybear327/go-netbox/netbox/models"
)

// CircuitsProvidersReadReader is a Reader for the CircuitsProvidersRead structure.
type CircuitsProvidersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProvidersReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersReadOK creates a CircuitsProvidersReadOK with default headers values
func NewCircuitsProvidersReadOK() *CircuitsProvidersReadOK {
	return &CircuitsProvidersReadOK{}
}

/*
CircuitsProvidersReadOK describes a response with status code 200, with default header values.

CircuitsProvidersReadOK circuits providers read o k
*/
type CircuitsProvidersReadOK struct {
	Payload *models.Provider
}

// IsSuccess returns true when this circuits providers read o k response has a 2xx status code
func (o *CircuitsProvidersReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits providers read o k response has a 3xx status code
func (o *CircuitsProvidersReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers read o k response has a 4xx status code
func (o *CircuitsProvidersReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits providers read o k response has a 5xx status code
func (o *CircuitsProvidersReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers read o k response a status code equal to that given
func (o *CircuitsProvidersReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsProvidersReadOK) Error() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuitsProvidersReadOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersReadOK) String() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuitsProvidersReadOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersReadOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProvidersReadDefault creates a CircuitsProvidersReadDefault with default headers values
func NewCircuitsProvidersReadDefault(code int) *CircuitsProvidersReadDefault {
	return &CircuitsProvidersReadDefault{
		_statusCode: code,
	}
}

/*
CircuitsProvidersReadDefault describes a response with status code -1, with default header values.

CircuitsProvidersReadDefault circuits providers read default
*/
type CircuitsProvidersReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits providers read default response
func (o *CircuitsProvidersReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits providers read default response has a 2xx status code
func (o *CircuitsProvidersReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits providers read default response has a 3xx status code
func (o *CircuitsProvidersReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits providers read default response has a 4xx status code
func (o *CircuitsProvidersReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits providers read default response has a 5xx status code
func (o *CircuitsProvidersReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits providers read default response a status code equal to that given
func (o *CircuitsProvidersReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsProvidersReadDefault) Error() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuits_providers_read default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersReadDefault) String() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuits_providers_read default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
