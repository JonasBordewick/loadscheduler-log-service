package service

import (
	"context"

	"github.com/JonasBordewick/loadscheduler-log-service/database"
	"github.com/JonasBordewick/loadscheduler-log-service/service/api"
)

type LogClient struct {
	client  api.LogServiceClient
	context context.Context
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
