package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/onerciller/portm/ui/modal"
	"github.com/onerciller/portm/ui/ports"

	"github.com/onerciller/portm/command"
	"github.com/rivo/tview"
)

type UI struct {
	command                    command.Command
	app                        *tview.Application
	isPortKillConfirmModalOpen bool
	selectedPort               *command.Port
	portKillConfirmModal       *modal.PortKillConfirmModal
	ports                      *ports.Ports
}

func New(command command.Command) UI {
	return UI{
		command:                    command,
		app:                        tview.NewApplication(),
		isPortKillConfirmModalOpen: false,
		portKillConfirmModal:       modal.NewPortKillConfirmModal(),
		ports:                      ports.New(command),
	}
}

func (ui *UI) onConfirmSelected(buttonIndex int, buttonLabel string) {
	if buttonLabel == "Kill" {
		ui.command.ExecutePortKill(ui.selectedPort)
		ui.isPortKillConfirmModalOpen = false
		ui.Render()
		return
	}

	if buttonLabel == "Cancel" {
		ui.isPortKillConfirmModalOpen = false
		ui.Render()
		return
	}
}

func (ui *UI) onPortSelected(row int, column int) {
	ui.selectedPort = ui.ports.GetPorts()[row-1]
	ui.isPortKillConfirmModalOpen = true
	ui.Render()
}

func (ui *UI) Render() {

	portsTable := ui.ports
	portsTable.SetSelectedFunc(ui.onPortSelected)
	renderedPortsTable := portsTable.Render()

	confirmModal := ui.portKillConfirmModal
	confirmModal.SetDoneFunc(ui.onConfirmSelected)
	confirmModal.SetPort(ui.selectedPort)
	renderedConfirmModal := confirmModal.Render()

	menu := ui.menu()

	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	if ui.isPortKillConfirmModalOpen {
		flex.AddItem(renderedPortsTable, 0, 16, false)
		flex.AddItem(renderedConfirmModal, 0, 0, true)
	} else {
		flex.AddItem(renderedPortsTable, 0, 16, true)
	}

	flex.AddItem(menu, 0, 1, false)

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			ui.app.Stop()
		}
		return event
	})

	if err := ui.app.SetRoot(flex, true).
		EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func (ui *UI) menu() *tview.TextView {
	menuItemMap := map[string]string{
		"C-c | Esc": "Quit",
		"Up":        "Go to Up",
		"Down":      "Go to Down",
		"Enter":     "Select Port",
	}

	mText := ""
	for key, description := range menuItemMap {

		mText += fmt.Sprintf("[%s:%s:b] <%s>", GetColorName(tcell.ColorWhite), GetColorName(tcell.ColorBlack), key)
		mText += fmt.Sprintf("[%s:%s:b] %s", GetColorName(tcell.ColorBlack), GetColorName(tcell.ColorSteelBlue), description)

		mText += "  "
	}

	menu := tview.NewTextView()
	menu.
		SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetText(mText).
		SetBorder(false).
		SetBackgroundColor(tcell.ColorBlack)
	return menu
}

func GetColorName(color tcell.Color) string {
	for name, c := range tcell.ColorNames {
		if c == color {
			return name
		}
	}
	return ""
}
