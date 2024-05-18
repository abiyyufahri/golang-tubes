package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app  *tview.Application
	flex *tview.Flex
	menu *tview.List
)

func main() {
	// Create a new application
	app = tview.NewApplication()

	// Create a new flex layout
	flex = tview.NewFlex().SetDirection(tview.FlexRow)

	// Add a title to the menu
	title := tview.NewTextView().
		SetText("Welcome to My App").
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tview.Styles.TitleColor)

	// Create a list of menu items
	menu = tview.NewList().
		AddItem("Option 1", "", '1', func() {
			// Upon selecting Option 1, display a table with video learning data
			showVideoLearningData()
		}).
		AddItem("Option 2", "", '2', nil).
		AddItem("Option 3", "", '3', nil).
		AddItem("Exit", "", 'q', func() {
			app.Stop()
		})

	// Add the title and menu to the flex layout
	flex.AddItem(title, 0, 1, false).
		AddItem(menu, 0, 2, true)

	// Set the root layout for the application
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

// Function to show video learning data
func showVideoLearningData() {
	// Create a new table
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false) // Make the table rows selectable but not the columns

	// Add table rows (for demonstration purposes)
	table.SetCell(0, 0, tview.NewTableCell("Video Title").SetAlign(tview.AlignCenter).SetTextColor(tview.Styles.SecondaryTextColor))
	table.SetCell(0, 1, tview.NewTableCell("Duration").SetAlign(tview.AlignCenter).SetTextColor(tview.Styles.SecondaryTextColor))
	table.SetCell(1, 0, tview.NewTableCell("Introduction to Go").SetAlign(tview.AlignCenter))
	table.SetCell(1, 1, tview.NewTableCell("30 minutes").SetAlign(tview.AlignCenter))
	table.SetCell(2, 0, tview.NewTableCell("Advanced Topics in Go").SetAlign(tview.AlignCenter))
	table.SetCell(2, 1, tview.NewTableCell("45 minutes").SetAlign(tview.AlignCenter))

	// Handle table selection
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Get the currently selected row and column
		row, column := table.GetSelection()
		if event.Key() == tcell.KeyEnter {
			// Handle the Enter key
			cell := table.GetCell(row, column)
			// Display the selected cell's text
			app.SetRoot(tview.NewTextView().SetText(fmt.Sprintf("You selected: %s", cell.Text)), true)
		}
		return event
	})

	// Clear the flex layout and add the table
	flex.Clear()
	flex.AddItem(table, 0, 1, true)

	// Refresh the application's root
	app.SetRoot(flex, true).SetFocus(table)
}
