// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/ia/rules.proto

package iav1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Rule) Validate() error {
	for _, item := range this.Params {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
			}
		}
	}
	if this.DefaultFor != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DefaultFor); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DefaultFor", err)
		}
	}
	if this.For != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.For); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("For", err)
		}
	}
	return nil
}
func (this *Rule_Param) Validate() error {
	return nil
}
func (this *ListAlertingRulesRequest) Validate() error {
	return nil
}
func (this *ListAlertingRulesResponse) Validate() error {
	for _, item := range this.Rules {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
			}
		}
	}
	return nil
}
func (this *ChangeAlertingRulesRequest) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	for _, item := range this.Params {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
			}
		}
	}
	if this.For != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.For); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("For", err)
		}
	}
	return nil
}
func (this *ChangeAlertingRulesRequest_Param) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *ChangeAlertingRulesResponse) Validate() error {
	if this.Rules != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Rules); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Rules", err)
		}
	}
	return nil
}
