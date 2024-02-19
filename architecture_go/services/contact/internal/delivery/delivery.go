package delivery

import (
	"architecture_go/services/contact/internal/usecase"
	"log/slog"
)

type HTTPDelivery struct {
	UseCase usecase.ContactUseCase
	Logger  *slog.Logger
}

func NewDelivery(useCase usecase.ContactUseCase, logger *slog.Logger) HTTPDelivery {
	return HTTPDelivery{
		UseCase: useCase,
		Logger:  logger,
	}
}
