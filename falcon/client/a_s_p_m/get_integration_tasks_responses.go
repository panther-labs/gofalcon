// Code generated by go-swagger; DO NOT EDIT.

package a_s_p_m

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

// GetIntegrationTasksReader is a Reader for the GetIntegrationTasks structure.
type GetIntegrationTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationTasksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /aspm-api-gateway/api/v1/integration_tasks] GetIntegrationTasks", response, response.Code())
	}
}

// NewGetIntegrationTasksOK creates a GetIntegrationTasksOK with default headers values
func NewGetIntegrationTasksOK() *GetIntegrationTasksOK {
	return &GetIntegrationTasksOK{}
}

/*
GetIntegrationTasksOK describes a response with status code 200, with default header values.

OK
*/
type GetIntegrationTasksOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesListIntegrationTasksResponse
}

// IsSuccess returns true when this get integration tasks o k response has a 2xx status code
func (o *GetIntegrationTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integration tasks o k response has a 3xx status code
func (o *GetIntegrationTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration tasks o k response has a 4xx status code
func (o *GetIntegrationTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integration tasks o k response has a 5xx status code
func (o *GetIntegrationTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration tasks o k response a status code equal to that given
func (o *GetIntegrationTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get integration tasks o k response
func (o *GetIntegrationTasksOK) Code() int {
	return 200
}

func (o *GetIntegrationTasksOK) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationTasksOK) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationTasksOK) GetPayload() *models.TypesListIntegrationTasksResponse {
	return o.Payload
}

func (o *GetIntegrationTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.TypesListIntegrationTasksResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationTasksBadRequest creates a GetIntegrationTasksBadRequest with default headers values
func NewGetIntegrationTasksBadRequest() *GetIntegrationTasksBadRequest {
	return &GetIntegrationTasksBadRequest{}
}

/*
GetIntegrationTasksBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetIntegrationTasksBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesErrorMessage
}

// IsSuccess returns true when this get integration tasks bad request response has a 2xx status code
func (o *GetIntegrationTasksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration tasks bad request response has a 3xx status code
func (o *GetIntegrationTasksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration tasks bad request response has a 4xx status code
func (o *GetIntegrationTasksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration tasks bad request response has a 5xx status code
func (o *GetIntegrationTasksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration tasks bad request response a status code equal to that given
func (o *GetIntegrationTasksBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get integration tasks bad request response
func (o *GetIntegrationTasksBadRequest) Code() int {
	return 400
}

func (o *GetIntegrationTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationTasksBadRequest) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationTasksBadRequest) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *GetIntegrationTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.TypesErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationTasksUnauthorized creates a GetIntegrationTasksUnauthorized with default headers values
func NewGetIntegrationTasksUnauthorized() *GetIntegrationTasksUnauthorized {
	return &GetIntegrationTasksUnauthorized{}
}

/*
GetIntegrationTasksUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetIntegrationTasksUnauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesErrorMessage
}

// IsSuccess returns true when this get integration tasks unauthorized response has a 2xx status code
func (o *GetIntegrationTasksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration tasks unauthorized response has a 3xx status code
func (o *GetIntegrationTasksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration tasks unauthorized response has a 4xx status code
func (o *GetIntegrationTasksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration tasks unauthorized response has a 5xx status code
func (o *GetIntegrationTasksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration tasks unauthorized response a status code equal to that given
func (o *GetIntegrationTasksUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get integration tasks unauthorized response
func (o *GetIntegrationTasksUnauthorized) Code() int {
	return 401
}

func (o *GetIntegrationTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationTasksUnauthorized) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationTasksUnauthorized) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *GetIntegrationTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.TypesErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationTasksForbidden creates a GetIntegrationTasksForbidden with default headers values
func NewGetIntegrationTasksForbidden() *GetIntegrationTasksForbidden {
	return &GetIntegrationTasksForbidden{}
}

/*
GetIntegrationTasksForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetIntegrationTasksForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesErrorMessage
}

// IsSuccess returns true when this get integration tasks forbidden response has a 2xx status code
func (o *GetIntegrationTasksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration tasks forbidden response has a 3xx status code
func (o *GetIntegrationTasksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration tasks forbidden response has a 4xx status code
func (o *GetIntegrationTasksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration tasks forbidden response has a 5xx status code
func (o *GetIntegrationTasksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration tasks forbidden response a status code equal to that given
func (o *GetIntegrationTasksForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get integration tasks forbidden response
func (o *GetIntegrationTasksForbidden) Code() int {
	return 403
}

func (o *GetIntegrationTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationTasksForbidden) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationTasksForbidden) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *GetIntegrationTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.TypesErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationTasksTooManyRequests creates a GetIntegrationTasksTooManyRequests with default headers values
func NewGetIntegrationTasksTooManyRequests() *GetIntegrationTasksTooManyRequests {
	return &GetIntegrationTasksTooManyRequests{}
}

/*
GetIntegrationTasksTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetIntegrationTasksTooManyRequests struct {

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

// IsSuccess returns true when this get integration tasks too many requests response has a 2xx status code
func (o *GetIntegrationTasksTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration tasks too many requests response has a 3xx status code
func (o *GetIntegrationTasksTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration tasks too many requests response has a 4xx status code
func (o *GetIntegrationTasksTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration tasks too many requests response has a 5xx status code
func (o *GetIntegrationTasksTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration tasks too many requests response a status code equal to that given
func (o *GetIntegrationTasksTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get integration tasks too many requests response
func (o *GetIntegrationTasksTooManyRequests) Code() int {
	return 429
}

func (o *GetIntegrationTasksTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationTasksTooManyRequests) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationTasksTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIntegrationTasksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIntegrationTasksInternalServerError creates a GetIntegrationTasksInternalServerError with default headers values
func NewGetIntegrationTasksInternalServerError() *GetIntegrationTasksInternalServerError {
	return &GetIntegrationTasksInternalServerError{}
}

/*
GetIntegrationTasksInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetIntegrationTasksInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesErrorMessage
}

// IsSuccess returns true when this get integration tasks internal server error response has a 2xx status code
func (o *GetIntegrationTasksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration tasks internal server error response has a 3xx status code
func (o *GetIntegrationTasksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration tasks internal server error response has a 4xx status code
func (o *GetIntegrationTasksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integration tasks internal server error response has a 5xx status code
func (o *GetIntegrationTasksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integration tasks internal server error response a status code equal to that given
func (o *GetIntegrationTasksInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get integration tasks internal server error response
func (o *GetIntegrationTasksInternalServerError) Code() int {
	return 500
}

func (o *GetIntegrationTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationTasksInternalServerError) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/integration_tasks][%d] getIntegrationTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationTasksInternalServerError) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *GetIntegrationTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.TypesErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
