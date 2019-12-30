package dependency

import (
	"context"

	d "github.com/rockspoon/rs.cor.device-model/model"
	"github.com/rockspoon/rs.cor.printer-ms/model"
)

// DeviceMS devices microservice
type DeviceMS interface {
	GetDefaultPrinter(ctx context.Context) (*d.Printer, error)
}

// PrintService is the interface for the printer controller
type PrintService interface {
	KitchenReceipt(ctx context.Context, request model.KitchenReceipt) (*model.Payload, error)
	TableBill(ctx context.Context, request model.Bill) (*model.Payload, error)
	PaymentReceipt(ctx context.Context, request model.PaymentReceipt) (*model.Payload, error)
}
