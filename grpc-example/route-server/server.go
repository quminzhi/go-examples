package main

import (
	"context"
	"google.golang.org/protobuf/proto"
	"log"
	"math"
	"net"

	pb "github.com/quminzhi/go-examples/grpc-example/route"
	"google.golang.org/grpc"
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	features []*pb.Feature // A fake db
}

// GetFeature is a unary
func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.features {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil, nil
}

// inRange checks if a point is in the range of rectangle
func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

// ListFeatures is a server-side streaming
func (s *routeGuideServer) ListFeatures(rectangle *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	for _, feature := range s.features {
		if inRange(feature.Location, rectangle) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

// RecordRoute is a client-side streaming
func (s *routeGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error {
	return nil
}

// Recommend is a bidirectional streaming
func (s *routeGuideServer) Recommend(pb.RouteGuide_RecommendServer) error {
	return nil
}

func newRouteGuideServer() pb.RouteGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{
			{Name: "上海交通大学闵行校区 上海市闵行区东川路800号", Location: &pb.Point{
				Latitude:  310235000,
				Longitude: 121437403,
			}},
			{Name: "复旦大学 上海市杨浦区五角场邯郸路220号", Location: &pb.Point{
				Latitude:  312978870,
				Longitude: 121503457,
			}},
			{Name: "华东理工大学 上海市徐汇区梅陇路130号", Location: &pb.Point{
				Latitude:  311416130,
				Longitude: 121424904,
			}},
		},
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newRouteGuideServer())
	log.Fatalln(grpcServer.Serve(lis))
}
