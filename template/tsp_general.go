package template

import (
	"fmt"
	"sort"
	"time"

	"github.com/rockspoon/rs.cor.printer-ms/command"
	"github.com/rockspoon/rs.cor.printer-ms/helper"
	"github.com/rockspoon/rs.cor.printer-ms/model"
)

const (
	dateFormat            = "Jan 2, 2006               03:04:05 PM"
	billItemFormat        = "\n%2dx %-16s %8.2f %8.2f\n"
	itemLine2Format       = "    %-16s\n"
	billSubEntryFormat    = "    * %-14s %+8.2f %+8.2f\n"
	subEntryLine2Format   = "      %-14s\n"
	totalsFormat          = "%-25s %-3s %8.2f\n"
	kitchenItemFormat     = "\n%2dx %-16s  %16s\n"
	kitchenSubEntryFormat = "    * %-14s\n"
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
	sort.Sort(check.Charges)
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

// AddItemsBill add items
func AddItemsBill(entryItems []model.EntryItem, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.Text("                         Unity   Final\n"))
	commands = append(commands, command.Text("QTY Item                 Price   Price\n"))

	for _, item := range entryItems {
		itemName := helper.WarpString(item.Name, 16)
		line := fmt.Sprintf(billItemFormat, item.Quantity, itemName[0], item.UnityPrice.Price, item.FinalPrice.Price)
		commands = append(commands, command.Text(line))
		for i := range itemName[1:] {
			line = fmt.Sprintf(itemLine2Format, itemName[i+1])
			commands = append(commands, command.Text(line))
		}

		sort.Sort(item.SubEntries)

		for _, subEntry := range item.SubEntries {
			title := subEntry.Name
			if subEntry.Description != "" {
				title = title + "(" + subEntry.Description + ")"
			}
			subEntryName := helper.WarpString(title, 16)
			line := fmt.Sprintf(billSubEntryFormat, subEntryName[0], subEntry.UnityPrice.Price, subEntry.FinalPrice.Price)
			commands = append(commands, command.Text(line))
			for i := range subEntryName[1:] {
				line = fmt.Sprintf(subEntryLine2Format, subEntryName[i+1])
				commands = append(commands, command.Text(line))
			}
		}
	}
	return commands
}

// AddServiceInfoBill adds Table Information
func AddServiceInfoBill(attendantName string, orderType model.TypesOfOrder, createdAt time.Time, check model.Check, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands,
		command.Text(createdAt.Format(dateFormat)),
		command.NewLine{},
		command.Text("Attendant: "+attendantName),
		command.NewLine{},
		command.NewLine{},
		command.Text("Order Type: "+model.TypesOfOrderMap[orderType]),
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

// AddServiceInfoKitchen does something
func AddServiceInfoKitchen(receipt model.KitchenReceipt, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands,
		command.Text(receipt.CreatedAt.Format(dateFormat)),
		command.NewLine{},
		command.Text("Kitchen: "+receipt.Kitchen),
		command.NewLine{},
		command.Text("Order Type: "+model.TypesOfOrderMap[receipt.OrderType]),
		command.NewLine{},
		command.NewLine{},
	)

	if receipt.OrderType == model.TypesOfOrderDinein {
		commands = append(commands,
			// command.Bold{Enabled: true},
			command.Text("Runner: "+receipt.DineInInfo.RunnerName),
			command.NewLine{},
			command.Text("Section: "+receipt.DineInInfo.SectionName),
			command.NewLine{},
			command.Text("Tables: "+receipt.DineInInfo.Tables),
			command.NewLine{},
		)
	} else {
		commands = append(commands,
			command.Text("Customer: "+receipt.CustomerInfo.Name),
			command.NewLine{},
			command.Text("Phone number: "+receipt.CustomerInfo.Phone),
			command.NewLine{},
		)
		if receipt.CustomerInfo.Address != nil {
			commands = append(commands, command.Text("Address: "+receipt.CustomerInfo.Address.Address1), command.NewLine{})
		}
	}

	return commands
}

// AddItemsKitchen does other thing
func AddItemsKitchen(kitchenItems []model.KitchenItem, commands []command.PrinterCommand) []command.PrinterCommand {
	commands = append(commands, command.Text("QTY Item                     Fire Type\n"))

	for _, item := range kitchenItems {
		itemName := helper.WarpString(item.Name, 16)
		line := fmt.Sprintf(kitchenItemFormat, item.Quantity, itemName[0], model.TypesOfFireMap[item.FireType])
		commands = append(commands, command.Text(line))
		for i := range itemName[1:] {
			line = fmt.Sprintf(itemLine2Format, itemName[i+1])
			commands = append(commands, command.Text(line))
		}

		sort.Sort(item.SubEntries)
		for _, subEntry := range item.SubEntries {
			title := subEntry.Name
			if subEntry.Description != "" {
				title = title + "(" + subEntry.Description + ")"
			}
			subEntryName := helper.WarpString(title, 16)
			line := fmt.Sprintf(kitchenSubEntryFormat, subEntryName[0])
			commands = append(commands, command.Text(line))
			for i := range subEntryName[1:] {
				line = fmt.Sprintf(subEntryLine2Format, subEntryName[i+1])
				commands = append(commands, command.Text(line))
			}
		}

		commands = append(commands, command.Text(" Seats: "+item.Seats))
		commands = append(commands, command.Text("\n\n---\n"))
	}
	commands = commands[:len(commands)-1]
	return commands
}
