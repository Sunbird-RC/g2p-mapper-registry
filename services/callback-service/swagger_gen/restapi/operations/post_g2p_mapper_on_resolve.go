// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/sunbirdrc/callback-service/swagger_gen/models"
)

// PostG2pMapperOnResolveHandlerFunc turns a function with the right signature into a post g2p mapper on resolve handler
type PostG2pMapperOnResolveHandlerFunc func(PostG2pMapperOnResolveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostG2pMapperOnResolveHandlerFunc) Handle(params PostG2pMapperOnResolveParams) middleware.Responder {
	return fn(params)
}

// PostG2pMapperOnResolveHandler interface for that can handle valid post g2p mapper on resolve params
type PostG2pMapperOnResolveHandler interface {
	Handle(PostG2pMapperOnResolveParams) middleware.Responder
}

// NewPostG2pMapperOnResolve creates a new http.Handler for the post g2p mapper on resolve operation
func NewPostG2pMapperOnResolve(ctx *middleware.Context, handler PostG2pMapperOnResolveHandler) *PostG2pMapperOnResolve {
	return &PostG2pMapperOnResolve{Context: ctx, Handler: handler}
}

/*
	PostG2pMapperOnResolve swagger:route POST /mapper/on-resolve postG2pMapperOnResolve

FAMAP-ON-RSLV : /mapper/on-resolve

Resolve response through callback end point
*/
type PostG2pMapperOnResolve struct {
	Context *middleware.Context
	Handler PostG2pMapperOnResolveHandler
}

func (o *PostG2pMapperOnResolve) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostG2pMapperOnResolveParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostG2pMapperOnResolveBody post g2p mapper on resolve body
//
// swagger:model PostG2pMapperOnResolveBody
type PostG2pMapperOnResolveBody struct {

	// header
	// Required: true
	Header struct {
		models.MsgCallbackHeader
	} `json:"header"`

	// message
	Message *PostG2pMapperOnResolveParamsBodyMessage `json:"message,omitempty"`

	// signature
	Signature models.MsgSignature `json:"signature,omitempty"`
}

// Validate validates this post g2p mapper on resolve body
func (o *PostG2pMapperOnResolveBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveBody) validateHeader(formats strfmt.Registry) error {

	return nil
}

func (o *PostG2pMapperOnResolveBody) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(o.Message) { // not required
		return nil
	}

	if o.Message != nil {
		if err := o.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "message")
			}
			return err
		}
	}

	return nil
}

func (o *PostG2pMapperOnResolveBody) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if err := o.Signature.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "signature")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "signature")
		}
		return err
	}

	return nil
}

// ContextValidate validate this post g2p mapper on resolve body based on the context it is used
func (o *PostG2pMapperOnResolveBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveBody) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (o *PostG2pMapperOnResolveBody) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if o.Message != nil {
		if err := o.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "message")
			}
			return err
		}
	}

	return nil
}

func (o *PostG2pMapperOnResolveBody) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Signature.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "signature")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "signature")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveBody) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostG2pMapperOnResolveDefaultBody post g2p mapper on resolve default body
//
// swagger:model PostG2pMapperOnResolveDefaultBody
type PostG2pMapperOnResolveDefaultBody struct {

	// message
	Message *PostG2pMapperOnResolveDefaultBodyMessage `json:"message,omitempty"`
}

// Validate validates this post g2p mapper on resolve default body
func (o *PostG2pMapperOnResolveDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveDefaultBody) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(o.Message) { // not required
		return nil
	}

	if o.Message != nil {
		if err := o.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post g2p mapper on resolve default body based on the context it is used
func (o *PostG2pMapperOnResolveDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveDefaultBody) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if o.Message != nil {
		if err := o.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveDefaultBody) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostG2pMapperOnResolveDefaultBodyMessage post g2p mapper on resolve default body message
//
// swagger:model PostG2pMapperOnResolveDefaultBodyMessage
type PostG2pMapperOnResolveDefaultBodyMessage struct {

	// ack status
	AckStatus models.Ack `json:"ack_status,omitempty"`

	// error
	Error *models.Error `json:"error,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp models.Timestamp `json:"timestamp,omitempty"`
}

// Validate validates this post g2p mapper on resolve default body message
func (o *PostG2pMapperOnResolveDefaultBodyMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAckStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveDefaultBodyMessage) validateAckStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.AckStatus) { // not required
		return nil
	}

	if err := o.AckStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "ack_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "ack_status")
		}
		return err
	}

	return nil
}

func (o *PostG2pMapperOnResolveDefaultBodyMessage) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *PostG2pMapperOnResolveDefaultBodyMessage) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := o.Timestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "timestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "timestamp")
		}
		return err
	}

	return nil
}

// ContextValidate validate this post g2p mapper on resolve default body message based on the context it is used
func (o *PostG2pMapperOnResolveDefaultBodyMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAckStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveDefaultBodyMessage) contextValidateAckStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := o.AckStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "ack_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "ack_status")
		}
		return err
	}

	return nil
}

func (o *PostG2pMapperOnResolveDefaultBodyMessage) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *PostG2pMapperOnResolveDefaultBodyMessage) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Timestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "timestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("post_g2p_mapper_on-resolve default" + "." + "message" + "." + "timestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveDefaultBodyMessage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveDefaultBodyMessage) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveDefaultBodyMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostG2pMapperOnResolveForbiddenBody HTTP transport layer error codes. Used by components like gateways, LB responding with HTTP status codes 1xx, 2xx, 3xx, 4xx and 5xx
//
// swagger:model PostG2pMapperOnResolveForbiddenBody
type PostG2pMapperOnResolveForbiddenBody struct {

	// error
	Error *PostG2pMapperOnResolveForbiddenBodyError `json:"error,omitempty"`
}

// Validate validates this post g2p mapper on resolve forbidden body
func (o *PostG2pMapperOnResolveForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveForbiddenBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postG2pMapperOnResolveForbidden" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postG2pMapperOnResolveForbidden" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post g2p mapper on resolve forbidden body based on the context it is used
func (o *PostG2pMapperOnResolveForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveForbiddenBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postG2pMapperOnResolveForbidden" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postG2pMapperOnResolveForbidden" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostG2pMapperOnResolveForbiddenBodyError post g2p mapper on resolve forbidden body error
//
// swagger:model PostG2pMapperOnResolveForbiddenBodyError
type PostG2pMapperOnResolveForbiddenBodyError struct {

	// error code
	Code string `json:"code,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

// Validate validates this post g2p mapper on resolve forbidden body error
func (o *PostG2pMapperOnResolveForbiddenBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post g2p mapper on resolve forbidden body error based on context it is used
func (o *PostG2pMapperOnResolveForbiddenBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveForbiddenBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveForbiddenBodyError) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveForbiddenBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostG2pMapperOnResolveInternalServerErrorBody HTTP transport layer error codes. Used by components like gateways, LB responding with HTTP status codes 1xx, 2xx, 3xx, 4xx and 5xx
//
// swagger:model PostG2pMapperOnResolveInternalServerErrorBody
type PostG2pMapperOnResolveInternalServerErrorBody struct {

	// error
	Error *PostG2pMapperOnResolveInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this post g2p mapper on resolve internal server error body
func (o *PostG2pMapperOnResolveInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postG2pMapperOnResolveInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postG2pMapperOnResolveInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post g2p mapper on resolve internal server error body based on the context it is used
func (o *PostG2pMapperOnResolveInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postG2pMapperOnResolveInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postG2pMapperOnResolveInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostG2pMapperOnResolveInternalServerErrorBodyError post g2p mapper on resolve internal server error body error
//
// swagger:model PostG2pMapperOnResolveInternalServerErrorBodyError
type PostG2pMapperOnResolveInternalServerErrorBodyError struct {

	// error code
	Code string `json:"code,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

// Validate validates this post g2p mapper on resolve internal server error body error
func (o *PostG2pMapperOnResolveInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post g2p mapper on resolve internal server error body error based on context it is used
func (o *PostG2pMapperOnResolveInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostG2pMapperOnResolveParamsBodyMessage post g2p mapper on resolve params body message
//
// swagger:model PostG2pMapperOnResolveParamsBodyMessage
type PostG2pMapperOnResolveParamsBodyMessage struct {

	// resolve response
	// Required: true
	ResolveResponse []*models.ResolveResponse `json:"resolve_response"`

	// transaction id
	// Required: true
	TransactionID *models.TransactionID `json:"transaction_id"`
}

// Validate validates this post g2p mapper on resolve params body message
func (o *PostG2pMapperOnResolveParamsBodyMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResolveResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveParamsBodyMessage) validateResolveResponse(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"message"+"."+"resolve_response", "body", o.ResolveResponse); err != nil {
		return err
	}

	for i := 0; i < len(o.ResolveResponse); i++ {
		if swag.IsZero(o.ResolveResponse[i]) { // not required
			continue
		}

		if o.ResolveResponse[i] != nil {
			if err := o.ResolveResponse[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "message" + "." + "resolve_response" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "message" + "." + "resolve_response" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostG2pMapperOnResolveParamsBodyMessage) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"message"+"."+"transaction_id", "body", o.TransactionID); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"message"+"."+"transaction_id", "body", o.TransactionID); err != nil {
		return err
	}

	if o.TransactionID != nil {
		if err := o.TransactionID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "message" + "." + "transaction_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "message" + "." + "transaction_id")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post g2p mapper on resolve params body message based on the context it is used
func (o *PostG2pMapperOnResolveParamsBodyMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResolveResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTransactionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveParamsBodyMessage) contextValidateResolveResponse(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ResolveResponse); i++ {

		if o.ResolveResponse[i] != nil {
			if err := o.ResolveResponse[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "message" + "." + "resolve_response" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "message" + "." + "resolve_response" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostG2pMapperOnResolveParamsBodyMessage) contextValidateTransactionID(ctx context.Context, formats strfmt.Registry) error {

	if o.TransactionID != nil {
		if err := o.TransactionID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "message" + "." + "transaction_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "message" + "." + "transaction_id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveParamsBodyMessage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveParamsBodyMessage) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveParamsBodyMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostG2pMapperOnResolveUnauthorizedBody HTTP transport layer error codes. Used by components like gateways, LB responding with HTTP status codes 1xx, 2xx, 3xx, 4xx and 5xx
//
// swagger:model PostG2pMapperOnResolveUnauthorizedBody
type PostG2pMapperOnResolveUnauthorizedBody struct {

	// error
	Error *PostG2pMapperOnResolveUnauthorizedBodyError `json:"error,omitempty"`
}

// Validate validates this post g2p mapper on resolve unauthorized body
func (o *PostG2pMapperOnResolveUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveUnauthorizedBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postG2pMapperOnResolveUnauthorized" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postG2pMapperOnResolveUnauthorized" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post g2p mapper on resolve unauthorized body based on the context it is used
func (o *PostG2pMapperOnResolveUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostG2pMapperOnResolveUnauthorizedBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postG2pMapperOnResolveUnauthorized" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postG2pMapperOnResolveUnauthorized" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostG2pMapperOnResolveUnauthorizedBodyError post g2p mapper on resolve unauthorized body error
//
// swagger:model PostG2pMapperOnResolveUnauthorizedBodyError
type PostG2pMapperOnResolveUnauthorizedBodyError struct {

	// error code
	Code string `json:"code,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

// Validate validates this post g2p mapper on resolve unauthorized body error
func (o *PostG2pMapperOnResolveUnauthorizedBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post g2p mapper on resolve unauthorized body error based on context it is used
func (o *PostG2pMapperOnResolveUnauthorizedBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostG2pMapperOnResolveUnauthorizedBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostG2pMapperOnResolveUnauthorizedBodyError) UnmarshalBinary(b []byte) error {
	var res PostG2pMapperOnResolveUnauthorizedBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
