// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: proto/tunnel.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ControlMessage_Type int32

const (
	ControlMessage_UNKNOWN ControlMessage_Type = 0
	ControlMessage_PING    ControlMessage_Type = 1
	ControlMessage_PONG    ControlMessage_Type = 2
	ControlMessage_ERROR   ControlMessage_Type = 3
)

// Enum value maps for ControlMessage_Type.
var (
	ControlMessage_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "PING",
		2: "PONG",
		3: "ERROR",
	}
	ControlMessage_Type_value = map[string]int32{
		"UNKNOWN": 0,
		"PING":    1,
		"PONG":    2,
		"ERROR":   3,
	}
)

func (x ControlMessage_Type) Enum() *ControlMessage_Type {
	p := new(ControlMessage_Type)
	*p = x
	return p
}

func (x ControlMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControlMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_tunnel_proto_enumTypes[0].Descriptor()
}

func (ControlMessage_Type) Type() protoreflect.EnumType {
	return &file_proto_tunnel_proto_enumTypes[0]
}

func (x ControlMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ControlMessage_Type.Descriptor instead.
func (ControlMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_tunnel_proto_rawDescGZIP(), []int{3, 0}
}

type ConnectRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AgentId       string                 `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce         string                 `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Signature     string                 `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	mi := &file_proto_tunnel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tunnel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_proto_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *ConnectRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConnectRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ConnectRequest) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *ConnectRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type TunnelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Method        string                 `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Path          string                 `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Headers       map[string]string      `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body          []byte                 `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TunnelRequest) Reset() {
	*x = TunnelRequest{}
	mi := &file_proto_tunnel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TunnelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelRequest) ProtoMessage() {}

func (x *TunnelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tunnel_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelRequest.ProtoReflect.Descriptor instead.
func (*TunnelRequest) Descriptor() ([]byte, []int) {
	return file_proto_tunnel_proto_rawDescGZIP(), []int{1}
}

func (x *TunnelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TunnelRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *TunnelRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TunnelRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *TunnelRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type TunnelResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status        int32                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Headers       map[string]string      `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body          []byte                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TunnelResponse) Reset() {
	*x = TunnelResponse{}
	mi := &file_proto_tunnel_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TunnelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelResponse) ProtoMessage() {}

func (x *TunnelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tunnel_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelResponse.ProtoReflect.Descriptor instead.
func (*TunnelResponse) Descriptor() ([]byte, []int) {
	return file_proto_tunnel_proto_rawDescGZIP(), []int{2}
}

func (x *TunnelResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TunnelResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TunnelResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *TunnelResponse) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ControlMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          ControlMessage_Type    `protobuf:"varint,1,opt,name=type,proto3,enum=tunnel.v1.ControlMessage_Type" json:"type,omitempty"`
	Payload       string                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ControlMessage) Reset() {
	*x = ControlMessage{}
	mi := &file_proto_tunnel_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControlMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlMessage) ProtoMessage() {}

func (x *ControlMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tunnel_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlMessage.ProtoReflect.Descriptor instead.
func (*ControlMessage) Descriptor() ([]byte, []int) {
	return file_proto_tunnel_proto_rawDescGZIP(), []int{3}
}

func (x *ControlMessage) GetType() ControlMessage_Type {
	if x != nil {
		return x.Type
	}
	return ControlMessage_UNKNOWN
}

func (x *ControlMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type TunnelData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Chunk         []byte                 `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TunnelData) Reset() {
	*x = TunnelData{}
	mi := &file_proto_tunnel_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TunnelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelData) ProtoMessage() {}

func (x *TunnelData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tunnel_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelData.ProtoReflect.Descriptor instead.
func (*TunnelData) Descriptor() ([]byte, []int) {
	return file_proto_tunnel_proto_rawDescGZIP(), []int{4}
}

func (x *TunnelData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TunnelData) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type TunnelClose struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TunnelClose) Reset() {
	*x = TunnelClose{}
	mi := &file_proto_tunnel_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TunnelClose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelClose) ProtoMessage() {}

func (x *TunnelClose) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tunnel_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelClose.ProtoReflect.Descriptor instead.
func (*TunnelClose) Descriptor() ([]byte, []int) {
	return file_proto_tunnel_proto_rawDescGZIP(), []int{5}
}

func (x *TunnelClose) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Envelope struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Message:
	//
	//	*Envelope_Connect
	//	*Envelope_Request
	//	*Envelope_Response
	//	*Envelope_Control
	//	*Envelope_Data
	//	*Envelope_Close
	Message       isEnvelope_Message `protobuf_oneof:"message"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	mi := &file_proto_tunnel_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tunnel_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_proto_tunnel_proto_rawDescGZIP(), []int{6}
}

func (x *Envelope) GetMessage() isEnvelope_Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Envelope) GetConnect() *ConnectRequest {
	if x != nil {
		if x, ok := x.Message.(*Envelope_Connect); ok {
			return x.Connect
		}
	}
	return nil
}

func (x *Envelope) GetRequest() *TunnelRequest {
	if x != nil {
		if x, ok := x.Message.(*Envelope_Request); ok {
			return x.Request
		}
	}
	return nil
}

func (x *Envelope) GetResponse() *TunnelResponse {
	if x != nil {
		if x, ok := x.Message.(*Envelope_Response); ok {
			return x.Response
		}
	}
	return nil
}

func (x *Envelope) GetControl() *ControlMessage {
	if x != nil {
		if x, ok := x.Message.(*Envelope_Control); ok {
			return x.Control
		}
	}
	return nil
}

func (x *Envelope) GetData() *TunnelData {
	if x != nil {
		if x, ok := x.Message.(*Envelope_Data); ok {
			return x.Data
		}
	}
	return nil
}

func (x *Envelope) GetClose() *TunnelClose {
	if x != nil {
		if x, ok := x.Message.(*Envelope_Close); ok {
			return x.Close
		}
	}
	return nil
}

type isEnvelope_Message interface {
	isEnvelope_Message()
}

type Envelope_Connect struct {
	Connect *ConnectRequest `protobuf:"bytes,1,opt,name=connect,proto3,oneof"`
}

type Envelope_Request struct {
	Request *TunnelRequest `protobuf:"bytes,2,opt,name=request,proto3,oneof"`
}

type Envelope_Response struct {
	Response *TunnelResponse `protobuf:"bytes,3,opt,name=response,proto3,oneof"`
}

type Envelope_Control struct {
	Control *ControlMessage `protobuf:"bytes,4,opt,name=control,proto3,oneof"`
}

type Envelope_Data struct {
	Data *TunnelData `protobuf:"bytes,5,opt,name=data,proto3,oneof"`
}

type Envelope_Close struct {
	Close *TunnelClose `protobuf:"bytes,6,opt,name=close,proto3,oneof"`
}

func (*Envelope_Connect) isEnvelope_Message() {}

func (*Envelope_Request) isEnvelope_Message() {}

func (*Envelope_Response) isEnvelope_Message() {}

func (*Envelope_Control) isEnvelope_Message() {}

func (*Envelope_Data) isEnvelope_Message() {}

func (*Envelope_Close) isEnvelope_Message() {}

var File_proto_tunnel_proto protoreflect.FileDescriptor

var file_proto_tunnel_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x22,
	0x93, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x0d, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x3f, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xca, 0x01, 0x0a, 0x0e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x40, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x92, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x32, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x32, 0x0a, 0x0a, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x1d, 0x0a, 0x0b, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xcf, 0x02, 0x0a, 0x08, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x34, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x48, 0x00, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2d, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_tunnel_proto_rawDescOnce sync.Once
	file_proto_tunnel_proto_rawDescData []byte
)

func file_proto_tunnel_proto_rawDescGZIP() []byte {
	file_proto_tunnel_proto_rawDescOnce.Do(func() {
		file_proto_tunnel_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_tunnel_proto_rawDesc), len(file_proto_tunnel_proto_rawDesc)))
	})
	return file_proto_tunnel_proto_rawDescData
}

var file_proto_tunnel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_tunnel_proto_goTypes = []any{
	(ControlMessage_Type)(0), // 0: tunnel.v1.ControlMessage.Type
	(*ConnectRequest)(nil),   // 1: tunnel.v1.ConnectRequest
	(*TunnelRequest)(nil),    // 2: tunnel.v1.TunnelRequest
	(*TunnelResponse)(nil),   // 3: tunnel.v1.TunnelResponse
	(*ControlMessage)(nil),   // 4: tunnel.v1.ControlMessage
	(*TunnelData)(nil),       // 5: tunnel.v1.TunnelData
	(*TunnelClose)(nil),      // 6: tunnel.v1.TunnelClose
	(*Envelope)(nil),         // 7: tunnel.v1.Envelope
	nil,                      // 8: tunnel.v1.TunnelRequest.HeadersEntry
	nil,                      // 9: tunnel.v1.TunnelResponse.HeadersEntry
}
var file_proto_tunnel_proto_depIdxs = []int32{
	8, // 0: tunnel.v1.TunnelRequest.headers:type_name -> tunnel.v1.TunnelRequest.HeadersEntry
	9, // 1: tunnel.v1.TunnelResponse.headers:type_name -> tunnel.v1.TunnelResponse.HeadersEntry
	0, // 2: tunnel.v1.ControlMessage.type:type_name -> tunnel.v1.ControlMessage.Type
	1, // 3: tunnel.v1.Envelope.connect:type_name -> tunnel.v1.ConnectRequest
	2, // 4: tunnel.v1.Envelope.request:type_name -> tunnel.v1.TunnelRequest
	3, // 5: tunnel.v1.Envelope.response:type_name -> tunnel.v1.TunnelResponse
	4, // 6: tunnel.v1.Envelope.control:type_name -> tunnel.v1.ControlMessage
	5, // 7: tunnel.v1.Envelope.data:type_name -> tunnel.v1.TunnelData
	6, // 8: tunnel.v1.Envelope.close:type_name -> tunnel.v1.TunnelClose
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_proto_tunnel_proto_init() }
func file_proto_tunnel_proto_init() {
	if File_proto_tunnel_proto != nil {
		return
	}
	file_proto_tunnel_proto_msgTypes[6].OneofWrappers = []any{
		(*Envelope_Connect)(nil),
		(*Envelope_Request)(nil),
		(*Envelope_Response)(nil),
		(*Envelope_Control)(nil),
		(*Envelope_Data)(nil),
		(*Envelope_Close)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_tunnel_proto_rawDesc), len(file_proto_tunnel_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_tunnel_proto_goTypes,
		DependencyIndexes: file_proto_tunnel_proto_depIdxs,
		EnumInfos:         file_proto_tunnel_proto_enumTypes,
		MessageInfos:      file_proto_tunnel_proto_msgTypes,
	}.Build()
	File_proto_tunnel_proto = out.File
	file_proto_tunnel_proto_goTypes = nil
	file_proto_tunnel_proto_depIdxs = nil
}
