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

// GetSchemasReader is a Reader for the GetSchemas structure.
type GetSchemasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchemasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetSchemasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSchemasOK creates a GetSchemasOK with default headers values
func NewGetSchemasOK() *GetSchemasOK {
	return &GetSchemasOK{}
}

/* GetSchemasOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSchemasOK struct {
	Payload []*models.Schema
}

func (o *GetSchemasOK) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] getSchemasOK  %+v", 200, o.Payload)
}
func (o *GetSchemasOK) GetPayload() []*models.Schema {
	return o.Payload
}

func (o *GetSchemasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemasInternalServerError creates a GetSchemasInternalServerError with default headers values
func NewGetSchemasInternalServerError() *GetSchemasInternalServerError {
	return &GetSchemasInternalServerError{}
}

/* GetSchemasInternalServerError describes a response with status code 500, with default header values.

Error code 50001 -- Error in the backend data store

*/
type GetSchemasInternalServerError struct {
}

func (o *GetSchemasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] getSchemasInternalServerError ", 500)
}

func (o *GetSchemasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
