# go-cli-templates [![Build Status](https://travis-ci.org/luke-chisholm6/go-cli-templates.svg?branch=master)](https://travis-ci.org/luke-chisholm6/go-cli-templates) [![Coverage Status](https://coveralls.io/repos/github/luke-chisholm6/go-cli-templates/badge.svg?branch=master)](https://coveralls.io/github/luke-chisholm6/go-cli-templates?branch=master)
This is a simple command line utility that leverages the go text/template package to generate dynamic files. 
My use case is for generating AWS Cli json command files in a ci pipeline.
### Installation
```bash
go get github.com/luke-chisholm6/go-cli-templates
```
or download the applicable binary for your platform from github releases replace darwin_amd64 with your os/arch combo (linux_386)
```bash
# using wget
wget https://github.com/luke-chisholm6/go-cli-templates/releases/download/0.1.0/go-cli-templates_darwin_amd64 -O /usr/local/bin/go-cli-templates 
# using curl
curl https://github.com/luke-chisholm6/go-cli-templates/releases/download/0.1.0/go-cli-templates_darwin_amd64 -Lo /usr/local/bin/go-cli-templates 
# don't forget to make it executable
chmod +x /usr/local/bin/go-cli-templates
```
### Usage
```bash
echo "{{.key1}} {{.key2}}" | go-cli-templates key1=Hello key2=world!
# will output: "Hello world!"
```