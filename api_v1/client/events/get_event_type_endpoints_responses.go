// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/AISphere/ffdl-rest-apis/api_v1/restmodels"
)

// GetEventTypeEndpointsReader is a Reader for the GetEventTypeEndpoints structure.
type GetEventTypeEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventTypeEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventTypeEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetEventTypeEndpointsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetEventTypeEndpointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 410:
		result := NewGetEventTypeEndpointsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventTypeEndpointsOK creates a GetEventTypeEndpointsOK with default headers values
func NewGetEventTypeEndpointsOK() *GetEventTypeEndpointsOK {
	return &GetEventTypeEndpointsOK{}
}

/*GetEventTypeEndpointsOK handles this case with default header values.

Event endpoints
*/
type GetEventTypeEndpointsOK struct {
	Payload *restmodels.EndpointList
}

func (o *GetEventTypeEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/events/{event_type}][%d] getEventTypeEndpointsOK  %+v", 200, o.Payload)
}

func (o *GetEventTypeEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.EndpointList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventTypeEndpointsUnauthorized creates a GetEventTypeEndpointsUnauthorized with default headers values
func NewGetEventTypeEndpointsUnauthorized() *GetEventTypeEndpointsUnauthorized {
	return &GetEventTypeEndpointsUnauthorized{}
}

/*GetEventTypeEndpointsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEventTypeEndpointsUnauthorized struct {
	Payload *restmodels.Error
}

func (o *GetEventTypeEndpointsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/events/{event_type}][%d] getEventTypeEndpointsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetEventTypeEndpointsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventTypeEndpointsNotFound creates a GetEventTypeEndpointsNotFound with default headers values
func NewGetEventTypeEndpointsNotFound() *GetEventTypeEndpointsNotFound {
	return &GetEventTypeEndpointsNotFound{}
}

/*GetEventTypeEndpointsNotFound handles this case with default header values.

The model or event type cannot be found.
*/
type GetEventTypeEndpointsNotFound struct {
	Payload *restmodels.Error
}

func (o *GetEventTypeEndpointsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/events/{event_type}][%d] getEventTypeEndpointsNotFound  %+v", 404, o.Payload)
}

func (o *GetEventTypeEndpointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventTypeEndpointsGone creates a GetEventTypeEndpointsGone with default headers values
func NewGetEventTypeEndpointsGone() *GetEventTypeEndpointsGone {
	return &GetEventTypeEndpointsGone{}
}

/*GetEventTypeEndpointsGone handles this case with default header values.

If the trained model storage time has expired and it has been deleted. It only gets deleted if it not stored on an external data store.
*/
type GetEventTypeEndpointsGone struct {
	Payload *restmodels.Error
}

func (o *GetEventTypeEndpointsGone) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/events/{event_type}][%d] getEventTypeEndpointsGone  %+v", 410, o.Payload)
}

func (o *GetEventTypeEndpointsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
