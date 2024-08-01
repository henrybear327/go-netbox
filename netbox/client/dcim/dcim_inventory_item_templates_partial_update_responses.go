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

// DcimInventoryItemTemplatesPartialUpdateReader is a Reader for the DcimInventoryItemTemplatesPartialUpdate structure.
type DcimInventoryItemTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemTemplatesPartialUpdateOK creates a DcimInventoryItemTemplatesPartialUpdateOK with default headers values
func NewDcimInventoryItemTemplatesPartialUpdateOK() *DcimInventoryItemTemplatesPartialUpdateOK {
	return &DcimInventoryItemTemplatesPartialUpdateOK{}
}

/*
DcimInventoryItemTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemTemplatesPartialUpdateOK dcim inventory item templates partial update o k
*/
type DcimInventoryItemTemplatesPartialUpdateOK struct {
	Payload *models.InventoryItemTemplate
}

// IsSuccess returns true when this dcim inventory item templates partial update o k response has a 2xx status code
func (o *DcimInventoryItemTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item templates partial update o k response has a 3xx status code
func (o *DcimInventoryItemTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item templates partial update o k response has a 4xx status code
func (o *DcimInventoryItemTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item templates partial update o k response has a 5xx status code
func (o *DcimInventoryItemTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item templates partial update o k response a status code equal to that given
func (o *DcimInventoryItemTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInventoryItemTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/{id}/][%d] dcimInventoryItemTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/{id}/][%d] dcimInventoryItemTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesPartialUpdateOK) GetPayload() *models.InventoryItemTemplate {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemTemplatesPartialUpdateDefault creates a DcimInventoryItemTemplatesPartialUpdateDefault with default headers values
func NewDcimInventoryItemTemplatesPartialUpdateDefault(code int) *DcimInventoryItemTemplatesPartialUpdateDefault {
	return &DcimInventoryItemTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemTemplatesPartialUpdateDefault dcim inventory item templates partial update default
*/
type DcimInventoryItemTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item templates partial update default response
func (o *DcimInventoryItemTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim inventory item templates partial update default response has a 2xx status code
func (o *DcimInventoryItemTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory item templates partial update default response has a 3xx status code
func (o *DcimInventoryItemTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory item templates partial update default response has a 4xx status code
func (o *DcimInventoryItemTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory item templates partial update default response has a 5xx status code
func (o *DcimInventoryItemTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory item templates partial update default response a status code equal to that given
func (o *DcimInventoryItemTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInventoryItemTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/{id}/][%d] dcim_inventory-item-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/{id}/][%d] dcim_inventory-item-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
