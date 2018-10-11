package main

import (
	"github.com/rivo/tview"
)

func main() {
	grid := tview.NewGrid().
		SetRows(0, 0).
		SetColumns(0, 0).
		AddItem(tview.NewBox().SetBorder(true).SetTitleAlign(tview.AlignLeft).SetTitle("instance panel"), 0, 0, 1, 3, 0, 0, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitleAlign(tview.AlignLeft).SetTitle("SG panel"), 1, 0, 1, 3, 0, 0, false)

	//Layout for screens narrower than 100 cells (menu and side bar are hidden).
	if err := tview.NewApplication().SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
