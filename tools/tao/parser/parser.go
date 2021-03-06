package parser

import (
	"errors"
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/miraclew/tao/tools/tao/parser/proto3"
	"os"
	"strings"
)

var (
	ErrModelNotFound   = errors.New("model not found")
	ErrServiceNotFound = errors.New("service not found")
)

func ParseProto3(file string) (*Result, error) {
	var parser = participle.MustBuild(&proto3.Proto{}, participle.UseLookahead(2))
	proto := &proto3.Proto{}
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	err = parser.Parse(r, proto)
	if err != nil {
		return nil, err
	}

	return &Result{
		Proto: proto,
	}, nil
}

type Result struct {
	Proto           *proto3.Proto
	ResourceName    string
	APIService      *proto3.Service
	EventService    *proto3.Service
	ResourceMessage *proto3.Message
	APIMessages     []*proto3.Message
	EventMessages   []*proto3.Message
}

func (r *Result) FileOption(option string) (string, error) {
	return FileOption(r.Proto, option)
}

func ResourceName(proto *proto3.Proto) (string, error) {
	return FileOption(proto, "resource")
}

func FileOption(proto *proto3.Proto, option string) (string, error) {
	for _, entry := range proto.Entries {
		if entry.Option != nil && entry.Option.Name == option {
			return *entry.Option.Value.String, nil
		}
	}
	return "", fmt.Errorf("option %s undefined", option)
}

func Package(proto *proto3.Proto) string {
	for _, entry := range proto.Entries {
		if entry.Package != "" {
			return entry.Package
		}
	}
	return ""
}

func ResourceMessage(proto *proto3.Proto) (*proto3.Message, error) {
	resourceName, err := ResourceName(proto)
	if err != nil {
		return nil, err
	}
	for _, entry := range proto.Entries {
		if entry.Message != nil {
			if entry.Message.Name == resourceName {
				return entry.Message, nil
			}
		}
	}
	return nil, ErrModelNotFound
}

func QueryMessages(proto *proto3.Proto, keys []string) []*proto3.Message {
	var result []*proto3.Message
	for _, entry := range proto.Entries {
		if entry.Message == nil {
			continue
		}

		for _, key := range keys {
			if strings.Contains(entry.Message.Name, key) {
				result = append(result, entry.Message)
			}
		}
	}

	return result
}

func IsPredefinedMessage(s string) bool {
	ms := []string{"Time", "Any", "Empty"}
	for _, v := range ms {
		if v == s {
			return true
		}
	}
	return false
}

func Service(proto *proto3.Proto, name string) (*proto3.Service, error) {
	for _, entry := range proto.Entries {
		if entry.Service != nil {
			if entry.Service.Name == name {
				return entry.Service, nil
			}
		}
	}
	return nil, ErrServiceNotFound
}

func ResourceFields(message *proto3.Message, quoted bool) []string {
	var fields []string
	for _, entry := range message.Entries {
		if entry.Field == nil {
			continue
		}
		if entry.Field.Name != "Id" && entry.Field.Name != "CreatedAt" && entry.Field.Name != "UpdatedAt" {
			var fieldName = entry.Field.Name
			if quoted {
				fieldName = fmt.Sprintf("\"%s\"", fieldName)
			}
			fields = append(fields, fieldName)
		}
	}
	return fields
}
