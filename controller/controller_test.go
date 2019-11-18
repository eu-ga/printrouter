package controller

import (
	"errors"
	"testing"

	d "github.com/rockspoon/rs.cor.device-model/model"
	s "github.com/rockspoon/rs.cor.middleware/model"
	"github.com/rockspoon/rs.cor.printer-ms/converter"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	b "github.com/rockspoon/rs.cor.printer-ms/template/bill"
	"github.com/rockspoon/rs.cor.printer-ms/template/kitchen"
	"github.com/stretchr/testify/require"
)

type MockDeviceMS struct {
	Printer *d.Printer
	Error   error
}

func (m MockDeviceMS) GetDefaultPrinter(path, key string) (*d.Printer, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Printer, nil
}

func TestController_KitchenReceipt(t *testing.T) {
	receipt := model.KitchenReceiptRequest{}.ToKitchenReceipt()
	cmdrs := kitchen.Generator{}.Generate(receipt, d.TSPPrinterType)
	strCmdrs := converter.ByteCodeGenerator{}.Convert(cmdrs, d.TSPPrinterType)

	tt := []struct {
		name         string
		data         *s.ContextData
		printer      *d.Printer
		printerError error
		request      model.KitchenReceiptRequest
		payload      *model.Payload
		expErr       string
	}{
		{
			name:         "Printer MS Error",
			printerError: errors.New("no default printer"),
			data: &s.ContextData{
				//nolint:staticcheck
				Tenant: s.Tenant{Key: "1"},
				Paths:  map[string]string{s.DEVICE: "device"},
			},
			expErr: "no default printer",
		},
		{
			name:    "success",
			printer: &d.Printer{IPAddress: "123", PrinterSettings: d.PrinterSettings{PrinterType: d.TSPPrinterType}},
			data: &s.ContextData{
				//nolint:staticcheck
				Tenant: s.Tenant{Key: "1"},
				Paths:  map[string]string{s.DEVICE: "device"},
			},
			payload: &model.Payload{IPAddress: "123", PrintPayload: strCmdrs, DescribeMessage: "[Printing Job] Kitchen Receipt"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deviceMS := MockDeviceMS{Printer: tc.printer, Error: tc.printerError}
			controller := NewPrintController(deviceMS)

			payload, err := controller.KitchenReceipt(tc.request, tc.data)
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
	cmdrs := b.Generator{}.Generate(bill, d.TSPPrinterType)
	strCmdrs := converter.ByteCodeGenerator{}.Convert(cmdrs, d.TSPPrinterType)

	tt := []struct {
		name         string
		data         *s.ContextData
		printer      *d.Printer
		printerError error
		bill         model.Bill
		payload      *model.Payload
		expErr       string
	}{
		{
			name:         "Printer MS Error",
			printerError: errors.New("no default printer"),
			data: &s.ContextData{
				//nolint:staticcheck
				Tenant: s.Tenant{Key: "1"},
				Paths:  map[string]string{s.DEVICE: "device"},
			},
			expErr: "no default printer",
		},
		{
			name:    "success",
			printer: &d.Printer{IPAddress: "123", PrinterSettings: d.PrinterSettings{PrinterType: d.TSPPrinterType}},
			data: &s.ContextData{
				//nolint:staticcheck
				Tenant: s.Tenant{Key: "1"},
				Paths:  map[string]string{s.DEVICE: "device"},
			},
			payload: &model.Payload{IPAddress: "123", PrintPayload: strCmdrs, DescribeMessage: "[Printing Job] Table Bill"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			deviceMS := MockDeviceMS{Printer: tc.printer, Error: tc.printerError}
			controller := NewPrintController(deviceMS)

			payload, err := controller.TableBill(tc.bill, tc.data)
			if tc.name == "success" {
				require.NoError(t, err)
				require.Equal(t, tc.payload, payload)
			} else {
				require.EqualError(t, err, tc.expErr)
			}
		})
	}
}
