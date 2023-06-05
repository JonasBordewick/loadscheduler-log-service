package service

import (
	"context"
	"fmt"

	"github.com/JonasBordewick/loadscheduler-log-service/database"
	"github.com/JonasBordewick/loadscheduler-log-service/service/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LogClient struct {
	client  api.LogServiceClient
	context context.Context
}
var instance *LogClient

func GetClient(address string, port string) (*LogClient, error) {
	if instance == nil {
		connection, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		client := api.NewLogServiceClient(connection)
		context := context.Background()

		instance = &LogClient{
			client: client,
			context: context,
		}
	}
	return instance, nil
}

func (client *LogClient) GetAllLogs() ([]*database.Log, error) {
	logs, err := client.client.GetAllLogs(client.context, &api.Empty{})
	if err != nil {
		return nil, err
	}
	parsed_logs := []*database.Log{}

	for _, log := range logs.Logs {
		parsed_logs = append(parsed_logs, database.FromGRPC(log))
	}

	return parsed_logs, nil
}

func (client *LogClient) GetLogsFromApplicant(applicant string) ([]*database.Log, error) {
	logs, err := client.client.GetLogsFromApplicant(client.context, &api.Request{
		RequestedApplicant: applicant,
	})
	if err != nil {
		return nil, err
	}
	parsed_logs := []*database.Log{}

	for _, log := range logs.Logs {
		parsed_logs = append(parsed_logs, database.FromGRPC(log))
	}

	return parsed_logs, nil
}

func (client *LogClient) WriteLog(applicant string, message string) error {
	_, err := client.client.WriteLog(client.context, &api.Log{Applicant: applicant, Message: message})
	return err
}
