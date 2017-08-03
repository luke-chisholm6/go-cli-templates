package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"errors"
)

func splitKV(keyValue string) (key string, value string, err error) {
	keyValueArray := strings.Split(keyValue, "=")

	if len(keyValueArray) != 2 {
		err = errors.New("bad formatting kv pair")
		return
	}

	key, value = keyValueArray[0], keyValueArray[1]

	return
}

func ExtractKeyValueFromArray(keyValueArray[]string) (map[string]string, error) {
	keyValueMap := make(map[string]string)

	for _, keyValue := range keyValueArray {
		key, value, err := splitKV(keyValue)
		if err != nil {
			return keyValueMap, err
		}
		keyValueMap[key] = value
	}

	return keyValueMap, nil
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	switch {
	case err != nil:
		panic(err)
	case len(bytes) <= 0:
		panic("no template context given")
	}
	inputTemplate := string(bytes)
	compiledTemplate, err := template.New("").Parse(inputTemplate)
	if err != nil {
		panic(err)
	}

	templateContext, err := ExtractKeyValueFromArray(os.Args[1:])
	if err != nil {
		panic(err)
	}


	compiledTemplate.Execute(os.Stdout, templateContext)
}
