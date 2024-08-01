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

// DcimDeviceTypesCreateReader is a Reader for the DcimDeviceTypesCreate structure.
type DcimDeviceTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimDeviceTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceTypesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceTypesCreateCreated creates a DcimDeviceTypesCreateCreated with default headers values
func NewDcimDeviceTypesCreateCreated() *DcimDeviceTypesCreateCreated {
	return &DcimDeviceTypesCreateCreated{}
}

/*
DcimDeviceTypesCreateCreated describes a response with status code 201, with default header values.

DcimDeviceTypesCreateCreated dcim device types create created
*/
type DcimDeviceTypesCreateCreated struct {
	Payload *models.DeviceType
}

// IsSuccess returns true when this dcim device types create created response has a 2xx status code
func (o *DcimDeviceTypesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device types create created response has a 3xx status code
func (o *DcimDeviceTypesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device types create created response has a 4xx status code
func (o *DcimDeviceTypesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device types create created response has a 5xx status code
func (o *DcimDeviceTypesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device types create created response a status code equal to that given
func (o *DcimDeviceTypesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimDeviceTypesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/device-types/][%d] dcimDeviceTypesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimDeviceTypesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/device-types/][%d] dcimDeviceTypesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimDeviceTypesCreateCreated) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceTypesCreateDefault creates a DcimDeviceTypesCreateDefault with default headers values
func NewDcimDeviceTypesCreateDefault(code int) *DcimDeviceTypesCreateDefault {
	return &DcimDeviceTypesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceTypesCreateDefault describes a response with status code -1, with default header values.

DcimDeviceTypesCreateDefault dcim device types create default
*/
type DcimDeviceTypesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device types create default response
func (o *DcimDeviceTypesCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device types create default response has a 2xx status code
func (o *DcimDeviceTypesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device types create default response has a 3xx status code
func (o *DcimDeviceTypesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device types create default response has a 4xx status code
func (o *DcimDeviceTypesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device types create default response has a 5xx status code
func (o *DcimDeviceTypesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device types create default response a status code equal to that given
func (o *DcimDeviceTypesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceTypesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/device-types/][%d] dcim_device-types_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/device-types/][%d] dcim_device-types_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceTypesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
