package domain

import (
	"github.com/nafisalfiani/ketson-go-lib/broker"
	"github.com/nafisalfiani/ketson-go-lib/email"
	"github.com/nafisalfiani/ketson-go-lib/log"
	"github.com/nafisalfiani/ketson-go-lib/parser"
	"github.com/nafisalfiani/ketson-notification-service/domain/mailer"
)

type Domains struct {
	Mailer mailer.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, broker broker.Interface, mail email.Interface) *Domains {
	return &Domains{
		Mailer: mailer.Init(logger, mail),
	}
}
