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

// DcimModuleTypesPartialUpdateReader is a Reader for the DcimModuleTypesPartialUpdate structure.
type DcimModuleTypesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleTypesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleTypesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleTypesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleTypesPartialUpdateOK creates a DcimModuleTypesPartialUpdateOK with default headers values
func NewDcimModuleTypesPartialUpdateOK() *DcimModuleTypesPartialUpdateOK {
	return &DcimModuleTypesPartialUpdateOK{}
}

/*
DcimModuleTypesPartialUpdateOK describes a response with status code 200, with default header values.

DcimModuleTypesPartialUpdateOK dcim module types partial update o k
*/
type DcimModuleTypesPartialUpdateOK struct {
	Payload *models.ModuleType
}

// IsSuccess returns true when this dcim module types partial update o k response has a 2xx status code
func (o *DcimModuleTypesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module types partial update o k response has a 3xx status code
func (o *DcimModuleTypesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module types partial update o k response has a 4xx status code
func (o *DcimModuleTypesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module types partial update o k response has a 5xx status code
func (o *DcimModuleTypesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module types partial update o k response a status code equal to that given
func (o *DcimModuleTypesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleTypesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-types/{id}/][%d] dcimModuleTypesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleTypesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/module-types/{id}/][%d] dcimModuleTypesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleTypesPartialUpdateOK) GetPayload() *models.ModuleType {
	return o.Payload
}

func (o *DcimModuleTypesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleTypesPartialUpdateDefault creates a DcimModuleTypesPartialUpdateDefault with default headers values
func NewDcimModuleTypesPartialUpdateDefault(code int) *DcimModuleTypesPartialUpdateDefault {
	return &DcimModuleTypesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleTypesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimModuleTypesPartialUpdateDefault dcim module types partial update default
*/
type DcimModuleTypesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module types partial update default response
func (o *DcimModuleTypesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module types partial update default response has a 2xx status code
func (o *DcimModuleTypesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module types partial update default response has a 3xx status code
func (o *DcimModuleTypesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module types partial update default response has a 4xx status code
func (o *DcimModuleTypesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module types partial update default response has a 5xx status code
func (o *DcimModuleTypesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module types partial update default response a status code equal to that given
func (o *DcimModuleTypesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleTypesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-types/{id}/][%d] dcim_module-types_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/module-types/{id}/][%d] dcim_module-types_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleTypesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
