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

// DcimPlatformsBulkPartialUpdateReader is a Reader for the DcimPlatformsBulkPartialUpdate structure.
type DcimPlatformsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPlatformsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPlatformsBulkPartialUpdateOK creates a DcimPlatformsBulkPartialUpdateOK with default headers values
func NewDcimPlatformsBulkPartialUpdateOK() *DcimPlatformsBulkPartialUpdateOK {
	return &DcimPlatformsBulkPartialUpdateOK{}
}

/*
DcimPlatformsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPlatformsBulkPartialUpdateOK dcim platforms bulk partial update o k
*/
type DcimPlatformsBulkPartialUpdateOK struct {
	Payload *models.Platform
}

// IsSuccess returns true when this dcim platforms bulk partial update o k response has a 2xx status code
func (o *DcimPlatformsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim platforms bulk partial update o k response has a 3xx status code
func (o *DcimPlatformsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms bulk partial update o k response has a 4xx status code
func (o *DcimPlatformsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim platforms bulk partial update o k response has a 5xx status code
func (o *DcimPlatformsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms bulk partial update o k response a status code equal to that given
func (o *DcimPlatformsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPlatformsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/][%d] dcimPlatformsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPlatformsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/][%d] dcimPlatformsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPlatformsBulkPartialUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPlatformsBulkPartialUpdateDefault creates a DcimPlatformsBulkPartialUpdateDefault with default headers values
func NewDcimPlatformsBulkPartialUpdateDefault(code int) *DcimPlatformsBulkPartialUpdateDefault {
	return &DcimPlatformsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPlatformsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPlatformsBulkPartialUpdateDefault dcim platforms bulk partial update default
*/
type DcimPlatformsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim platforms bulk partial update default response
func (o *DcimPlatformsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim platforms bulk partial update default response has a 2xx status code
func (o *DcimPlatformsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim platforms bulk partial update default response has a 3xx status code
func (o *DcimPlatformsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim platforms bulk partial update default response has a 4xx status code
func (o *DcimPlatformsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim platforms bulk partial update default response has a 5xx status code
func (o *DcimPlatformsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim platforms bulk partial update default response a status code equal to that given
func (o *DcimPlatformsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPlatformsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/][%d] dcim_platforms_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/][%d] dcim_platforms_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
