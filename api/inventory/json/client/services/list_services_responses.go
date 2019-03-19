// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// ListServicesReader is a Reader for the ListServices structure.
type ListServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListServicesOK creates a ListServicesOK with default headers values
func NewListServicesOK() *ListServicesOK {
	return &ListServicesOK{}
}

/*ListServicesOK handles this case with default header values.

A successful response.
*/
type ListServicesOK struct {
	Payload *ListServicesOKBody
}

func (o *ListServicesOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/List][%d] listServicesOk  %+v", 200, o.Payload)
}

func (o *ListServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesDefault creates a ListServicesDefault with default headers values
func NewListServicesDefault(code int) *ListServicesDefault {
	return &ListServicesDefault{
		_statusCode: code,
	}
}

/*ListServicesDefault handles this case with default header values.

An error response.
*/
type ListServicesDefault struct {
	_statusCode int

	Payload *ListServicesDefaultBody
}

// Code gets the status code for the list services default response
func (o *ListServicesDefault) Code() int {
	return o._statusCode
}

func (o *ListServicesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/List][%d] ListServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListServicesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AmazonRDSMysqlItems0 AmazonRDSMySQLService represents a MySQL instance running on a single RemoteAmazonRDS Node
swagger:model AmazonRDSMysqlItems0
*/
type AmazonRDSMysqlItems0 struct {

	// Instance endpoint (full DNS name).
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Instance port.
	Port int64 `json:"port,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this amazon RDS mysql items0
func (o *AmazonRDSMysqlItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AmazonRDSMysqlItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AmazonRDSMysqlItems0) UnmarshalBinary(b []byte) error {
	var res AmazonRDSMysqlItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesBody list services body
swagger:model ListServicesBody
*/
type ListServicesBody struct {

	// Return only Services running on that Node.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this list services body
func (o *ListServicesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesBody) UnmarshalBinary(b []byte) error {
	var res ListServicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ListServicesDefaultBody
*/
type ListServicesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this list services default body
func (o *ListServicesDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListServicesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListServicesOKBody list services OK body
swagger:model ListServicesOKBody
*/
type ListServicesOKBody struct {

	// amazon rds mysql
	AmazonRDSMysql []*AmazonRDSMysqlItems0 `json:"amazon_rds_mysql"`

	// mongodb
	Mongodb []*MongodbItems0 `json:"mongodb"`

	// mysql
	Mysql []*MysqlItems0 `json:"mysql"`

	// postgresql
	Postgresql []*PostgresqlItems0 `json:"postgresql"`
}

// Validate validates this list services OK body
func (o *ListServicesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmazonRDSMysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongodb(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePostgresql(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListServicesOKBody) validateAmazonRDSMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.AmazonRDSMysql) { // not required
		return nil
	}

	for i := 0; i < len(o.AmazonRDSMysql); i++ {
		if swag.IsZero(o.AmazonRDSMysql[i]) { // not required
			continue
		}

		if o.AmazonRDSMysql[i] != nil {
			if err := o.AmazonRDSMysql[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "amazon_rds_mysql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListServicesOKBody) validateMongodb(formats strfmt.Registry) error {

	if swag.IsZero(o.Mongodb) { // not required
		return nil
	}

	for i := 0; i < len(o.Mongodb); i++ {
		if swag.IsZero(o.Mongodb[i]) { // not required
			continue
		}

		if o.Mongodb[i] != nil {
			if err := o.Mongodb[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "mongodb" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListServicesOKBody) validateMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Mysql) { // not required
		return nil
	}

	for i := 0; i < len(o.Mysql); i++ {
		if swag.IsZero(o.Mysql[i]) { // not required
			continue
		}

		if o.Mysql[i] != nil {
			if err := o.Mysql[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "mysql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListServicesOKBody) validatePostgresql(formats strfmt.Registry) error {

	if swag.IsZero(o.Postgresql) { // not required
		return nil
	}

	for i := 0; i < len(o.Postgresql); i++ {
		if swag.IsZero(o.Postgresql[i]) { // not required
			continue
		}

		if o.Postgresql[i] != nil {
			if err := o.Postgresql[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listServicesOk" + "." + "postgresql" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListServicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListServicesOKBody) UnmarshalBinary(b []byte) error {
	var res ListServicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MongodbItems0 MongoDBService represents a generic MongoDB instance.
swagger:model MongodbItems0
*/
type MongodbItems0 struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this mongodb items0
func (o *MongodbItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MongodbItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MongodbItems0) UnmarshalBinary(b []byte) error {
	var res MongodbItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MysqlItems0 MySQLService represents a generic MySQL instance.
swagger:model MysqlItems0
*/
type MysqlItems0 struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this mysql items0
func (o *MysqlItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MysqlItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MysqlItems0) UnmarshalBinary(b []byte) error {
	var res MysqlItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostgresqlItems0 PostgreSQLService represents a generic PostgreSQL instance.
swagger:model PostgresqlItems0
*/
type PostgresqlItems0 struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this postgresql items0
func (o *PostgresqlItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostgresqlItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostgresqlItems0) UnmarshalBinary(b []byte) error {
	var res PostgresqlItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
