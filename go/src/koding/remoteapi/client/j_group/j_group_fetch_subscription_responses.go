package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JGroupFetchSubscriptionReader is a Reader for the JGroupFetchSubscription structure.
type JGroupFetchSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JGroupFetchSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJGroupFetchSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJGroupFetchSubscriptionOK creates a JGroupFetchSubscriptionOK with default headers values
func NewJGroupFetchSubscriptionOK() *JGroupFetchSubscriptionOK {
	return &JGroupFetchSubscriptionOK{}
}

/*JGroupFetchSubscriptionOK handles this case with default header values.

OK
*/
type JGroupFetchSubscriptionOK struct {
	Payload JGroupFetchSubscriptionOKBody
}

func (o *JGroupFetchSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.fetchSubscription/{id}][%d] jGroupFetchSubscriptionOK  %+v", 200, o.Payload)
}

func (o *JGroupFetchSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JGroupFetchSubscriptionOKBody j group fetch subscription o k body
swagger:model JGroupFetchSubscriptionOKBody
*/
type JGroupFetchSubscriptionOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JGroupFetchSubscriptionOKBody) UnmarshalJSON(raw []byte) error {

	var jGroupFetchSubscriptionOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &jGroupFetchSubscriptionOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = jGroupFetchSubscriptionOKBodyAO0

	var jGroupFetchSubscriptionOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jGroupFetchSubscriptionOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jGroupFetchSubscriptionOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JGroupFetchSubscriptionOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jGroupFetchSubscriptionOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupFetchSubscriptionOKBodyAO0)

	jGroupFetchSubscriptionOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupFetchSubscriptionOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j group fetch subscription o k body
func (o *JGroupFetchSubscriptionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JGroup.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
