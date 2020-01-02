package controller

import (
	"context"
	"errors"
	"testing"

	orderModel "github.com/rockspoon/rs.com.order-model/model"
	d "github.com/rockspoon/rs.cor.device-model/model"
	"github.com/rockspoon/rs.cor.printer-ms/converter"
	"github.com/rockspoon/rs.cor.printer-ms/mocks"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	billTemplate "github.com/rockspoon/rs.cor.printer-ms/template/bill"
	kitchenTemplate "github.com/rockspoon/rs.cor.printer-ms/template/kitchen"
	receiptTemplate "github.com/rockspoon/rs.cor.printer-ms/template/receipt"
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
			printer: &d.Printer{IPAddress: "123", PrinterModel: "TSPP"},
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
			printer: &d.Printer{IPAddress: "123", PrinterModel: "TSPP"},
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
			printer: &d.Printer{IPAddress: "123", PrinterModel: "TSPP"},
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
