package controller

import (
	d "github.com/rockspoon/rs.cor.device-model/model"
	s "github.com/rockspoon/rs.cor.middleware/soajs"
	"github.com/rockspoon/rs.cor.printer-ms/converter"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template/kitchen"
)

// DeviceMS devices microservice
type DeviceMS interface {
	GetDefaultPrinter(path, key string) (*d.Printer, error)
}

// PrintController Print functions
type PrintController struct {
	KitchenReceiptGenerator kitchen.Generator
	Converter               converter.ByteCodeGenerator
	DeviceMS                DeviceMS
}

// NewPrintController creates a new print controller
func NewPrintController(deviceMS DeviceMS) PrintController {
	return PrintController{
		KitchenReceiptGenerator: kitchen.Generator{},
		Converter:               converter.NewByteCodeGenerator(),
		DeviceMS:                deviceMS,
	}
}

// KitchenReceipt print a kitchen receipt
func (c PrintController) KitchenReceipt(request model.KitchenReceiptRequest, cData *s.ContextData) (*model.Payload, error) {
	printer, err := c.DeviceMS.GetDefaultPrinter(cData.Tenant.Key, cData.Paths[s.DEVICE])
	if err != nil {
		return nil, err
	}

	receipt := request.ToKitchenReceipt()
	commands := c.KitchenReceiptGenerator.Generate(receipt, printer.PrinterSettings.PrinterType)

	payload := model.Payload{
		PrintPayload:    c.Converter.Convert(commands, printer.PrinterSettings.PrinterType),
		IPAddress:       printer.IPAddress,
		PrinterModel:    printer.PrinterModel,
		DescribeMessage: "[Printing Job] Kitchen Receipt",
	}
	return &payload, nil
}
