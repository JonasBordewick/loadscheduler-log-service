package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/JonasBordewick/loadscheduler-log-service/database"
	"github.com/JonasBordewick/loadscheduler-log-service/service/api"
	"google.golang.org/grpc"
)

type LogServer struct {
	api.UnimplementedLogServiceServer
	database *database.DB
}

func (server *LogServer) mustEmbedUnimplementedConfiguratorAPIServer() {}

func (server *LogServer) GetAllLogs(req *api.Empty, stream api.LogService_GetAllLogsServer) error {
	logs, err := server.database.GetAllLogs()
	if err != nil {
		stream.Context().Done()
	}
	for _, l := range logs {
		stream.Send(l.ToGRPC())
	}
	stream.Context().Done()
	return nil
}

func (server *LogServer) GetLogsFromApplicant(req *api.Request, stream api.LogService_GetLogsFromApplicantServer) error {
	logs, err := server.database.GetLogsFromApplicant(req.RequestedApplicant)
	if err != nil {
		stream.Context().Done()
	}
	for _, l := range logs {
		stream.Send(l.ToGRPC())
	}
	stream.Context().Done()
	return nil
}

func (server *LogServer) WriteLog(ctx context.Context, log *api.Log) (*api.Empty, error) {
	server.database.WriteLog(database.FromGRPC(log))
	return &api.Empty{}, nil
}

func create_new_server(database *database.DB) *LogServer {
	var server = &LogServer{
		database: database,
	}
	return server
}

func StartServer(wg *sync.WaitGroup, port int, database *database.DB) {
	defer wg.Done()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterLogServiceServer(grpcServer, create_new_server(database))
	grpcServer.Serve(lis)
}
