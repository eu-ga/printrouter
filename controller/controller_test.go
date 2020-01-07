package controller

import (
	"context"
	"errors"
	"testing"

	orderModel "github.com/rockspoon/rs.com.order-model/model"
	"github.com/rockspoon/rs.cor.common-model/address"
	d "github.com/rockspoon/rs.cor.printer-ms/controller/integration/model"
	"github.com/rockspoon/rs.cor.printer-ms/converter"
	ec "github.com/rockspoon/rs.cor.printer-ms/errors"
	"github.com/rockspoon/rs.cor.printer-ms/mocks"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	billTemplate "github.com/rockspoon/rs.cor.printer-ms/template/bill"
	kitchenTemplate "github.com/rockspoon/rs.cor.printer-ms/template/kitchen"
	receiptTemplate "github.com/rockspoon/rs.cor.printer-ms/template/receipt"
	testTemplate "github.com/rockspoon/rs.cor.printer-ms/template/test"
	"github.com/stretchr/testify/require"
)

func TestController_KitchenReceipt(t *testing.T) {
	receipt := model.KitchenReceipt{}
	cmdrs := kitchenTemplate.Generator{}.Generate(receipt)
	strCmdrs := converter.ByteCodeGenerator{}.Convert(cmdrs)

	tt := []struct {
		name         string
		printer      *d.Printer
		printerError error
		payload      *model.Payload
		expErr       string
	}{
		{
			name:         "Printer MS Error",
			printerError: errors.New("no default printer"),
			expErr:       "no default printer",
		},
		{
			name:    "success",
			printer: &d.Printer{IP: "123", Model: "TSPP"},
			payload: &model.Payload{IPAddress: "123", PrinterModel: "TSPP", PrintPayload: strCmdrs, DescribeMessage: "[Printing Job] Kitchen Receipt"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deviceMS := new(mocks.DeviceMS)
			controller := NewPrintController(deviceMS)
			ctx := context.Background()

			deviceMS.On("GetDefaultPrinter", ctx).Return(tc.printer, tc.printerError)
			payload, err := controller.KitchenReceipt(ctx, receipt)
			if tc.name == "success" {
				require.NoError(t, err)
				require.Equal(t, tc.payload, payload)
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}

func TestController_TableBill(t *testing.T) {
	bill := model.Bill{}
	cmdrs := billTemplate.Generator{}.Generate(bill)
	strCmdrs := converter.ByteCodeGenerator{}.Convert(cmdrs)

	tt := []struct {
		name         string
		printer      *d.Printer
		printerError error
		bill         model.Bill
		payload      *model.Payload
		expErr       string
	}{
		{
			name:         "Printer MS Error",
			printerError: errors.New("no default printer"),
			expErr:       "no default printer",
		},
		{
			name:    "success",
			printer: &d.Printer{IP: "123", Model: "TSPP"},
			payload: &model.Payload{IPAddress: "123", PrinterModel: "TSPP", PrintPayload: strCmdrs, DescribeMessage: "[Printing Job] Table Bill"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deviceMS := new(mocks.DeviceMS)
			controller := NewPrintController(deviceMS)
			ctx := context.Background()

			deviceMS.On("GetDefaultPrinter", ctx).Return(tc.printer, tc.printerError)

			payload, err := controller.TableBill(ctx, tc.bill)
			if tc.name == "success" {
				require.NoError(t, err)
				require.Equal(t, tc.payload, payload)
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}

func TestController_PaymentReceipt(t *testing.T) {
	receipt := model.PaymentReceipt{
		OrderType: orderModel.OrderTypeDineIn,
		Check: model.Check{
			DineInOptions: &model.DineInOptions{},
		},
	}
	cmdrs := receiptTemplate.Generator{}.Generate(receipt)
	strCmdrs := converter.ByteCodeGenerator{}.Convert(cmdrs)

	tt := []struct {
		name         string
		printer      *d.Printer
		printerError error
		receipt      model.PaymentReceipt
		payload      *model.Payload
		expErr       string
	}{
		{
			name:         "Printer MS Error",
			printerError: errors.New("no default printer"),
			expErr:       "no default printer",
		},
		{
			name:    "success",
			printer: &d.Printer{IP: "123", Model: "TSPP"},
			receipt: receipt,
			payload: &model.Payload{IPAddress: "123", PrinterModel: "TSPP", PrintPayload: strCmdrs, DescribeMessage: "[Printing Job] Payment Receipt"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deviceMS := new(mocks.DeviceMS)
			controller := NewPrintController(deviceMS)
			ctx := context.Background()

			deviceMS.On("GetDefaultPrinter", ctx).Return(tc.printer, tc.printerError)

			payload, err := controller.PaymentReceipt(ctx, tc.receipt)
			if tc.name == "success" {
				require.NoError(t, err)
				require.Equal(t, tc.payload, payload)
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}

func TestController_TestPayload(t *testing.T) {
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
	cmdrs := testTemplate.Generator{}.Generate(test)
	strCmdrs := converter.ByteCodeGenerator{}.Convert(cmdrs)

	tt := []struct {
		name         string
		test         model.TestPayload
		payload      *model.Payload
		ipAddress    string
		printerModel string
		expErr       error
	}{
		{
			name:   "missing ip address",
			expErr: ec.InvalidIPAddress(),
		},
		{
			name:      "not valid ip address",
			expErr:    ec.InvalidIPAddress(),
			ipAddress: "123",
		},
		{
			name:      "missing printer model",
			ipAddress: "192.168.0.1",
			expErr:    ec.InvalidPrinterModel(),
		},
		{
			name:         "success",
			test:         test,
			ipAddress:    "192.168.0.1",
			printerModel: "TSPP",
			payload:      &model.Payload{IPAddress: "192.168.0.1", PrinterModel: "TSPP", PrintPayload: strCmdrs, DescribeMessage: "[Printing Job] Test"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			controller := NewPrintController(nil)
			ctx := context.Background()

			payload, err := controller.TestPayload(ctx, tc.ipAddress, tc.printerModel)
			if tc.name == "success" {
				require.NoError(t, err)
				require.Equal(t, tc.payload, payload)
			} else {
				require.Equal(t, err, tc.expErr)
			}
		})
	}
}
