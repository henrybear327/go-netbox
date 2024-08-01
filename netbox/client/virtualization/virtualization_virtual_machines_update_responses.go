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

// VirtualizationVirtualMachinesUpdateReader is a Reader for the VirtualizationVirtualMachinesUpdate structure.
type VirtualizationVirtualMachinesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualMachinesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualMachinesUpdateOK creates a VirtualizationVirtualMachinesUpdateOK with default headers values
func NewVirtualizationVirtualMachinesUpdateOK() *VirtualizationVirtualMachinesUpdateOK {
	return &VirtualizationVirtualMachinesUpdateOK{}
}

/*
VirtualizationVirtualMachinesUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesUpdateOK virtualization virtual machines update o k
*/
type VirtualizationVirtualMachinesUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

// IsSuccess returns true when this virtualization virtual machines update o k response has a 2xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual machines update o k response has a 3xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual machines update o k response has a 4xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual machines update o k response has a 5xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual machines update o k response a status code equal to that given
func (o *VirtualizationVirtualMachinesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationVirtualMachinesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualMachinesUpdateDefault creates a VirtualizationVirtualMachinesUpdateDefault with default headers values
func NewVirtualizationVirtualMachinesUpdateDefault(code int) *VirtualizationVirtualMachinesUpdateDefault {
	return &VirtualizationVirtualMachinesUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationVirtualMachinesUpdateDefault describes a response with status code -1, with default header values.

VirtualizationVirtualMachinesUpdateDefault virtualization virtual machines update default
*/
type VirtualizationVirtualMachinesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization virtual machines update default response
func (o *VirtualizationVirtualMachinesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization virtual machines update default response has a 2xx status code
func (o *VirtualizationVirtualMachinesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization virtual machines update default response has a 3xx status code
func (o *VirtualizationVirtualMachinesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization virtual machines update default response has a 4xx status code
func (o *VirtualizationVirtualMachinesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization virtual machines update default response has a 5xx status code
func (o *VirtualizationVirtualMachinesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization virtual machines update default response a status code equal to that given
func (o *VirtualizationVirtualMachinesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationVirtualMachinesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualization_virtual-machines_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualization_virtual-machines_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
