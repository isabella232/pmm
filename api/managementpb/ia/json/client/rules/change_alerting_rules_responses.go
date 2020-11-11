// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeAlertingRulesReader is a Reader for the ChangeAlertingRules structure.
type ChangeAlertingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeAlertingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeAlertingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeAlertingRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeAlertingRulesOK creates a ChangeAlertingRulesOK with default headers values
func NewChangeAlertingRulesOK() *ChangeAlertingRulesOK {
	return &ChangeAlertingRulesOK{}
}

/*ChangeAlertingRulesOK handles this case with default header values.

A successful response.
*/
type ChangeAlertingRulesOK struct {
	Payload *ChangeAlertingRulesOKBody
}

func (o *ChangeAlertingRulesOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Rules/Change][%d] changeAlertingRulesOk  %+v", 200, o.Payload)
}

func (o *ChangeAlertingRulesOK) GetPayload() *ChangeAlertingRulesOKBody {
	return o.Payload
}

func (o *ChangeAlertingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeAlertingRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeAlertingRulesDefault creates a ChangeAlertingRulesDefault with default headers values
func NewChangeAlertingRulesDefault(code int) *ChangeAlertingRulesDefault {
	return &ChangeAlertingRulesDefault{
		_statusCode: code,
	}
}

/*ChangeAlertingRulesDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeAlertingRulesDefault struct {
	_statusCode int

	Payload *ChangeAlertingRulesDefaultBody
}

// Code gets the status code for the change alerting rules default response
func (o *ChangeAlertingRulesDefault) Code() int {
	return o._statusCode
}

func (o *ChangeAlertingRulesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ia/Rules/Change][%d] ChangeAlertingRules default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeAlertingRulesDefault) GetPayload() *ChangeAlertingRulesDefaultBody {
	return o.Payload
}

func (o *ChangeAlertingRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeAlertingRulesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeAlertingRulesBody change alerting rules body
swagger:model ChangeAlertingRulesBody
*/
type ChangeAlertingRulesBody struct {

	// Rule name.
	Name string `json:"name,omitempty"`

	// BooleanFlag represent a command to enable some boolean property (set to true),
	// disable some boolean property (set to false), or avoid changing that property.
	//
	//  - DO_NOT_CHANGE: Do not change boolean property. Default value.
	//  - ENABLE: Enable boolean property.
	//  - DISABLE: Disable boolean property.
	// Enum: [DO_NOT_CHANGE ENABLE DISABLE]
	Enabled1 *string `json:"enabled1,omitempty"`

	// New rule status.
	Enabled2 bool `json:"enabled2,omitempty"`

	// Parameters to change. Missing parameters will not be changed.
	Params []*ParamsItems0 `json:"params"`

	// Rule set duration. Zero value will not change it.
	// FIXME ?
	For string `json:"for,omitempty"`

	// Remove set duration (reset to default).
	RemoveFor bool `json:"remove_for,omitempty"`
}

// Validate validates this change alerting rules body
func (o *ChangeAlertingRulesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnabled1(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeAlertingRulesBodyTypeEnabled1PropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DO_NOT_CHANGE","ENABLE","DISABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeAlertingRulesBodyTypeEnabled1PropEnum = append(changeAlertingRulesBodyTypeEnabled1PropEnum, v)
	}
}

const (

	// ChangeAlertingRulesBodyEnabled1DONOTCHANGE captures enum value "DO_NOT_CHANGE"
	ChangeAlertingRulesBodyEnabled1DONOTCHANGE string = "DO_NOT_CHANGE"

	// ChangeAlertingRulesBodyEnabled1ENABLE captures enum value "ENABLE"
	ChangeAlertingRulesBodyEnabled1ENABLE string = "ENABLE"

	// ChangeAlertingRulesBodyEnabled1DISABLE captures enum value "DISABLE"
	ChangeAlertingRulesBodyEnabled1DISABLE string = "DISABLE"
)

// prop value enum
func (o *ChangeAlertingRulesBody) validateEnabled1Enum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeAlertingRulesBodyTypeEnabled1PropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeAlertingRulesBody) validateEnabled1(formats strfmt.Registry) error {

	if swag.IsZero(o.Enabled1) { // not required
		return nil
	}

	// value enum
	if err := o.validateEnabled1Enum("body"+"."+"enabled1", "body", *o.Enabled1); err != nil {
		return err
	}

	return nil
}

func (o *ChangeAlertingRulesBody) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	for i := 0; i < len(o.Params); i++ {
		if swag.IsZero(o.Params[i]) { // not required
			continue
		}

		if o.Params[i] != nil {
			if err := o.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAlertingRulesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAlertingRulesBody) UnmarshalBinary(b []byte) error {
	var res ChangeAlertingRulesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAlertingRulesDefaultBody change alerting rules default body
swagger:model ChangeAlertingRulesDefaultBody
*/
type ChangeAlertingRulesDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change alerting rules default body
func (o *ChangeAlertingRulesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAlertingRulesDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeAlertingRules default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAlertingRulesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAlertingRulesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeAlertingRulesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAlertingRulesOKBody change alerting rules OK body
swagger:model ChangeAlertingRulesOKBody
*/
type ChangeAlertingRulesOKBody struct {

	// rules
	Rules *ChangeAlertingRulesOKBodyRules `json:"rules,omitempty"`
}

// Validate validates this change alerting rules OK body
func (o *ChangeAlertingRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAlertingRulesOKBody) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(o.Rules) { // not required
		return nil
	}

	if o.Rules != nil {
		if err := o.Rules.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeAlertingRulesOk" + "." + "rules")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAlertingRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAlertingRulesOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeAlertingRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAlertingRulesOKBodyRules Rule represents Integrated Alerting rule.
swagger:model ChangeAlertingRulesOKBodyRules
*/
type ChangeAlertingRulesOKBodyRules struct {

	// Rule name.
	Name string `json:"name,omitempty"`

	// Rules status: enabled or disabled.
	Enabled bool `json:"enabled,omitempty"`

	// Rule description.
	Help string `json:"help,omitempty"`

	// Rule parameters.
	Params []*ChangeAlertingRulesOKBodyRulesParamsItems0 `json:"params"`

	// Rule default duration.
	DefaultFor string `json:"default_for,omitempty"`

	// Rule set duration.
	For string `json:"for,omitempty"`
}

// Validate validates this change alerting rules OK body rules
func (o *ChangeAlertingRulesOKBodyRules) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeAlertingRulesOKBodyRules) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	for i := 0; i < len(o.Params); i++ {
		if swag.IsZero(o.Params[i]) { // not required
			continue
		}

		if o.Params[i] != nil {
			if err := o.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changeAlertingRulesOk" + "." + "rules" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAlertingRulesOKBodyRules) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAlertingRulesOKBodyRules) UnmarshalBinary(b []byte) error {
	var res ChangeAlertingRulesOKBodyRules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeAlertingRulesOKBodyRulesParamsItems0 Param repsesents a single Integrated Alerting rule parameter.
swagger:model ChangeAlertingRulesOKBodyRulesParamsItems0
*/
type ChangeAlertingRulesOKBodyRulesParamsItems0 struct {

	// Parameter name.
	Name string `json:"name,omitempty"`

	// Parameter description.
	Help string `json:"help,omitempty"`

	// ParamUnit Integrated Alerting rule parameter unit.
	// Enum: [PARAM_UNIT_INVALID PERCENTAGE]
	Unit *string `json:"unit,omitempty"`

	// ParamType Integrated Alerting rule parameter type.
	// Enum: [PARAM_TYPE_INVALID FLOAT]
	Type *string `json:"type,omitempty"`

	// Parameter minimal value (float).
	MinValue float32 `json:"min_value,omitempty"`

	// Parameter maximum value (float).
	MaxValue float32 `json:"max_value,omitempty"`

	// Parameter set value (float).
	DefaultValue float32 `json:"default_value,omitempty"`

	// Parameter set value (float).
	Value float32 `json:"value,omitempty"`
}

// Validate validates this change alerting rules OK body rules params items0
func (o *ChangeAlertingRulesOKBodyRulesParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeAlertingRulesOkBodyRulesParamsItems0TypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PARAM_UNIT_INVALID","PERCENTAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeAlertingRulesOkBodyRulesParamsItems0TypeUnitPropEnum = append(changeAlertingRulesOkBodyRulesParamsItems0TypeUnitPropEnum, v)
	}
}

const (

	// ChangeAlertingRulesOKBodyRulesParamsItems0UnitPARAMUNITINVALID captures enum value "PARAM_UNIT_INVALID"
	ChangeAlertingRulesOKBodyRulesParamsItems0UnitPARAMUNITINVALID string = "PARAM_UNIT_INVALID"

	// ChangeAlertingRulesOKBodyRulesParamsItems0UnitPERCENTAGE captures enum value "PERCENTAGE"
	ChangeAlertingRulesOKBodyRulesParamsItems0UnitPERCENTAGE string = "PERCENTAGE"
)

// prop value enum
func (o *ChangeAlertingRulesOKBodyRulesParamsItems0) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeAlertingRulesOkBodyRulesParamsItems0TypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeAlertingRulesOKBodyRulesParamsItems0) validateUnit(formats strfmt.Registry) error {

	if swag.IsZero(o.Unit) { // not required
		return nil
	}

	// value enum
	if err := o.validateUnitEnum("unit", "body", *o.Unit); err != nil {
		return err
	}

	return nil
}

var changeAlertingRulesOkBodyRulesParamsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PARAM_TYPE_INVALID","FLOAT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeAlertingRulesOkBodyRulesParamsItems0TypeTypePropEnum = append(changeAlertingRulesOkBodyRulesParamsItems0TypeTypePropEnum, v)
	}
}

const (

	// ChangeAlertingRulesOKBodyRulesParamsItems0TypePARAMTYPEINVALID captures enum value "PARAM_TYPE_INVALID"
	ChangeAlertingRulesOKBodyRulesParamsItems0TypePARAMTYPEINVALID string = "PARAM_TYPE_INVALID"

	// ChangeAlertingRulesOKBodyRulesParamsItems0TypeFLOAT captures enum value "FLOAT"
	ChangeAlertingRulesOKBodyRulesParamsItems0TypeFLOAT string = "FLOAT"
)

// prop value enum
func (o *ChangeAlertingRulesOKBodyRulesParamsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, changeAlertingRulesOkBodyRulesParamsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChangeAlertingRulesOKBodyRulesParamsItems0) validateType(formats strfmt.Registry) error {

	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeAlertingRulesOKBodyRulesParamsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeAlertingRulesOKBodyRulesParamsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeAlertingRulesOKBodyRulesParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ParamsItems0 Param repsesents a single Integrated Alerting rule parameter change.
swagger:model ParamsItems0
*/
type ParamsItems0 struct {

	// Parameter name.
	Name string `json:"name,omitempty"`

	// Parameter set value (float).
	Value float32 `json:"value,omitempty"`

	// Remove set value (reset to default).
	RemoveValue bool `json:"remove_value,omitempty"`
}

// Validate validates this params items0
func (o *ParamsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ParamsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ParamsItems0) UnmarshalBinary(b []byte) error {
	var res ParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
