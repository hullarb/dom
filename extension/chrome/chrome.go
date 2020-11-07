package chrome

import "github.com/hullarb/dom/js"

var (
	chrome  = js.Get("chrome")
	runtime = chrome.Get("runtime")
)

type WindowID int

const CurrentWindow = WindowID(0)

func lastError() error {
	e := runtime.Get("lastError")
	if !e.Valid() {
		return nil
	}
	return js.Error{Value: e.JSValue()}
}
