package Models

import (
	"fmt"
	"os"
)

var envStruct *Env

type Env struct {
	Port string
	Url string
	TimeShiftGrpcServer string
	TimeShiftGrpcPort string
	Env string
	TracingConnection string
}

func InitEnv()  {
	envStruct = &Env{
		Env: 			  			os.Getenv("ENV"),
		Port:						os.Getenv("PORT"),
		Url:						os.Getenv("URL"),
		TimeShiftGrpcServer:		os.Getenv("TIMESHIFT_GRPC_SERVER"),
		TimeShiftGrpcPort:			os.Getenv("TIMESHIFT_GRPC_PORT"),
		TracingConnection: 			os.Getenv("TRACING_CONNECTION"),
	}
	fmt.Println(envStruct)
}

func GetEnvStruct() *Env  {
	return  envStruct
}