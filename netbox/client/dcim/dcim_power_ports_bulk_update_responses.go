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

// DcimPowerPortsBulkUpdateReader is a Reader for the DcimPowerPortsBulkUpdate structure.
type DcimPowerPortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsBulkUpdateOK creates a DcimPowerPortsBulkUpdateOK with default headers values
func NewDcimPowerPortsBulkUpdateOK() *DcimPowerPortsBulkUpdateOK {
	return &DcimPowerPortsBulkUpdateOK{}
}

/*
DcimPowerPortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortsBulkUpdateOK dcim power ports bulk update o k
*/
type DcimPowerPortsBulkUpdateOK struct {
	Payload *models.PowerPort
}

// IsSuccess returns true when this dcim power ports bulk update o k response has a 2xx status code
func (o *DcimPowerPortsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports bulk update o k response has a 3xx status code
func (o *DcimPowerPortsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports bulk update o k response has a 4xx status code
func (o *DcimPowerPortsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports bulk update o k response has a 5xx status code
func (o *DcimPowerPortsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports bulk update o k response a status code equal to that given
func (o *DcimPowerPortsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerPortsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-ports/][%d] dcimPowerPortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-ports/][%d] dcimPowerPortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsBulkUpdateOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortsBulkUpdateDefault creates a DcimPowerPortsBulkUpdateDefault with default headers values
func NewDcimPowerPortsBulkUpdateDefault(code int) *DcimPowerPortsBulkUpdateDefault {
	return &DcimPowerPortsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimPowerPortsBulkUpdateDefault dcim power ports bulk update default
*/
type DcimPowerPortsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power ports bulk update default response
func (o *DcimPowerPortsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power ports bulk update default response has a 2xx status code
func (o *DcimPowerPortsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power ports bulk update default response has a 3xx status code
func (o *DcimPowerPortsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power ports bulk update default response has a 4xx status code
func (o *DcimPowerPortsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power ports bulk update default response has a 5xx status code
func (o *DcimPowerPortsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power ports bulk update default response a status code equal to that given
func (o *DcimPowerPortsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerPortsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-ports/][%d] dcim_power-ports_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/power-ports/][%d] dcim_power-ports_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
