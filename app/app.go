/*
BarCode scanner app
*/
package app

import (
	"barcode-keyboard-emulator/app/keyboard"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// GET "/" handler
// Display main page
func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	// Load template app/template.go MainPage variable
	var data = "Main page"
	tmpl := template.Must(template.New(data).Parse(MainPage))
	tmpl.Execute(w, data)
}

// POST "/code" handler
// Parse form and send code to system (keyboard emulate)
func CodeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// "code" form field processing
	if code := r.FormValue("code"); code != "" {
		keyboard.SendString(code, true)
		log.Printf("Received code '%s'\r\n", code)
	}
	// "c_code" form field processing
	if c_code := r.FormValue("c_code"); c_code != "" {
		keyboard.SendString(c_code, false)
		log.Printf("Received code '%s'\r\n", c_code)
	}
	// Redirect 301
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Run(version, addr, port string) (err error) {
	// Main page
	http.HandleFunc("/", HomeRouterHandler)
	// POST form
	http.HandleFunc("/code", CodeRouterHandler)
	// Start server
	addrString := fmt.Sprintf("%s:%s", addr, port)
	motd(version, addr, port)
	err = http.ListenAndServe(addrString, nil)
	return
}

// Start message
func motd(version, addr, port string) {
	if addr == "" {
		addr = "*Any_host_ip_address*"
	}
	text := "BarCode keyboard emulator v%s\r\nUse URL: http://%s:%s\r\n"
	fmt.Printf(text, version, addr, port)
}
