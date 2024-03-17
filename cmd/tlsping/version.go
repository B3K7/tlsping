package main

import (
	"os"
)

//  printVersion prints the version information about this application
func printVersion(f *os.File) {
	const template = `
{{.AppName}} {{.AppVersion}}

Compiler:
{{.Tab1}}{{.GoVersion}} ({{.Os}}/{{.Arch}})

Built on:
{{.Tab1}}{{.BuildTime}}

Author:
{{.Tab1}}Brent Kimberley   (C) 2024
{{.Tab1}}Fabio Hernandez, 2016-2020
{{.Tab1}}IN2P3/CNRS computing center, Lyon (France)

Source code and documentation:
{{.Tab1}}https://github.com/B3K7/{{.AppName}}
`
	render(template, tmplFields, f)
}
