package main

import (
	"bytes"
	"github.com/luke-chisholm6/go-cli-templates/readers"
	"github.com/luke-chisholm6/go-cli-templates/writers"
	"io"
	"reflect"
	"strings"
	"testing"
)

const (
	InvalidTemplateErrorString          = "\"%v\" is an invalid template"
	ValidTemplateErrorString            = "\"%v\" is a valid template"
	UnexpectedResultErrorString         = "Expected: %+v\nGot: %+v"
	InvalidTemplate_NonexistentFunction = "a template {{nonexistentfunction}}"
)

type TemplateTest struct {
	templateString string
	context        map[string]string
	contextRaw     []string
	expected       string
}

func NewLegitimateTemplateTest() *TemplateTest {
	return &TemplateTest{
		templateString: "legit template {{.test}} {{.key}}",
		context: map[string]string{
			"test": "test",
			"key":  "value",
		},
		contextRaw: []string{
			"test=test",
			"key=value",
		},
		expected: "legit template test value",
	}
}

func TestGetTemplateContext_SliceWithInvalidStrings(t *testing.T) {
	kvSlice := []string{
		"test",
	}

	if _, err := getTemplateContext(kvSlice); err == nil {
		t.Error("Cannot split a string into a kv pair that is not in the format of k=v")
	}
}

func TestGetTemplateContext_SliceWithValidStrings(t *testing.T) {
	templateTest := NewLegitimateTemplateTest()

	kvMap, err := getTemplateContext(templateTest.contextRaw)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(kvMap, templateTest.context) {
		t.Errorf(UnexpectedResultErrorString, templateTest.context, kvMap)
	}
}

func TestCompileTemplate_invalid(t *testing.T) {
	templateString := InvalidTemplate_NonexistentFunction
	if _, err := compileTemplate(strings.NewReader(templateString)); err == nil {
		t.Errorf(InvalidTemplateErrorString, templateString)
	}

	templateString = ""
	if _, err := compileTemplate(strings.NewReader(templateString)); err == nil {
		t.Errorf(InvalidTemplateErrorString, templateString)
	}

	alwaysErrReader := readers.NewErrorReader()
	if _, err := compileTemplate(alwaysErrReader); err == nil {
		t.Errorf("%v should always error", alwaysErrReader)
	}
}

func TestCompileTemplate_valid(t *testing.T) {
	templateTest := NewLegitimateTemplateTest()
	if _, err := compileTemplate(strings.NewReader(templateTest.templateString)); err != nil {
		t.Errorf(ValidTemplateErrorString, templateTest.templateString)
	}
}

func TestRender(t *testing.T) {
	templateTest := NewLegitimateTemplateTest()
	compiledTemplate, err := compileTemplate(strings.NewReader(templateTest.templateString))
	if err != nil {
		t.Errorf(ValidTemplateErrorString, templateTest.templateString)
	}

	writer := new(bytes.Buffer)
	err = render(compiledTemplate, templateTest.context, writer)
	if err != nil {
		t.Errorf(UnexpectedResultErrorString, templateTest.expected, err)
	}
	if got := writer.String(); templateTest.expected != got {
		t.Errorf(UnexpectedResultErrorString, templateTest.expected, got)
	}
}

func TestRun(t *testing.T) {
	templateTest := NewLegitimateTemplateTest()
	input := strings.NewReader(templateTest.templateString)
	writer := new(bytes.Buffer)
	run(input, templateTest.contextRaw, writer)

	got := writer.String()
	if got != templateTest.expected {
		t.Errorf(UnexpectedResultErrorString, templateTest.expected, got)
	}
}

func TestRun_Invalid(t *testing.T) {
	var writer io.Writer = new(bytes.Buffer)
	context := []string{
		"test",
		"key=world!",
	}
	if err := run(strings.NewReader(InvalidTemplate_NonexistentFunction), context, writer); err == nil {
		t.Errorf("\"%v\" is an invalid template", InvalidTemplate_NonexistentFunction)
	}

	templateTest := NewLegitimateTemplateTest()
	if err := run(strings.NewReader(templateTest.templateString), context, writer); err == nil {
		t.Errorf("\"%v\" is invalid context", context)
	}

	writer = writers.NewErrorWriter()
	if err := run(strings.NewReader(templateTest.templateString), templateTest.contextRaw, writer); err == nil {
		t.Errorf("\"%+v\" should always error", writer)
	}
}
