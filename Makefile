.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o wasm/out.wasm wasm/wasm.go
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" wasm
	python -m http.server --directory=wasm