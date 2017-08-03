package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func main() {
	argsWithoutProg := os.Args[1:]

	templateContext := make(map[string]string)

	for _, arg := range argsWithoutProg {
		parsedArg := strings.Split(arg, "=")
		templateContext[parsedArg[0]] = parsedArg[1]
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	inputTemplate := string(bytes)

	compiledTemplate, err := template.New("").Parse(inputTemplate)
	if err != nil {
		panic(err)
	}
	compiledTemplate.Execute(os.Stdout, templateContext)
}
