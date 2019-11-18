package template

import (
	"fmt"
	"time"

	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/helper"
	"github.com/rockspoon/rs.cor.printer-ms/model"
)

const (
	serverLineFormat    = "Attendant: %-25s\n"
	dateFormat          = "Jan 2, 2006               03:04:05 PM"
	itemLine1Format     = "\n%2dx %-16s %8.2f %8.2f\n"
	itemLine2Format     = "    %-16s\n"
	subEntryFormat      = "    * %-14s %+8.2f %+8.2f\n"
	subEntryLine2Format = "      %-14s\n"
	totalsFormat        = "%-25s %-3s %8.2f\n"
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

// AddCheckTotal adds pricing information
func AddCheckTotal(check model.Check, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.Text(fmt.Sprintf(totalsFormat, "Subtotal:", check.Subtotal.Symbol, check.Subtotal.Price)))
	for _, charge := range check.Charges {
		title := charge.Name
		if charge.Description != "" {
			title = title + "(" + charge.Description + ")"
		}
		chargeTitle := helper.WarpString(title, 27)
		line := fmt.Sprintf(totalsFormat, chargeTitle[0], check.Subtotal.Symbol, charge.FinalPrice.Price)
		commands = append(commands, command.Text(line))
		for i := range chargeTitle[1:] {
			line = fmt.Sprintf(subEntryLine2Format, chargeTitle[i+1])
			commands = append(commands, command.Text(line))
		}
	}
	commands = append(commands, command.Text(fmt.Sprintf(totalsFormat, "Total:", check.Subtotal.Symbol, check.Total.Price)))

	return commands
}

// AddItems add items
func AddItems(entryItems []model.EntryItem, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.Text("                         Unity   Final\n"))
	commands = append(commands, command.Text("QTY Item                 Price   Price\n"))

	for _, item := range entryItems {
		itemName := helper.WarpString(item.Name, 16)
		line := fmt.Sprintf(itemLine1Format, item.Quantity, itemName[0], item.UnityPrice.Price, item.FinalPrice.Price)
		commands = append(commands, command.Text(line))
		for i := range itemName[1:] {
			line = fmt.Sprintf(itemLine2Format, itemName[i+1])
			commands = append(commands, command.Text(line))
		}

		for _, subEntry := range item.SubEntries {
			title := subEntry.Name
			if subEntry.Description != "" {
				title = title + "(" + subEntry.Description + ")"
			}
			subEntryName := helper.WarpString(title, 16)
			line := fmt.Sprintf(subEntryFormat, subEntryName[0], subEntry.UnityPrice.Price, subEntry.FinalPrice.Price)
			commands = append(commands, command.Text(line))
			for i := range subEntryName[1:] {
				line = fmt.Sprintf(subEntryLine2Format, subEntryName[i+1])
				commands = append(commands, command.Text(line))
			}
		}
	}
	return commands
}

// AddServiceInfo adds Table Information
func AddServiceInfo(attendantName string, orderType model.TypesOfOrder, createdAt time.Time, check model.Check, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands,
		command.Text(createdAt.Format(dateFormat)),
		command.NewLine{},
		command.Text(fmt.Sprintf(serverLineFormat, attendantName)),
		command.NewLine{},
	)

	if orderType == model.TypesOfOrderDinein {
		commands = append(commands,
			// command.Bold{Enabled: true},
			command.Text("Section: "+check.DineInOptions.SectionName),
			command.NewLine{},
			command.Text("Tables: "+check.DineInOptions.Tables),
			command.NewLine{},
			command.Text("Seats: "+check.DineInOptions.Seats),
			command.NewLine{},
		)
	} else {
		commands = append(commands,
			command.Text("Customer: "+check.CustomerInfo.Name),
			command.NewLine{},
			command.Text("Phone number: "+check.CustomerInfo.Phone),
			command.NewLine{},
		)
		if check.CustomerInfo.Address != nil {
			commands = append(commands, command.Text("Address: "+check.CustomerInfo.Address.Address1), command.NewLine{})
		}
	}
	return commands
}
