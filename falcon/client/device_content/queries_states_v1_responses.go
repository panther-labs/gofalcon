// Code generated by go-swagger; DO NOT EDIT.

package device_content

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

// QueriesStatesV1Reader is a Reader for the QueriesStatesV1 structure.
type QueriesStatesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueriesStatesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueriesStatesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueriesStatesV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueriesStatesV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueriesStatesV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /device-content/queries/states/v1] queries.states.v1", response, response.Code())
	}
}

// NewQueriesStatesV1OK creates a QueriesStatesV1OK with default headers values
func NewQueriesStatesV1OK() *QueriesStatesV1OK {
	return &QueriesStatesV1OK{}
}

/*
QueriesStatesV1OK describes a response with status code 200, with default header values.

OK
*/
type QueriesStatesV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DevicecontentapiQueryResponseV1
}

// IsSuccess returns true when this queries states v1 o k response has a 2xx status code
func (o *QueriesStatesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this queries states v1 o k response has a 3xx status code
func (o *QueriesStatesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this queries states v1 o k response has a 4xx status code
func (o *QueriesStatesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this queries states v1 o k response has a 5xx status code
func (o *QueriesStatesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this queries states v1 o k response a status code equal to that given
func (o *QueriesStatesV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the queries states v1 o k response
func (o *QueriesStatesV1OK) Code() int {
	return 200
}

func (o *QueriesStatesV1OK) Error() string {
	return fmt.Sprintf("[GET /device-content/queries/states/v1][%d] queriesStatesV1OK  %+v", 200, o.Payload)
}

func (o *QueriesStatesV1OK) String() string {
	return fmt.Sprintf("[GET /device-content/queries/states/v1][%d] queriesStatesV1OK  %+v", 200, o.Payload)
}

func (o *QueriesStatesV1OK) GetPayload() *models.DevicecontentapiQueryResponseV1 {
	return o.Payload
}

func (o *QueriesStatesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DevicecontentapiQueryResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueriesStatesV1Forbidden creates a QueriesStatesV1Forbidden with default headers values
func NewQueriesStatesV1Forbidden() *QueriesStatesV1Forbidden {
	return &QueriesStatesV1Forbidden{}
}

/*
QueriesStatesV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueriesStatesV1Forbidden struct {

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

// IsSuccess returns true when this queries states v1 forbidden response has a 2xx status code
func (o *QueriesStatesV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this queries states v1 forbidden response has a 3xx status code
func (o *QueriesStatesV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this queries states v1 forbidden response has a 4xx status code
func (o *QueriesStatesV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this queries states v1 forbidden response has a 5xx status code
func (o *QueriesStatesV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this queries states v1 forbidden response a status code equal to that given
func (o *QueriesStatesV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the queries states v1 forbidden response
func (o *QueriesStatesV1Forbidden) Code() int {
	return 403
}

func (o *QueriesStatesV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /device-content/queries/states/v1][%d] queriesStatesV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueriesStatesV1Forbidden) String() string {
	return fmt.Sprintf("[GET /device-content/queries/states/v1][%d] queriesStatesV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueriesStatesV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueriesStatesV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueriesStatesV1TooManyRequests creates a QueriesStatesV1TooManyRequests with default headers values
func NewQueriesStatesV1TooManyRequests() *QueriesStatesV1TooManyRequests {
	return &QueriesStatesV1TooManyRequests{}
}

/*
QueriesStatesV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueriesStatesV1TooManyRequests struct {

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

// IsSuccess returns true when this queries states v1 too many requests response has a 2xx status code
func (o *QueriesStatesV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this queries states v1 too many requests response has a 3xx status code
func (o *QueriesStatesV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this queries states v1 too many requests response has a 4xx status code
func (o *QueriesStatesV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this queries states v1 too many requests response has a 5xx status code
func (o *QueriesStatesV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this queries states v1 too many requests response a status code equal to that given
func (o *QueriesStatesV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the queries states v1 too many requests response
func (o *QueriesStatesV1TooManyRequests) Code() int {
	return 429
}

func (o *QueriesStatesV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /device-content/queries/states/v1][%d] queriesStatesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueriesStatesV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /device-content/queries/states/v1][%d] queriesStatesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueriesStatesV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueriesStatesV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueriesStatesV1InternalServerError creates a QueriesStatesV1InternalServerError with default headers values
func NewQueriesStatesV1InternalServerError() *QueriesStatesV1InternalServerError {
	return &QueriesStatesV1InternalServerError{}
}

/*
QueriesStatesV1InternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type QueriesStatesV1InternalServerError struct {

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

// IsSuccess returns true when this queries states v1 internal server error response has a 2xx status code
func (o *QueriesStatesV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this queries states v1 internal server error response has a 3xx status code
func (o *QueriesStatesV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this queries states v1 internal server error response has a 4xx status code
func (o *QueriesStatesV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this queries states v1 internal server error response has a 5xx status code
func (o *QueriesStatesV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this queries states v1 internal server error response a status code equal to that given
func (o *QueriesStatesV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the queries states v1 internal server error response
func (o *QueriesStatesV1InternalServerError) Code() int {
	return 500
}

func (o *QueriesStatesV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /device-content/queries/states/v1][%d] queriesStatesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *QueriesStatesV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /device-content/queries/states/v1][%d] queriesStatesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *QueriesStatesV1InternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueriesStatesV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
