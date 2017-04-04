package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// SocialMessageFetchReader is a Reader for the SocialMessageFetch structure.
type SocialMessageFetchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SocialMessageFetchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSocialMessageFetchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSocialMessageFetchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSocialMessageFetchOK creates a SocialMessageFetchOK with default headers values
func NewSocialMessageFetchOK() *SocialMessageFetchOK {
	return &SocialMessageFetchOK{}
}

/*SocialMessageFetchOK handles this case with default header values.

Request processed successfully
*/
type SocialMessageFetchOK struct {
	Payload *models.DefaultResponse
}

func (o *SocialMessageFetchOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.fetch][%d] socialMessageFetchOK  %+v", 200, o.Payload)
}

func (o *SocialMessageFetchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSocialMessageFetchUnauthorized creates a SocialMessageFetchUnauthorized with default headers values
func NewSocialMessageFetchUnauthorized() *SocialMessageFetchUnauthorized {
	return &SocialMessageFetchUnauthorized{}
}

/*SocialMessageFetchUnauthorized handles this case with default header values.

Unauthorized request
*/
type SocialMessageFetchUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *SocialMessageFetchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.fetch][%d] socialMessageFetchUnauthorized  %+v", 401, o.Payload)
}

func (o *SocialMessageFetchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
