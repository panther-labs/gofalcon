// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// ConnectCSPMGCPAccountReader is a Reader for the ConnectCSPMGCPAccount structure.
type ConnectCSPMGCPAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCSPMGCPAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewConnectCSPMGCPAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewConnectCSPMGCPAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConnectCSPMGCPAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewConnectCSPMGCPAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewConnectCSPMGCPAccountConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewConnectCSPMGCPAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConnectCSPMGCPAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloud-connect-cspm-gcp/entities/account/v2] ConnectCSPMGCPAccount", response, response.Code())
	}
}

// NewConnectCSPMGCPAccountCreated creates a ConnectCSPMGCPAccountCreated with default headers values
func NewConnectCSPMGCPAccountCreated() *ConnectCSPMGCPAccountCreated {
	return &ConnectCSPMGCPAccountCreated{}
}

/*
ConnectCSPMGCPAccountCreated describes a response with status code 201, with default header values.

Created
*/
type ConnectCSPMGCPAccountCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseExtV2
}

// IsSuccess returns true when this connect c s p m g c p account created response has a 2xx status code
func (o *ConnectCSPMGCPAccountCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this connect c s p m g c p account created response has a 3xx status code
func (o *ConnectCSPMGCPAccountCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this connect c s p m g c p account created response has a 4xx status code
func (o *ConnectCSPMGCPAccountCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this connect c s p m g c p account created response has a 5xx status code
func (o *ConnectCSPMGCPAccountCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this connect c s p m g c p account created response a status code equal to that given
func (o *ConnectCSPMGCPAccountCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the connect c s p m g c p account created response
func (o *ConnectCSPMGCPAccountCreated) Code() int {
	return 201
}

func (o *ConnectCSPMGCPAccountCreated) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountCreated  %+v", 201, o.Payload)
}

func (o *ConnectCSPMGCPAccountCreated) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountCreated  %+v", 201, o.Payload)
}

func (o *ConnectCSPMGCPAccountCreated) GetPayload() *models.RegistrationGCPAccountResponseExtV2 {
	return o.Payload
}

func (o *ConnectCSPMGCPAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseExtV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCSPMGCPAccountMultiStatus creates a ConnectCSPMGCPAccountMultiStatus with default headers values
func NewConnectCSPMGCPAccountMultiStatus() *ConnectCSPMGCPAccountMultiStatus {
	return &ConnectCSPMGCPAccountMultiStatus{}
}

/*
ConnectCSPMGCPAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type ConnectCSPMGCPAccountMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseExtV2
}

// IsSuccess returns true when this connect c s p m g c p account multi status response has a 2xx status code
func (o *ConnectCSPMGCPAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this connect c s p m g c p account multi status response has a 3xx status code
func (o *ConnectCSPMGCPAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this connect c s p m g c p account multi status response has a 4xx status code
func (o *ConnectCSPMGCPAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this connect c s p m g c p account multi status response has a 5xx status code
func (o *ConnectCSPMGCPAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this connect c s p m g c p account multi status response a status code equal to that given
func (o *ConnectCSPMGCPAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the connect c s p m g c p account multi status response
func (o *ConnectCSPMGCPAccountMultiStatus) Code() int {
	return 207
}

func (o *ConnectCSPMGCPAccountMultiStatus) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *ConnectCSPMGCPAccountMultiStatus) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *ConnectCSPMGCPAccountMultiStatus) GetPayload() *models.RegistrationGCPAccountResponseExtV2 {
	return o.Payload
}

func (o *ConnectCSPMGCPAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseExtV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCSPMGCPAccountBadRequest creates a ConnectCSPMGCPAccountBadRequest with default headers values
func NewConnectCSPMGCPAccountBadRequest() *ConnectCSPMGCPAccountBadRequest {
	return &ConnectCSPMGCPAccountBadRequest{}
}

/*
ConnectCSPMGCPAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ConnectCSPMGCPAccountBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseExtV2
}

// IsSuccess returns true when this connect c s p m g c p account bad request response has a 2xx status code
func (o *ConnectCSPMGCPAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this connect c s p m g c p account bad request response has a 3xx status code
func (o *ConnectCSPMGCPAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this connect c s p m g c p account bad request response has a 4xx status code
func (o *ConnectCSPMGCPAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this connect c s p m g c p account bad request response has a 5xx status code
func (o *ConnectCSPMGCPAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this connect c s p m g c p account bad request response a status code equal to that given
func (o *ConnectCSPMGCPAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the connect c s p m g c p account bad request response
func (o *ConnectCSPMGCPAccountBadRequest) Code() int {
	return 400
}

func (o *ConnectCSPMGCPAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountBadRequest  %+v", 400, o.Payload)
}

func (o *ConnectCSPMGCPAccountBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountBadRequest  %+v", 400, o.Payload)
}

func (o *ConnectCSPMGCPAccountBadRequest) GetPayload() *models.RegistrationGCPAccountResponseExtV2 {
	return o.Payload
}

func (o *ConnectCSPMGCPAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseExtV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCSPMGCPAccountForbidden creates a ConnectCSPMGCPAccountForbidden with default headers values
func NewConnectCSPMGCPAccountForbidden() *ConnectCSPMGCPAccountForbidden {
	return &ConnectCSPMGCPAccountForbidden{}
}

/*
ConnectCSPMGCPAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ConnectCSPMGCPAccountForbidden struct {

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

// IsSuccess returns true when this connect c s p m g c p account forbidden response has a 2xx status code
func (o *ConnectCSPMGCPAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this connect c s p m g c p account forbidden response has a 3xx status code
func (o *ConnectCSPMGCPAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this connect c s p m g c p account forbidden response has a 4xx status code
func (o *ConnectCSPMGCPAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this connect c s p m g c p account forbidden response has a 5xx status code
func (o *ConnectCSPMGCPAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this connect c s p m g c p account forbidden response a status code equal to that given
func (o *ConnectCSPMGCPAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the connect c s p m g c p account forbidden response
func (o *ConnectCSPMGCPAccountForbidden) Code() int {
	return 403
}

func (o *ConnectCSPMGCPAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountForbidden  %+v", 403, o.Payload)
}

func (o *ConnectCSPMGCPAccountForbidden) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountForbidden  %+v", 403, o.Payload)
}

func (o *ConnectCSPMGCPAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ConnectCSPMGCPAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewConnectCSPMGCPAccountConflict creates a ConnectCSPMGCPAccountConflict with default headers values
func NewConnectCSPMGCPAccountConflict() *ConnectCSPMGCPAccountConflict {
	return &ConnectCSPMGCPAccountConflict{}
}

/*
ConnectCSPMGCPAccountConflict describes a response with status code 409, with default header values.

Conflict
*/
type ConnectCSPMGCPAccountConflict struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseExtV2
}

// IsSuccess returns true when this connect c s p m g c p account conflict response has a 2xx status code
func (o *ConnectCSPMGCPAccountConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this connect c s p m g c p account conflict response has a 3xx status code
func (o *ConnectCSPMGCPAccountConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this connect c s p m g c p account conflict response has a 4xx status code
func (o *ConnectCSPMGCPAccountConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this connect c s p m g c p account conflict response has a 5xx status code
func (o *ConnectCSPMGCPAccountConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this connect c s p m g c p account conflict response a status code equal to that given
func (o *ConnectCSPMGCPAccountConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the connect c s p m g c p account conflict response
func (o *ConnectCSPMGCPAccountConflict) Code() int {
	return 409
}

func (o *ConnectCSPMGCPAccountConflict) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountConflict  %+v", 409, o.Payload)
}

func (o *ConnectCSPMGCPAccountConflict) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountConflict  %+v", 409, o.Payload)
}

func (o *ConnectCSPMGCPAccountConflict) GetPayload() *models.RegistrationGCPAccountResponseExtV2 {
	return o.Payload
}

func (o *ConnectCSPMGCPAccountConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseExtV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCSPMGCPAccountTooManyRequests creates a ConnectCSPMGCPAccountTooManyRequests with default headers values
func NewConnectCSPMGCPAccountTooManyRequests() *ConnectCSPMGCPAccountTooManyRequests {
	return &ConnectCSPMGCPAccountTooManyRequests{}
}

/*
ConnectCSPMGCPAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ConnectCSPMGCPAccountTooManyRequests struct {

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

// IsSuccess returns true when this connect c s p m g c p account too many requests response has a 2xx status code
func (o *ConnectCSPMGCPAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this connect c s p m g c p account too many requests response has a 3xx status code
func (o *ConnectCSPMGCPAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this connect c s p m g c p account too many requests response has a 4xx status code
func (o *ConnectCSPMGCPAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this connect c s p m g c p account too many requests response has a 5xx status code
func (o *ConnectCSPMGCPAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this connect c s p m g c p account too many requests response a status code equal to that given
func (o *ConnectCSPMGCPAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the connect c s p m g c p account too many requests response
func (o *ConnectCSPMGCPAccountTooManyRequests) Code() int {
	return 429
}

func (o *ConnectCSPMGCPAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ConnectCSPMGCPAccountTooManyRequests) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ConnectCSPMGCPAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ConnectCSPMGCPAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewConnectCSPMGCPAccountInternalServerError creates a ConnectCSPMGCPAccountInternalServerError with default headers values
func NewConnectCSPMGCPAccountInternalServerError() *ConnectCSPMGCPAccountInternalServerError {
	return &ConnectCSPMGCPAccountInternalServerError{}
}

/*
ConnectCSPMGCPAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ConnectCSPMGCPAccountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseExtV2
}

// IsSuccess returns true when this connect c s p m g c p account internal server error response has a 2xx status code
func (o *ConnectCSPMGCPAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this connect c s p m g c p account internal server error response has a 3xx status code
func (o *ConnectCSPMGCPAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this connect c s p m g c p account internal server error response has a 4xx status code
func (o *ConnectCSPMGCPAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this connect c s p m g c p account internal server error response has a 5xx status code
func (o *ConnectCSPMGCPAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this connect c s p m g c p account internal server error response a status code equal to that given
func (o *ConnectCSPMGCPAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the connect c s p m g c p account internal server error response
func (o *ConnectCSPMGCPAccountInternalServerError) Code() int {
	return 500
}

func (o *ConnectCSPMGCPAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *ConnectCSPMGCPAccountInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-gcp/entities/account/v2][%d] connectCSPMGCPAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *ConnectCSPMGCPAccountInternalServerError) GetPayload() *models.RegistrationGCPAccountResponseExtV2 {
	return o.Payload
}

func (o *ConnectCSPMGCPAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseExtV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
