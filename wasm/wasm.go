package main

import (
	"fmt"
	"syscall/js"
)

// this will return a js function that will be exported as part of the program
func greeter(this js.Value, args []js.Value) interface{} {
	fmt.Println("Hello " + args[0].String())
	return nil
}

// a more complicated example that adds an event listener onto any element
func eventListener(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		errorConstructor := js.Global().Get("Error")
		errorObject := errorConstructor.New("Missing argument")
		return errorObject
	}

	document := js.Global().Get("document")

	element := document.Call("getElementById", args[0].String())
	element.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("Event listener called!")
		return nil
	}))

	return nil
}

func main() {
	js.Global().Set("greeter", js.FuncOf(greeter))
	js.Global().Set("eventListener", js.FuncOf(eventListener))
	<-make(chan bool)
}
