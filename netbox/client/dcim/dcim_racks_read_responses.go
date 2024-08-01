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

// DcimRacksReadReader is a Reader for the DcimRacksRead structure.
type DcimRacksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRacksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRacksReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksReadOK creates a DcimRacksReadOK with default headers values
func NewDcimRacksReadOK() *DcimRacksReadOK {
	return &DcimRacksReadOK{}
}

/*
DcimRacksReadOK describes a response with status code 200, with default header values.

DcimRacksReadOK dcim racks read o k
*/
type DcimRacksReadOK struct {
	Payload *models.Rack
}

// IsSuccess returns true when this dcim racks read o k response has a 2xx status code
func (o *DcimRacksReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim racks read o k response has a 3xx status code
func (o *DcimRacksReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks read o k response has a 4xx status code
func (o *DcimRacksReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim racks read o k response has a 5xx status code
func (o *DcimRacksReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks read o k response a status code equal to that given
func (o *DcimRacksReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRacksReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/][%d] dcimRacksReadOK  %+v", 200, o.Payload)
}

func (o *DcimRacksReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/][%d] dcimRacksReadOK  %+v", 200, o.Payload)
}

func (o *DcimRacksReadOK) GetPayload() *models.Rack {
	return o.Payload
}

func (o *DcimRacksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRacksReadDefault creates a DcimRacksReadDefault with default headers values
func NewDcimRacksReadDefault(code int) *DcimRacksReadDefault {
	return &DcimRacksReadDefault{
		_statusCode: code,
	}
}

/*
DcimRacksReadDefault describes a response with status code -1, with default header values.

DcimRacksReadDefault dcim racks read default
*/
type DcimRacksReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim racks read default response
func (o *DcimRacksReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim racks read default response has a 2xx status code
func (o *DcimRacksReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim racks read default response has a 3xx status code
func (o *DcimRacksReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim racks read default response has a 4xx status code
func (o *DcimRacksReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim racks read default response has a 5xx status code
func (o *DcimRacksReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim racks read default response a status code equal to that given
func (o *DcimRacksReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRacksReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/][%d] dcim_racks_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/][%d] dcim_racks_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
