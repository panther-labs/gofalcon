// Code generated by go-swagger; DO NOT EDIT.

package foundry_logscale

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

// ExecuteDynamicReader is a Reader for the ExecuteDynamic structure.
type ExecuteDynamicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteDynamicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecuteDynamicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExecuteDynamicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExecuteDynamicForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExecuteDynamicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewExecuteDynamicTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExecuteDynamicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1] ExecuteDynamic", response, response.Code())
	}
}

// NewExecuteDynamicOK creates a ExecuteDynamicOK with default headers values
func NewExecuteDynamicOK() *ExecuteDynamicOK {
	return &ExecuteDynamicOK{}
}

/*
ExecuteDynamicOK describes a response with status code 200, with default header values.

OK
*/
type ExecuteDynamicOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ApidomainQueryResponseWrapperV1
}

// IsSuccess returns true when this execute dynamic o k response has a 2xx status code
func (o *ExecuteDynamicOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this execute dynamic o k response has a 3xx status code
func (o *ExecuteDynamicOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute dynamic o k response has a 4xx status code
func (o *ExecuteDynamicOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this execute dynamic o k response has a 5xx status code
func (o *ExecuteDynamicOK) IsServerError() bool {
	return false
}

// IsCode returns true when this execute dynamic o k response a status code equal to that given
func (o *ExecuteDynamicOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the execute dynamic o k response
func (o *ExecuteDynamicOK) Code() int {
	return 200
}

func (o *ExecuteDynamicOK) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicOK  %+v", 200, o.Payload)
}

func (o *ExecuteDynamicOK) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicOK  %+v", 200, o.Payload)
}

func (o *ExecuteDynamicOK) GetPayload() *models.ApidomainQueryResponseWrapperV1 {
	return o.Payload
}

func (o *ExecuteDynamicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ApidomainQueryResponseWrapperV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteDynamicBadRequest creates a ExecuteDynamicBadRequest with default headers values
func NewExecuteDynamicBadRequest() *ExecuteDynamicBadRequest {
	return &ExecuteDynamicBadRequest{}
}

/*
ExecuteDynamicBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExecuteDynamicBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this execute dynamic bad request response has a 2xx status code
func (o *ExecuteDynamicBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute dynamic bad request response has a 3xx status code
func (o *ExecuteDynamicBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute dynamic bad request response has a 4xx status code
func (o *ExecuteDynamicBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this execute dynamic bad request response has a 5xx status code
func (o *ExecuteDynamicBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this execute dynamic bad request response a status code equal to that given
func (o *ExecuteDynamicBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the execute dynamic bad request response
func (o *ExecuteDynamicBadRequest) Code() int {
	return 400
}

func (o *ExecuteDynamicBadRequest) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicBadRequest  %+v", 400, o.Payload)
}

func (o *ExecuteDynamicBadRequest) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicBadRequest  %+v", 400, o.Payload)
}

func (o *ExecuteDynamicBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *ExecuteDynamicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteDynamicForbidden creates a ExecuteDynamicForbidden with default headers values
func NewExecuteDynamicForbidden() *ExecuteDynamicForbidden {
	return &ExecuteDynamicForbidden{}
}

/*
ExecuteDynamicForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExecuteDynamicForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this execute dynamic forbidden response has a 2xx status code
func (o *ExecuteDynamicForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute dynamic forbidden response has a 3xx status code
func (o *ExecuteDynamicForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute dynamic forbidden response has a 4xx status code
func (o *ExecuteDynamicForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this execute dynamic forbidden response has a 5xx status code
func (o *ExecuteDynamicForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this execute dynamic forbidden response a status code equal to that given
func (o *ExecuteDynamicForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the execute dynamic forbidden response
func (o *ExecuteDynamicForbidden) Code() int {
	return 403
}

func (o *ExecuteDynamicForbidden) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicForbidden  %+v", 403, o.Payload)
}

func (o *ExecuteDynamicForbidden) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicForbidden  %+v", 403, o.Payload)
}

func (o *ExecuteDynamicForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *ExecuteDynamicForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteDynamicNotFound creates a ExecuteDynamicNotFound with default headers values
func NewExecuteDynamicNotFound() *ExecuteDynamicNotFound {
	return &ExecuteDynamicNotFound{}
}

/*
ExecuteDynamicNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ExecuteDynamicNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this execute dynamic not found response has a 2xx status code
func (o *ExecuteDynamicNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute dynamic not found response has a 3xx status code
func (o *ExecuteDynamicNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute dynamic not found response has a 4xx status code
func (o *ExecuteDynamicNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this execute dynamic not found response has a 5xx status code
func (o *ExecuteDynamicNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this execute dynamic not found response a status code equal to that given
func (o *ExecuteDynamicNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the execute dynamic not found response
func (o *ExecuteDynamicNotFound) Code() int {
	return 404
}

func (o *ExecuteDynamicNotFound) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicNotFound  %+v", 404, o.Payload)
}

func (o *ExecuteDynamicNotFound) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicNotFound  %+v", 404, o.Payload)
}

func (o *ExecuteDynamicNotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *ExecuteDynamicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteDynamicTooManyRequests creates a ExecuteDynamicTooManyRequests with default headers values
func NewExecuteDynamicTooManyRequests() *ExecuteDynamicTooManyRequests {
	return &ExecuteDynamicTooManyRequests{}
}

/*
ExecuteDynamicTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ExecuteDynamicTooManyRequests struct {

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

// IsSuccess returns true when this execute dynamic too many requests response has a 2xx status code
func (o *ExecuteDynamicTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute dynamic too many requests response has a 3xx status code
func (o *ExecuteDynamicTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute dynamic too many requests response has a 4xx status code
func (o *ExecuteDynamicTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this execute dynamic too many requests response has a 5xx status code
func (o *ExecuteDynamicTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this execute dynamic too many requests response a status code equal to that given
func (o *ExecuteDynamicTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the execute dynamic too many requests response
func (o *ExecuteDynamicTooManyRequests) Code() int {
	return 429
}

func (o *ExecuteDynamicTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExecuteDynamicTooManyRequests) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExecuteDynamicTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExecuteDynamicTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExecuteDynamicInternalServerError creates a ExecuteDynamicInternalServerError with default headers values
func NewExecuteDynamicInternalServerError() *ExecuteDynamicInternalServerError {
	return &ExecuteDynamicInternalServerError{}
}

/*
ExecuteDynamicInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExecuteDynamicInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this execute dynamic internal server error response has a 2xx status code
func (o *ExecuteDynamicInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute dynamic internal server error response has a 3xx status code
func (o *ExecuteDynamicInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute dynamic internal server error response has a 4xx status code
func (o *ExecuteDynamicInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this execute dynamic internal server error response has a 5xx status code
func (o *ExecuteDynamicInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this execute dynamic internal server error response a status code equal to that given
func (o *ExecuteDynamicInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the execute dynamic internal server error response
func (o *ExecuteDynamicInternalServerError) Code() int {
	return 500
}

func (o *ExecuteDynamicInternalServerError) Error() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecuteDynamicInternalServerError) String() string {
	return fmt.Sprintf("[POST /loggingapi/entities/saved-searches/execute-dynamic/v1][%d] executeDynamicInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecuteDynamicInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *ExecuteDynamicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
