package config

import (
	"github.com/nafisalfiani/ketson-notification-service/handler/grpc"
	"github.com/nafisalfiani/ketson-notification-service/lib/auth"
	"github.com/nafisalfiani/ketson-notification-service/lib/broker"
	"github.com/nafisalfiani/ketson-notification-service/lib/email"
	"github.com/nafisalfiani/ketson-notification-service/lib/log"
	"github.com/nafisalfiani/ketson-notification-service/lib/security"
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
