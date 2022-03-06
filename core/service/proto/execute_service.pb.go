// Code generated by protoc-gen-go. DO NOT EDIT.
// source: execute_service.proto

package proto // import "./"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetJobRequest struct {
	// * 任务名称
	JobName string `protobuf:"bytes,1,opt,name=JobName,proto3" json:"JobName"`
	// * 任务不存在时是否创建
	Create               bool     `protobuf:"varint,2,opt,name=Create,proto3" json:"Create"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobRequest) Reset()         { *m = GetJobRequest{} }
func (m *GetJobRequest) String() string { return proto.CompactTextString(m) }
func (*GetJobRequest) ProtoMessage()    {}
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_execute_service_0bf2a36d18009366, []int{0}
}
func (m *GetJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobRequest.Unmarshal(m, b)
}
func (m *GetJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobRequest.Marshal(b, m, deterministic)
}
func (dst *GetJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobRequest.Merge(dst, src)
}
func (m *GetJobRequest) XXX_Size() int {
	return xxx_messageInfo_GetJobRequest.Size(m)
}
func (m *GetJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobRequest proto.InternalMessageInfo

func (m *GetJobRequest) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *GetJobRequest) GetCreate() bool {
	if m != nil {
		return m.Create
	}
	return false
}

// * UpdateCursorRequest
type UpdateCursorRequest struct {
	// * 任务名称
	JobName string `protobuf:"bytes,1,opt,name=JobName,proto3" json:"JobName"`
	// * 记录编号
	RecordId             int64    `protobuf:"varint,2,opt,name=RecordId,proto3" json:"RecordId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCursorRequest) Reset()         { *m = UpdateCursorRequest{} }
func (m *UpdateCursorRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCursorRequest) ProtoMessage()    {}
func (*UpdateCursorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_execute_service_0bf2a36d18009366, []int{1}
}
func (m *UpdateCursorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCursorRequest.Unmarshal(m, b)
}
func (m *UpdateCursorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCursorRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateCursorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCursorRequest.Merge(dst, src)
}
func (m *UpdateCursorRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCursorRequest.Size(m)
}
func (m *UpdateCursorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCursorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCursorRequest proto.InternalMessageInfo

func (m *UpdateCursorRequest) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *UpdateCursorRequest) GetRecordId() int64 {
	if m != nil {
		return m.RecordId
	}
	return 0
}

type AddFailRequest struct {
	// * 任务名称
	JobName string `protobuf:"bytes,1,opt,name=JobName,proto3" json:"JobName"`
	// * 记录编号
	RecordId             int64    `protobuf:"varint,2,opt,name=RecordId,proto3" json:"RecordId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFailRequest) Reset()         { *m = AddFailRequest{} }
func (m *AddFailRequest) String() string { return proto.CompactTextString(m) }
func (*AddFailRequest) ProtoMessage()    {}
func (*AddFailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_execute_service_0bf2a36d18009366, []int{2}
}
func (m *AddFailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFailRequest.Unmarshal(m, b)
}
func (m *AddFailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFailRequest.Marshal(b, m, deterministic)
}
func (dst *AddFailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFailRequest.Merge(dst, src)
}
func (m *AddFailRequest) XXX_Size() int {
	return xxx_messageInfo_AddFailRequest.Size(m)
}
func (m *AddFailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddFailRequest proto.InternalMessageInfo

func (m *AddFailRequest) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *AddFailRequest) GetRecordId() int64 {
	if m != nil {
		return m.RecordId
	}
	return 0
}

// * JobExecData
type SExecData struct {
	// * 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	// * 任务名称
	JobName string `protobuf:"bytes,2,opt,name=JobName,proto3" json:"JobName"`
	// * 上次执行位置索引
	LastExecIndex int64 `protobuf:"varint,3,opt,name=LastExecIndex,proto3" json:"LastExecIndex"`
	// * 最后执行时间
	LastExecTime         int64    `protobuf:"varint,4,opt,name=LastExecTime,proto3" json:"LastExecTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SExecData) Reset()         { *m = SExecData{} }
func (m *SExecData) String() string { return proto.CompactTextString(m) }
func (*SExecData) ProtoMessage()    {}
func (*SExecData) Descriptor() ([]byte, []int) {
	return fileDescriptor_execute_service_0bf2a36d18009366, []int{3}
}
func (m *SExecData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SExecData.Unmarshal(m, b)
}
func (m *SExecData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SExecData.Marshal(b, m, deterministic)
}
func (dst *SExecData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SExecData.Merge(dst, src)
}
func (m *SExecData) XXX_Size() int {
	return xxx_messageInfo_SExecData.Size(m)
}
func (m *SExecData) XXX_DiscardUnknown() {
	xxx_messageInfo_SExecData.DiscardUnknown(m)
}

var xxx_messageInfo_SExecData proto.InternalMessageInfo

func (m *SExecData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SExecData) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *SExecData) GetLastExecIndex() int64 {
	if m != nil {
		return m.LastExecIndex
	}
	return 0
}

func (m *SExecData) GetLastExecTime() int64 {
	if m != nil {
		return m.LastExecTime
	}
	return 0
}

func init() {
	proto.RegisterType((*GetJobRequest)(nil), "GetJobRequest")
	proto.RegisterType((*UpdateCursorRequest)(nil), "UpdateCursorRequest")
	proto.RegisterType((*AddFailRequest)(nil), "AddFailRequest")
	proto.RegisterType((*SExecData)(nil), "SExecData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecuteServiceClient is the client API for ExecuteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecuteServiceClient interface {
	// * 获取JobExecData
	GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*SExecData, error)
	// * 更新任务执行游标
	UpdateExecCursor(ctx context.Context, in *UpdateCursorRequest, opts ...grpc.CallOption) (*Result, error)
	// * 记录失败
	AddFail(ctx context.Context, in *AddFailRequest, opts ...grpc.CallOption) (*Result, error)
}

type executeServiceClient struct {
	cc *grpc.ClientConn
}

func NewExecuteServiceClient(cc *grpc.ClientConn) ExecuteServiceClient {
	return &executeServiceClient{cc}
}

func (c *executeServiceClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*SExecData, error) {
	out := new(SExecData)
	err := c.cc.Invoke(ctx, "/ExecuteService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeServiceClient) UpdateExecCursor(ctx context.Context, in *UpdateCursorRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExecuteService/UpdateExecCursor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executeServiceClient) AddFail(ctx context.Context, in *AddFailRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExecuteService/AddFail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecuteServiceServer is the server API for ExecuteService service.
type ExecuteServiceServer interface {
	// * 获取JobExecData
	GetJob(context.Context, *GetJobRequest) (*SExecData, error)
	// * 更新任务执行游标
	UpdateExecCursor(context.Context, *UpdateCursorRequest) (*Result, error)
	// * 记录失败
	AddFail(context.Context, *AddFailRequest) (*Result, error)
}

func RegisterExecuteServiceServer(s *grpc.Server, srv ExecuteServiceServer) {
	s.RegisterService(&_ExecuteService_serviceDesc, srv)
}

func _ExecuteService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExecuteService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteServiceServer).GetJob(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteService_UpdateExecCursor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCursorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteServiceServer).UpdateExecCursor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExecuteService/UpdateExecCursor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteServiceServer).UpdateExecCursor(ctx, req.(*UpdateCursorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecuteService_AddFail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteServiceServer).AddFail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExecuteService/AddFail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteServiceServer).AddFail(ctx, req.(*AddFailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecuteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExecuteService",
	HandlerType: (*ExecuteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJob",
			Handler:    _ExecuteService_GetJob_Handler,
		},
		{
			MethodName: "UpdateExecCursor",
			Handler:    _ExecuteService_UpdateExecCursor_Handler,
		},
		{
			MethodName: "AddFail",
			Handler:    _ExecuteService_AddFail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "execute_service.proto",
}

func init() {
	proto.RegisterFile("execute_service.proto", fileDescriptor_execute_service_0bf2a36d18009366)
}

var fileDescriptor_execute_service_0bf2a36d18009366 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0xd7, 0xee, 0xc7, 0xfe, 0x3c, 0x6c, 0xfd, 0x49, 0xfc, 0x43, 0xe9, 0x69, 0x94, 0x29,
	0x3b, 0x45, 0xd8, 0x8e, 0x9e, 0xe6, 0xdc, 0xa4, 0x53, 0x3c, 0x64, 0x7a, 0xf1, 0x22, 0x69, 0xf3,
	0x30, 0x0b, 0x9d, 0x99, 0x69, 0x2a, 0xbd, 0xf9, 0x26, 0x7c, 0xc1, 0xb2, 0x64, 0x1b, 0x2b, 0x08,
	0x1e, 0x3c, 0x85, 0xef, 0x97, 0xe7, 0xf9, 0x10, 0x3e, 0x09, 0x9c, 0x62, 0x89, 0x49, 0xa1, 0xf1,
	0x25, 0x47, 0xf5, 0x91, 0x26, 0x48, 0xd7, 0x4a, 0x6a, 0x19, 0x74, 0x96, 0x99, 0x8c, 0x79, 0x66,
	0x53, 0x38, 0x86, 0xee, 0x2d, 0xea, 0xb9, 0x8c, 0x19, 0xbe, 0x17, 0x98, 0x6b, 0xe2, 0x43, 0x73,
	0x2e, 0xe3, 0x07, 0xbe, 0x42, 0xdf, 0xe9, 0x39, 0x83, 0x36, 0xdb, 0x45, 0x72, 0x06, 0x8d, 0x89,
	0x42, 0xae, 0xd1, 0x77, 0x7b, 0xce, 0xa0, 0xc5, 0xb6, 0x29, 0xbc, 0x83, 0xe3, 0xa7, 0xb5, 0xe0,
	0x1a, 0x27, 0x85, 0xca, 0xa5, 0xfa, 0x1d, 0x14, 0x40, 0x8b, 0x61, 0x22, 0x95, 0x88, 0x84, 0x41,
	0xd5, 0xd9, 0x3e, 0x87, 0x33, 0xf0, 0xc6, 0x42, 0xcc, 0x78, 0x9a, 0xfd, 0x8d, 0xf3, 0x09, 0xed,
	0xc5, 0xb4, 0xc4, 0xe4, 0x86, 0x6b, 0x4e, 0x3c, 0x70, 0x23, 0x61, 0xb6, 0xeb, 0xcc, 0x8d, 0xc4,
	0x21, 0xd2, 0xad, 0x22, 0xfb, 0xd0, 0xbd, 0xe7, 0xb9, 0xde, 0x6c, 0x46, 0x6f, 0x02, 0x4b, 0xbf,
	0x6e, 0x96, 0xaa, 0x25, 0x09, 0xa1, 0xb3, 0x2b, 0x1e, 0xd3, 0x15, 0xfa, 0xff, 0xcc, 0x50, 0xa5,
	0x1b, 0x7e, 0x39, 0xe0, 0x4d, 0xed, 0x03, 0x2c, 0xac, 0x7f, 0x72, 0x01, 0x0d, 0xeb, 0x9a, 0x78,
	0xb4, 0x22, 0x3d, 0x00, 0xba, 0xbf, 0x6c, 0x58, 0x23, 0x23, 0x38, 0xb2, 0x42, 0x37, 0x9d, 0x95,
	0x4a, 0x4e, 0xe8, 0x0f, 0x8e, 0x83, 0x26, 0x65, 0x98, 0x17, 0x99, 0x0e, 0x6b, 0xe4, 0x1c, 0x9a,
	0x5b, 0x71, 0xe4, 0x3f, 0xad, 0x2a, 0x3c, 0x18, 0xbb, 0xee, 0x83, 0x9f, 0xc8, 0x15, 0x5d, 0xa6,
	0xfa, 0xb5, 0x88, 0x69, 0x5a, 0x2a, 0xa4, 0x4b, 0x39, 0x94, 0x54, 0xad, 0x93, 0xe7, 0x16, 0xbd,
	0xbc, 0x32, 0xbf, 0x22, 0x6e, 0x98, 0x63, 0xf4, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x23, 0xdd,
	0xdd, 0x43, 0x02, 0x00, 0x00,
}
