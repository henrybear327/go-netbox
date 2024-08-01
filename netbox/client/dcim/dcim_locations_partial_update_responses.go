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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/henrybear327/go-netbox/netbox/models"
)

// DcimLocationsPartialUpdateReader is a Reader for the DcimLocationsPartialUpdate structure.
type DcimLocationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimLocationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimLocationsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsPartialUpdateOK creates a DcimLocationsPartialUpdateOK with default headers values
func NewDcimLocationsPartialUpdateOK() *DcimLocationsPartialUpdateOK {
	return &DcimLocationsPartialUpdateOK{}
}

/*
DcimLocationsPartialUpdateOK describes a response with status code 200, with default header values.

DcimLocationsPartialUpdateOK dcim locations partial update o k
*/
type DcimLocationsPartialUpdateOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this dcim locations partial update o k response has a 2xx status code
func (o *DcimLocationsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations partial update o k response has a 3xx status code
func (o *DcimLocationsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations partial update o k response has a 4xx status code
func (o *DcimLocationsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations partial update o k response has a 5xx status code
func (o *DcimLocationsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations partial update o k response a status code equal to that given
func (o *DcimLocationsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimLocationsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/locations/{id}/][%d] dcimLocationsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimLocationsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/locations/{id}/][%d] dcimLocationsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimLocationsPartialUpdateOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *DcimLocationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimLocationsPartialUpdateDefault creates a DcimLocationsPartialUpdateDefault with default headers values
func NewDcimLocationsPartialUpdateDefault(code int) *DcimLocationsPartialUpdateDefault {
	return &DcimLocationsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimLocationsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimLocationsPartialUpdateDefault dcim locations partial update default
*/
type DcimLocationsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim locations partial update default response
func (o *DcimLocationsPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim locations partial update default response has a 2xx status code
func (o *DcimLocationsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim locations partial update default response has a 3xx status code
func (o *DcimLocationsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim locations partial update default response has a 4xx status code
func (o *DcimLocationsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim locations partial update default response has a 5xx status code
func (o *DcimLocationsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim locations partial update default response a status code equal to that given
func (o *DcimLocationsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimLocationsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/locations/{id}/][%d] dcim_locations_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/locations/{id}/][%d] dcim_locations_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
