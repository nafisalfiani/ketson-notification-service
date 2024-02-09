package config

import (
	"github.com/nafisalfiani/ketson-go-lib/auth"
	"github.com/nafisalfiani/ketson-go-lib/broker"
	"github.com/nafisalfiani/ketson-go-lib/email"
	"github.com/nafisalfiani/ketson-go-lib/log"
	"github.com/nafisalfiani/ketson-go-lib/security"
	"github.com/nafisalfiani/ketson-notification-service/handler/grpc"
)

type Application struct {
	ApiGateway ApiGateway
	Auth       auth.Config
	Log        log.Config
	Security   security.Config
	Mail       email.Config
	Broker     broker.Config
	Grpc       grpc.Config
}

type ApiGateway struct {
	Url string
}
