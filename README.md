# RAW Image EXIF Viewer (WebAssembly + Go)

This is a privacy-focused web application that allows users to select a RAW image from their local file system and view its **EXIF metadata** directly in the browser. It uses **Go + WebAssembly (WASM)** to parse the image, with all processing happening locally — **no data is uploaded** anywhere.

---

## Key Features

- **Select RAW images** directly from your device
- **Privacy-first**: All image processing happens **completely offline** in your browser
- Powered by **Go + WebAssembly**

---

## Supported RAW Formats

Currently, the following RAW file formats are supported:

- `.cr2` (Canon)
- `.nef` (Nikon)
- `.arw` (Sony)
- `.dng` (Adobe)
- `.rw2` (Panasonic)
- `.raf` (Fujifilm)

---

## How It Works (Technical Overview)

### WebAssembly + Go

- The core EXIF decoding logic is written in **Go**, using [`goexif`](https://github.com/rwcarlsen/goexif) to parse EXIF metadata.
- The Go code is compiled to **WebAssembly** (`main.wasm`) using:
  ```bash
  GOOS=js GOARCH=wasm go build -o main.wasm main.go
  ```

### JS ↔ Go Interop

- In the browser, `wasm_exec.js` (provided by Go) bridges JavaScript and WebAssembly.
- A `process(arrayBuffer)` function is exposed to JS via `js.FuncOf()` in Go.
- JavaScript reads the selected image using `FileReader.readAsArrayBuffer()`, and the byte array is passed into Go using:
  ```js
  js.CopyBytesToGo(imgBytes, p[0])
  ```

### Data Flow

1. User selects a file → JS reads it as `ArrayBuffer`
2. JS passes the bytes to WASM-Go via `process()`
3. Go decodes EXIF using `goexif.Decode` and returns JSON
4. JS parses JSON and renders a styled table

