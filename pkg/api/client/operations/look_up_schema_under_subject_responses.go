// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hineso11/go-schema-registry/pkg/api/models"
)

// LookUpSchemaUnderSubjectReader is a Reader for the LookUpSchemaUnderSubject structure.
type LookUpSchemaUnderSubjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookUpSchemaUnderSubjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 404:
		result := NewLookUpSchemaUnderSubjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLookUpSchemaUnderSubjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLookUpSchemaUnderSubjectNotFound creates a LookUpSchemaUnderSubjectNotFound with default headers values
func NewLookUpSchemaUnderSubjectNotFound() *LookUpSchemaUnderSubjectNotFound {
	return &LookUpSchemaUnderSubjectNotFound{}
}

/* LookUpSchemaUnderSubjectNotFound describes a response with status code 404, with default header values.

 Error code 40401 -- Subject not found
Error code 40403 -- Schema not found
*/
type LookUpSchemaUnderSubjectNotFound struct {
}

func (o *LookUpSchemaUnderSubjectNotFound) Error() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] lookUpSchemaUnderSubjectNotFound ", 404)
}

func (o *LookUpSchemaUnderSubjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLookUpSchemaUnderSubjectInternalServerError creates a LookUpSchemaUnderSubjectInternalServerError with default headers values
func NewLookUpSchemaUnderSubjectInternalServerError() *LookUpSchemaUnderSubjectInternalServerError {
	return &LookUpSchemaUnderSubjectInternalServerError{}
}

/* LookUpSchemaUnderSubjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LookUpSchemaUnderSubjectInternalServerError struct {
	Payload *models.Schema
}

func (o *LookUpSchemaUnderSubjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /subjects/{subject}][%d] lookUpSchemaUnderSubjectInternalServerError  %+v", 500, o.Payload)
}
func (o *LookUpSchemaUnderSubjectInternalServerError) GetPayload() *models.Schema {
	return o.Payload
}

func (o *LookUpSchemaUnderSubjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Schema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
