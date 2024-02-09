package main

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nafisalfiani/ketson-go-lib/auth"
	"github.com/nafisalfiani/ketson-go-lib/broker"
	"github.com/nafisalfiani/ketson-go-lib/configreader"
	"github.com/nafisalfiani/ketson-go-lib/email"
	"github.com/nafisalfiani/ketson-go-lib/log"
	"github.com/nafisalfiani/ketson-go-lib/parser"
	"github.com/nafisalfiani/ketson-go-lib/security"
	"github.com/nafisalfiani/ketson-notification-service/config"
	"github.com/nafisalfiani/ketson-notification-service/domain"
	"github.com/nafisalfiani/ketson-notification-service/handler/consumer"
	"github.com/nafisalfiani/ketson-notification-service/handler/grpc"
	"github.com/nafisalfiani/ketson-notification-service/usecase"
)

func main() {
	log.DefaultLogger().Info(context.Background(), "starting server...")

	// init config file
	cfg := configreader.Init(configreader.Options{
		Type:       configreader.Viper,
		ConfigFile: "./config.json",
	})

	// read from config
	config := &config.Application{}
	if err := cfg.ReadConfig(config); err != nil {
		log.DefaultLogger().Fatal(context.Background(), err)
	}

	log.DefaultLogger().Info(context.Background(), config)

	// init logger
	logger := log.Init(config.Log, log.Zerolog)

	// init parser
	parser := parser.InitParser(logger, parser.Options{})

	// init validator
	validator := validator.New(validator.WithRequiredStructEnabled())

	// init security
	sec := security.Init(logger, config.Security)

	// init mail
	mail := email.Init(config.Mail, logger)

	// init broker
	broker, err := broker.Init(config.Broker, logger, parser.JSONParser())
	if err != nil {
		logger.Fatal(context.Background(), err)
	}
	defer broker.Close()

	// init auth
	auth := auth.Init(config.Auth, logger, parser.JSONParser(), http.DefaultClient)

	// init domain
	dom := domain.Init(logger, parser.JSONParser(), broker, mail)

	// init usecase
	uc := usecase.Init(config.ApiGateway.Url, logger, dom)

	// init consumer
	consumer := consumer.Init(logger, broker, parser.JSONParser(), uc.Mailer)
	go consumer.StartRegistrationConsumer(context.Background())
	go consumer.StartTransactionConsumer(context.Background())

	// TODO: init scheduler

	// init grpc
	grpc := grpc.Init(config.Grpc, logger, uc, sec, auth, validator)

	// start grpc server
	grpc.Run()
}
