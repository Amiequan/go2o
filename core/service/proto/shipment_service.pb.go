// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shipment_service.proto

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

// ShipmentServiceClient is the client API for ShipmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShipmentServiceClient interface {
	// 创建一个配送覆盖的区域
	CreateCoverageArea_(ctx context.Context, in *SCoverageValue, opts ...grpc.CallOption) (*Result, error)
	// 获取订单的发货单信息
	GetShipOrderOfOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*ShipmentOrderListResponse, error)
	// * 物流追踪
	GetLogisticFlowTrack(ctx context.Context, in *LogisticFlowTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error)
	// * 获取发货单的物流追踪信息,$shipOrderId:发货单编号 $invert:是否倒序排列
	ShipOrderLogisticTrack(ctx context.Context, in *OrderLogisticTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error)
}

type shipmentServiceClient struct {
	cc *grpc.ClientConn
}

func NewShipmentServiceClient(cc *grpc.ClientConn) ShipmentServiceClient {
	return &shipmentServiceClient{cc}
}

func (c *shipmentServiceClient) CreateCoverageArea_(ctx context.Context, in *SCoverageValue, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShipmentService/CreateCoverageArea_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetShipOrderOfOrder(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*ShipmentOrderListResponse, error) {
	out := new(ShipmentOrderListResponse)
	err := c.cc.Invoke(ctx, "/ShipmentService/GetShipOrderOfOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) GetLogisticFlowTrack(ctx context.Context, in *LogisticFlowTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error) {
	out := new(SShipOrderTrack)
	err := c.cc.Invoke(ctx, "/ShipmentService/GetLogisticFlowTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shipmentServiceClient) ShipOrderLogisticTrack(ctx context.Context, in *OrderLogisticTrackRequest, opts ...grpc.CallOption) (*SShipOrderTrack, error) {
	out := new(SShipOrderTrack)
	err := c.cc.Invoke(ctx, "/ShipmentService/ShipOrderLogisticTrack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipmentServiceServer is the server API for ShipmentService service.
type ShipmentServiceServer interface {
	// 创建一个配送覆盖的区域
	CreateCoverageArea_(context.Context, *SCoverageValue) (*Result, error)
	// 获取订单的发货单信息
	GetShipOrderOfOrder(context.Context, *OrderId) (*ShipmentOrderListResponse, error)
	// * 物流追踪
	GetLogisticFlowTrack(context.Context, *LogisticFlowTrackRequest) (*SShipOrderTrack, error)
	// * 获取发货单的物流追踪信息,$shipOrderId:发货单编号 $invert:是否倒序排列
	ShipOrderLogisticTrack(context.Context, *OrderLogisticTrackRequest) (*SShipOrderTrack, error)
}

func RegisterShipmentServiceServer(s *grpc.Server, srv ShipmentServiceServer) {
	s.RegisterService(&_ShipmentService_serviceDesc, srv)
}

func _ShipmentService_CreateCoverageArea__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SCoverageValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).CreateCoverageArea_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/CreateCoverageArea_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).CreateCoverageArea_(ctx, req.(*SCoverageValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetShipOrderOfOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetShipOrderOfOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/GetShipOrderOfOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetShipOrderOfOrder(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_GetLogisticFlowTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogisticFlowTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).GetLogisticFlowTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/GetLogisticFlowTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).GetLogisticFlowTrack(ctx, req.(*LogisticFlowTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShipmentService_ShipOrderLogisticTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderLogisticTrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipmentServiceServer).ShipOrderLogisticTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShipmentService/ShipOrderLogisticTrack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipmentServiceServer).ShipOrderLogisticTrack(ctx, req.(*OrderLogisticTrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShipmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShipmentService",
	HandlerType: (*ShipmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoverageArea_",
			Handler:    _ShipmentService_CreateCoverageArea__Handler,
		},
		{
			MethodName: "GetShipOrderOfOrder",
			Handler:    _ShipmentService_GetShipOrderOfOrder_Handler,
		},
		{
			MethodName: "GetLogisticFlowTrack",
			Handler:    _ShipmentService_GetLogisticFlowTrack_Handler,
		},
		{
			MethodName: "ShipOrderLogisticTrack",
			Handler:    _ShipmentService_ShipOrderLogisticTrack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipment_service.proto",
}

func init() {
	proto.RegisterFile("shipment_service.proto", fileDescriptor_shipment_service_6847cf98d37302ee)
}

var fileDescriptor_shipment_service_6847cf98d37302ee = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xcb, 0xbf, 0xf8, 0x2d, 0x83, 0x50, 0x99, 0x48, 0xc5, 0xd9, 0x08, 0x79, 0x80, 0x5b,
	0xac, 0x4b, 0x71, 0xa1, 0x45, 0xab, 0x50, 0x28, 0x24, 0xe2, 0xc2, 0x4d, 0x98, 0x24, 0xd7, 0xe9,
	0x60, 0xd2, 0x1b, 0x67, 0x6e, 0xea, 0x8b, 0xf8, 0xc0, 0xd2, 0x69, 0x92, 0x4d, 0xc1, 0xd5, 0x61,
	0xbe, 0xc3, 0x7c, 0x07, 0xae, 0x98, 0xfa, 0x8d, 0x6d, 0x6a, 0xdc, 0x72, 0xe6, 0xd1, 0xed, 0x6c,
	0x81, 0xd0, 0x38, 0x62, 0x52, 0xa7, 0xa6, 0xa2, 0x5c, 0x57, 0xdd, 0xeb, 0xa2, 0x46, 0xef, 0xb5,
	0xc1, 0x19, 0xb9, 0x12, 0x5d, 0x56, 0x32, 0x75, 0x85, 0xea, 0x8b, 0x41, 0x33, 0x74, 0xf3, 0x9f,
	0x7f, 0x62, 0x92, 0x76, 0x38, 0x3d, 0xc8, 0xe5, 0xb5, 0x88, 0x16, 0x0e, 0x35, 0xe3, 0x82, 0x76,
	0xe8, 0xb4, 0xc1, 0x7b, 0x87, 0x3a, 0x93, 0x13, 0x48, 0x7b, 0xf0, 0xa6, 0xab, 0x16, 0xd5, 0x09,
	0x24, 0xe8, 0xdb, 0x8a, 0xe3, 0x91, 0xbc, 0x13, 0xd1, 0x12, 0x79, 0x2f, 0x5a, 0xef, 0xc7, 0xd7,
	0x1f, 0x21, 0xe4, 0x18, 0x42, 0xbe, 0x94, 0x4a, 0x41, 0xbf, 0x12, 0xc8, 0xca, 0x7a, 0x4e, 0xd0,
	0x37, 0xb4, 0xf5, 0x18, 0x8f, 0xe4, 0xa3, 0x38, 0x5f, 0x22, 0xaf, 0xc8, 0x58, 0xcf, 0xb6, 0x78,
	0xaa, 0xe8, 0xfb, 0xd5, 0xe9, 0xe2, 0x53, 0x5e, 0xc2, 0x11, 0x4b, 0xf0, 0xab, 0x45, 0xcf, 0xea,
	0x0c, 0xd2, 0x61, 0x2e, 0x14, 0xf1, 0x48, 0x3e, 0x8b, 0xe9, 0xc0, 0xfa, 0x8f, 0x07, 0x91, 0x82,
	0x63, 0xf8, 0x87, 0xe9, 0xe1, 0x4a, 0x44, 0x05, 0xd5, 0x60, 0x2c, 0x6f, 0xda, 0x1c, 0x0c, 0xcd,
	0x09, 0x5c, 0x53, 0xbc, 0x8f, 0x61, 0x76, 0x1b, 0xee, 0x96, 0xff, 0x0f, 0x71, 0xf3, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x52, 0x00, 0x93, 0x3e, 0x9b, 0x01, 0x00, 0x00,
}
