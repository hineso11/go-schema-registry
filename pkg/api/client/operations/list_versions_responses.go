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

// ListVersionsReader is a Reader for the ListVersions structure.
type ListVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListVersionsOK creates a ListVersionsOK with default headers values
func NewListVersionsOK() *ListVersionsOK {
	return &ListVersionsOK{}
}

/* ListVersionsOK describes a response with status code 200, with default header values.

successful operation
*/
type ListVersionsOK struct {
	Payload []int32
}

func (o *ListVersionsOK) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions][%d] listVersionsOK  %+v", 200, o.Payload)
}
func (o *ListVersionsOK) GetPayload() []int32 {
	return o.Payload
}

func (o *ListVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVersionsNotFound creates a ListVersionsNotFound with default headers values
func NewListVersionsNotFound() *ListVersionsNotFound {
	return &ListVersionsNotFound{}
}

/* ListVersionsNotFound describes a response with status code 404, with default header values.

Error code 40401 -- Subject not found
*/
type ListVersionsNotFound struct {
}

func (o *ListVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions][%d] listVersionsNotFound ", 404)
}

func (o *ListVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListVersionsInternalServerError creates a ListVersionsInternalServerError with default headers values
func NewListVersionsInternalServerError() *ListVersionsInternalServerError {
	return &ListVersionsInternalServerError{}
}

/* ListVersionsInternalServerError describes a response with status code 500, with default header values.

Error code 50001 -- Error in the backend data store
*/
type ListVersionsInternalServerError struct {
}

func (o *ListVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions][%d] listVersionsInternalServerError ", 500)
}

func (o *ListVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
