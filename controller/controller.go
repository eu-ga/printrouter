package controller

import (
	d "github.com/rockspoon/rs.cor.device-model/model"
	s "github.com/rockspoon/rs.cor.middleware/model"
	"github.com/rockspoon/rs.cor.printer-ms/converter"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	billTemplate "github.com/rockspoon/rs.cor.printer-ms/template/bill"
	kitchenTemplate "github.com/rockspoon/rs.cor.printer-ms/template/kitchen"
	receiptTemplate "github.com/rockspoon/rs.cor.printer-ms/template/receipt"
)

// DeviceMS devices microservice
type DeviceMS interface {
	GetDefaultPrinter(path, key string) (*d.Printer, error)
}

// PrintController Print functions
type PrintController struct {
	KitchenReceiptGenerator kitchenTemplate.Generator
	TableBillGenerator      billTemplate.Generator
	PaymentReceiptGenerator receiptTemplate.Generator
	Converter               converter.ByteCodeGenerator
	DeviceMS                DeviceMS
}

// NewPrintController creates a new print controller
func NewPrintController(deviceMS DeviceMS) PrintController {
	return PrintController{
		KitchenReceiptGenerator: kitchenTemplate.Generator{},
		Converter:               converter.NewByteCodeGenerator(),
		DeviceMS:                deviceMS,
	}
}

// KitchenReceipt print a kitchen receipt
func (c PrintController) KitchenReceipt(receipt model.KitchenReceipt, cData *s.ContextData) (*model.Payload, error) {
	printer, err := c.DeviceMS.GetDefaultPrinter(cData.Tenant.Key, cData.Paths[s.DEVICE])
	if err != nil {
		return nil, err
	}

	commands := c.KitchenReceiptGenerator.Generate(receipt, printer.PrinterModel)

	payload := model.Payload{
		PrintPayload:    c.Converter.Convert(commands, printer.PrinterModel),
		IPAddress:       printer.IPAddress,
		PrinterModel:    printer.PrinterModel,
		DescribeMessage: "[Printing Job] Kitchen Receipt",
	}
	return &payload, nil
}

// TableBill prints a table bill
func (c PrintController) TableBill(bill model.Bill, cData *s.ContextData) (*model.Payload, error) {
	printer, err := c.DeviceMS.GetDefaultPrinter(cData.Tenant.Key, cData.Paths[s.DEVICE])
	if err != nil {
		return nil, err
	}

	commands := c.TableBillGenerator.Generate(bill, printer.PrinterModel)

	payload := model.Payload{
		PrintPayload:    c.Converter.Convert(commands, printer.PrinterModel),
		IPAddress:       printer.IPAddress,
		PrinterModel:    printer.PrinterModel,
		DescribeMessage: "[Printing Job] Table Bill",
	}
	return &payload, nil
}

// PaymentReceipt prints a payment receipt
func (c PrintController) PaymentReceipt(receipt model.PaymentReceipt, cData *s.ContextData) (*model.Payload, error) {
	printer, err := c.DeviceMS.GetDefaultPrinter(cData.Tenant.Key, cData.Paths[s.DEVICE])
	if err != nil {
		return nil, err
	}

	commands := c.PaymentReceiptGenerator.Generate(receipt, printer.PrinterModel)

	payload := model.Payload{
		PrintPayload:    c.Converter.Convert(commands, printer.PrinterModel),
		IPAddress:       printer.IPAddress,
		PrinterModel:    printer.PrinterModel,
		DescribeMessage: "[Printing Job] Payment Receipt",
	}
	return &payload, nil
}
