// Code generated by go-swagger; DO NOT EDIT.

package overwatch_dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// AggregatesEventsReader is a Reader for the AggregatesEvents structure.
type AggregatesEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatesEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatesEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregatesEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregatesEventsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregatesEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /overwatch-dashboards/aggregates/events/GET/v1] AggregatesEvents", response, response.Code())
	}
}

// NewAggregatesEventsOK creates a AggregatesEventsOK with default headers values
func NewAggregatesEventsOK() *AggregatesEventsOK {
	return &AggregatesEventsOK{}
}

/*
AggregatesEventsOK describes a response with status code 200, with default header values.

OK
*/
type AggregatesEventsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this aggregates events o k response has a 2xx status code
func (o *AggregatesEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregates events o k response has a 3xx status code
func (o *AggregatesEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates events o k response has a 4xx status code
func (o *AggregatesEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregates events o k response has a 5xx status code
func (o *AggregatesEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregates events o k response a status code equal to that given
func (o *AggregatesEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregates events o k response
func (o *AggregatesEventsOK) Code() int {
	return 200
}

func (o *AggregatesEventsOK) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events/GET/v1][%d] aggregatesEventsOK  %+v", 200, o.Payload)
}

func (o *AggregatesEventsOK) String() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events/GET/v1][%d] aggregatesEventsOK  %+v", 200, o.Payload)
}

func (o *AggregatesEventsOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregatesEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatesEventsForbidden creates a AggregatesEventsForbidden with default headers values
func NewAggregatesEventsForbidden() *AggregatesEventsForbidden {
	return &AggregatesEventsForbidden{}
}

/*
AggregatesEventsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregatesEventsForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this aggregates events forbidden response has a 2xx status code
func (o *AggregatesEventsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregates events forbidden response has a 3xx status code
func (o *AggregatesEventsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates events forbidden response has a 4xx status code
func (o *AggregatesEventsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregates events forbidden response has a 5xx status code
func (o *AggregatesEventsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregates events forbidden response a status code equal to that given
func (o *AggregatesEventsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregates events forbidden response
func (o *AggregatesEventsForbidden) Code() int {
	return 403
}

func (o *AggregatesEventsForbidden) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events/GET/v1][%d] aggregatesEventsForbidden  %+v", 403, o.Payload)
}

func (o *AggregatesEventsForbidden) String() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events/GET/v1][%d] aggregatesEventsForbidden  %+v", 403, o.Payload)
}

func (o *AggregatesEventsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatesEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatesEventsTooManyRequests creates a AggregatesEventsTooManyRequests with default headers values
func NewAggregatesEventsTooManyRequests() *AggregatesEventsTooManyRequests {
	return &AggregatesEventsTooManyRequests{}
}

/*
AggregatesEventsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregatesEventsTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this aggregates events too many requests response has a 2xx status code
func (o *AggregatesEventsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregates events too many requests response has a 3xx status code
func (o *AggregatesEventsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates events too many requests response has a 4xx status code
func (o *AggregatesEventsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregates events too many requests response has a 5xx status code
func (o *AggregatesEventsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregates events too many requests response a status code equal to that given
func (o *AggregatesEventsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregates events too many requests response
func (o *AggregatesEventsTooManyRequests) Code() int {
	return 429
}

func (o *AggregatesEventsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events/GET/v1][%d] aggregatesEventsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatesEventsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events/GET/v1][%d] aggregatesEventsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatesEventsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatesEventsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatesEventsInternalServerError creates a AggregatesEventsInternalServerError with default headers values
func NewAggregatesEventsInternalServerError() *AggregatesEventsInternalServerError {
	return &AggregatesEventsInternalServerError{}
}

/*
AggregatesEventsInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type AggregatesEventsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this aggregates events internal server error response has a 2xx status code
func (o *AggregatesEventsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregates events internal server error response has a 3xx status code
func (o *AggregatesEventsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates events internal server error response has a 4xx status code
func (o *AggregatesEventsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregates events internal server error response has a 5xx status code
func (o *AggregatesEventsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregates events internal server error response a status code equal to that given
func (o *AggregatesEventsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregates events internal server error response
func (o *AggregatesEventsInternalServerError) Code() int {
	return 500
}

func (o *AggregatesEventsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events/GET/v1][%d] aggregatesEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregatesEventsInternalServerError) String() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events/GET/v1][%d] aggregatesEventsInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregatesEventsInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatesEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
