// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/checks.proto

package managementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *StartSecurityChecksRequest) Validate() error {
	return nil
}
func (this *StartSecurityChecksResponse) Validate() error {
	return nil
}
func (this *GetSecurityCheckResultsRequest) Validate() error {
	return nil
}
func (this *SecurityCheckResult) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GetSecurityCheckResultsResponse) Validate() error {
	for _, item := range this.Results {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Results", err)
			}
		}
	}
	return nil
}
