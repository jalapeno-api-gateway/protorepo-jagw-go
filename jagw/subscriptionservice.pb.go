// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: subscriptionservice/subscriptionservice.proto

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

type TopologySubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys       []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Properties []string `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty"`
}

func (x *TopologySubscription) Reset() {
	*x = TopologySubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopologySubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopologySubscription) ProtoMessage() {}

func (x *TopologySubscription) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopologySubscription.ProtoReflect.Descriptor instead.
func (*TopologySubscription) Descriptor() ([]byte, []int) {
	return file_subscriptionservice_subscriptionservice_proto_rawDescGZIP(), []int{0}
}

func (x *TopologySubscription) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *TopologySubscription) GetProperties() []string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type TelemetrySubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorPath    *string         `protobuf:"bytes,1,req,name=sensor_path,json=sensorPath" json:"sensor_path,omitempty"`
	Properties    []string        `protobuf:"bytes,2,rep,name=properties" json:"properties,omitempty"`
	Unflatten     *bool           `protobuf:"varint,3,opt,name=Unflatten" json:"Unflatten,omitempty"`
	StringFilters []*StringFilter `protobuf:"bytes,4,rep,name=string_filters,json=stringFilters" json:"string_filters,omitempty"`
}

func (x *TelemetrySubscription) Reset() {
	*x = TelemetrySubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetrySubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetrySubscription) ProtoMessage() {}

func (x *TelemetrySubscription) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetrySubscription.ProtoReflect.Descriptor instead.
func (*TelemetrySubscription) Descriptor() ([]byte, []int) {
	return file_subscriptionservice_subscriptionservice_proto_rawDescGZIP(), []int{1}
}

func (x *TelemetrySubscription) GetSensorPath() string {
	if x != nil && x.SensorPath != nil {
		return *x.SensorPath
	}
	return ""
}

func (x *TelemetrySubscription) GetProperties() []string {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *TelemetrySubscription) GetUnflatten() bool {
	if x != nil && x.Unflatten != nil {
		return *x.Unflatten
	}
	return false
}

func (x *TelemetrySubscription) GetStringFilters() []*StringFilter {
	if x != nil {
		return x.StringFilters
	}
	return nil
}

type LsNodeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action *string `protobuf:"bytes,1,req,name=action" json:"action,omitempty"`
	Key    *string `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	LsNode *LsNode `protobuf:"bytes,3,opt,name=ls_node,json=lsNode" json:"ls_node,omitempty"`
}

func (x *LsNodeEvent) Reset() {
	*x = LsNodeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsNodeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsNodeEvent) ProtoMessage() {}

func (x *LsNodeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsNodeEvent.ProtoReflect.Descriptor instead.
func (*LsNodeEvent) Descriptor() ([]byte, []int) {
	return file_subscriptionservice_subscriptionservice_proto_rawDescGZIP(), []int{2}
}

func (x *LsNodeEvent) GetAction() string {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return ""
}

func (x *LsNodeEvent) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *LsNodeEvent) GetLsNode() *LsNode {
	if x != nil {
		return x.LsNode
	}
	return nil
}

type LsLinkEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action *string `protobuf:"bytes,1,req,name=action" json:"action,omitempty"`
	Key    *string `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	LsLink *LsLink `protobuf:"bytes,3,opt,name=ls_link,json=lsLink" json:"ls_link,omitempty"`
}

func (x *LsLinkEvent) Reset() {
	*x = LsLinkEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsLinkEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsLinkEvent) ProtoMessage() {}

func (x *LsLinkEvent) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsLinkEvent.ProtoReflect.Descriptor instead.
func (*LsLinkEvent) Descriptor() ([]byte, []int) {
	return file_subscriptionservice_subscriptionservice_proto_rawDescGZIP(), []int{3}
}

func (x *LsLinkEvent) GetAction() string {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return ""
}

func (x *LsLinkEvent) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *LsLinkEvent) GetLsLink() *LsLink {
	if x != nil {
		return x.LsLink
	}
	return nil
}

type LsPrefixEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action   *string   `protobuf:"bytes,1,req,name=action" json:"action,omitempty"`
	Key      *string   `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	LsPrefix *LsPrefix `protobuf:"bytes,3,opt,name=ls_prefix,json=lsPrefix" json:"ls_prefix,omitempty"`
}

func (x *LsPrefixEvent) Reset() {
	*x = LsPrefixEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsPrefixEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsPrefixEvent) ProtoMessage() {}

func (x *LsPrefixEvent) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsPrefixEvent.ProtoReflect.Descriptor instead.
func (*LsPrefixEvent) Descriptor() ([]byte, []int) {
	return file_subscriptionservice_subscriptionservice_proto_rawDescGZIP(), []int{4}
}

func (x *LsPrefixEvent) GetAction() string {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return ""
}

func (x *LsPrefixEvent) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *LsPrefixEvent) GetLsPrefix() *LsPrefix {
	if x != nil {
		return x.LsPrefix
	}
	return nil
}

type LsSrv6SidEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    *string    `protobuf:"bytes,1,req,name=action" json:"action,omitempty"`
	Key       *string    `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	LsSrv6Sid *LsSrv6Sid `protobuf:"bytes,3,opt,name=ls_srv6_sid,json=lsSrv6Sid" json:"ls_srv6_sid,omitempty"`
}

func (x *LsSrv6SidEvent) Reset() {
	*x = LsSrv6SidEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsSrv6SidEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsSrv6SidEvent) ProtoMessage() {}

func (x *LsSrv6SidEvent) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsSrv6SidEvent.ProtoReflect.Descriptor instead.
func (*LsSrv6SidEvent) Descriptor() ([]byte, []int) {
	return file_subscriptionservice_subscriptionservice_proto_rawDescGZIP(), []int{5}
}

func (x *LsSrv6SidEvent) GetAction() string {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return ""
}

func (x *LsSrv6SidEvent) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *LsSrv6SidEvent) GetLsSrv6Sid() *LsSrv6Sid {
	if x != nil {
		return x.LsSrv6Sid
	}
	return nil
}

type LsNodeEdgeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action     *string     `protobuf:"bytes,1,req,name=action" json:"action,omitempty"`
	Key        *string     `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	LsNodeEdge *LsNodeEdge `protobuf:"bytes,3,opt,name=ls_node_edge,json=lsNodeEdge" json:"ls_node_edge,omitempty"`
}

func (x *LsNodeEdgeEvent) Reset() {
	*x = LsNodeEdgeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LsNodeEdgeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LsNodeEdgeEvent) ProtoMessage() {}

func (x *LsNodeEdgeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LsNodeEdgeEvent.ProtoReflect.Descriptor instead.
func (*LsNodeEdgeEvent) Descriptor() ([]byte, []int) {
	return file_subscriptionservice_subscriptionservice_proto_rawDescGZIP(), []int{6}
}

func (x *LsNodeEdgeEvent) GetAction() string {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return ""
}

func (x *LsNodeEdgeEvent) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *LsNodeEdgeEvent) GetLsNodeEdge() *LsNodeEdge {
	if x != nil {
		return x.LsNodeEdge
	}
	return nil
}

type PeerEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action *string `protobuf:"bytes,1,req,name=action" json:"action,omitempty"`
	Key    *string `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Peer   *Peer   `protobuf:"bytes,3,opt,name=peer" json:"peer,omitempty"`
}

func (x *PeerEvent) Reset() {
	*x = PeerEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerEvent) ProtoMessage() {}

func (x *PeerEvent) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerEvent.ProtoReflect.Descriptor instead.
func (*PeerEvent) Descriptor() ([]byte, []int) {
	return file_subscriptionservice_subscriptionservice_proto_rawDescGZIP(), []int{7}
}

func (x *PeerEvent) GetAction() string {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return ""
}

func (x *PeerEvent) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *PeerEvent) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

type TelemetryEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TelemetryData *string `protobuf:"bytes,1,req,name=telemetry_data,json=telemetryData" json:"telemetry_data,omitempty"`
}

func (x *TelemetryEvent) Reset() {
	*x = TelemetryEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryEvent) ProtoMessage() {}

func (x *TelemetryEvent) ProtoReflect() protoreflect.Message {
	mi := &file_subscriptionservice_subscriptionservice_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryEvent.ProtoReflect.Descriptor instead.
func (*TelemetryEvent) Descriptor() ([]byte, []int) {
	return file_subscriptionservice_subscriptionservice_proto_rawDescGZIP(), []int{8}
}

func (x *TelemetryEvent) GetTelemetryData() string {
	if x != nil && x.TelemetryData != nil {
		return *x.TelemetryData
	}
	return ""
}

var File_subscriptionservice_subscriptionservice_proto protoreflect.FileDescriptor

var file_subscriptionservice_subscriptionservice_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6a, 0x61, 0x67, 0x77, 0x1a, 0x13, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a,
	0x0a, 0x14, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x15, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x6e, 0x66, 0x6c, 0x61, 0x74, 0x74,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x55, 0x6e, 0x66, 0x6c, 0x61, 0x74,
	0x74, 0x65, 0x6e, 0x12, 0x39, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6a, 0x61,
	0x67, 0x77, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x5e,
	0x0a, 0x0b, 0x4c, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x6c, 0x73, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e,
	0x4c, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x6c, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x5e,
	0x0a, 0x0b, 0x4c, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x6c, 0x73, 0x5f, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e,
	0x4c, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x06, 0x6c, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x66,
	0x0a, 0x0d, 0x4c, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x09, 0x6c, 0x73, 0x5f,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6a,
	0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x08, 0x6c, 0x73,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x6b, 0x0a, 0x0e, 0x4c, 0x73, 0x53, 0x72, 0x76, 0x36,
	0x53, 0x69, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2f, 0x0a, 0x0b, 0x6c, 0x73, 0x5f, 0x73, 0x72, 0x76, 0x36, 0x5f, 0x73, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c,
	0x73, 0x53, 0x72, 0x76, 0x36, 0x53, 0x69, 0x64, 0x52, 0x09, 0x6c, 0x73, 0x53, 0x72, 0x76, 0x36,
	0x53, 0x69, 0x64, 0x22, 0x6f, 0x0a, 0x0f, 0x4c, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x64, 0x67,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x32, 0x0a, 0x0c, 0x6c, 0x73, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x64, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73,
	0x4e, 0x6f, 0x64, 0x65, 0x45, 0x64, 0x67, 0x65, 0x52, 0x0a, 0x6c, 0x73, 0x4e, 0x6f, 0x64, 0x65,
	0x45, 0x64, 0x67, 0x65, 0x22, 0x55, 0x0a, 0x09, 0x50, 0x65, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x04, 0x70,
	0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6a, 0x61, 0x67, 0x77,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x0e, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x32, 0xad, 0x04, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x12,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4c, 0x73, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x1a, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11,
	0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x6f, 0x4c, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1a, 0x2e, 0x6a, 0x61,
	0x67, 0x77, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c,
	0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4c,
	0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4c, 0x73, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54,
	0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x15,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4c, 0x73, 0x53, 0x72, 0x76,
	0x36, 0x53, 0x69, 0x64, 0x73, 0x12, 0x1a, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x14, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73, 0x53, 0x72, 0x76, 0x36, 0x53,
	0x69, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x16, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4c, 0x73, 0x4e, 0x6f, 0x64, 0x65,
	0x45, 0x64, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x15, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x4c, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x45,
	0x64, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x10,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x12, 0x1a, 0x2e, 0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x6a,
	0x61, 0x67, 0x77, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x51, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x2e,
	0x6a, 0x61, 0x67, 0x77, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x6a, 0x61, 0x67,
	0x77, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6c, 0x61, 0x70, 0x65, 0x6e, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6a, 0x61, 0x67, 0x77, 0x3b, 0x6a, 0x61, 0x67, 0x77,
}

var (
	file_subscriptionservice_subscriptionservice_proto_rawDescOnce sync.Once
	file_subscriptionservice_subscriptionservice_proto_rawDescData = file_subscriptionservice_subscriptionservice_proto_rawDesc
)

func file_subscriptionservice_subscriptionservice_proto_rawDescGZIP() []byte {
	file_subscriptionservice_subscriptionservice_proto_rawDescOnce.Do(func() {
		file_subscriptionservice_subscriptionservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscriptionservice_subscriptionservice_proto_rawDescData)
	})
	return file_subscriptionservice_subscriptionservice_proto_rawDescData
}

var file_subscriptionservice_subscriptionservice_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_subscriptionservice_subscriptionservice_proto_goTypes = []interface{}{
	(*TopologySubscription)(nil),  // 0: jagw.TopologySubscription
	(*TelemetrySubscription)(nil), // 1: jagw.TelemetrySubscription
	(*LsNodeEvent)(nil),           // 2: jagw.LsNodeEvent
	(*LsLinkEvent)(nil),           // 3: jagw.LsLinkEvent
	(*LsPrefixEvent)(nil),         // 4: jagw.LsPrefixEvent
	(*LsSrv6SidEvent)(nil),        // 5: jagw.LsSrv6SidEvent
	(*LsNodeEdgeEvent)(nil),       // 6: jagw.LsNodeEdgeEvent
	(*PeerEvent)(nil),             // 7: jagw.PeerEvent
	(*TelemetryEvent)(nil),        // 8: jagw.TelemetryEvent
	(*StringFilter)(nil),          // 9: jagw.StringFilter
	(*LsNode)(nil),                // 10: jagw.LsNode
	(*LsLink)(nil),                // 11: jagw.LsLink
	(*LsPrefix)(nil),              // 12: jagw.LsPrefix
	(*LsSrv6Sid)(nil),             // 13: jagw.LsSrv6Sid
	(*LsNodeEdge)(nil),            // 14: jagw.LsNodeEdge
	(*Peer)(nil),                  // 15: jagw.Peer
}
var file_subscriptionservice_subscriptionservice_proto_depIdxs = []int32{
	9,  // 0: jagw.TelemetrySubscription.string_filters:type_name -> jagw.StringFilter
	10, // 1: jagw.LsNodeEvent.ls_node:type_name -> jagw.LsNode
	11, // 2: jagw.LsLinkEvent.ls_link:type_name -> jagw.LsLink
	12, // 3: jagw.LsPrefixEvent.ls_prefix:type_name -> jagw.LsPrefix
	13, // 4: jagw.LsSrv6SidEvent.ls_srv6_sid:type_name -> jagw.LsSrv6Sid
	14, // 5: jagw.LsNodeEdgeEvent.ls_node_edge:type_name -> jagw.LsNodeEdge
	15, // 6: jagw.PeerEvent.peer:type_name -> jagw.Peer
	0,  // 7: jagw.SubscriptionService.SubscribeToLsNodes:input_type -> jagw.TopologySubscription
	0,  // 8: jagw.SubscriptionService.SubscribeToLsLinks:input_type -> jagw.TopologySubscription
	0,  // 9: jagw.SubscriptionService.SubscribeToLsPrefixes:input_type -> jagw.TopologySubscription
	0,  // 10: jagw.SubscriptionService.SubscribeToLsSrv6Sids:input_type -> jagw.TopologySubscription
	0,  // 11: jagw.SubscriptionService.SubscribeToLsNodeEdges:input_type -> jagw.TopologySubscription
	0,  // 12: jagw.SubscriptionService.SubscribeToPeers:input_type -> jagw.TopologySubscription
	1,  // 13: jagw.SubscriptionService.SubscribeToTelemetryData:input_type -> jagw.TelemetrySubscription
	2,  // 14: jagw.SubscriptionService.SubscribeToLsNodes:output_type -> jagw.LsNodeEvent
	3,  // 15: jagw.SubscriptionService.SubscribeToLsLinks:output_type -> jagw.LsLinkEvent
	4,  // 16: jagw.SubscriptionService.SubscribeToLsPrefixes:output_type -> jagw.LsPrefixEvent
	5,  // 17: jagw.SubscriptionService.SubscribeToLsSrv6Sids:output_type -> jagw.LsSrv6SidEvent
	6,  // 18: jagw.SubscriptionService.SubscribeToLsNodeEdges:output_type -> jagw.LsNodeEdgeEvent
	7,  // 19: jagw.SubscriptionService.SubscribeToPeers:output_type -> jagw.PeerEvent
	8,  // 20: jagw.SubscriptionService.SubscribeToTelemetryData:output_type -> jagw.TelemetryEvent
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_subscriptionservice_subscriptionservice_proto_init() }
func file_subscriptionservice_subscriptionservice_proto_init() {
	if File_subscriptionservice_subscriptionservice_proto != nil {
		return
	}
	file_core_topology_proto_init()
	file_core_filters_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_subscriptionservice_subscriptionservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopologySubscription); i {
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
		file_subscriptionservice_subscriptionservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetrySubscription); i {
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
		file_subscriptionservice_subscriptionservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsNodeEvent); i {
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
		file_subscriptionservice_subscriptionservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsLinkEvent); i {
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
		file_subscriptionservice_subscriptionservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsPrefixEvent); i {
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
		file_subscriptionservice_subscriptionservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsSrv6SidEvent); i {
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
		file_subscriptionservice_subscriptionservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LsNodeEdgeEvent); i {
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
		file_subscriptionservice_subscriptionservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerEvent); i {
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
		file_subscriptionservice_subscriptionservice_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryEvent); i {
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
			RawDescriptor: file_subscriptionservice_subscriptionservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscriptionservice_subscriptionservice_proto_goTypes,
		DependencyIndexes: file_subscriptionservice_subscriptionservice_proto_depIdxs,
		MessageInfos:      file_subscriptionservice_subscriptionservice_proto_msgTypes,
	}.Build()
	File_subscriptionservice_subscriptionservice_proto = out.File
	file_subscriptionservice_subscriptionservice_proto_rawDesc = nil
	file_subscriptionservice_subscriptionservice_proto_goTypes = nil
	file_subscriptionservice_subscriptionservice_proto_depIdxs = nil
}
