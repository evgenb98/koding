package j_name

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JNameClaimNamesReader is a Reader for the JNameClaimNames structure.
type JNameClaimNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JNameClaimNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJNameClaimNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJNameClaimNamesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJNameClaimNamesOK creates a JNameClaimNamesOK with default headers values
func NewJNameClaimNamesOK() *JNameClaimNamesOK {
	return &JNameClaimNamesOK{}
}

/*JNameClaimNamesOK handles this case with default header values.

Request processed successfully
*/
type JNameClaimNamesOK struct {
	Payload *models.DefaultResponse
}

func (o *JNameClaimNamesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JName.claimNames][%d] jNameClaimNamesOK  %+v", 200, o.Payload)
}

func (o *JNameClaimNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJNameClaimNamesUnauthorized creates a JNameClaimNamesUnauthorized with default headers values
func NewJNameClaimNamesUnauthorized() *JNameClaimNamesUnauthorized {
	return &JNameClaimNamesUnauthorized{}
}

/*JNameClaimNamesUnauthorized handles this case with default header values.

Unauthorized request
*/
type JNameClaimNamesUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JNameClaimNamesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JName.claimNames][%d] jNameClaimNamesUnauthorized  %+v", 401, o.Payload)
}

func (o *JNameClaimNamesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
