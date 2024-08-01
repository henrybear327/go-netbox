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

// DcimPowerFeedsBulkUpdateReader is a Reader for the DcimPowerFeedsBulkUpdate structure.
type DcimPowerFeedsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsBulkUpdateOK creates a DcimPowerFeedsBulkUpdateOK with default headers values
func NewDcimPowerFeedsBulkUpdateOK() *DcimPowerFeedsBulkUpdateOK {
	return &DcimPowerFeedsBulkUpdateOK{}
}

/*
DcimPowerFeedsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerFeedsBulkUpdateOK dcim power feeds bulk update o k
*/
type DcimPowerFeedsBulkUpdateOK struct {
	Payload *models.PowerFeed
}

// IsSuccess returns true when this dcim power feeds bulk update o k response has a 2xx status code
func (o *DcimPowerFeedsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power feeds bulk update o k response has a 3xx status code
func (o *DcimPowerFeedsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds bulk update o k response has a 4xx status code
func (o *DcimPowerFeedsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power feeds bulk update o k response has a 5xx status code
func (o *DcimPowerFeedsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds bulk update o k response a status code equal to that given
func (o *DcimPowerFeedsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerFeedsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcimPowerFeedsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcimPowerFeedsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsBulkUpdateDefault creates a DcimPowerFeedsBulkUpdateDefault with default headers values
func NewDcimPowerFeedsBulkUpdateDefault(code int) *DcimPowerFeedsBulkUpdateDefault {
	return &DcimPowerFeedsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerFeedsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimPowerFeedsBulkUpdateDefault dcim power feeds bulk update default
*/
type DcimPowerFeedsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds bulk update default response
func (o *DcimPowerFeedsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power feeds bulk update default response has a 2xx status code
func (o *DcimPowerFeedsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power feeds bulk update default response has a 3xx status code
func (o *DcimPowerFeedsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power feeds bulk update default response has a 4xx status code
func (o *DcimPowerFeedsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power feeds bulk update default response has a 5xx status code
func (o *DcimPowerFeedsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power feeds bulk update default response a status code equal to that given
func (o *DcimPowerFeedsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerFeedsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcim_power-feeds_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcim_power-feeds_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
