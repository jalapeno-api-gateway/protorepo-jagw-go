// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: requestservice/requestservice.proto

package jagw

import (
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

type TopologyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys          []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	PropertyNames []string `protobuf:"bytes,2,rep,name=property_names,json=propertyNames" json:"property_names,omitempty"`
}

func (x *TopologyRequest) Reset() {
	*x = TopologyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestservice_requestservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopologyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopologyRequest) ProtoMessage() {}

func (x *TopologyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requestservice_requestservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopologyRequest.ProtoReflect.Descriptor instead.
func (*TopologyRequest) Descriptor() ([]byte, []int) {
	return file_requestservice_requestservice_proto_rawDescGZIP(), []int{0}
}

func (x *TopologyRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *TopologyRequest) GetPropertyNames() []string {
	if x != nil {
		return x.PropertyNames
	}
	return nil
}

type TelemetryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceIds  []*InterfaceIdentifier `protobuf:"bytes,1,rep,name=interface_ids,json=interfaceIds" json:"interface_ids,omitempty"`
	PropertyNames []string               `protobuf:"bytes,2,rep,name=property_names,json=propertyNames" json:"property_names,omitempty"`
}

func (x *TelemetryRequest) Reset() {
	*x = TelemetryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestservice_requestservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryRequest) ProtoMessage() {}

func (x *TelemetryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requestservice_requestservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryRequest.ProtoReflect.Descriptor instead.
func (*TelemetryRequest) Descriptor() ([]byte, []int) {
	return file_requestservice_requestservice_proto_rawDescGZIP(), []int{1}
}

func (x *TelemetryRequest) GetInterfaceIds() []*InterfaceIdentifier {
	if x != nil {
		return x.InterfaceIds
	}
	return nil
}

func (x *TelemetryRequest) GetPropertyNames() []string {
	if x != nil {
		return x.PropertyNames
	}
	return nil
}

type LsNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LsNodes []*LsNode `protobuf:"bytes,1,rep,name=ls_nodes,json=lsNodes" json:"ls_nodes,omitempty"`
}

func (x *LsNodeResponse) Reset() {
	*x = LsNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestservice_requestservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsNodeResponse) ProtoMessage() {}

func (x *LsNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requestservice_requestservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsNodeResponse.ProtoReflect.Descriptor instead.
func (*LsNodeResponse) Descriptor() ([]byte, []int) {
	return file_requestservice_requestservice_proto_rawDescGZIP(), []int{2}
}

func (x *LsNodeResponse) GetLsNodes() []*LsNode {
	if x != nil {
		return x.LsNodes
	}
	return nil
}

type LsLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LsLinks []*LsLink `protobuf:"bytes,1,rep,name=ls_links,json=lsLinks" json:"ls_links,omitempty"`
}

func (x *LsLinkResponse) Reset() {
	*x = LsLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestservice_requestservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsLinkResponse) ProtoMessage() {}

func (x *LsLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requestservice_requestservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsLinkResponse.ProtoReflect.Descriptor instead.
func (*LsLinkResponse) Descriptor() ([]byte, []int) {
	return file_requestservice_requestservice_proto_rawDescGZIP(), []int{3}
}

func (x *LsLinkResponse) GetLsLinks() []*LsLink {
	if x != nil {
		return x.LsLinks
	}
	return nil
}

type LsPrefixResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LsPrefixes []*LsPrefix `protobuf:"bytes,1,rep,name=ls_prefixes,json=lsPrefixes" json:"ls_prefixes,omitempty"`
}

func (x *LsPrefixResponse) Reset() {
	*x = LsPrefixResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestservice_requestservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsPrefixResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsPrefixResponse) ProtoMessage() {}

func (x *LsPrefixResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requestservice_requestservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsPrefixResponse.ProtoReflect.Descriptor instead.
func (*LsPrefixResponse) Descriptor() ([]byte, []int) {
	return file_requestservice_requestservice_proto_rawDescGZIP(), []int{4}
}

func (x *LsPrefixResponse) GetLsPrefixes() []*LsPrefix {
	if x != nil {
		return x.LsPrefixes
	}
	return nil
}

type LsSrv6SidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LsSrv6Sids []*LsSrv6Sid `protobuf:"bytes,1,rep,name=ls_srv6_sids,json=lsSrv6Sids" json:"ls_srv6_sids,omitempty"`
}

func (x *LsSrv6SidResponse) Reset() {
	*x = LsSrv6SidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestservice_requestservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsSrv6SidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsSrv6SidResponse) ProtoMessage() {}

func (x *LsSrv6SidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requestservice_requestservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsSrv6SidResponse.ProtoReflect.Descriptor instead.
func (*LsSrv6SidResponse) Descriptor() ([]byte, []int) {
	return file_requestservice_requestservice_proto_rawDescGZIP(), []int{5}
}

func (x *LsSrv6SidResponse) GetLsSrv6Sids() []*LsSrv6Sid {
	if x != nil {
		return x.LsSrv6Sids
	}
	return nil
}

type TelemetryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TelemetryData []*TelemetryData `protobuf:"bytes,1,rep,name=telemetry_data,json=telemetryData" json:"telemetry_data,omitempty"`
}

func (x *TelemetryResponse) Reset() {
	*x = TelemetryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestservice_requestservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryResponse) ProtoMessage() {}

func (x *TelemetryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requestservice_requestservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryResponse.ProtoReflect.Descriptor instead.
func (*TelemetryResponse) Descriptor() ([]byte, []int) {
	return file_requestservice_requestservice_proto_rawDescGZIP(), []int{6}
}

func (x *TelemetryResponse) GetTelemetryData() []*TelemetryData {
	if x != nil {
		return x.TelemetryData
	}
	return nil
}

type TelemetryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceId             *InterfaceIdentifier `protobuf:"bytes,1,req,name=interface_id,json=interfaceId" json:"interface_id,omitempty"`
	Ipv4Address             *string              `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	DataRate                *int64               `protobuf:"varint,3,opt,name=data_rate,json=dataRate" json:"data_rate,omitempty"`
	PacketsSent             *int64               `protobuf:"varint,4,opt,name=packets_sent,json=packetsSent" json:"packets_sent,omitempty"`
	PacketsReceived         *int64               `protobuf:"varint,5,opt,name=packets_received,json=packetsReceived" json:"packets_received,omitempty"`
	State                   *string              `protobuf:"bytes,6,opt,name=state" json:"state,omitempty"`
	LastStateTransitionTime *int64               `protobuf:"varint,7,opt,name=last_state_transition_time,json=lastStateTransitionTime" json:"last_state_transition_time,omitempty"`
}

func (x *TelemetryData) Reset() {
	*x = TelemetryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestservice_requestservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryData) ProtoMessage() {}

func (x *TelemetryData) ProtoReflect() protoreflect.Message {
	mi := &file_requestservice_requestservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryData.ProtoReflect.Descriptor instead.
func (*TelemetryData) Descriptor() ([]byte, []int) {
	return file_requestservice_requestservice_proto_rawDescGZIP(), []int{7}
}

func (x *TelemetryData) GetInterfaceId() *InterfaceIdentifier {
	if x != nil {
		return x.InterfaceId
	}
	return nil
}

func (x *TelemetryData) GetIpv4Address() string {
	if x != nil && x.Ipv4Address != nil {
		return *x.Ipv4Address
	}
	return ""
}

func (x *TelemetryData) GetDataRate() int64 {
	if x != nil && x.DataRate != nil {
		return *x.DataRate
	}
	return 0
}

func (x *TelemetryData) GetPacketsSent() int64 {
	if x != nil && x.PacketsSent != nil {
		return *x.PacketsSent
	}
	return 0
}

func (x *TelemetryData) GetPacketsReceived() int64 {
	if x != nil && x.PacketsReceived != nil {
		return *x.PacketsReceived
	}
	return 0
}

func (x *TelemetryData) GetState() string {
	if x != nil && x.State != nil {
		return *x.State
	}
	return ""
}

func (x *TelemetryData) GetLastStateTransitionTime() int64 {
	if x != nil && x.LastStateTransitionTime != nil {
		return *x.LastStateTransitionTime
	}
	return 0
}

var File_requestservice_requestservice_proto protoreflect.FileDescriptor

var file_requestservice_requestservice_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6a, 0x61, 0x67, 0x77, 0x1a, 0x13, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4c, 0x0a, 0x0f, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x79, 0x0a, 0x10, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6a, 0x61, 0x67,
	0x77, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x0e, 0x4c, 0x73,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08,
	0x6c, 0x73, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x6c, 0x73,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x0e, 0x4c, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x6c, 0x73, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6a, 0x61, 0x67, 0x77,
	0x2e, 0x4c, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x07, 0x6c, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x22, 0x43, 0x0a, 0x10, 0x4c, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x6c, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6a, 0x61, 0x67, 0x77,
	0x2e, 0x4c, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x0a, 0x6c, 0x73, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x65, 0x73, 0x22, 0x46, 0x0a, 0x11, 0x4c, 0x73, 0x53, 0x72, 0x76, 0x36, 0x53,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x6c, 0x73,
	0x5f, 0x73, 0x72, 0x76, 0x36, 0x5f, 0x73, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73, 0x53, 0x72, 0x76, 0x36, 0x53, 0x69,
	0x64, 0x52, 0x0a, 0x6c, 0x73, 0x53, 0x72, 0x76, 0x36, 0x53, 0x69, 0x64, 0x73, 0x22, 0x4f, 0x0a,
	0x11, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6a, 0x61, 0x67,
	0x77, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0d, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x22, 0xae,
	0x02, 0x0a, 0x0d, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x3c, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x52, 0x61, 0x74, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x53, 0x65, 0x6e,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32,
	0xd6, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x12, 0x15, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c,
	0x73, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x15, 0x2e,
	0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4c, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x15, 0x2e,
	0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x73, 0x53, 0x72, 0x76, 0x36, 0x53, 0x69, 0x64, 0x73, 0x12,
	0x15, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73,
	0x53, 0x72, 0x76, 0x36, 0x53, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6c, 0x61, 0x70, 0x65, 0x6e, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6a, 0x61, 0x67, 0x77, 0x3b, 0x6a, 0x61, 0x67, 0x77,
}

var (
	file_requestservice_requestservice_proto_rawDescOnce sync.Once
	file_requestservice_requestservice_proto_rawDescData = file_requestservice_requestservice_proto_rawDesc
)

func file_requestservice_requestservice_proto_rawDescGZIP() []byte {
	file_requestservice_requestservice_proto_rawDescOnce.Do(func() {
		file_requestservice_requestservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_requestservice_requestservice_proto_rawDescData)
	})
	return file_requestservice_requestservice_proto_rawDescData
}

var file_requestservice_requestservice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_requestservice_requestservice_proto_goTypes = []interface{}{
	(*TopologyRequest)(nil),     // 0: jagw.TopologyRequest
	(*TelemetryRequest)(nil),    // 1: jagw.TelemetryRequest
	(*LsNodeResponse)(nil),      // 2: jagw.LsNodeResponse
	(*LsLinkResponse)(nil),      // 3: jagw.LsLinkResponse
	(*LsPrefixResponse)(nil),    // 4: jagw.LsPrefixResponse
	(*LsSrv6SidResponse)(nil),   // 5: jagw.LsSrv6SidResponse
	(*TelemetryResponse)(nil),   // 6: jagw.TelemetryResponse
	(*TelemetryData)(nil),       // 7: jagw.TelemetryData
	(*InterfaceIdentifier)(nil), // 8: jagw.InterfaceIdentifier
	(*LsNode)(nil),              // 9: jagw.LsNode
	(*LsLink)(nil),              // 10: jagw.LsLink
	(*LsPrefix)(nil),            // 11: jagw.LsPrefix
	(*LsSrv6Sid)(nil),           // 12: jagw.LsSrv6Sid
}
var file_requestservice_requestservice_proto_depIdxs = []int32{
	8,  // 0: jagw.TelemetryRequest.interface_ids:type_name -> jagw.InterfaceIdentifier
	9,  // 1: jagw.LsNodeResponse.ls_nodes:type_name -> jagw.LsNode
	10, // 2: jagw.LsLinkResponse.ls_links:type_name -> jagw.LsLink
	11, // 3: jagw.LsPrefixResponse.ls_prefixes:type_name -> jagw.LsPrefix
	12, // 4: jagw.LsSrv6SidResponse.ls_srv6_sids:type_name -> jagw.LsSrv6Sid
	7,  // 5: jagw.TelemetryResponse.telemetry_data:type_name -> jagw.TelemetryData
	8,  // 6: jagw.TelemetryData.interface_id:type_name -> jagw.InterfaceIdentifier
	0,  // 7: jagw.RequestService.GetLsNodes:input_type -> jagw.TopologyRequest
	0,  // 8: jagw.RequestService.GetLsLinks:input_type -> jagw.TopologyRequest
	0,  // 9: jagw.RequestService.GetLsPrefixes:input_type -> jagw.TopologyRequest
	0,  // 10: jagw.RequestService.GetLsSrv6Sids:input_type -> jagw.TopologyRequest
	1,  // 11: jagw.RequestService.GetTelemetryData:input_type -> jagw.TelemetryRequest
	2,  // 12: jagw.RequestService.GetLsNodes:output_type -> jagw.LsNodeResponse
	3,  // 13: jagw.RequestService.GetLsLinks:output_type -> jagw.LsLinkResponse
	4,  // 14: jagw.RequestService.GetLsPrefixes:output_type -> jagw.LsPrefixResponse
	5,  // 15: jagw.RequestService.GetLsSrv6Sids:output_type -> jagw.LsSrv6SidResponse
	6,  // 16: jagw.RequestService.GetTelemetryData:output_type -> jagw.TelemetryResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_requestservice_requestservice_proto_init() }
func file_requestservice_requestservice_proto_init() {
	if File_requestservice_requestservice_proto != nil {
		return
	}
	file_core_topology_proto_init()
	file_core_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_requestservice_requestservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopologyRequest); i {
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
		file_requestservice_requestservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryRequest); i {
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
		file_requestservice_requestservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsNodeResponse); i {
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
		file_requestservice_requestservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsLinkResponse); i {
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
		file_requestservice_requestservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsPrefixResponse); i {
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
		file_requestservice_requestservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsSrv6SidResponse); i {
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
		file_requestservice_requestservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryResponse); i {
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
		file_requestservice_requestservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryData); i {
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
			RawDescriptor: file_requestservice_requestservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_requestservice_requestservice_proto_goTypes,
		DependencyIndexes: file_requestservice_requestservice_proto_depIdxs,
		MessageInfos:      file_requestservice_requestservice_proto_msgTypes,
	}.Build()
	File_requestservice_requestservice_proto = out.File
	file_requestservice_requestservice_proto_rawDesc = nil
	file_requestservice_requestservice_proto_goTypes = nil
	file_requestservice_requestservice_proto_depIdxs = nil
}
