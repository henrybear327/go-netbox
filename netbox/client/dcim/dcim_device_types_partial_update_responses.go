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

// DcimDeviceTypesPartialUpdateReader is a Reader for the DcimDeviceTypesPartialUpdate structure.
type DcimDeviceTypesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceTypesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceTypesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceTypesPartialUpdateOK creates a DcimDeviceTypesPartialUpdateOK with default headers values
func NewDcimDeviceTypesPartialUpdateOK() *DcimDeviceTypesPartialUpdateOK {
	return &DcimDeviceTypesPartialUpdateOK{}
}

/*
DcimDeviceTypesPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceTypesPartialUpdateOK dcim device types partial update o k
*/
type DcimDeviceTypesPartialUpdateOK struct {
	Payload *models.DeviceType
}

// IsSuccess returns true when this dcim device types partial update o k response has a 2xx status code
func (o *DcimDeviceTypesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device types partial update o k response has a 3xx status code
func (o *DcimDeviceTypesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device types partial update o k response has a 4xx status code
func (o *DcimDeviceTypesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device types partial update o k response has a 5xx status code
func (o *DcimDeviceTypesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device types partial update o k response a status code equal to that given
func (o *DcimDeviceTypesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceTypesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/{id}/][%d] dcimDeviceTypesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceTypesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/{id}/][%d] dcimDeviceTypesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceTypesPartialUpdateOK) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceTypesPartialUpdateDefault creates a DcimDeviceTypesPartialUpdateDefault with default headers values
func NewDcimDeviceTypesPartialUpdateDefault(code int) *DcimDeviceTypesPartialUpdateDefault {
	return &DcimDeviceTypesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceTypesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceTypesPartialUpdateDefault dcim device types partial update default
*/
type DcimDeviceTypesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device types partial update default response
func (o *DcimDeviceTypesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device types partial update default response has a 2xx status code
func (o *DcimDeviceTypesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device types partial update default response has a 3xx status code
func (o *DcimDeviceTypesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device types partial update default response has a 4xx status code
func (o *DcimDeviceTypesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device types partial update default response has a 5xx status code
func (o *DcimDeviceTypesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device types partial update default response a status code equal to that given
func (o *DcimDeviceTypesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceTypesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/{id}/][%d] dcim_device-types_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/{id}/][%d] dcim_device-types_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceTypesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
