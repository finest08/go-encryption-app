package gui

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "github.com/finest08/go-encryption-app/gui"
)

func FileSelWindow() {
	// New App
	a := app.New()
	// New window and title
	w := a.NewWindow("Go Encryption Program")
	// Resize window
	w.Resize(fyne.NewSize(500, 400))
	// title of our form
	title := canvas.NewText("Go Encrytion Program", color.Black)
	title.TextSize = 20 // text font size is 20
	title.TextStyle = fyne.TextStyle{Bold: true}
	// I need it bold.. you can make it italic
	title.Resize(fyne.NewSize(300, 35)) // 300 is width & 35 is height
	title.Move(fyne.NewPos(80, 30))
	//position my widget
	//50 px from left and 10 px from top
	// copy / the setting(resize/ postion)
	// Name field
	secret := widget.NewEntry()
	secret.SetPlaceHolder("Secret Key")
	secret.Resize(fyne.NewSize(300, 38))
	secret.Move(fyne.NewPos(50, 100))
	// copy / paste for next widget email
	file := widget.NewEntry()
	file.SetPlaceHolder("File")
	file.Resize(fyne.NewSize(300, 38))
	file.Move(fyne.NewPos(50, 150))
	submit_btn := widget.NewButton("SUBMIT", nil)
	submit_btn.Resize(fyne.NewSize(110, 60))
	// button need to be small as compared to textarea
	submit_btn.Move(fyne.NewPos(140, 220))

	// setup content
	// we are going to use container without layout// my favorite
	w.SetContent(
		container.NewWithoutLayout(
			title,
			secret,
			file,
			submit_btn,
		),
	)
	// Show and run
	w.ShowAndRun()

}

func GenKeyWindow() {

	a := app.New()
	w := a.NewWindow("Generate Key")
	w.Resize(fyne.NewSize(500, 400))

	text4 := canvas.NewText("centered", color.Black)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	w.SetContent(container.New(layout.NewVBoxLayout(), centered))
	w.ShowAndRun()

}

func Form() {
	a := app.New()
	w := a.NewWindow("Grid Layout")
	w.Resize(fyne.NewSize(500, 400))
	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Secret", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
			w.Close()
		},
	}

	// we can also append items
	form.Append("Text", textArea)

	w.SetContent(form)
	w.ShowAndRun()
}
