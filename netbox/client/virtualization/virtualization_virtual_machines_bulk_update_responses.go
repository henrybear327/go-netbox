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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/henrybear327/go-netbox/netbox/models"
)

// VirtualizationVirtualMachinesBulkUpdateReader is a Reader for the VirtualizationVirtualMachinesBulkUpdate structure.
type VirtualizationVirtualMachinesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualMachinesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualMachinesBulkUpdateOK creates a VirtualizationVirtualMachinesBulkUpdateOK with default headers values
func NewVirtualizationVirtualMachinesBulkUpdateOK() *VirtualizationVirtualMachinesBulkUpdateOK {
	return &VirtualizationVirtualMachinesBulkUpdateOK{}
}

/*
VirtualizationVirtualMachinesBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesBulkUpdateOK virtualization virtual machines bulk update o k
*/
type VirtualizationVirtualMachinesBulkUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

// IsSuccess returns true when this virtualization virtual machines bulk update o k response has a 2xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual machines bulk update o k response has a 3xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual machines bulk update o k response has a 4xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual machines bulk update o k response has a 5xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual machines bulk update o k response a status code equal to that given
func (o *VirtualizationVirtualMachinesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationVirtualMachinesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/][%d] virtualizationVirtualMachinesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/][%d] virtualizationVirtualMachinesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesBulkUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualMachinesBulkUpdateDefault creates a VirtualizationVirtualMachinesBulkUpdateDefault with default headers values
func NewVirtualizationVirtualMachinesBulkUpdateDefault(code int) *VirtualizationVirtualMachinesBulkUpdateDefault {
	return &VirtualizationVirtualMachinesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationVirtualMachinesBulkUpdateDefault describes a response with status code -1, with default header values.

VirtualizationVirtualMachinesBulkUpdateDefault virtualization virtual machines bulk update default
*/
type VirtualizationVirtualMachinesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization virtual machines bulk update default response
func (o *VirtualizationVirtualMachinesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization virtual machines bulk update default response has a 2xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization virtual machines bulk update default response has a 3xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization virtual machines bulk update default response has a 4xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization virtual machines bulk update default response has a 5xx status code
func (o *VirtualizationVirtualMachinesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization virtual machines bulk update default response a status code equal to that given
func (o *VirtualizationVirtualMachinesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationVirtualMachinesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/][%d] virtualization_virtual-machines_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/][%d] virtualization_virtual-machines_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
