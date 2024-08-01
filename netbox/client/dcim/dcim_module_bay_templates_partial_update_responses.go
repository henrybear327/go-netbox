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

// DcimModuleBayTemplatesPartialUpdateReader is a Reader for the DcimModuleBayTemplatesPartialUpdate structure.
type DcimModuleBayTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBayTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBayTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBayTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBayTemplatesPartialUpdateOK creates a DcimModuleBayTemplatesPartialUpdateOK with default headers values
func NewDcimModuleBayTemplatesPartialUpdateOK() *DcimModuleBayTemplatesPartialUpdateOK {
	return &DcimModuleBayTemplatesPartialUpdateOK{}
}

/*
DcimModuleBayTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimModuleBayTemplatesPartialUpdateOK dcim module bay templates partial update o k
*/
type DcimModuleBayTemplatesPartialUpdateOK struct {
	Payload *models.ModuleBayTemplate
}

// IsSuccess returns true when this dcim module bay templates partial update o k response has a 2xx status code
func (o *DcimModuleBayTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bay templates partial update o k response has a 3xx status code
func (o *DcimModuleBayTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bay templates partial update o k response has a 4xx status code
func (o *DcimModuleBayTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bay templates partial update o k response has a 5xx status code
func (o *DcimModuleBayTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bay templates partial update o k response a status code equal to that given
func (o *DcimModuleBayTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleBayTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBayTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBayTemplatesPartialUpdateOK) GetPayload() *models.ModuleBayTemplate {
	return o.Payload
}

func (o *DcimModuleBayTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBayTemplatesPartialUpdateDefault creates a DcimModuleBayTemplatesPartialUpdateDefault with default headers values
func NewDcimModuleBayTemplatesPartialUpdateDefault(code int) *DcimModuleBayTemplatesPartialUpdateDefault {
	return &DcimModuleBayTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBayTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimModuleBayTemplatesPartialUpdateDefault dcim module bay templates partial update default
*/
type DcimModuleBayTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bay templates partial update default response
func (o *DcimModuleBayTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bay templates partial update default response has a 2xx status code
func (o *DcimModuleBayTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bay templates partial update default response has a 3xx status code
func (o *DcimModuleBayTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bay templates partial update default response has a 4xx status code
func (o *DcimModuleBayTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bay templates partial update default response has a 5xx status code
func (o *DcimModuleBayTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bay templates partial update default response a status code equal to that given
func (o *DcimModuleBayTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBayTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-bay-templates/{id}/][%d] dcim_module-bay-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/module-bay-templates/{id}/][%d] dcim_module-bay-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBayTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
