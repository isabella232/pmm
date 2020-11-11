// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: managementpb/dbaas/kubernetes.proto

package dbaasv1beta1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// KubeAuth represents Kubernetes / kubectl authentication and authorization information.
type KubeAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubeconfig file content.
	Kubeconfig string `protobuf:"bytes,1,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
}

func (x *KubeAuth) Reset() {
	*x = KubeAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeAuth) ProtoMessage() {}

func (x *KubeAuth) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeAuth.ProtoReflect.Descriptor instead.
func (*KubeAuth) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_kubernetes_proto_rawDescGZIP(), []int{0}
}

func (x *KubeAuth) GetKubeconfig() string {
	if x != nil {
		return x.Kubeconfig
	}
	return ""
}

type ListKubernetesClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListKubernetesClustersRequest) Reset() {
	*x = ListKubernetesClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKubernetesClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKubernetesClustersRequest) ProtoMessage() {}

func (x *ListKubernetesClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKubernetesClustersRequest.ProtoReflect.Descriptor instead.
func (*ListKubernetesClustersRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_kubernetes_proto_rawDescGZIP(), []int{1}
}

type ListKubernetesClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes clusters.
	KubernetesClusters []*ListKubernetesClustersResponse_Cluster `protobuf:"bytes,1,rep,name=kubernetes_clusters,json=kubernetesClusters,proto3" json:"kubernetes_clusters,omitempty"`
}

func (x *ListKubernetesClustersResponse) Reset() {
	*x = ListKubernetesClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKubernetesClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKubernetesClustersResponse) ProtoMessage() {}

func (x *ListKubernetesClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKubernetesClustersResponse.ProtoReflect.Descriptor instead.
func (*ListKubernetesClustersResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_kubernetes_proto_rawDescGZIP(), []int{2}
}

func (x *ListKubernetesClustersResponse) GetKubernetesClusters() []*ListKubernetesClustersResponse_Cluster {
	if x != nil {
		return x.KubernetesClusters
	}
	return nil
}

type RegisterKubernetesClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes cluster name.
	KubernetesClusterName string `protobuf:"bytes,1,opt,name=kubernetes_cluster_name,json=kubernetesClusterName,proto3" json:"kubernetes_cluster_name,omitempty"`
	// Kubernetes auth.
	KubeAuth *KubeAuth `protobuf:"bytes,2,opt,name=kube_auth,json=kubeAuth,proto3" json:"kube_auth,omitempty"`
	// PMM server public address URL
	PublicAddressUrl string `protobuf:"bytes,3,opt,name=public_address_url,json=publicAddressUrl,proto3" json:"public_address_url,omitempty"`
}

func (x *RegisterKubernetesClusterRequest) Reset() {
	*x = RegisterKubernetesClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterKubernetesClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterKubernetesClusterRequest) ProtoMessage() {}

func (x *RegisterKubernetesClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterKubernetesClusterRequest.ProtoReflect.Descriptor instead.
func (*RegisterKubernetesClusterRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_kubernetes_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterKubernetesClusterRequest) GetKubernetesClusterName() string {
	if x != nil {
		return x.KubernetesClusterName
	}
	return ""
}

func (x *RegisterKubernetesClusterRequest) GetKubeAuth() *KubeAuth {
	if x != nil {
		return x.KubeAuth
	}
	return nil
}

func (x *RegisterKubernetesClusterRequest) GetPublicAddressUrl() string {
	if x != nil {
		return x.PublicAddressUrl
	}
	return ""
}

type RegisterKubernetesClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterKubernetesClusterResponse) Reset() {
	*x = RegisterKubernetesClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterKubernetesClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterKubernetesClusterResponse) ProtoMessage() {}

func (x *RegisterKubernetesClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterKubernetesClusterResponse.ProtoReflect.Descriptor instead.
func (*RegisterKubernetesClusterResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_kubernetes_proto_rawDescGZIP(), []int{4}
}

type UnregisterKubernetesClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes cluster name.
	KubernetesClusterName string `protobuf:"bytes,1,opt,name=kubernetes_cluster_name,json=kubernetesClusterName,proto3" json:"kubernetes_cluster_name,omitempty"`
}

func (x *UnregisterKubernetesClusterRequest) Reset() {
	*x = UnregisterKubernetesClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterKubernetesClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterKubernetesClusterRequest) ProtoMessage() {}

func (x *UnregisterKubernetesClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterKubernetesClusterRequest.ProtoReflect.Descriptor instead.
func (*UnregisterKubernetesClusterRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_kubernetes_proto_rawDescGZIP(), []int{5}
}

func (x *UnregisterKubernetesClusterRequest) GetKubernetesClusterName() string {
	if x != nil {
		return x.KubernetesClusterName
	}
	return ""
}

type UnregisterKubernetesClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnregisterKubernetesClusterResponse) Reset() {
	*x = UnregisterKubernetesClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterKubernetesClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterKubernetesClusterResponse) ProtoMessage() {}

func (x *UnregisterKubernetesClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterKubernetesClusterResponse.ProtoReflect.Descriptor instead.
func (*UnregisterKubernetesClusterResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_kubernetes_proto_rawDescGZIP(), []int{6}
}

// Cluster contains public info about kubernetes cluster.
type ListKubernetesClustersResponse_Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes cluster name.
	KubernetesClusterName string `protobuf:"bytes,1,opt,name=kubernetes_cluster_name,json=kubernetesClusterName,proto3" json:"kubernetes_cluster_name,omitempty"`
}

func (x *ListKubernetesClustersResponse_Cluster) Reset() {
	*x = ListKubernetesClustersResponse_Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKubernetesClustersResponse_Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKubernetesClustersResponse_Cluster) ProtoMessage() {}

func (x *ListKubernetesClustersResponse_Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_kubernetes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKubernetesClustersResponse_Cluster.ProtoReflect.Descriptor instead.
func (*ListKubernetesClustersResponse_Cluster) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_kubernetes_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListKubernetesClustersResponse_Cluster) GetKubernetesClusterName() string {
	if x != nil {
		return x.KubernetesClusterName
	}
	return ""
}

var File_managementpb_dbaas_kubernetes_proto protoreflect.FileDescriptor

var file_managementpb_dbaas_kubernetes_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64,
	0x62, 0x61, 0x61, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x08, 0x4b, 0x75,
	0x62, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x1f,
	0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xcb, 0x01, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x66, 0x0a, 0x13, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x41, 0x0a, 0x07, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x17, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd6, 0x01,
	0x0a, 0x20, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3e, 0x0a, 0x17, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x15, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x41, 0x75, 0x74, 0x68, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x08, 0x6b, 0x75, 0x62, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x34, 0x0a, 0x12, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x58, 0x01, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x22, 0x23, 0x0a, 0x21, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64, 0x0a, 0x22, 0x55,
	0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3e, 0x0a, 0x17, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x15, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x25, 0x0a, 0x23, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa9, 0x04, 0x0a, 0x0a, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0xa6, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x2c, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x44, 0x42, 0x61, 0x61, 0x53, 0x2f, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0xb3, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2f,
	0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x22, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x44, 0x42, 0x61, 0x61, 0x53, 0x2f,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0xbb, 0x01, 0x0a, 0x1b, 0x55, 0x6e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x64, 0x62, 0x61, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x22, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x44, 0x42, 0x61, 0x61, 0x53, 0x2f, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x3a, 0x01, 0x2a, 0x42, 0x25, 0x5a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64, 0x62, 0x61, 0x61, 0x73, 0x3b, 0x64,
	0x62, 0x61, 0x61, 0x73, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_managementpb_dbaas_kubernetes_proto_rawDescOnce sync.Once
	file_managementpb_dbaas_kubernetes_proto_rawDescData = file_managementpb_dbaas_kubernetes_proto_rawDesc
)

func file_managementpb_dbaas_kubernetes_proto_rawDescGZIP() []byte {
	file_managementpb_dbaas_kubernetes_proto_rawDescOnce.Do(func() {
		file_managementpb_dbaas_kubernetes_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_dbaas_kubernetes_proto_rawDescData)
	})
	return file_managementpb_dbaas_kubernetes_proto_rawDescData
}

var file_managementpb_dbaas_kubernetes_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_managementpb_dbaas_kubernetes_proto_goTypes = []interface{}{
	(*KubeAuth)(nil),                               // 0: dbaas.v1beta1.KubeAuth
	(*ListKubernetesClustersRequest)(nil),          // 1: dbaas.v1beta1.ListKubernetesClustersRequest
	(*ListKubernetesClustersResponse)(nil),         // 2: dbaas.v1beta1.ListKubernetesClustersResponse
	(*RegisterKubernetesClusterRequest)(nil),       // 3: dbaas.v1beta1.RegisterKubernetesClusterRequest
	(*RegisterKubernetesClusterResponse)(nil),      // 4: dbaas.v1beta1.RegisterKubernetesClusterResponse
	(*UnregisterKubernetesClusterRequest)(nil),     // 5: dbaas.v1beta1.UnregisterKubernetesClusterRequest
	(*UnregisterKubernetesClusterResponse)(nil),    // 6: dbaas.v1beta1.UnregisterKubernetesClusterResponse
	(*ListKubernetesClustersResponse_Cluster)(nil), // 7: dbaas.v1beta1.ListKubernetesClustersResponse.Cluster
}
var file_managementpb_dbaas_kubernetes_proto_depIdxs = []int32{
	7, // 0: dbaas.v1beta1.ListKubernetesClustersResponse.kubernetes_clusters:type_name -> dbaas.v1beta1.ListKubernetesClustersResponse.Cluster
	0, // 1: dbaas.v1beta1.RegisterKubernetesClusterRequest.kube_auth:type_name -> dbaas.v1beta1.KubeAuth
	1, // 2: dbaas.v1beta1.Kubernetes.ListKubernetesClusters:input_type -> dbaas.v1beta1.ListKubernetesClustersRequest
	3, // 3: dbaas.v1beta1.Kubernetes.RegisterKubernetesCluster:input_type -> dbaas.v1beta1.RegisterKubernetesClusterRequest
	5, // 4: dbaas.v1beta1.Kubernetes.UnregisterKubernetesCluster:input_type -> dbaas.v1beta1.UnregisterKubernetesClusterRequest
	2, // 5: dbaas.v1beta1.Kubernetes.ListKubernetesClusters:output_type -> dbaas.v1beta1.ListKubernetesClustersResponse
	4, // 6: dbaas.v1beta1.Kubernetes.RegisterKubernetesCluster:output_type -> dbaas.v1beta1.RegisterKubernetesClusterResponse
	6, // 7: dbaas.v1beta1.Kubernetes.UnregisterKubernetesCluster:output_type -> dbaas.v1beta1.UnregisterKubernetesClusterResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_managementpb_dbaas_kubernetes_proto_init() }
func file_managementpb_dbaas_kubernetes_proto_init() {
	if File_managementpb_dbaas_kubernetes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_dbaas_kubernetes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeAuth); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_dbaas_kubernetes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKubernetesClustersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_dbaas_kubernetes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKubernetesClustersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_dbaas_kubernetes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterKubernetesClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_dbaas_kubernetes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterKubernetesClusterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_dbaas_kubernetes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterKubernetesClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_dbaas_kubernetes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterKubernetesClusterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_dbaas_kubernetes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKubernetesClustersResponse_Cluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_dbaas_kubernetes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_dbaas_kubernetes_proto_goTypes,
		DependencyIndexes: file_managementpb_dbaas_kubernetes_proto_depIdxs,
		MessageInfos:      file_managementpb_dbaas_kubernetes_proto_msgTypes,
	}.Build()
	File_managementpb_dbaas_kubernetes_proto = out.File
	file_managementpb_dbaas_kubernetes_proto_rawDesc = nil
	file_managementpb_dbaas_kubernetes_proto_goTypes = nil
	file_managementpb_dbaas_kubernetes_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KubernetesClient is the client API for Kubernetes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KubernetesClient interface {
	// ListKubernetesClusters returns a list of all registered Kubernetes clusters.
	ListKubernetesClusters(ctx context.Context, in *ListKubernetesClustersRequest, opts ...grpc.CallOption) (*ListKubernetesClustersResponse, error)
	// RegisterKubernetesCluster registers an existing Kubernetes cluster in PMM.
	RegisterKubernetesCluster(ctx context.Context, in *RegisterKubernetesClusterRequest, opts ...grpc.CallOption) (*RegisterKubernetesClusterResponse, error)
	// UnregisterKubernetesCluster removes a registered Kubernetes cluster from PMM.
	UnregisterKubernetesCluster(ctx context.Context, in *UnregisterKubernetesClusterRequest, opts ...grpc.CallOption) (*UnregisterKubernetesClusterResponse, error)
}

type kubernetesClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesClient(cc grpc.ClientConnInterface) KubernetesClient {
	return &kubernetesClient{cc}
}

func (c *kubernetesClient) ListKubernetesClusters(ctx context.Context, in *ListKubernetesClustersRequest, opts ...grpc.CallOption) (*ListKubernetesClustersResponse, error) {
	out := new(ListKubernetesClustersResponse)
	err := c.cc.Invoke(ctx, "/dbaas.v1beta1.Kubernetes/ListKubernetesClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) RegisterKubernetesCluster(ctx context.Context, in *RegisterKubernetesClusterRequest, opts ...grpc.CallOption) (*RegisterKubernetesClusterResponse, error) {
	out := new(RegisterKubernetesClusterResponse)
	err := c.cc.Invoke(ctx, "/dbaas.v1beta1.Kubernetes/RegisterKubernetesCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) UnregisterKubernetesCluster(ctx context.Context, in *UnregisterKubernetesClusterRequest, opts ...grpc.CallOption) (*UnregisterKubernetesClusterResponse, error) {
	out := new(UnregisterKubernetesClusterResponse)
	err := c.cc.Invoke(ctx, "/dbaas.v1beta1.Kubernetes/UnregisterKubernetesCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesServer is the server API for Kubernetes service.
type KubernetesServer interface {
	// ListKubernetesClusters returns a list of all registered Kubernetes clusters.
	ListKubernetesClusters(context.Context, *ListKubernetesClustersRequest) (*ListKubernetesClustersResponse, error)
	// RegisterKubernetesCluster registers an existing Kubernetes cluster in PMM.
	RegisterKubernetesCluster(context.Context, *RegisterKubernetesClusterRequest) (*RegisterKubernetesClusterResponse, error)
	// UnregisterKubernetesCluster removes a registered Kubernetes cluster from PMM.
	UnregisterKubernetesCluster(context.Context, *UnregisterKubernetesClusterRequest) (*UnregisterKubernetesClusterResponse, error)
}

// UnimplementedKubernetesServer can be embedded to have forward compatible implementations.
type UnimplementedKubernetesServer struct {
}

func (*UnimplementedKubernetesServer) ListKubernetesClusters(context.Context, *ListKubernetesClustersRequest) (*ListKubernetesClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKubernetesClusters not implemented")
}
func (*UnimplementedKubernetesServer) RegisterKubernetesCluster(context.Context, *RegisterKubernetesClusterRequest) (*RegisterKubernetesClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterKubernetesCluster not implemented")
}
func (*UnimplementedKubernetesServer) UnregisterKubernetesCluster(context.Context, *UnregisterKubernetesClusterRequest) (*UnregisterKubernetesClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterKubernetesCluster not implemented")
}

func RegisterKubernetesServer(s *grpc.Server, srv KubernetesServer) {
	s.RegisterService(&_Kubernetes_serviceDesc, srv)
}

func _Kubernetes_ListKubernetesClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKubernetesClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).ListKubernetesClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbaas.v1beta1.Kubernetes/ListKubernetesClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).ListKubernetesClusters(ctx, req.(*ListKubernetesClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_RegisterKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterKubernetesClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).RegisterKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbaas.v1beta1.Kubernetes/RegisterKubernetesCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).RegisterKubernetesCluster(ctx, req.(*RegisterKubernetesClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_UnregisterKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterKubernetesClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).UnregisterKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbaas.v1beta1.Kubernetes/UnregisterKubernetesCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).UnregisterKubernetesCluster(ctx, req.(*UnregisterKubernetesClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kubernetes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbaas.v1beta1.Kubernetes",
	HandlerType: (*KubernetesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKubernetesClusters",
			Handler:    _Kubernetes_ListKubernetesClusters_Handler,
		},
		{
			MethodName: "RegisterKubernetesCluster",
			Handler:    _Kubernetes_RegisterKubernetesCluster_Handler,
		},
		{
			MethodName: "UnregisterKubernetesCluster",
			Handler:    _Kubernetes_UnregisterKubernetesCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/dbaas/kubernetes.proto",
}
