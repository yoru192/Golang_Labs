package main

import (
	"fmt"
	"reflect"
	"strings"
)

func ToJSON(v any) (string, error) {
	val := reflect.ValueOf(v)
	t := val.Type()

	var parts []string

	for i := 0; i < val.NumField(); i++ {
		field := t.Field(i)
		fieldVal := val.Field(i)

		key := strings.Split(field.Tag.Get("json"), ",")[0]
		if key == "" {
			key = field.Name
		}

		var jsonVal string
		switch fieldVal.Kind() {
		case reflect.String:
			jsonVal = fmt.Sprintf("%q", fieldVal.String())
		case reflect.Int:
			jsonVal = fmt.Sprintf("%d", fieldVal.Int())
		case reflect.Bool:
			jsonVal = fmt.Sprintf("%t", fieldVal.Bool())
		case reflect.Slice:
			var elems []string
			for j := 0; j < fieldVal.Len(); j++ {
				elems = append(elems, fmt.Sprintf("    %q", fieldVal.Index(j).String()))
			}
			jsonVal = "[\n" + strings.Join(elems, ",\n") + "\n  ]"
		default:
			return "null", fmt.Errorf("unsupported type: %s", val.Kind())
		}

		parts = append(parts, fmt.Sprintf("  %q: %s", key, jsonVal))
	}

	return "{\n" + strings.Join(parts, ",\n") + "\n}", nil
}

type Server struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
}

func main() {
	s := Server{
		Host:       "localhost",
		Port:       8080,
		Debug:      true,
		AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
	}

	result, err := ToJSON(s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
