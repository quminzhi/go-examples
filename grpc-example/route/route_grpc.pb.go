// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: route.proto

package route

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

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuideClient interface {
	// unary
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	// server-side streaming
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error)
	// client-side streaming
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	// bidirectional streaming
	Recommend(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecommendClient, error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/route.RouteGuide/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[0], "/route.RouteGuide/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[1], "/route.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) Recommend(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecommendClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[2], "/route.RouteGuide/Recommend", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecommendClient{stream}
	return x, nil
}

type RouteGuide_RecommendClient interface {
	Send(*RecommendationRequest) error
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideRecommendClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecommendClient) Send(m *RecommendationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecommendClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServer is the server API for RouteGuide service.
// All implementations must embed UnimplementedRouteGuideServer
// for forward compatibility
type RouteGuideServer interface {
	// unary
	GetFeature(context.Context, *Point) (*Feature, error)
	// server-side streaming
	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	// client-side streaming
	RecordRoute(RouteGuide_RecordRouteServer) error
	// bidirectional streaming
	Recommend(RouteGuide_RecommendServer) error
	mustEmbedUnimplementedRouteGuideServer()
}

// UnimplementedRouteGuideServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (UnimplementedRouteGuideServer) GetFeature(context.Context, *Point) (*Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedRouteGuideServer) ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}
func (UnimplementedRouteGuideServer) RecordRoute(RouteGuide_RecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}
func (UnimplementedRouteGuideServer) Recommend(RouteGuide_RecommendServer) error {
	return status.Errorf(codes.Unimplemented, "method Recommend not implemented")
}
func (UnimplementedRouteGuideServer) mustEmbedUnimplementedRouteGuideServer() {}

// UnsafeRouteGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuideServer will
// result in compilation errors.
type UnsafeRouteGuideServer interface {
	mustEmbedUnimplementedRouteGuideServer()
}

func RegisterRouteGuideServer(s grpc.ServiceRegistrar, srv RouteGuideServer) {
	s.RegisterService(&RouteGuide_ServiceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).ListFeatures(m, &routeGuideListFeaturesServer{stream})
}

type RouteGuide_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordRoute(&routeGuideRecordRouteServer{stream})
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_Recommend_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).Recommend(&routeGuideRecommendServer{stream})
}

type RouteGuide_RecommendServer interface {
	Send(*Feature) error
	Recv() (*RecommendationRequest, error)
	grpc.ServerStream
}

type routeGuideRecommendServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecommendServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecommendServer) Recv() (*RecommendationRequest, error) {
	m := new(RecommendationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuide_ServiceDesc is the grpc.ServiceDesc for RouteGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuide_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Recommend",
			Handler:       _RouteGuide_Recommend_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "route.proto",
}
