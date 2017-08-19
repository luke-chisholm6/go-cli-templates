package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
	"github.com/luke-chisholm6/go-cli-templates/readers"
	"github.com/luke-chisholm6/go-cli-templates/writers"
	"io"
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

func TestRun(t *testing.T) {
	input := strings.NewReader("{{.test}} {{.key}}")
	context := []string{
		"test=Hello",
		"key=world!",
	}
	writer := new(bytes.Buffer)
	run(input, context, writer)

	expected := "Hello world!"
	got := writer.String()
	if got != expected {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}

func TestRun_Invalid(t *testing.T) {
	templateString := "a template {{nonexistentfunction}}"
	var writer io.Writer = new(bytes.Buffer)
	context := []string{
		"test",
		"key=world!",
	}
	if err := run(strings.NewReader(templateString), context, writer); err == nil {
		t.Errorf("\"%v\" is an invalid template", templateString)
	}

	templateString = "a legit template {{.key}}"
	if err := run(strings.NewReader(templateString), context, writer); err == nil {
		t.Errorf("\"%v\" is invalid context", context)
	}

	context = []string{
		"test=success",
	}
	writer = writers.NewErrorWriter()
	if err := run(strings.NewReader(templateString), context, writer); err == nil {
		t.Errorf("\"%v\" is invalid context", context)
	}

}