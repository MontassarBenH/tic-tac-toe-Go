package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var (
	player     = 1
	markings   = [9]string{"", "", "", "", "", "", "", "", ""}
	winMessage = ""
)

func main() {
	a := app.New()
	w := a.NewWindow("Tic-Tac-Toe")
	w.Resize(fyne.NewSize(300, 300))

	// Create the grid of buttons
	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(3),
		newButton(0), newButton(1), newButton(2),
		newButton(3), newButton(4), newButton(5),
		newButton(6), newButton(7), newButton(8),
	)
	w.SetContent(grid)

	w.ShowAndRun()
}

func newButton(i int) *widget.Button {
	button := widget.NewButton("", func() {
		// Handle button click
		if markings[i] == "" {
			if player == 1 {
				markings[i] = "X"
			} else {
				markings[i] = "O"
			}
			button.SetText(markings[i])
			checkWin()
			player *= -1
		}
	})
	button.Resize(fyne.NewSize(100, 100))
	button.SetIcon(canvas.NewText(markings[i], theme.TextColor()))
	return button
}

func checkWin() {
	// Check for horizontal wins
	for i := 0; i < 9; i += 3 {
		if markings[i] != "" && markings[i] == markings[i+1] && markings[i] == markings[i+2] {
			winMessage = "Player " + markings[i] + " wins!"
		}
	}

	// Check for vertical wins
	for i := 0; i < 3; i++ {
		if markings[i] != "" && markings[i] == markings[i+3] && markings[i] == markings[i+6] {
			winMessage = "Player " + markings[i] + " wins!"
		}
	}

	// Check for diagonal wins
	if markings[0] != "" && markings[0] == markings[4] && markings[0] == markings[8] {
		winMessage = "Player " + markings[0] + " wins!"
	}
	if markings[2] != "" && markings[2] == markings[4] && markings[2] == markings[6] {
		winMessage = "Player " + markings[2] + " wins!"
	}

	if winMessage != "" {
		// Show win message
		dialog := dialog.NewInformation("Game Over", winMessage, nil)
		dialog.Show()
	}
}
