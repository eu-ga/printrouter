package controller

import (
	"context"

	"github.com/rockspoon/rs.cor.common-model/address"
	"github.com/rockspoon/rs.cor.printer-ms/converter"
	"github.com/rockspoon/rs.cor.printer-ms/dependency"
	"github.com/rockspoon/rs.cor.printer-ms/errors"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	billTemplate "github.com/rockspoon/rs.cor.printer-ms/template/bill"
	kitchenTemplate "github.com/rockspoon/rs.cor.printer-ms/template/kitchen"
	receiptTemplate "github.com/rockspoon/rs.cor.printer-ms/template/receipt"
	testTemplate "github.com/rockspoon/rs.cor.printer-ms/template/test"
)

// PrintController Print functions
type PrintController struct {
	KitchenReceiptGenerator kitchenTemplate.Generator
	TableBillGenerator      billTemplate.Generator
	PaymentReceiptGenerator receiptTemplate.Generator
	TestGenerator           testTemplate.Generator
	Converter               converter.ByteCodeGenerator
	DeviceMS                dependency.DeviceMS
}

// NewPrintController creates a new print controller
func NewPrintController(deviceMS dependency.DeviceMS) PrintController {
	return PrintController{
		KitchenReceiptGenerator: kitchenTemplate.Generator{},
		TestGenerator:           testTemplate.Generator{},
		Converter:               converter.NewByteCodeGenerator(),
		DeviceMS:                deviceMS,
	}
}

// KitchenReceipt print a kitchen receipt
func (c PrintController) KitchenReceipt(ctx context.Context, receipt model.KitchenReceipt) (*model.Payload, error) {
	printer, err := c.DeviceMS.GetDefaultPrinter(ctx)
	if err != nil {
		return nil, err
	}

	commands := c.KitchenReceiptGenerator.Generate(receipt)

	payload := model.Payload{
		PrintPayload:    c.Converter.Convert(commands),
		IPAddress:       printer.IP,
		PrinterModel:    printer.Model,
		DescribeMessage: "[Printing Job] Kitchen Receipt",
	}
	return &payload, nil
}

// TableBill prints a table bill
func (c PrintController) TableBill(ctx context.Context, bill model.Bill) (*model.Payload, error) {
	printer, err := c.DeviceMS.GetDefaultPrinter(ctx)
	if err != nil {
		return nil, err
	}

	commands := c.TableBillGenerator.Generate(bill)

	payload := model.Payload{
		PrintPayload:    c.Converter.Convert(commands),
		IPAddress:       printer.IP,
		PrinterModel:    printer.Model,
		DescribeMessage: "[Printing Job] Table Bill",
	}
	return &payload, nil
}

// PaymentReceipt prints a payment receipt
func (c PrintController) PaymentReceipt(ctx context.Context, receipt model.PaymentReceipt) (*model.Payload, error) {
	printer, err := c.DeviceMS.GetDefaultPrinter(ctx)
	if err != nil {
		return nil, err
	}

	commands := c.PaymentReceiptGenerator.Generate(receipt)

	payload := model.Payload{
		PrintPayload:    c.Converter.Convert(commands),
		IPAddress:       printer.IP,
		PrinterModel:    printer.Model,
		DescribeMessage: "[Printing Job] Payment Receipt",
	}
	return &payload, nil
}

// TestPayload prints a test payload
func (c PrintController) TestPayload(ctx context.Context, ipAddress, printerModel string) (*model.Payload, error) {
	test := model.TestPayload{
		Restaurant: model.RestaurantInfo{
			Name: "Alvorada",
			Address: address.Address{
				Name:     "Restaurante Alvorada",
				City:     "São Paulo",
				State:    "SP",
				Country:  "Brasil",
				Address1: "Rua Alvorada, 154",
				Region:   "Vila Olímpia",
				ZipCode:  "05593010",
			},
		},
	}

	if ipAddress == "" {
		return nil, errors.InvalidIPAddress()
	}

	if printerModel == "" {
		return nil, errors.InvalidPrinterModel()
	}

	commands := c.TestGenerator.Generate(test)

	payload := model.Payload{
		PrintPayload:    c.Converter.Convert(commands),
		IPAddress:       ipAddress,
		PrinterModel:    printerModel,
		DescribeMessage: "[Printing Job] Test",
	}
	return &payload, nil
}
