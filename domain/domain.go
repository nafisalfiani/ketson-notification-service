package domain

import (
	"github.com/nafisalfiani/ketson-notification-service/domain/mailer"
	"github.com/nafisalfiani/ketson-notification-service/lib/broker"
	"github.com/nafisalfiani/ketson-notification-service/lib/email"
	"github.com/nafisalfiani/ketson-notification-service/lib/log"
	"github.com/nafisalfiani/ketson-notification-service/lib/parser"
)

type Domains struct {
	Mailer mailer.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, broker broker.Interface, mail email.Interface) *Domains {
	return &Domains{
		Mailer: mailer.Init(logger, mail),
	}
}
