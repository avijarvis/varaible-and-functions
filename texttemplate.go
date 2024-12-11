package main

import (
	"os"
	"text/template"
)

func main() {
	// Define a simple template string
	tmpl := "Hello, {{.Name}}! Welcome to {{.Place}}."

	// Create and parse the template
	t, err := template.New("greeting").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Data to fill the template
	data := struct {
		Name  string
		Place string
	}{
		Name:  "Avinash",
		Place: "GoLang",
	}

	// Execute the template and print to standard output
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
