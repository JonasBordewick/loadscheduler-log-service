package database

import "os"

var (
	DATABASE_HOST     string = os.Getenv("DATABASE_HOST") // Gets Internal Docker IP of DATABASE
	DATABASE_PORT            = os.Getenv("DATABASE_PORT") // Gets Internal Docker PORT of DATABASE
	DATABASE_USER            = os.Getenv("DATABASE_USER") // USERNAME
	DATABASE_PASSWORD        = os.Getenv("DATABASE_PAWD") // PASSWORD
	DATABASE_NAME            = os.Getenv("DATABASE_NAME") // DATABASE NAME
)