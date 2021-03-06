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

// JGroupFetchModeratorsReader is a Reader for the JGroupFetchModerators structure.
type JGroupFetchModeratorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JGroupFetchModeratorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJGroupFetchModeratorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJGroupFetchModeratorsOK creates a JGroupFetchModeratorsOK with default headers values
func NewJGroupFetchModeratorsOK() *JGroupFetchModeratorsOK {
	return &JGroupFetchModeratorsOK{}
}

/*JGroupFetchModeratorsOK handles this case with default header values.

OK
*/
type JGroupFetchModeratorsOK struct {
	Payload JGroupFetchModeratorsOKBody
}

func (o *JGroupFetchModeratorsOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.fetchModerators/{id}][%d] jGroupFetchModeratorsOK  %+v", 200, o.Payload)
}

func (o *JGroupFetchModeratorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JGroupFetchModeratorsOKBody j group fetch moderators o k body
swagger:model JGroupFetchModeratorsOKBody
*/
type JGroupFetchModeratorsOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JGroupFetchModeratorsOKBody) UnmarshalJSON(raw []byte) error {

	var jGroupFetchModeratorsOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &jGroupFetchModeratorsOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = jGroupFetchModeratorsOKBodyAO0

	var jGroupFetchModeratorsOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jGroupFetchModeratorsOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jGroupFetchModeratorsOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JGroupFetchModeratorsOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jGroupFetchModeratorsOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupFetchModeratorsOKBodyAO0)

	jGroupFetchModeratorsOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jGroupFetchModeratorsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j group fetch moderators o k body
func (o *JGroupFetchModeratorsOKBody) Validate(formats strfmt.Registry) error {
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
