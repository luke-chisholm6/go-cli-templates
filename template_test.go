package main

import (
	"bytes"
	"github.com/luke-chisholm6/go-cli-templates/readers"
	"reflect"
	"strings"
	"testing"
)

func TestGetTemplateContext_SliceWithInvalidStrings(t *testing.T) {
	kvSlice := []string{
		"test",
	}

	if _, err := getTemplateContext(kvSlice); err == nil {
		t.Error("Cannot split a string into a kv pair that is not in the format of k=v")
	}
}

func TestGetTemplateContext_SliceWithValidStrings(t *testing.T) {
	kvSlice := []string{
		"test=test",
		"key=value",
	}

	kvMap, err := getTemplateContext(kvSlice)
	if err != nil {
		t.Error("Cannot split a string into a kv pair that is not in the format of k=v")
	}

	kvMapComparison := map[string]string{
		"test": "test",
		"key":  "value",
	}
	if !reflect.DeepEqual(kvMap, kvMapComparison) {
		t.Errorf("Expected: %+v\nGot: %+v", kvMapComparison, kvMap)
	}

}

func TestCompileTemplate_invalid(t *testing.T) {
	templateString := "a template {{nonexistentfunction}}"
	if _, err := compileTemplate(strings.NewReader(templateString)); err == nil {
		t.Errorf("\"%v\" is an invalid template", templateString)
	}

	templateString = ""
	if _, err := compileTemplate(strings.NewReader(templateString)); err == nil {
		t.Errorf("\"%v\" is an invalid template", templateString)
	}

	alwaysErrReader := readers.NewErrorReader()
	if _, err := compileTemplate(alwaysErrReader); err == nil {
		t.Errorf("%v should always error", alwaysErrReader)
	}
}

func TestCompileTemplate_valid(t *testing.T) {
	templateString := "legit template {{.test}}"
	if _, err := compileTemplate(strings.NewReader(templateString)); err != nil {
		t.Errorf("\"%v\" is a valid template", templateString)
	}
}

func TestRender(t *testing.T) {
	templateString := "legit template {{.test}} {{.key}}"
	compiledTemplate, err := compileTemplate(strings.NewReader(templateString))
	if err != nil {
		t.Errorf("\"%v\" is a valid template", templateString)
	}

	context := map[string]string{
		"test": "test",
		"key":  "value",
	}

	writer := new(bytes.Buffer)
	err = render(compiledTemplate, context, writer)
	expected := "legit template test value"
	if err != nil {
		t.Errorf("Expected: %v, Got: %v", expected, err)
	}
	if got := writer.String(); expected != got {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}
