// Code generated by go-swagger; DO NOT EDIT.

package psmdb_cluster

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

// ListPSMDBClustersReader is a Reader for the ListPSMDBClusters structure.
type ListPSMDBClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPSMDBClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPSMDBClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListPSMDBClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPSMDBClustersOK creates a ListPSMDBClustersOK with default headers values
func NewListPSMDBClustersOK() *ListPSMDBClustersOK {
	return &ListPSMDBClustersOK{}
}

/*ListPSMDBClustersOK handles this case with default header values.

A successful response.
*/
type ListPSMDBClustersOK struct {
	Payload *ListPSMDBClustersOKBody
}

func (o *ListPSMDBClustersOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PSMDBClusters/List][%d] listPsmdbClustersOk  %+v", 200, o.Payload)
}

func (o *ListPSMDBClustersOK) GetPayload() *ListPSMDBClustersOKBody {
	return o.Payload
}

func (o *ListPSMDBClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListPSMDBClustersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPSMDBClustersDefault creates a ListPSMDBClustersDefault with default headers values
func NewListPSMDBClustersDefault(code int) *ListPSMDBClustersDefault {
	return &ListPSMDBClustersDefault{
		_statusCode: code,
	}
}

/*ListPSMDBClustersDefault handles this case with default header values.

An unexpected error response.
*/
type ListPSMDBClustersDefault struct {
	_statusCode int

	Payload *ListPSMDBClustersDefaultBody
}

// Code gets the status code for the list PSMDB clusters default response
func (o *ListPSMDBClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListPSMDBClustersDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/PSMDBClusters/List][%d] ListPSMDBClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListPSMDBClustersDefault) GetPayload() *ListPSMDBClustersDefaultBody {
	return o.Payload
}

func (o *ListPSMDBClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListPSMDBClustersDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ClustersItems0 Cluster represents PSMDB cluster information.
swagger:model ClustersItems0
*/
type ClustersItems0 struct {

	// Cluster name.
	Name string `json:"name,omitempty"`

	// PSMDBClusterState represents PSMDB cluster CR state.
	//
	//  - PSMDB_CLUSTER_STATE_INVALID: PSMDB_CLUSTER_STATE_INVALID represents unknown state.
	//  - PSMDB_CLUSTER_STATE_CHANGING: PSMDB_CLUSTER_STATE_CHANGING represents a cluster being changed.
	//  - PSMDB_CLUSTER_STATE_READY: PSMDB_CLUSTER_STATE_READY represents a cluster without pending changes.
	//  - PSMDB_CLUSTER_STATE_FAILED: PSMDB_CLUSTER_STATE_FAILED represents a failed cluster.
	//  - PSMDB_CLUSTER_STATE_DELETING: PSMDB_CLUSTER_STATE_DELETING represents a cluster being deleting.
	// Enum: [PSMDB_CLUSTER_STATE_INVALID PSMDB_CLUSTER_STATE_CHANGING PSMDB_CLUSTER_STATE_READY PSMDB_CLUSTER_STATE_FAILED PSMDB_CLUSTER_STATE_DELETING]
	State *string `json:"state,omitempty"`

	// operation
	Operation *ClustersItems0Operation `json:"operation,omitempty"`

	// params
	Params *ClustersItems0Params `json:"params,omitempty"`
}

// Validate validates this clusters items0
func (o *ClustersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOperation(formats); err != nil {
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

var clustersItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PSMDB_CLUSTER_STATE_INVALID","PSMDB_CLUSTER_STATE_CHANGING","PSMDB_CLUSTER_STATE_READY","PSMDB_CLUSTER_STATE_FAILED","PSMDB_CLUSTER_STATE_DELETING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clustersItems0TypeStatePropEnum = append(clustersItems0TypeStatePropEnum, v)
	}
}

const (

	// ClustersItems0StatePSMDBCLUSTERSTATEINVALID captures enum value "PSMDB_CLUSTER_STATE_INVALID"
	ClustersItems0StatePSMDBCLUSTERSTATEINVALID string = "PSMDB_CLUSTER_STATE_INVALID"

	// ClustersItems0StatePSMDBCLUSTERSTATECHANGING captures enum value "PSMDB_CLUSTER_STATE_CHANGING"
	ClustersItems0StatePSMDBCLUSTERSTATECHANGING string = "PSMDB_CLUSTER_STATE_CHANGING"

	// ClustersItems0StatePSMDBCLUSTERSTATEREADY captures enum value "PSMDB_CLUSTER_STATE_READY"
	ClustersItems0StatePSMDBCLUSTERSTATEREADY string = "PSMDB_CLUSTER_STATE_READY"

	// ClustersItems0StatePSMDBCLUSTERSTATEFAILED captures enum value "PSMDB_CLUSTER_STATE_FAILED"
	ClustersItems0StatePSMDBCLUSTERSTATEFAILED string = "PSMDB_CLUSTER_STATE_FAILED"

	// ClustersItems0StatePSMDBCLUSTERSTATEDELETING captures enum value "PSMDB_CLUSTER_STATE_DELETING"
	ClustersItems0StatePSMDBCLUSTERSTATEDELETING string = "PSMDB_CLUSTER_STATE_DELETING"
)

// prop value enum
func (o *ClustersItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clustersItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ClustersItems0) validateState(formats strfmt.Registry) error {

	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(o.Operation) { // not required
		return nil
	}

	if o.Operation != nil {
		if err := o.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (o *ClustersItems0) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0) UnmarshalBinary(b []byte) error {
	var res ClustersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0Operation RunningOperation respresents a long-running operation.
swagger:model ClustersItems0Operation
*/
type ClustersItems0Operation struct {

	// Progress from 0.0 to 1.0; can decrease compared to the previous value.
	Progress float32 `json:"progress,omitempty"`

	// Text describing the current operation progress step.
	Message string `json:"message,omitempty"`
}

// Validate validates this clusters items0 operation
func (o *ClustersItems0Operation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0Operation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0Operation) UnmarshalBinary(b []byte) error {
	var res ClustersItems0Operation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0Params PSMDBClusterParams represents PSMDB cluster parameters that can be updated.
swagger:model ClustersItems0Params
*/
type ClustersItems0Params struct {

	// Cluster size.
	ClusterSize int32 `json:"cluster_size,omitempty"`

	// replicaset
	Replicaset *ClustersItems0ParamsReplicaset `json:"replicaset,omitempty"`
}

// Validate validates this clusters items0 params
func (o *ClustersItems0Params) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReplicaset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClustersItems0Params) validateReplicaset(formats strfmt.Registry) error {

	if swag.IsZero(o.Replicaset) { // not required
		return nil
	}

	if o.Replicaset != nil {
		if err := o.Replicaset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params" + "." + "replicaset")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0Params) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0Params) UnmarshalBinary(b []byte) error {
	var res ClustersItems0Params
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0ParamsReplicaset ReplicaSet container parameters.
swagger:model ClustersItems0ParamsReplicaset
*/
type ClustersItems0ParamsReplicaset struct {

	// Disk size in megabytes.
	DiskSize string `json:"disk_size,omitempty"`

	// compute resources
	ComputeResources *ClustersItems0ParamsReplicasetComputeResources `json:"compute_resources,omitempty"`
}

// Validate validates this clusters items0 params replicaset
func (o *ClustersItems0ParamsReplicaset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComputeResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClustersItems0ParamsReplicaset) validateComputeResources(formats strfmt.Registry) error {

	if swag.IsZero(o.ComputeResources) { // not required
		return nil
	}

	if o.ComputeResources != nil {
		if err := o.ComputeResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params" + "." + "replicaset" + "." + "compute_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0ParamsReplicaset) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0ParamsReplicaset) UnmarshalBinary(b []byte) error {
	var res ClustersItems0ParamsReplicaset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ClustersItems0ParamsReplicasetComputeResources ComputeResources represents container computer resources requests or limits.
swagger:model ClustersItems0ParamsReplicasetComputeResources
*/
type ClustersItems0ParamsReplicasetComputeResources struct {

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CPUm int32 `json:"cpu_m,omitempty"`

	// Memory in bytes.
	MemoryBytes string `json:"memory_bytes,omitempty"`
}

// Validate validates this clusters items0 params replicaset compute resources
func (o *ClustersItems0ParamsReplicasetComputeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClustersItems0ParamsReplicasetComputeResources) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClustersItems0ParamsReplicasetComputeResources) UnmarshalBinary(b []byte) error {
	var res ClustersItems0ParamsReplicasetComputeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListPSMDBClustersBody list PSMDB clusters body
swagger:model ListPSMDBClustersBody
*/
type ListPSMDBClustersBody struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`
}

// Validate validates this list PSMDB clusters body
func (o *ListPSMDBClustersBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListPSMDBClustersBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPSMDBClustersBody) UnmarshalBinary(b []byte) error {
	var res ListPSMDBClustersBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListPSMDBClustersDefaultBody list PSMDB clusters default body
swagger:model ListPSMDBClustersDefaultBody
*/
type ListPSMDBClustersDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list PSMDB clusters default body
func (o *ListPSMDBClustersDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPSMDBClustersDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ListPSMDBClusters default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListPSMDBClustersDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPSMDBClustersDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListPSMDBClustersDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListPSMDBClustersOKBody list PSMDB clusters OK body
swagger:model ListPSMDBClustersOKBody
*/
type ListPSMDBClustersOKBody struct {

	// PSMDB clusters information.
	Clusters []*ClustersItems0 `json:"clusters"`
}

// Validate validates this list PSMDB clusters OK body
func (o *ListPSMDBClustersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListPSMDBClustersOKBody) validateClusters(formats strfmt.Registry) error {

	if swag.IsZero(o.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(o.Clusters); i++ {
		if swag.IsZero(o.Clusters[i]) { // not required
			continue
		}

		if o.Clusters[i] != nil {
			if err := o.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listPsmdbClustersOk" + "." + "clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListPSMDBClustersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListPSMDBClustersOKBody) UnmarshalBinary(b []byte) error {
	var res ListPSMDBClustersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
