package tsp

import (
	"fmt"
	"strings"

	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/model"
	"github.com/rockspoon/rs.cor.printer-ms/template"
)

const (
	serverLineFormat      = "Server: %-25s\n"
	stationLineFormat     = "Station: %-25s\n"
	timeLineFormat        = "%-12s    %22s\n"
	invoiceLineFormat     = "Order: %31s\n"
	invoiceLine2Format    = "Table: %31s\n"
	deliveryAddressFormat = "Address: %-25s\n"
	itemLine1Format       = "%3dx %-23s%10s\n"
	itemLine2Format       = "    %s\n"
	dateFormat            = "Jan 2, 2006"
	timeFormat            = "03:04:05 PM"
)

// ReceiptGenerator generates a receipt for tsp printer
type ReceiptGenerator struct {
}

// Generate converts a model into a command list
func (gen ReceiptGenerator) Generate(receipt model.KitchenReceipt) []command.PrinterCommand {
	cmdrs := make([]command.PrinterCommand, 0)
	cmdrs = append(cmdrs,
		command.NewLine{},
		command.NewLine{},
		command.Text(fmt.Sprintf(serverLineFormat, receipt.Server)),
	)

	if !receipt.IsPrintedForRunner {
		cmdrs = append(cmdrs, command.Text(fmt.Sprintf(stationLineFormat, receipt.Station)))
	}
	cmdrs = append(cmdrs,
		command.Text(fmt.Sprintf(timeLineFormat, receipt.Timestamp.Format(dateFormat), receipt.Timestamp.Format(timeFormat))),
		command.Text(fmt.Sprintf(invoiceLineFormat, "#"+receipt.InvoiceNumber)),
		command.Text(fmt.Sprintf(invoiceLine2Format, gen.getTableNumber(receipt))),
	)

	if receipt.DeliveryAddress != "" {
		cmdrs = append(cmdrs, command.Text(fmt.Sprintf(deliveryAddressFormat, receipt.DeliveryAddress)))
	}
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = append(cmdrs,
		command.Text(fmt.Sprintf("%38s\n", receipt.FireType)),
		command.Text("QTY  Item                 Seat      \n"),
	)

	for _, item := range receipt.Items {
		itemName := item.Name
		if len(itemName) > 23 {
			itemName = itemName[:20] + "..."
		}
		var builder strings.Builder
		if len(item.SeatNumber) == 1 {
			if !item.IsAllSeats {
				builder.WriteString(item.SeatNumber[0])
			}
		} else {
			if !item.IsAllSeats {
				if item.IsSplit {
					builder.WriteString("Split ")
				}
				for i := 0; i < len(item.SeatNumber); i++ {
					if i > 0 {
						builder.WriteString(",")
					}
					builder.WriteString(item.SeatNumber[i])
				}
			} else if item.IsSplit {
				builder.WriteString("Split ")
			}
		}
		cmdrs = append(cmdrs, command.Text(fmt.Sprintf(
			itemLine1Format,
			item.Quantity,
			itemName,
			builder.String(),
		)))
		if item.Modifiers != "" {
			cmdrs = append(cmdrs, command.Text(fmt.Sprintf(
				itemLine2Format,
				item.Modifiers,
			)))
		}
	}
	cmdrs = template.LineSeparator(cmdrs)
	cmdrs = template.Footer(cmdrs)
	cmdrs = append(cmdrs, command.Cut{})
	return cmdrs
}

func (ReceiptGenerator) getTableNumber(receipt model.KitchenReceipt) string {
	if receipt.TableNumber != "" {
		return receipt.TableNumber
	}
	return receipt.OrderType
}
