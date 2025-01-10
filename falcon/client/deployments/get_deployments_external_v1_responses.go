// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// GetDeploymentsExternalV1Reader is a Reader for the GetDeploymentsExternalV1 structure.
type GetDeploymentsExternalV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentsExternalV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentsExternalV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentsExternalV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeploymentsExternalV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentsExternalV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDeploymentsExternalV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeploymentsExternalV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /deployment-coordinator/entities/deployments/external/v1] GetDeploymentsExternalV1", response, response.Code())
	}
}

// NewGetDeploymentsExternalV1OK creates a GetDeploymentsExternalV1OK with default headers values
func NewGetDeploymentsExternalV1OK() *GetDeploymentsExternalV1OK {
	return &GetDeploymentsExternalV1OK{}
}

/*
GetDeploymentsExternalV1OK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentsExternalV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DeploymentsAPIDeploymentViewWrapper
}

// IsSuccess returns true when this get deployments external v1 o k response has a 2xx status code
func (o *GetDeploymentsExternalV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployments external v1 o k response has a 3xx status code
func (o *GetDeploymentsExternalV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments external v1 o k response has a 4xx status code
func (o *GetDeploymentsExternalV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployments external v1 o k response has a 5xx status code
func (o *GetDeploymentsExternalV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments external v1 o k response a status code equal to that given
func (o *GetDeploymentsExternalV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deployments external v1 o k response
func (o *GetDeploymentsExternalV1OK) Code() int {
	return 200
}

func (o *GetDeploymentsExternalV1OK) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentsExternalV1OK) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentsExternalV1OK) GetPayload() *models.DeploymentsAPIDeploymentViewWrapper {
	return o.Payload
}

func (o *GetDeploymentsExternalV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DeploymentsAPIDeploymentViewWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsExternalV1BadRequest creates a GetDeploymentsExternalV1BadRequest with default headers values
func NewGetDeploymentsExternalV1BadRequest() *GetDeploymentsExternalV1BadRequest {
	return &GetDeploymentsExternalV1BadRequest{}
}

/*
GetDeploymentsExternalV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDeploymentsExternalV1BadRequest struct {

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

// IsSuccess returns true when this get deployments external v1 bad request response has a 2xx status code
func (o *GetDeploymentsExternalV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments external v1 bad request response has a 3xx status code
func (o *GetDeploymentsExternalV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments external v1 bad request response has a 4xx status code
func (o *GetDeploymentsExternalV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments external v1 bad request response has a 5xx status code
func (o *GetDeploymentsExternalV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments external v1 bad request response a status code equal to that given
func (o *GetDeploymentsExternalV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get deployments external v1 bad request response
func (o *GetDeploymentsExternalV1BadRequest) Code() int {
	return 400
}

func (o *GetDeploymentsExternalV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetDeploymentsExternalV1BadRequest) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetDeploymentsExternalV1BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetDeploymentsExternalV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentsExternalV1Forbidden creates a GetDeploymentsExternalV1Forbidden with default headers values
func NewGetDeploymentsExternalV1Forbidden() *GetDeploymentsExternalV1Forbidden {
	return &GetDeploymentsExternalV1Forbidden{}
}

/*
GetDeploymentsExternalV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDeploymentsExternalV1Forbidden struct {

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

// IsSuccess returns true when this get deployments external v1 forbidden response has a 2xx status code
func (o *GetDeploymentsExternalV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments external v1 forbidden response has a 3xx status code
func (o *GetDeploymentsExternalV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments external v1 forbidden response has a 4xx status code
func (o *GetDeploymentsExternalV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments external v1 forbidden response has a 5xx status code
func (o *GetDeploymentsExternalV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments external v1 forbidden response a status code equal to that given
func (o *GetDeploymentsExternalV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get deployments external v1 forbidden response
func (o *GetDeploymentsExternalV1Forbidden) Code() int {
	return 403
}

func (o *GetDeploymentsExternalV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetDeploymentsExternalV1Forbidden) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetDeploymentsExternalV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetDeploymentsExternalV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentsExternalV1NotFound creates a GetDeploymentsExternalV1NotFound with default headers values
func NewGetDeploymentsExternalV1NotFound() *GetDeploymentsExternalV1NotFound {
	return &GetDeploymentsExternalV1NotFound{}
}

/*
GetDeploymentsExternalV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeploymentsExternalV1NotFound struct {

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

// IsSuccess returns true when this get deployments external v1 not found response has a 2xx status code
func (o *GetDeploymentsExternalV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments external v1 not found response has a 3xx status code
func (o *GetDeploymentsExternalV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments external v1 not found response has a 4xx status code
func (o *GetDeploymentsExternalV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments external v1 not found response has a 5xx status code
func (o *GetDeploymentsExternalV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments external v1 not found response a status code equal to that given
func (o *GetDeploymentsExternalV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get deployments external v1 not found response
func (o *GetDeploymentsExternalV1NotFound) Code() int {
	return 404
}

func (o *GetDeploymentsExternalV1NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1NotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentsExternalV1NotFound) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1NotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentsExternalV1NotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetDeploymentsExternalV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentsExternalV1TooManyRequests creates a GetDeploymentsExternalV1TooManyRequests with default headers values
func NewGetDeploymentsExternalV1TooManyRequests() *GetDeploymentsExternalV1TooManyRequests {
	return &GetDeploymentsExternalV1TooManyRequests{}
}

/*
GetDeploymentsExternalV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeploymentsExternalV1TooManyRequests struct {

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

// IsSuccess returns true when this get deployments external v1 too many requests response has a 2xx status code
func (o *GetDeploymentsExternalV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments external v1 too many requests response has a 3xx status code
func (o *GetDeploymentsExternalV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments external v1 too many requests response has a 4xx status code
func (o *GetDeploymentsExternalV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments external v1 too many requests response has a 5xx status code
func (o *GetDeploymentsExternalV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments external v1 too many requests response a status code equal to that given
func (o *GetDeploymentsExternalV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get deployments external v1 too many requests response
func (o *GetDeploymentsExternalV1TooManyRequests) Code() int {
	return 429
}

func (o *GetDeploymentsExternalV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeploymentsExternalV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeploymentsExternalV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetDeploymentsExternalV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentsExternalV1InternalServerError creates a GetDeploymentsExternalV1InternalServerError with default headers values
func NewGetDeploymentsExternalV1InternalServerError() *GetDeploymentsExternalV1InternalServerError {
	return &GetDeploymentsExternalV1InternalServerError{}
}

/*
GetDeploymentsExternalV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDeploymentsExternalV1InternalServerError struct {

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

// IsSuccess returns true when this get deployments external v1 internal server error response has a 2xx status code
func (o *GetDeploymentsExternalV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments external v1 internal server error response has a 3xx status code
func (o *GetDeploymentsExternalV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments external v1 internal server error response has a 4xx status code
func (o *GetDeploymentsExternalV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployments external v1 internal server error response has a 5xx status code
func (o *GetDeploymentsExternalV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get deployments external v1 internal server error response a status code equal to that given
func (o *GetDeploymentsExternalV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get deployments external v1 internal server error response
func (o *GetDeploymentsExternalV1InternalServerError) Code() int {
	return 500
}

func (o *GetDeploymentsExternalV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentsExternalV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/entities/deployments/external/v1][%d] getDeploymentsExternalV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentsExternalV1InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetDeploymentsExternalV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
