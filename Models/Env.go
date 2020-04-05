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
}

func InitEnv()  {
	envStruct = &Env{
		Env: 			  			os.Getenv("ENV"),
		Port:						os.Getenv("PORT"),
		Url:						os.Getenv("URL"),
		TimeShiftGrpcServer:		os.Getenv("TIMESHIFT_GRPC_SERVER"),
		TimeShiftGrpcPort:			os.Getenv("TIMESHIFT_GRPC_PORT"),
	}
	fmt.Println(envStruct)
}

func GetEnvStruct() *Env  {
	return  envStruct
}