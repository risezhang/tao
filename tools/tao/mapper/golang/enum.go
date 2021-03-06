package golang

import (
	"github.com/miraclew/tao/tools/tao/parser/proto3"
)

type EnumMapper struct {
}

func (e2 EnumMapper) Map(e *proto3.Enum) (*Enum, error) {
	var enum Enum
	enum.Name = e.Name
	for _, value := range e.Values {
		enum.Values = append(enum.Values, Value{
			Name:  value.Value.Key,
			Text:  EnumEntryTextOption(value.Value),
			Value: value.Value.Value,
		})
	}
	return &enum, nil
}

func EnumEntryTextOption(v *proto3.EnumValue) string {
	for _, option := range v.Options {
		if option.Name == "text" {
			return *option.Value.String
		}
	}
	return ""
}
