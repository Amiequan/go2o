// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_service.proto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	// 获取通知项,key
	GetNotifyItem(ctx context.Context, in *String, opts ...grpc.CallOption) (*SNotifyItem, error)
	// 发送短信
	SendPhoneMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Result, error)
	// 获取所有通知项
	GetAllNotifyItem(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NotifyItemListResponse, error)
	// 保存通知项设置
	SaveNotifyItem(ctx context.Context, in *SNotifyItem, opts ...grpc.CallOption) (*Result, error)
	// 获取邮件模版
	GetMailTemplate(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SMailTemplate, error)
	// 保存邮件模板
	SaveMailTemplate(ctx context.Context, in *SMailTemplate, opts ...grpc.CallOption) (*Result, error)
	// 获取邮件模板
	GetMailTemplates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MailTemplateListResponse, error)
	// 删除邮件模板
	DeleteMailTemplate(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 获取邮件绑定
	// rpc GetConfig() mss.Config
	// 保存邮件
	// rpc SaveConfig(conf *mss.Config) error
	// 发送站内信
	SendSiteMessage(ctx context.Context, in *SendSiteMessageRequest, opts ...grpc.CallOption) (*Result, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) GetNotifyItem(ctx context.Context, in *String, opts ...grpc.CallOption) (*SNotifyItem, error) {
	out := new(SNotifyItem)
	err := c.cc.Invoke(ctx, "/MessageService/GetNotifyItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendPhoneMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/SendPhoneMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetAllNotifyItem(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NotifyItemListResponse, error) {
	out := new(NotifyItemListResponse)
	err := c.cc.Invoke(ctx, "/MessageService/GetAllNotifyItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SaveNotifyItem(ctx context.Context, in *SNotifyItem, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/SaveNotifyItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMailTemplate(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SMailTemplate, error) {
	out := new(SMailTemplate)
	err := c.cc.Invoke(ctx, "/MessageService/GetMailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SaveMailTemplate(ctx context.Context, in *SMailTemplate, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/SaveMailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMailTemplates(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MailTemplateListResponse, error) {
	out := new(MailTemplateListResponse)
	err := c.cc.Invoke(ctx, "/MessageService/GetMailTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) DeleteMailTemplate(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/DeleteMailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendSiteMessage(ctx context.Context, in *SendSiteMessageRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/MessageService/SendSiteMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	// 获取通知项,key
	GetNotifyItem(context.Context, *String) (*SNotifyItem, error)
	// 发送短信
	SendPhoneMessage(context.Context, *SendMessageRequest) (*Result, error)
	// 获取所有通知项
	GetAllNotifyItem(context.Context, *Empty) (*NotifyItemListResponse, error)
	// 保存通知项设置
	SaveNotifyItem(context.Context, *SNotifyItem) (*Result, error)
	// 获取邮件模版
	GetMailTemplate(context.Context, *Int64) (*SMailTemplate, error)
	// 保存邮件模板
	SaveMailTemplate(context.Context, *SMailTemplate) (*Result, error)
	// 获取邮件模板
	GetMailTemplates(context.Context, *Empty) (*MailTemplateListResponse, error)
	// 删除邮件模板
	DeleteMailTemplate(context.Context, *Int64) (*Result, error)
	// 获取邮件绑定
	// rpc GetConfig() mss.Config
	// 保存邮件
	// rpc SaveConfig(conf *mss.Config) error
	// 发送站内信
	SendSiteMessage(context.Context, *SendSiteMessageRequest) (*Result, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_GetNotifyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetNotifyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetNotifyItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetNotifyItem(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendPhoneMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendPhoneMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SendPhoneMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendPhoneMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetAllNotifyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetAllNotifyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetAllNotifyItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetAllNotifyItem(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SaveNotifyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SNotifyItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SaveNotifyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SaveNotifyItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SaveNotifyItem(ctx, req.(*SNotifyItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetMailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMailTemplate(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SaveMailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMailTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SaveMailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SaveMailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SaveMailTemplate(ctx, req.(*SMailTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMailTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMailTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/GetMailTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMailTemplates(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_DeleteMailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).DeleteMailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/DeleteMailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).DeleteMailTemplate(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendSiteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSiteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendSiteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageService/SendSiteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendSiteMessage(ctx, req.(*SendSiteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotifyItem",
			Handler:    _MessageService_GetNotifyItem_Handler,
		},
		{
			MethodName: "SendPhoneMessage",
			Handler:    _MessageService_SendPhoneMessage_Handler,
		},
		{
			MethodName: "GetAllNotifyItem",
			Handler:    _MessageService_GetAllNotifyItem_Handler,
		},
		{
			MethodName: "SaveNotifyItem",
			Handler:    _MessageService_SaveNotifyItem_Handler,
		},
		{
			MethodName: "GetMailTemplate",
			Handler:    _MessageService_GetMailTemplate_Handler,
		},
		{
			MethodName: "SaveMailTemplate",
			Handler:    _MessageService_SaveMailTemplate_Handler,
		},
		{
			MethodName: "GetMailTemplates",
			Handler:    _MessageService_GetMailTemplates_Handler,
		},
		{
			MethodName: "DeleteMailTemplate",
			Handler:    _MessageService_DeleteMailTemplate_Handler,
		},
		{
			MethodName: "SendSiteMessage",
			Handler:    _MessageService_SendSiteMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_service.proto",
}

func init() {
	proto.RegisterFile("message_service.proto", fileDescriptor_message_service_4b57e4408139eb5d)
}

var fileDescriptor_message_service_4b57e4408139eb5d = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0x03, 0x2f, 0xb4, 0x2f, 0x43, 0x6d, 0xcb, 0x14, 0x29, 0xcd, 0x46, 0xc8, 0x46, 0x45,
	0xbc, 0x85, 0xfa, 0xb5, 0x70, 0xa5, 0x28, 0xa5, 0x60, 0x45, 0x3a, 0xae, 0xdc, 0x48, 0x9a, 0x5e,
	0xd3, 0x81, 0x49, 0x26, 0x66, 0x6e, 0x0a, 0xfd, 0x81, 0xfe, 0x2f, 0xc9, 0x87, 0xed, 0x24, 0xba,
	0x0a, 0xf7, 0xb9, 0x27, 0xe7, 0x1e, 0xce, 0xb0, 0xc3, 0x08, 0x8d, 0xf1, 0x43, 0x7c, 0x37, 0x98,
	0x6e, 0x64, 0x80, 0x90, 0xa4, 0x9a, 0xb4, 0xdb, 0x09, 0x95, 0x5e, 0xfa, 0xaa, 0x9a, 0x46, 0x95,
	0x68, 0xfc, 0x23, 0x5e, 0x91, 0x2e, 0x57, 0x93, 0xaf, 0x7f, 0xac, 0x3b, 0x2f, 0xa9, 0x28, 0x1d,
	0xf8, 0x09, 0x3b, 0x98, 0x22, 0x3d, 0x6b, 0x92, 0x1f, 0xdb, 0x19, 0x61, 0xc4, 0xdb, 0x20, 0x28,
	0x95, 0x71, 0xe8, 0x76, 0x40, 0xec, 0xb1, 0xe7, 0xf0, 0x09, 0xeb, 0x0b, 0x8c, 0x57, 0x2f, 0x6b,
	0x1d, 0x63, 0x65, 0xc2, 0x07, 0x90, 0xa3, 0x6a, 0x5a, 0xe0, 0x67, 0x86, 0x86, 0xdc, 0x36, 0x2c,
	0xd0, 0x64, 0x8a, 0x3c, 0x87, 0x5f, 0xb1, 0xfe, 0x14, 0xe9, 0x4e, 0x29, 0xeb, 0x40, 0x0b, 0x1e,
	0xa3, 0x84, 0xb6, 0xee, 0x10, 0xf6, 0xf0, 0x49, 0x1a, 0x5a, 0xa0, 0x49, 0x74, 0x6c, 0xd0, 0x73,
	0xf8, 0x29, 0xeb, 0x0a, 0x7f, 0x83, 0xd6, 0x4f, 0xb5, 0x30, 0xf6, 0x85, 0x33, 0xd6, 0x9b, 0x22,
	0xcd, 0x7d, 0xa9, 0x5e, 0x31, 0x4a, 0x94, 0x4f, 0xc8, 0x5b, 0x30, 0x8b, 0xe9, 0xfa, 0xd2, 0xed,
	0x82, 0xb0, 0xb9, 0xe7, 0xf0, 0x73, 0xd6, 0xcf, 0x7d, 0x6b, 0xea, 0x86, 0xca, 0xf6, 0xbe, 0x29,
	0xd2, 0xdb, 0x5b, 0xb3, 0x4b, 0x3f, 0x02, 0x9b, 0x37, 0xf2, 0x1f, 0x33, 0xfe, 0x80, 0x0a, 0x09,
	0xff, 0xcc, 0x55, 0xeb, 0xa7, 0x97, 0x17, 0x28, 0x24, 0xed, 0x2a, 0x1d, 0x42, 0x83, 0xfc, 0xae,
	0xf5, 0xfe, 0x88, 0x0d, 0x02, 0x1d, 0x41, 0x28, 0x69, 0x9d, 0x2d, 0x21, 0xd4, 0x13, 0x0d, 0x69,
	0x12, 0xbc, 0xfd, 0x87, 0xf1, 0x6d, 0xf1, 0xd0, 0xcb, 0x56, 0xf1, 0xb9, 0xf8, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x27, 0x5f, 0xf8, 0xfd, 0x31, 0x02, 0x00, 0x00,
}
