package template

import (
	"fmt"
	"strconv"
	"time"

	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/helper"
	"github.com/rockspoon/rs.cor.printer-ms/model"
)

const (
	invoiceLineFormat    = "Order: %30s\n"
	dateFormat           = "Jan 2, 2006               03:04:05 PM"
	serverLineFormat     = "Server: %-25s\n"
	itemLine1Format      = "%3dx %-22s $ %6.2f\n"
	itemLine2Format      = "     %s\n"
	totalsFormat         = "%-27s $ %6.2f\n"
	discountTotalsFormat = "%-27s $ %6.2f\n"
)

// LineSeparator adds a horizontal line separator
func LineSeparator(commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands,
		command.NewLine{},
		command.FontB,
		command.DashedLine{Font: command.FontB},
		command.NewLine{},
	)
	return commands
}

// Footer adds Rockspoon footer
func Footer(commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.NewLine{},
		command.NewLine{},
		command.Text(helper.Center("Thank You!", " ", 38)),
		command.NewLine{},
		command.Text(helper.Center("Powered by Rockspoon", " ", 38)),
		command.NewLine{},
		command.Text(helper.Center("www.rockspoon.com", " ", 38)),
		command.NewLine{},
	)
	return commands
}

// AddRestaurantInfo adds Restaurant Information
func AddRestaurantInfo(info model.RestaurantInfo, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.NewLine{},
		command.NewLine{},
		command.Text(helper.Center(info.Name, " ", 38)),
		command.NewLine{},
		command.Text(helper.Center(info.Address.Address1, " ", 38)),
		command.NewLine{},
		command.Text(helper.Center(info.Address.City+" "+info.Address.ZipCode, " ", 38)),
		command.NewLine{},
		command.Text(helper.Center(info.Address.Region+" "+info.Address.Country, " ", 38)),
		command.NewLine{},
	)
	return commands
}

// AddInvoiceCheck adds pricing information
func AddInvoiceCheck(invoice model.InvoiceCheck, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.Text(fmt.Sprintf(totalsFormat, "Subtotal:", invoice.SubTotal.Price())))

	if invoice.DeliveryFeeAmount.Value > 0 {
		commands = append(commands, command.Text(fmt.Sprintf(totalsFormat, "Delivery:", invoice.DeliveryFeeAmount.Price())))
	}
	if invoice.DiscountAmount.Value > 0 {
		commands = append(commands, command.Text(fmt.Sprintf(discountTotalsFormat, "Discount:", invoice.DiscountAmount.Price()*float64(-1))))
	}
	if invoice.MandatoryGratuityAmount.Value > 0 {
		str := "Gratuity (" + strconv.FormatFloat(float64(invoice.MandatoryGratuityRate)*float64(100), 'f', 2, 32) + "):"
		commands = append(commands, command.Text(fmt.Sprintf(discountTotalsFormat, str, invoice.MandatoryGratuityAmount.Price())))
	}
	commands = append(commands, command.Text(fmt.Sprintf(totalsFormat,
		invoice.SalesTaxDescription, invoice.TaxAmount.Price())),
		command.Text(fmt.Sprintf(totalsFormat, "Total:", invoice.Total.Price())),
	)

	return commands
}

// AddItems add items
// func AddItems(entryItems []model.EntryItem, commands []command.PrinterCommand) []command.PrinterCommand {
// 	commands = append(commands, command.Text("                         Unity  Final \n"))
// 	commands = append(commands, command.Text("QTY Item                 Price  Price \n\n"))

// 	for i := range items {

// 	}
// }

// AddInvoiceItems add invoice items
func AddInvoiceItems(items []model.InvoiceItem, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.Text("QTY  Item                      Price\n\n"))
	for _, item := range items {
		itemName := item.ItemName
		if len(itemName) > 23 {
			itemName = itemName[:20] + "..."
		}

		amount := item.Amount.Price() * float64(item.Quantity)
		commands = append(commands, command.Text(fmt.Sprintf(itemLine1Format, item.Quantity, itemName, amount)))
		if item.Modifiers != "" {
			commands = append(commands, command.FontB)

			for i := 0; i*31 < len(item.Modifiers); i++ {
				rightBound := (i + 1) * 31
				if rightBound > len(item.Modifiers) {
					rightBound = len(item.Modifiers)
				}
				commands = append(commands, command.Text(fmt.Sprintf(itemLine2Format, item.Modifiers[i*31:rightBound])))
			}

		}
	}
	return commands
}

// AddTableInfo adds Table Information
func AddTableInfo(table model.TableInfo, invoiceNumber string, billTime time.Time, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.Text(fmt.Sprintf(invoiceLineFormat, "#"+invoiceNumber)),
		command.NewLine{},
		command.Text(billTime.Format(dateFormat)),
		command.NewLine{},
		command.Text(fmt.Sprintf(serverLineFormat, table.ServerName)),
		command.NewLine{},
	)
	if table.DiningPartyType == "dinein" {
		commands = append(commands,
			command.Bold{Enabled: true},
			command.Text("Table: "+table.TableNumber),
			command.NewLine{},
		)
	} else {
		commands = append(commands,
			command.Text("Customer: "+table.CustomerName),
			command.NewLine{},
			command.NewLine{},
			command.Text(helper.Center(table.DiningPartyType, " ", 38)),
			command.NewLine{},
		)
	}
	return commands
}
