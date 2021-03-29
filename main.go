package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	SetCorrectTheme()

	a := app.New()
	w := a.NewWindow("AutoShutdown")
	entry := widget.NewEntry()
	list := []string{
		"min",
		"hour",
	}

	sel := widget.NewSelectEntry(list)
	sel.SetText(list[0])

	entry.Validator = fyne.StringValidator(ValidateNumber)

	w.SetContent(container.NewVBox(
		container.NewHBox(
			widget.NewButton("Shutdown in", func() {
				i, err := strconv.Atoi(entry.Text)
				if err != nil || i == 0 {
					return
				}

				switch sel.Text {
				case "min":
					ShutdownIn(i)
				case "hour":
					ShutdownIn(i * 60)
				}

			}),
			widget.NewButton("Disable", func() {
				DisableShutdown()
			}),
		),
		container.NewGridWithColumns(2,
			entry,
			sel,
		),
	))
	w.ShowAndRun()
}

func ValidateNumber(s string) error {
	if _, err := strconv.Atoi(s); len(s) > 0 && err == nil{
		return nil
	}
	return errors.New("incorrect")
}

func SetCorrectTheme() {
	t := time.Now()
	h := t.Hour()

	if h >= 20 || h <= 6{
		os.Setenv("FYNE_THEME", "dark")
	} else {
		os.Setenv("FYNE_THEME", "light")
	}
}

func ShutdownIn(minutes int) error {
	inSec := minutes * 60
	str := strconv.Itoa(inSec)
	com := exec.Command("shutdown", "-s", "-t", str)
	return com.Run()
}

func DisableShutdown() error {
	com := exec.Command("shutdown", "/a")
	return com.Run()
}
