package keyboard

import (
	"github.com/go-vgo/robotgo"
)

/*
func SendString
@data (string) - barcode string
@tapeEnter (bool) - send "Enter" char code after sending barcode sting
*/
func SendString(data string, tapEnter bool) {
	robotgo.TypeStr(data,5)
	if tapEnter {
		robotgo.KeyTap("enter")
	}
}
