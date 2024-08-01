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

// DcimInterfaceTemplatesReadReader is a Reader for the DcimInterfaceTemplatesRead structure.
type DcimInterfaceTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfaceTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfaceTemplatesReadOK creates a DcimInterfaceTemplatesReadOK with default headers values
func NewDcimInterfaceTemplatesReadOK() *DcimInterfaceTemplatesReadOK {
	return &DcimInterfaceTemplatesReadOK{}
}

/*
DcimInterfaceTemplatesReadOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesReadOK dcim interface templates read o k
*/
type DcimInterfaceTemplatesReadOK struct {
	Payload *models.InterfaceTemplate
}

// IsSuccess returns true when this dcim interface templates read o k response has a 2xx status code
func (o *DcimInterfaceTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interface templates read o k response has a 3xx status code
func (o *DcimInterfaceTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates read o k response has a 4xx status code
func (o *DcimInterfaceTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interface templates read o k response has a 5xx status code
func (o *DcimInterfaceTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates read o k response a status code equal to that given
func (o *DcimInterfaceTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInterfaceTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimInterfaceTemplatesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/interface-templates/{id}/][%d] dcimInterfaceTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimInterfaceTemplatesReadOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfaceTemplatesReadDefault creates a DcimInterfaceTemplatesReadDefault with default headers values
func NewDcimInterfaceTemplatesReadDefault(code int) *DcimInterfaceTemplatesReadDefault {
	return &DcimInterfaceTemplatesReadDefault{
		_statusCode: code,
	}
}

/*
DcimInterfaceTemplatesReadDefault describes a response with status code -1, with default header values.

DcimInterfaceTemplatesReadDefault dcim interface templates read default
*/
type DcimInterfaceTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interface templates read default response
func (o *DcimInterfaceTemplatesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interface templates read default response has a 2xx status code
func (o *DcimInterfaceTemplatesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interface templates read default response has a 3xx status code
func (o *DcimInterfaceTemplatesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interface templates read default response has a 4xx status code
func (o *DcimInterfaceTemplatesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interface templates read default response has a 5xx status code
func (o *DcimInterfaceTemplatesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interface templates read default response a status code equal to that given
func (o *DcimInterfaceTemplatesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfaceTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/interface-templates/{id}/][%d] dcim_interface-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/interface-templates/{id}/][%d] dcim_interface-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfaceTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
