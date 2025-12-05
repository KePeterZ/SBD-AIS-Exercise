package server

import (
	"context"
	"exc8/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer

	drinks []*pb.Drink
	orders map[int32]int32
}

func StartGrpcServer() error {
	grpcService := &GRPCService{
		drinks: []*pb.Drink{
			{Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"},
			{Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"},
			{Id: 3, Name: "Coffee", Price: 0, Description: "Mifare isn't that secure"},
		},
		orders: make(map[int32]int32),
	}

	srv := grpc.NewServer()
	pb.RegisterOrderServiceServer(srv, grpcService)

	fmt.Println("Protobuf Server running on :4444")
	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		return err
	}
	return srv.Serve(lis)
}

func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.DrinkList, error) {
	return &pb.DrinkList{Drinks: s.drinks}, nil
}

func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderRequest) (*wrapperspb.BoolValue, error) {
	for _, item := range req.Items {
		s.orders[item.DrinkId] += item.Quantity
	}

	return wrapperspb.Bool(true), nil
}

func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.OrderResponse, error) {
	var totals []*pb.OrderItem
	for id, qty := range s.orders {
		totals = append(totals, &pb.OrderItem{
			DrinkId:  id,
			Quantity: qty,
		})
	}

	return &pb.OrderResponse{Totals: totals}, nil
}
