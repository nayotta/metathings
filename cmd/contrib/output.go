package cmd_contrib

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jeremywohl/flatten"
	"gopkg.in/yaml.v2"
)

type OutputOptioner interface {
	GetOutputP() *string
	GetOutput() string
	SetOutput(string)
}

type OutputOption struct {
	Output string
}

func (self *OutputOption) GetOutputP() *string {
	return &self.Output
}

func (self *OutputOption) GetOutput() string {
	return self.Output
}

func (self *OutputOption) SetOutput(o string) {
	self.Output = o
}

func ProcessOutput(o OutputOptioner, in interface{}) error {
	switch o.GetOutput() {
	case "json":
		return processJsonOutput(in)
	case "yaml":
		return processYamlOutput(in)
	case "shell":
		return processShellOutput(in)
	default:
		return processYamlOutput(in)
	}
}

func processJsonOutput(in interface{}) error {
	buf, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return err
	}
	fmt.Print(string(buf))
	return nil
}

func processYamlOutput(in interface{}) error {
	buf, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	fmt.Print(string(buf))
	return nil
}

func processShellOutput(in interface{}) error {
	fltm, err := flatten.Flatten(in.(map[string]interface{}), "mt_", flatten.UnderscoreStyle)
	if err != nil {
		return err
	}

	for k, v := range fltm {
		fmt.Printf("export %v=\"%v\"\n", strings.ToUpper(k), v)
	}
	return nil
}
