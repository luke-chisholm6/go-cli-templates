# go-cli-templates [![Build Status](https://travis-ci.org/luke-chisholm6/go-cli-templates.svg?branch=master)](https://travis-ci.org/luke-chisholm6/go-cli-templates) [![Coverage Status](https://coveralls.io/repos/github/luke-chisholm6/go-cli-templates/badge.svg?branch=master)](https://coveralls.io/github/luke-chisholm6/go-cli-templates?branch=master)
This is a simple command line utility that leverages the go text/template package to generate dynamic files. 
My use case is for generating AWS Cli json command files in a ci pipeline.
### Installation
```bash
go get github.com/luke-chisholm6/go-cli-templates
```
### Usage
```bash
echo "{{.key1}} {{.key2}}" | go-cli-templates key1=Hello key2=world!
# will output: "Hello world!"
```