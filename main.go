package main

import (
	"strconv"
	"syscall/js"
)

var c chan bool

func add(this js.Value, i []js.Value) interface{} {
    document := js.Global().Get("document")
    value1 := document.Call("getElementById", "value1").Get("value").String()
    value2 := document.Call("getElementById", "value2").Get("value").String()
    int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)
    sum := int1 + int2
    document.Call("getElementById", "result").Set("innerText", sum)
    return nil
}

func substract(this js.Value, i []js.Value) interface{} {
    document := js.Global().Get("document")
    value1 := document.Call("getElementById", "value1").Get("value").String()
    value2 := document.Call("getElementById", "value2").Get("value").String()
    int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)
    substract := int1 - int2
    document.Call("getElementById", "result").Set("innerText", substract)
    return nil
}

func registerCallbacks() {
    js.Global().Set("add", js.FuncOf(add))
    js.Global().Set("substract", js.FuncOf(substract))
}

func main() {
    println("Go WebAssembly Initialized")
    registerCallbacks()

    <-c
}

func init() {
    c = make(chan bool)
}
