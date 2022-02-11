package modal

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/onerciller/portm/command"
	"github.com/rivo/tview"
)

type PortKillConfirmModal struct {
	modal *tview.Modal
	port  string
}

func NewPortKillConfirmModal() *PortKillConfirmModal {
	return &PortKillConfirmModal{
		modal: tview.NewModal(),
	}
}

func (ui *PortKillConfirmModal) Render() *tview.Modal {
	ui.modal.ClearButtons()
	return ui.modal.
		SetBackgroundColor(tcell.ColorSteelBlue).
		SetText(fmt.Sprintf("Do you want to kill %s port?", ui.port)).
		AddButtons([]string{"Kill", "Cancel"})
}

func (ui *PortKillConfirmModal) SetPort(selectedPort *command.Port) {
	if selectedPort != nil {
		ui.port = selectedPort.Port
	}
}

func (ui *PortKillConfirmModal) SetDoneFunc(handler func(buttonIndex int, buttonLabel string)) {
	ui.modal.SetDoneFunc(handler)
}
