package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"logger-service/data"
	"logger-service/logs"
	"net"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, request *logs.LogRequest) (*logs.LogResponse, error) {
	input := request.GetLogEntry()

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)

	if err != nil {
		response := &logs.LogResponse{
			Result: "Failed to write log",
		}

		return response, err
	}

	response := &logs.LogResponse{
		Result: "Logged!",
	}

	return response, nil
}

func (app *Config) gRPCListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))

	if err != nil {
		log.Fatalf("Failed to listen for gPRC: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{
		Models: app.Models,
	})

	log.Printf("gRPC Server started on port %s", gRpcPort)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to listen for gPRC: %v", err)
	}
}