package main

import (
	"image/color"
	"os"
	"time"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var GlobalTimeVar int = 111

func main() {
	a := app.NewWithID("com.timer.app")
	w := a.NewWindow("Timer")

	w.Resize(fyne.NewSize(600, 500))

	//welcome Section
	Welcome_text := canvas.NewText("Weclome To the Timer App!!", color.White)
	Welcome_text.Alignment = fyne.TextAlignCenter
	Welcome_text.TextStyle = fyne.TextStyle{Bold: true, Monospace: true}
	Welcome_text.TextSize = 30

	// Time Choose Section
	timerInt := 5
	GlobalTimeVar = timerInt
	timeValue := strconv.Itoa(timerInt)
	ChoseTimeText := canvas.NewText("", color.White)
	ChoseTimeText.TextSize = 50
	ChoseTimeText.TextStyle = fyne.TextStyle{Bold: true}
	ChoseTimeText.Text = (timeValue)
	ChoseTimeText.Refresh()

	text := canvas.NewText("MINUTES", color.White)
	text.TextSize = 25

	button_Decrease := widget.NewButton("     -     ", func() {
		if timerInt == 5 {
			return
		}
		timerInt = timerInt - 5
		GlobalTimeVar = timerInt
		timeValue := strconv.Itoa(timerInt)
		ChoseTimeText.Text = (timeValue)
		ChoseTimeText.Refresh()
	})

	button_Increase := widget.NewButton("     +     ", func() {
		if timerInt == 60 {
			return
		}
		timerInt = timerInt + 5
		GlobalTimeVar = timerInt
		timeValue := strconv.Itoa(timerInt)
		ChoseTimeText.Text = (timeValue)
		ChoseTimeText.Refresh()

	})

	//Timer Running Section
	HourText := canvas.NewText("25", color.White)
	HourText.TextStyle = fyne.TextStyle{Bold: true}
	HourText.TextSize = 100
	hourVariableText := strconv.Itoa(GlobalTimeVar)
	HourText.Text = (hourVariableText)
	HourText.Refresh()

	// //Image
	// img := canvas.NewImageFromFile("Image/ClockU.png")
	// img.FillMode = canvas.ImageFillContain

	timerRunningSeperator := canvas.NewText(":", color.White)
	timerRunningSeperator.TextStyle = fyne.TextStyle{Bold: true}
	timerRunningSeperator.TextSize = 50

	minutetext := canvas.NewText("60", color.White)
	minutetext.TextStyle = fyne.TextStyle{Bold: true}
	minutetext.TextSize = 100
	minutetextVariable := strconv.Itoa(60)
	minutetext.Text = (minutetextVariable)
	minutetext.Refresh()

	TimeRunningSection := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), HourText, timerRunningSeperator, minutetext, layout.NewSpacer())

	TimeRunningSection.Hide()

	TimeChoserSection := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button_Decrease, layout.NewSpacer(), ChoseTimeText, text, layout.NewSpacer(), button_Increase, layout.NewSpacer())

	button_Stop := widget.NewButton("Stop", func() {

	})

	button_Stop.Hide()

	button_Start := widget.NewButton("START", func() {
		button_Stop.Show()
		TimeChoserSection.Hide()
		TimeRunningSection.Show()

		for i := GlobalTimeVar - 1; i >= 0; i-- {

			GlobalTimeVar = i
			for j := 59; j >= 0; j-- {
				hourVariableText := strconv.Itoa(GlobalTimeVar)
				valueI := j
				minutetextVariable := strconv.Itoa(valueI)
				minutetext.Text = (minutetextVariable)
				minutetext.Refresh()

				HourText.Text = (hourVariableText)
				HourText.Refresh()
				SecondDelay()
			}
		}
		if GlobalTimeVar == 0 {

			TimeChoserSection.Show()
			TimeRunningSection.Hide()
			button_Stop.Hide()
			go showNotification(a)
			f, _ := os.Open("Hen.mp3")
			// decoding mp3 file
			// 3 outputs
			// stream , format and error
			streamer, format, _ := mp3.Decode(f)
			// activate speakers
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			// play
			speaker.Play(streamer)
			GlobalTimeVar = timerInt
			// looping
			// select {}
		}
	})

	vBOX := container.NewVBox(Welcome_text, widget.NewSeparator(), layout.NewSpacer(), TimeChoserSection, TimeRunningSection, layout.NewSpacer(), button_Start, button_Stop, layout.NewSpacer())
	w.SetContent(vBOX)
	w.ShowAndRun()
}

func showNotification(a fyne.App) {
	time.Sleep(time.Second * 1)
	a.SendNotification(fyne.NewNotification("Timer", "Your Timer Is Stoped!!!"))
}

func SecondDelay() {
	timeValue := 1
	for range time.Tick(time.Second * 1) {
		timeValue--
		if timeValue == 0 {
			return
		}
	}
}
