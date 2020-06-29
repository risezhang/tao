package generator

import (
	"testing"
)

func TestEngine_GenerateAPI(t *testing.T) {
	protoFile := "../testdata/demo.proto"
	e, err := NewEngine()
	if err != nil {
		t.Error(err)
	}
	e.Config = &Config{
		GoOutputDir:    "../testdata",
		SwiftOutputDir: "",
		DartOutputDir:  "",
		UseSnackCase:   false,
		Dependencies:   nil,
	}
	e.TemplateDir = "../templates"
	err = e.GenerateAPI(protoFile)
	if err != nil {
		t.Error(err)
	}
}
