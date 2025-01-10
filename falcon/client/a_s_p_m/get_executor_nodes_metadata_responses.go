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

// GetExecutorNodesMetadataReader is a Reader for the GetExecutorNodesMetadata structure.
type GetExecutorNodesMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExecutorNodesMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExecutorNodesMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExecutorNodesMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExecutorNodesMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExecutorNodesMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExecutorNodesMetadataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExecutorNodesMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata] GetExecutorNodesMetadata", response, response.Code())
	}
}

// NewGetExecutorNodesMetadataOK creates a GetExecutorNodesMetadataOK with default headers values
func NewGetExecutorNodesMetadataOK() *GetExecutorNodesMetadataOK {
	return &GetExecutorNodesMetadataOK{}
}

/*
GetExecutorNodesMetadataOK describes a response with status code 200, with default header values.

OK
*/
type GetExecutorNodesMetadataOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.TypesGetExecutorNodesMetadataResponse
}

// IsSuccess returns true when this get executor nodes metadata o k response has a 2xx status code
func (o *GetExecutorNodesMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get executor nodes metadata o k response has a 3xx status code
func (o *GetExecutorNodesMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executor nodes metadata o k response has a 4xx status code
func (o *GetExecutorNodesMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get executor nodes metadata o k response has a 5xx status code
func (o *GetExecutorNodesMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get executor nodes metadata o k response a status code equal to that given
func (o *GetExecutorNodesMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get executor nodes metadata o k response
func (o *GetExecutorNodesMetadataOK) Code() int {
	return 200
}

func (o *GetExecutorNodesMetadataOK) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataOK  %+v", 200, o.Payload)
}

func (o *GetExecutorNodesMetadataOK) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataOK  %+v", 200, o.Payload)
}

func (o *GetExecutorNodesMetadataOK) GetPayload() *models.TypesGetExecutorNodesMetadataResponse {
	return o.Payload
}

func (o *GetExecutorNodesMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.TypesGetExecutorNodesMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutorNodesMetadataBadRequest creates a GetExecutorNodesMetadataBadRequest with default headers values
func NewGetExecutorNodesMetadataBadRequest() *GetExecutorNodesMetadataBadRequest {
	return &GetExecutorNodesMetadataBadRequest{}
}

/*
GetExecutorNodesMetadataBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetExecutorNodesMetadataBadRequest struct {

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

// IsSuccess returns true when this get executor nodes metadata bad request response has a 2xx status code
func (o *GetExecutorNodesMetadataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get executor nodes metadata bad request response has a 3xx status code
func (o *GetExecutorNodesMetadataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executor nodes metadata bad request response has a 4xx status code
func (o *GetExecutorNodesMetadataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get executor nodes metadata bad request response has a 5xx status code
func (o *GetExecutorNodesMetadataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get executor nodes metadata bad request response a status code equal to that given
func (o *GetExecutorNodesMetadataBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get executor nodes metadata bad request response
func (o *GetExecutorNodesMetadataBadRequest) Code() int {
	return 400
}

func (o *GetExecutorNodesMetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetExecutorNodesMetadataBadRequest) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetExecutorNodesMetadataBadRequest) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *GetExecutorNodesMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetExecutorNodesMetadataUnauthorized creates a GetExecutorNodesMetadataUnauthorized with default headers values
func NewGetExecutorNodesMetadataUnauthorized() *GetExecutorNodesMetadataUnauthorized {
	return &GetExecutorNodesMetadataUnauthorized{}
}

/*
GetExecutorNodesMetadataUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetExecutorNodesMetadataUnauthorized struct {

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

// IsSuccess returns true when this get executor nodes metadata unauthorized response has a 2xx status code
func (o *GetExecutorNodesMetadataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get executor nodes metadata unauthorized response has a 3xx status code
func (o *GetExecutorNodesMetadataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executor nodes metadata unauthorized response has a 4xx status code
func (o *GetExecutorNodesMetadataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get executor nodes metadata unauthorized response has a 5xx status code
func (o *GetExecutorNodesMetadataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get executor nodes metadata unauthorized response a status code equal to that given
func (o *GetExecutorNodesMetadataUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get executor nodes metadata unauthorized response
func (o *GetExecutorNodesMetadataUnauthorized) Code() int {
	return 401
}

func (o *GetExecutorNodesMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExecutorNodesMetadataUnauthorized) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExecutorNodesMetadataUnauthorized) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *GetExecutorNodesMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetExecutorNodesMetadataForbidden creates a GetExecutorNodesMetadataForbidden with default headers values
func NewGetExecutorNodesMetadataForbidden() *GetExecutorNodesMetadataForbidden {
	return &GetExecutorNodesMetadataForbidden{}
}

/*
GetExecutorNodesMetadataForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetExecutorNodesMetadataForbidden struct {

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

// IsSuccess returns true when this get executor nodes metadata forbidden response has a 2xx status code
func (o *GetExecutorNodesMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get executor nodes metadata forbidden response has a 3xx status code
func (o *GetExecutorNodesMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executor nodes metadata forbidden response has a 4xx status code
func (o *GetExecutorNodesMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get executor nodes metadata forbidden response has a 5xx status code
func (o *GetExecutorNodesMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get executor nodes metadata forbidden response a status code equal to that given
func (o *GetExecutorNodesMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get executor nodes metadata forbidden response
func (o *GetExecutorNodesMetadataForbidden) Code() int {
	return 403
}

func (o *GetExecutorNodesMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetExecutorNodesMetadataForbidden) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetExecutorNodesMetadataForbidden) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *GetExecutorNodesMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetExecutorNodesMetadataTooManyRequests creates a GetExecutorNodesMetadataTooManyRequests with default headers values
func NewGetExecutorNodesMetadataTooManyRequests() *GetExecutorNodesMetadataTooManyRequests {
	return &GetExecutorNodesMetadataTooManyRequests{}
}

/*
GetExecutorNodesMetadataTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetExecutorNodesMetadataTooManyRequests struct {

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

// IsSuccess returns true when this get executor nodes metadata too many requests response has a 2xx status code
func (o *GetExecutorNodesMetadataTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get executor nodes metadata too many requests response has a 3xx status code
func (o *GetExecutorNodesMetadataTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executor nodes metadata too many requests response has a 4xx status code
func (o *GetExecutorNodesMetadataTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get executor nodes metadata too many requests response has a 5xx status code
func (o *GetExecutorNodesMetadataTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get executor nodes metadata too many requests response a status code equal to that given
func (o *GetExecutorNodesMetadataTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get executor nodes metadata too many requests response
func (o *GetExecutorNodesMetadataTooManyRequests) Code() int {
	return 429
}

func (o *GetExecutorNodesMetadataTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExecutorNodesMetadataTooManyRequests) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExecutorNodesMetadataTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetExecutorNodesMetadataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetExecutorNodesMetadataInternalServerError creates a GetExecutorNodesMetadataInternalServerError with default headers values
func NewGetExecutorNodesMetadataInternalServerError() *GetExecutorNodesMetadataInternalServerError {
	return &GetExecutorNodesMetadataInternalServerError{}
}

/*
GetExecutorNodesMetadataInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetExecutorNodesMetadataInternalServerError struct {

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

// IsSuccess returns true when this get executor nodes metadata internal server error response has a 2xx status code
func (o *GetExecutorNodesMetadataInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get executor nodes metadata internal server error response has a 3xx status code
func (o *GetExecutorNodesMetadataInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executor nodes metadata internal server error response has a 4xx status code
func (o *GetExecutorNodesMetadataInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get executor nodes metadata internal server error response has a 5xx status code
func (o *GetExecutorNodesMetadataInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get executor nodes metadata internal server error response a status code equal to that given
func (o *GetExecutorNodesMetadataInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get executor nodes metadata internal server error response
func (o *GetExecutorNodesMetadataInternalServerError) Code() int {
	return 500
}

func (o *GetExecutorNodesMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExecutorNodesMetadataInternalServerError) String() string {
	return fmt.Sprintf("[GET /aspm-api-gateway/api/v1/executor_nodes/metadata][%d] getExecutorNodesMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExecutorNodesMetadataInternalServerError) GetPayload() *models.TypesErrorMessage {
	return o.Payload
}

func (o *GetExecutorNodesMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
