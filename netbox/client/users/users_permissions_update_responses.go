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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/henrybear327/go-netbox/netbox/models"
)

// UsersPermissionsUpdateReader is a Reader for the UsersPermissionsUpdate structure.
type UsersPermissionsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersPermissionsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersPermissionsUpdateOK creates a UsersPermissionsUpdateOK with default headers values
func NewUsersPermissionsUpdateOK() *UsersPermissionsUpdateOK {
	return &UsersPermissionsUpdateOK{}
}

/*
UsersPermissionsUpdateOK describes a response with status code 200, with default header values.

UsersPermissionsUpdateOK users permissions update o k
*/
type UsersPermissionsUpdateOK struct {
	Payload *models.ObjectPermission
}

// IsSuccess returns true when this users permissions update o k response has a 2xx status code
func (o *UsersPermissionsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users permissions update o k response has a 3xx status code
func (o *UsersPermissionsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users permissions update o k response has a 4xx status code
func (o *UsersPermissionsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users permissions update o k response has a 5xx status code
func (o *UsersPermissionsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users permissions update o k response a status code equal to that given
func (o *UsersPermissionsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersPermissionsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/permissions/{id}/][%d] usersPermissionsUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersPermissionsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /users/permissions/{id}/][%d] usersPermissionsUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersPermissionsUpdateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersPermissionsUpdateDefault creates a UsersPermissionsUpdateDefault with default headers values
func NewUsersPermissionsUpdateDefault(code int) *UsersPermissionsUpdateDefault {
	return &UsersPermissionsUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersPermissionsUpdateDefault describes a response with status code -1, with default header values.

UsersPermissionsUpdateDefault users permissions update default
*/
type UsersPermissionsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users permissions update default response
func (o *UsersPermissionsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users permissions update default response has a 2xx status code
func (o *UsersPermissionsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users permissions update default response has a 3xx status code
func (o *UsersPermissionsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users permissions update default response has a 4xx status code
func (o *UsersPermissionsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users permissions update default response has a 5xx status code
func (o *UsersPermissionsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users permissions update default response a status code equal to that given
func (o *UsersPermissionsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersPermissionsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /users/permissions/{id}/][%d] users_permissions_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /users/permissions/{id}/][%d] users_permissions_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersPermissionsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
