package ports

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/onerciller/portm/command"
	"github.com/rivo/tview"
)

type Ports struct {
	command command.Command
	table   *tview.Table
	ports   []*command.Port
}

func New(command command.Command) *Ports {
	return &Ports{
		command: command,
		table:   tview.NewTable(),
	}
}

func (ui *Ports) GetPorts() []*command.Port {
	return ui.ports
}

func (ui *Ports) Render() *tview.Table {
	ui.table.Clear()
	ui.ports = ui.command.ExecutePortList()

	ui.table.SetBorder(true).
		SetBorderColor(tcell.ColorSteelBlue).
		SetTitle("PORT MANAGER")

	ui.table.SetSelectable(true, false)
	ui.table.SetBorderPadding(1, 0, 1, 1)

	ui.renderHeader()
	ui.renderRows(ui.ports)

	return ui.table
}

func (ui *Ports) SetSelectedFunc(handler func(row, column int)) {
	ui.table.SetSelectedFunc(handler)
}

func (ui *Ports) renderHeader() {
	headerCols := []string{"NAME", "PID", "PORT", "NODE", "FD", "TYPE"}
	for key, colName := range headerCols {
		ui.table.SetCell(0, key, tview.NewTableCell(fmt.Sprintf("[black::b]%s", strings.ToUpper(colName))).
			SetTextColor(tcell.ColorWhite).
			SetExpansion(1).
			SetBackgroundColor(tcell.ColorSteelBlue).
			SetAttributes(tcell.AttrBold).SetSelectable(false),
		)
	}
}

func (ui *Ports) renderRows(ports []*command.Port) {
	for r, p := range ports {
		ui.table.SetCell(r+1, 0, tview.NewTableCell(p.Command).SetAlign(tview.AlignLeft))
		ui.table.SetCell(r+1, 1, tview.NewTableCell(p.PID).SetAlign(tview.AlignLeft))
		ui.table.SetCell(r+1, 2, tview.NewTableCell(p.Port).SetAlign(tview.AlignLeft).SetAttributes(tcell.AttrBold))
		ui.table.SetCell(r+1, 3, tview.NewTableCell(p.Node).SetAlign(tview.AlignLeft))
		ui.table.SetCell(r+1, 4, tview.NewTableCell(p.Fd).SetAlign(tview.AlignLeft))
		ui.table.SetCell(r+1, 5, tview.NewTableCell(p.PType).SetAlign(tview.AlignLeft))
	}
}
