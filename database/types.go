package database

import "github.com/JonasBordewick/loadscheduler-log-service/service/api"

type Log struct {
	Id        int
	Applicant string
	Timestamp string
	Message   string
}

func (log *Log) ToGRPC() *api.Log {
	return &api.Log{
		Id: int32(log.Id),
		Applicant: log.Applicant,
		Timestamp: log.Timestamp,
		Message: log.Message,
	}
}

func FromGRPC(log *api.Log) *Log {
	return &Log{
		Id: 0,
		Applicant: log.Applicant,
		Timestamp: "0",
		Message: log.Message,
	}
}