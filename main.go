//go:build wasip1

// Social Icons is a goblog WebAssembly plugin: it renders a row of
// Font Awesome profile links in the footer of every page, and exposes the
// same links to templates as .links for themes that want to place them.
//
// Build:  GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -ldflags="-s -w" -o plugin.wasm .
// Every export takes JSON on stdin (pdk.Input) and returns JSON or HTML
// (pdk.Output). See goblog.live/docs/plugin-api for the contract.
package main

import (
	"encoding/json"

	pdk "github.com/extism/go-pdk"
)

// hookInput is the ctx goblog passes to template hooks; only settings are
// needed here.
type hookInput struct {
	Settings map[string]string `json:"settings"`
}

func outputJSON(v any) int32 {
	if err := pdk.OutputJSON(v); err != nil {
		pdk.SetErrorString("encode output: " + err.Error())
		return 1
	}
	return 0
}

func readHook(name string) (hookInput, bool) {
	var in hookInput
	if err := json.Unmarshal(pdk.Input(), &in); err != nil {
		pdk.SetErrorString(name + ": " + err.Error())
		return in, false
	}
	return in, true
}

//go:wasmexport identity
func identity() int32 {
	return outputJSON(map[string]string{"name": "socialicons", "display_name": "Social Icons", "version": "2.0.0"})
}

//go:wasmexport settings
func settings() int32 {
	return outputJSON(settingDefs())
}

//go:wasmexport template_footer
func templateFooter() int32 {
	in, ok := readHook("template_footer")
	if !ok {
		return 1
	}
	pdk.OutputString(footerHTML(in.Settings))
	return 0
}

//go:wasmexport template_data
func templateData() int32 {
	in, ok := readHook("template_data")
	if !ok {
		return 1
	}
	ls := links(in.Settings)
	if ls == nil {
		return 0
	}
	return outputJSON(map[string]any{"links": ls})
}

func main() {}
