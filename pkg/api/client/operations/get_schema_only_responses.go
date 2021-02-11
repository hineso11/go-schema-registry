// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSchemaOnlyReader is a Reader for the GetSchemaOnly structure.
type GetSchemaOnlyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemaOnlyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchemaOnlyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSchemaOnlyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetSchemaOnlyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSchemaOnlyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSchemaOnlyOK creates a GetSchemaOnlyOK with default headers values
func NewGetSchemaOnlyOK() *GetSchemaOnlyOK {
	return &GetSchemaOnlyOK{}
}

/* GetSchemaOnlyOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSchemaOnlyOK struct {
	Payload string
}

func (o *GetSchemaOnlyOK) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions/{version}/schema][%d] getSchemaOnlyOK  %+v", 200, o.Payload)
}
func (o *GetSchemaOnlyOK) GetPayload() string {
	return o.Payload
}

func (o *GetSchemaOnlyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemaOnlyNotFound creates a GetSchemaOnlyNotFound with default headers values
func NewGetSchemaOnlyNotFound() *GetSchemaOnlyNotFound {
	return &GetSchemaOnlyNotFound{}
}

/* GetSchemaOnlyNotFound describes a response with status code 404, with default header values.

 Error code 40401 -- Subject not found
Error code 40402 -- Version not found
*/
type GetSchemaOnlyNotFound struct {
}

func (o *GetSchemaOnlyNotFound) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions/{version}/schema][%d] getSchemaOnlyNotFound ", 404)
}

func (o *GetSchemaOnlyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemaOnlyUnprocessableEntity creates a GetSchemaOnlyUnprocessableEntity with default headers values
func NewGetSchemaOnlyUnprocessableEntity() *GetSchemaOnlyUnprocessableEntity {
	return &GetSchemaOnlyUnprocessableEntity{}
}

/* GetSchemaOnlyUnprocessableEntity describes a response with status code 422, with default header values.

Error code 42202 -- Invalid version
*/
type GetSchemaOnlyUnprocessableEntity struct {
}

func (o *GetSchemaOnlyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions/{version}/schema][%d] getSchemaOnlyUnprocessableEntity ", 422)
}

func (o *GetSchemaOnlyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemaOnlyInternalServerError creates a GetSchemaOnlyInternalServerError with default headers values
func NewGetSchemaOnlyInternalServerError() *GetSchemaOnlyInternalServerError {
	return &GetSchemaOnlyInternalServerError{}
}

/* GetSchemaOnlyInternalServerError describes a response with status code 500, with default header values.

Error code 50001 -- Error in the backend data store
*/
type GetSchemaOnlyInternalServerError struct {
}

func (o *GetSchemaOnlyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions/{version}/schema][%d] getSchemaOnlyInternalServerError ", 500)
}

func (o *GetSchemaOnlyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}