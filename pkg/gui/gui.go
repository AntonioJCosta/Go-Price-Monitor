// package to create the GUI for the user to interact with the program
package gui

import (
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/utils"
)

// UserGUI is the main GUI for the user to interact with the program
func UserGUI() UserInput {

	a := app.New()
	w := a.NewWindow("Price Monitor using Go")
	w.SetPadded(true)
	w.Resize(fyne.NewSize(400, 300))
	w.CenterOnScreen()
	// In case user close the window (X button) the program will exit
	w.SetOnClosed(func() {
		a.Quit()
	})

	product := widget.NewEntry()
	tel := widget.NewEntry()
	value := widget.NewEntry()

	form := &widget.Form{
		BaseWidget: widget.BaseWidget{
			Hidden: false,
		},
		Items: []*widget.FormItem{{
			Text:     "Enter the product name",
			Widget:   product,
			HintText: "Ex: Microphone",
		}, {
			Text:     "Enter with your phone number",
			Widget:   tel,
			HintText: "Ex: 5519992685736",
		}, {
			Text:     "Enter the value",
			Widget:   value,
			HintText: "Ex: 1000",
		}},
		OnSubmit: func() {
			// Check if any field is empty
			isFieldsEmpty := product.Text == "" || tel.Text == "" || value.Text == ""
			if isFieldsEmpty {
				dialog.ShowError(fmt.Errorf("Please fill all the fields"), w)
				return
			}
			// Check if telephone is a number
			_, err := strconv.ParseInt(tel.Text, 10, 64)
			if err != nil {
				dialog.ShowError(fmt.Errorf("Please enter a valid telephone number"), w)
				return
			}
			// Check if the telephone has more than 8 digits
			isTelephoneValid := len(tel.Text) > 8
			if !isTelephoneValid {
				dialog.ShowError(fmt.Errorf("Please, the telephone number must have at least 9 digits"), w)
				return
			}
			// Check if value is a number
			val, err := utils.Float64Converter(value.Text)
			if err != nil {
				dialog.ShowError(fmt.Errorf("Please enter a valid value"), w)
				return
			}
			// Check if value is greater than 0
			if val <= 0 {
				dialog.ShowError(fmt.Errorf("Please enter a value greater than 0"), w)
				return
			}
			// If all checks are passed, the window will close and the program will continue
			w.Close()
		},
		OnCancel: func() {
			dialog.ShowConfirm("Exit", "Are you sure you want to exit?", func(b bool) {
				if b {
					os.Exit(0)
				}
			}, w)
		},
		SubmitText: "Submit",
		CancelText: "Cancel",
	}

	w.SetContent(form)
	w.ShowAndRun()
	val, _ := utils.Float64Converter(value.Text)

	return UserInput{
		Product:   product.Text,
		Telephone: tel.Text,
		Value:     val,
	}

}

func ErrorPopup(errMsg string) {
	a := app.New()
	w := a.NewWindow("Price Monitor using Go")
	w.SetPadded(true)
	w.Resize(fyne.NewSize(400, 400))
	w.CenterOnScreen()
	// In case user close the window (X button) the program will exit
	w.SetOnClosed(func() {
		a.Quit()
	})

	dialog.ShowError(fmt.Errorf(errMsg), w)
}
