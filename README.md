# Golang WebAssembly Calculator

Very basic calculator for learn go and webAssembly.
 
 Work in progress.

## Setup and run the project

`git clone https://github.com/Leos1113/golang-wasm-calculator.git`

`cd golang-wasm-calculator`

`cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`

`GOARCH=wasm GOOS=js go build -o lib.wasm main.go`

`go run server.go`
