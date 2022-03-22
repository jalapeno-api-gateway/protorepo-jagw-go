// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package jagw

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RequestServiceClient is the client API for RequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestServiceClient interface {
	GetLsNodes(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsNodeResponse, error)
	GetLsLinks(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsLinkResponse, error)
	GetLsPrefixes(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsPrefixResponse, error)
	GetLsSrv6Sids(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsSrv6SidResponse, error)
	GetLsNodeEdges(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsNodeEdgeResponse, error)
	GetPeers(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*PeerResponse, error)
	GetLsNodeCoordinates(ctx context.Context, in *LsNodeCoordinatesRequest, opts ...grpc.CallOption) (*LsNodeCoordinatesResponse, error)
	GetTelemetryData(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error)
	GetMeasurements(ctx context.Context, in *MeasurementsRequest, opts ...grpc.CallOption) (*MeasurementsResponse, error)
	GetMeasurementDetails(ctx context.Context, in *MeasurementDetailsRequest, opts ...grpc.CallOption) (*MeasurementDetailsResponse, error)
}

type requestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestServiceClient(cc grpc.ClientConnInterface) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) GetLsNodes(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsNodeResponse, error) {
	out := new(LsNodeResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetLsNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetLsLinks(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsLinkResponse, error) {
	out := new(LsLinkResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetLsLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetLsPrefixes(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsPrefixResponse, error) {
	out := new(LsPrefixResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetLsPrefixes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetLsSrv6Sids(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsSrv6SidResponse, error) {
	out := new(LsSrv6SidResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetLsSrv6Sids", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetLsNodeEdges(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*LsNodeEdgeResponse, error) {
	out := new(LsNodeEdgeResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetLsNodeEdges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetPeers(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*PeerResponse, error) {
	out := new(PeerResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetLsNodeCoordinates(ctx context.Context, in *LsNodeCoordinatesRequest, opts ...grpc.CallOption) (*LsNodeCoordinatesResponse, error) {
	out := new(LsNodeCoordinatesResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetLsNodeCoordinates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetTelemetryData(ctx context.Context, in *TelemetryRequest, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetTelemetryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetMeasurements(ctx context.Context, in *MeasurementsRequest, opts ...grpc.CallOption) (*MeasurementsResponse, error) {
	out := new(MeasurementsResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetMeasurements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) GetMeasurementDetails(ctx context.Context, in *MeasurementDetailsRequest, opts ...grpc.CallOption) (*MeasurementDetailsResponse, error) {
	out := new(MeasurementDetailsResponse)
	err := c.cc.Invoke(ctx, "/jagw.RequestService/GetMeasurementDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServiceServer is the server API for RequestService service.
// All implementations must embed UnimplementedRequestServiceServer
// for forward compatibility
type RequestServiceServer interface {
	GetLsNodes(context.Context, *TopologyRequest) (*LsNodeResponse, error)
	GetLsLinks(context.Context, *TopologyRequest) (*LsLinkResponse, error)
	GetLsPrefixes(context.Context, *TopologyRequest) (*LsPrefixResponse, error)
	GetLsSrv6Sids(context.Context, *TopologyRequest) (*LsSrv6SidResponse, error)
	GetLsNodeEdges(context.Context, *TopologyRequest) (*LsNodeEdgeResponse, error)
	GetPeers(context.Context, *TopologyRequest) (*PeerResponse, error)
	GetLsNodeCoordinates(context.Context, *LsNodeCoordinatesRequest) (*LsNodeCoordinatesResponse, error)
	GetTelemetryData(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
	GetMeasurements(context.Context, *MeasurementsRequest) (*MeasurementsResponse, error)
	GetMeasurementDetails(context.Context, *MeasurementDetailsRequest) (*MeasurementDetailsResponse, error)
	mustEmbedUnimplementedRequestServiceServer()
}

// UnimplementedRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRequestServiceServer struct {
}

func (UnimplementedRequestServiceServer) GetLsNodes(context.Context, *TopologyRequest) (*LsNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLsNodes not implemented")
}
func (UnimplementedRequestServiceServer) GetLsLinks(context.Context, *TopologyRequest) (*LsLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLsLinks not implemented")
}
func (UnimplementedRequestServiceServer) GetLsPrefixes(context.Context, *TopologyRequest) (*LsPrefixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLsPrefixes not implemented")
}
func (UnimplementedRequestServiceServer) GetLsSrv6Sids(context.Context, *TopologyRequest) (*LsSrv6SidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLsSrv6Sids not implemented")
}
func (UnimplementedRequestServiceServer) GetLsNodeEdges(context.Context, *TopologyRequest) (*LsNodeEdgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLsNodeEdges not implemented")
}
func (UnimplementedRequestServiceServer) GetPeers(context.Context, *TopologyRequest) (*PeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeers not implemented")
}
func (UnimplementedRequestServiceServer) GetLsNodeCoordinates(context.Context, *LsNodeCoordinatesRequest) (*LsNodeCoordinatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLsNodeCoordinates not implemented")
}
func (UnimplementedRequestServiceServer) GetTelemetryData(context.Context, *TelemetryRequest) (*TelemetryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTelemetryData not implemented")
}
func (UnimplementedRequestServiceServer) GetMeasurements(context.Context, *MeasurementsRequest) (*MeasurementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeasurements not implemented")
}
func (UnimplementedRequestServiceServer) GetMeasurementDetails(context.Context, *MeasurementDetailsRequest) (*MeasurementDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeasurementDetails not implemented")
}
func (UnimplementedRequestServiceServer) mustEmbedUnimplementedRequestServiceServer() {}

// UnsafeRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestServiceServer will
// result in compilation errors.
type UnsafeRequestServiceServer interface {
	mustEmbedUnimplementedRequestServiceServer()
}

func RegisterRequestServiceServer(s grpc.ServiceRegistrar, srv RequestServiceServer) {
	s.RegisterService(&RequestService_ServiceDesc, srv)
}

func _RequestService_GetLsNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetLsNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetLsNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetLsNodes(ctx, req.(*TopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetLsLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetLsLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetLsLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetLsLinks(ctx, req.(*TopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetLsPrefixes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetLsPrefixes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetLsPrefixes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetLsPrefixes(ctx, req.(*TopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetLsSrv6Sids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetLsSrv6Sids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetLsSrv6Sids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetLsSrv6Sids(ctx, req.(*TopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetLsNodeEdges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetLsNodeEdges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetLsNodeEdges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetLsNodeEdges(ctx, req.(*TopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetPeers(ctx, req.(*TopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetLsNodeCoordinates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LsNodeCoordinatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetLsNodeCoordinates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetLsNodeCoordinates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetLsNodeCoordinates(ctx, req.(*LsNodeCoordinatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetTelemetryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TelemetryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetTelemetryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetTelemetryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetTelemetryData(ctx, req.(*TelemetryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetMeasurements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasurementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetMeasurements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetMeasurements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetMeasurements(ctx, req.(*MeasurementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_GetMeasurementDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasurementDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetMeasurementDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jagw.RequestService/GetMeasurementDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetMeasurementDetails(ctx, req.(*MeasurementDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestService_ServiceDesc is the grpc.ServiceDesc for RequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jagw.RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLsNodes",
			Handler:    _RequestService_GetLsNodes_Handler,
		},
		{
			MethodName: "GetLsLinks",
			Handler:    _RequestService_GetLsLinks_Handler,
		},
		{
			MethodName: "GetLsPrefixes",
			Handler:    _RequestService_GetLsPrefixes_Handler,
		},
		{
			MethodName: "GetLsSrv6Sids",
			Handler:    _RequestService_GetLsSrv6Sids_Handler,
		},
		{
			MethodName: "GetLsNodeEdges",
			Handler:    _RequestService_GetLsNodeEdges_Handler,
		},
		{
			MethodName: "GetPeers",
			Handler:    _RequestService_GetPeers_Handler,
		},
		{
			MethodName: "GetLsNodeCoordinates",
			Handler:    _RequestService_GetLsNodeCoordinates_Handler,
		},
		{
			MethodName: "GetTelemetryData",
			Handler:    _RequestService_GetTelemetryData_Handler,
		},
		{
			MethodName: "GetMeasurements",
			Handler:    _RequestService_GetMeasurements_Handler,
		},
		{
			MethodName: "GetMeasurementDetails",
			Handler:    _RequestService_GetMeasurementDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "requestservice/requestservice.proto",
}
