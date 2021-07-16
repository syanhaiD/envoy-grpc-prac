// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: verify.proto

package verify

import (
	context "context"
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

type BiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Free1 string   `protobuf:"bytes,1,opt,name=free1,proto3" json:"free1,omitempty"`
	P1    *Packet1 `protobuf:"bytes,2,opt,name=p1,proto3" json:"p1,omitempty"`
	P2    *Packet2 `protobuf:"bytes,3,opt,name=p2,proto3" json:"p2,omitempty"`
}

func (x *BiRequest) Reset() {
	*x = BiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiRequest) ProtoMessage() {}

func (x *BiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiRequest.ProtoReflect.Descriptor instead.
func (*BiRequest) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{0}
}

func (x *BiRequest) GetFree1() string {
	if x != nil {
		return x.Free1
	}
	return ""
}

func (x *BiRequest) GetP1() *Packet1 {
	if x != nil {
		return x.P1
	}
	return nil
}

func (x *BiRequest) GetP2() *Packet2 {
	if x != nil {
		return x.P2
	}
	return nil
}

type Packet1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Free1 float32 `protobuf:"fixed32,1,opt,name=free1,proto3" json:"free1,omitempty"`
	Free2 float32 `protobuf:"fixed32,2,opt,name=free2,proto3" json:"free2,omitempty"`
	Free3 float32 `protobuf:"fixed32,3,opt,name=free3,proto3" json:"free3,omitempty"`
}

func (x *Packet1) Reset() {
	*x = Packet1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet1) ProtoMessage() {}

func (x *Packet1) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet1.ProtoReflect.Descriptor instead.
func (*Packet1) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{1}
}

func (x *Packet1) GetFree1() float32 {
	if x != nil {
		return x.Free1
	}
	return 0
}

func (x *Packet1) GetFree2() float32 {
	if x != nil {
		return x.Free2
	}
	return 0
}

func (x *Packet1) GetFree3() float32 {
	if x != nil {
		return x.Free3
	}
	return 0
}

type Packet2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Free1 string `protobuf:"bytes,1,opt,name=free1,proto3" json:"free1,omitempty"`
	Free2 int64  `protobuf:"varint,2,opt,name=free2,proto3" json:"free2,omitempty"`
}

func (x *Packet2) Reset() {
	*x = Packet2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet2) ProtoMessage() {}

func (x *Packet2) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet2.ProtoReflect.Descriptor instead.
func (*Packet2) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{2}
}

func (x *Packet2) GetFree1() string {
	if x != nil {
		return x.Free1
	}
	return ""
}

func (x *Packet2) GetFree2() int64 {
	if x != nil {
		return x.Free2
	}
	return 0
}

type BiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Free1 string   `protobuf:"bytes,1,opt,name=free1,proto3" json:"free1,omitempty"`
	Free2 int32    `protobuf:"varint,2,opt,name=free2,proto3" json:"free2,omitempty"`
	P1    *Packet1 `protobuf:"bytes,3,opt,name=p1,proto3" json:"p1,omitempty"`
	P2    *Packet2 `protobuf:"bytes,4,opt,name=p2,proto3" json:"p2,omitempty"`
}

func (x *BiResponse) Reset() {
	*x = BiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_verify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiResponse) ProtoMessage() {}

func (x *BiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_verify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiResponse.ProtoReflect.Descriptor instead.
func (*BiResponse) Descriptor() ([]byte, []int) {
	return file_verify_proto_rawDescGZIP(), []int{3}
}

func (x *BiResponse) GetFree1() string {
	if x != nil {
		return x.Free1
	}
	return ""
}

func (x *BiResponse) GetFree2() int32 {
	if x != nil {
		return x.Free2
	}
	return 0
}

func (x *BiResponse) GetP1() *Packet1 {
	if x != nil {
		return x.P1
	}
	return nil
}

func (x *BiResponse) GetP2() *Packet2 {
	if x != nil {
		return x.P2
	}
	return nil
}

var File_verify_proto protoreflect.FileDescriptor

var file_verify_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0x63, 0x0a, 0x09, 0x42, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x65, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x72, 0x65, 0x65, 0x31, 0x12, 0x1f, 0x0a, 0x02, 0x70, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x31, 0x52, 0x02, 0x70, 0x31, 0x12, 0x1f, 0x0a, 0x02, 0x70, 0x32,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x32, 0x52, 0x02, 0x70, 0x32, 0x22, 0x4b, 0x0a, 0x07, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x65, 0x65, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x72, 0x65, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x72, 0x65, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x72, 0x65,
	0x65, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x65, 0x65, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x66, 0x72, 0x65, 0x65, 0x33, 0x22, 0x35, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x65, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x72, 0x65, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x65,
	0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x72, 0x65, 0x65, 0x32, 0x22,
	0x7a, 0x0a, 0x0a, 0x42, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x72, 0x65, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x72,
	0x65, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x65, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x66, 0x72, 0x65, 0x65, 0x32, 0x12, 0x1f, 0x0a, 0x02, 0x70, 0x31, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x31, 0x52, 0x02, 0x70, 0x31, 0x12, 0x1f, 0x0a, 0x02, 0x70, 0x32,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x32, 0x52, 0x02, 0x70, 0x32, 0x32, 0x46, 0x0a, 0x06, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x3c, 0x0a, 0x0d, 0x42, 0x69, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x11, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e,
	0x42, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x2e, 0x42, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_verify_proto_rawDescOnce sync.Once
	file_verify_proto_rawDescData = file_verify_proto_rawDesc
)

func file_verify_proto_rawDescGZIP() []byte {
	file_verify_proto_rawDescOnce.Do(func() {
		file_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_verify_proto_rawDescData)
	})
	return file_verify_proto_rawDescData
}

var file_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_verify_proto_goTypes = []interface{}{
	(*BiRequest)(nil),  // 0: verify.BiRequest
	(*Packet1)(nil),    // 1: verify.Packet1
	(*Packet2)(nil),    // 2: verify.Packet2
	(*BiResponse)(nil), // 3: verify.BiResponse
}
var file_verify_proto_depIdxs = []int32{
	1, // 0: verify.BiRequest.p1:type_name -> verify.Packet1
	2, // 1: verify.BiRequest.p2:type_name -> verify.Packet2
	1, // 2: verify.BiResponse.p1:type_name -> verify.Packet1
	2, // 3: verify.BiResponse.p2:type_name -> verify.Packet2
	0, // 4: verify.Verify.BiDirectional:input_type -> verify.BiRequest
	3, // 5: verify.Verify.BiDirectional:output_type -> verify.BiResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_verify_proto_init() }
func file_verify_proto_init() {
	if File_verify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_verify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BiRequest); i {
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
		file_verify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet1); i {
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
		file_verify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet2); i {
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
		file_verify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BiResponse); i {
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
			RawDescriptor: file_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_verify_proto_goTypes,
		DependencyIndexes: file_verify_proto_depIdxs,
		MessageInfos:      file_verify_proto_msgTypes,
	}.Build()
	File_verify_proto = out.File
	file_verify_proto_rawDesc = nil
	file_verify_proto_goTypes = nil
	file_verify_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VerifyClient is the client API for Verify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerifyClient interface {
	BiDirectional(ctx context.Context, opts ...grpc.CallOption) (Verify_BiDirectionalClient, error)
}

type verifyClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifyClient(cc grpc.ClientConnInterface) VerifyClient {
	return &verifyClient{cc}
}

func (c *verifyClient) BiDirectional(ctx context.Context, opts ...grpc.CallOption) (Verify_BiDirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Verify_serviceDesc.Streams[0], "/verify.Verify/BiDirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &verifyBiDirectionalClient{stream}
	return x, nil
}

type Verify_BiDirectionalClient interface {
	Send(*BiRequest) error
	Recv() (*BiResponse, error)
	grpc.ClientStream
}

type verifyBiDirectionalClient struct {
	grpc.ClientStream
}

func (x *verifyBiDirectionalClient) Send(m *BiRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *verifyBiDirectionalClient) Recv() (*BiResponse, error) {
	m := new(BiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VerifyServer is the server API for Verify service.
type VerifyServer interface {
	BiDirectional(Verify_BiDirectionalServer) error
}

// UnimplementedVerifyServer can be embedded to have forward compatible implementations.
type UnimplementedVerifyServer struct {
}

func (*UnimplementedVerifyServer) BiDirectional(Verify_BiDirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectional not implemented")
}

func RegisterVerifyServer(s *grpc.Server, srv VerifyServer) {
	s.RegisterService(&_Verify_serviceDesc, srv)
}

func _Verify_BiDirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VerifyServer).BiDirectional(&verifyBiDirectionalServer{stream})
}

type Verify_BiDirectionalServer interface {
	Send(*BiResponse) error
	Recv() (*BiRequest, error)
	grpc.ServerStream
}

type verifyBiDirectionalServer struct {
	grpc.ServerStream
}

func (x *verifyBiDirectionalServer) Send(m *BiResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *verifyBiDirectionalServer) Recv() (*BiRequest, error) {
	m := new(BiRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Verify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "verify.Verify",
	HandlerType: (*VerifyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BiDirectional",
			Handler:       _Verify_BiDirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "verify.proto",
}