package usecase

import (
	"github.com/nafisalfiani/ketson-go-lib/log"
	"github.com/nafisalfiani/ketson-notification-service/domain"
	"github.com/nafisalfiani/ketson-notification-service/usecase/mailer"
)

type Usecases struct {
	Mailer mailer.Interface
}

func Init(apiGwUrl string, logger log.Interface, dom *domain.Domains) *Usecases {
	return &Usecases{
		Mailer: mailer.Init(mailer.Config{
			ApiGatewayUrl: apiGwUrl,
		}, dom.Mailer),
	}
}
