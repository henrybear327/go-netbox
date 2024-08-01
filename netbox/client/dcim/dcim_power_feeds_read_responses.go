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

// DcimPowerFeedsReadReader is a Reader for the DcimPowerFeedsRead structure.
type DcimPowerFeedsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsReadOK creates a DcimPowerFeedsReadOK with default headers values
func NewDcimPowerFeedsReadOK() *DcimPowerFeedsReadOK {
	return &DcimPowerFeedsReadOK{}
}

/*
DcimPowerFeedsReadOK describes a response with status code 200, with default header values.

DcimPowerFeedsReadOK dcim power feeds read o k
*/
type DcimPowerFeedsReadOK struct {
	Payload *models.PowerFeed
}

// IsSuccess returns true when this dcim power feeds read o k response has a 2xx status code
func (o *DcimPowerFeedsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power feeds read o k response has a 3xx status code
func (o *DcimPowerFeedsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds read o k response has a 4xx status code
func (o *DcimPowerFeedsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power feeds read o k response has a 5xx status code
func (o *DcimPowerFeedsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds read o k response a status code equal to that given
func (o *DcimPowerFeedsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerFeedsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/][%d] dcimPowerFeedsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/][%d] dcimPowerFeedsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsReadOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsReadDefault creates a DcimPowerFeedsReadDefault with default headers values
func NewDcimPowerFeedsReadDefault(code int) *DcimPowerFeedsReadDefault {
	return &DcimPowerFeedsReadDefault{
		_statusCode: code,
	}
}

/*
DcimPowerFeedsReadDefault describes a response with status code -1, with default header values.

DcimPowerFeedsReadDefault dcim power feeds read default
*/
type DcimPowerFeedsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds read default response
func (o *DcimPowerFeedsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power feeds read default response has a 2xx status code
func (o *DcimPowerFeedsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power feeds read default response has a 3xx status code
func (o *DcimPowerFeedsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power feeds read default response has a 4xx status code
func (o *DcimPowerFeedsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power feeds read default response has a 5xx status code
func (o *DcimPowerFeedsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power feeds read default response a status code equal to that given
func (o *DcimPowerFeedsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerFeedsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/][%d] dcim_power-feeds_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/][%d] dcim_power-feeds_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
