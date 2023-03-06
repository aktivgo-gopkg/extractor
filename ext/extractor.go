package ext

import (
	"errors"
	"fmt"
	"github.com/semichkin-gopkg/conf"
	"github.com/semichkin-gopkg/conv"
	"net/http"
)

type Extractor interface {
	ExtractUrl(Request, string) (string, error)
	ExtractQuery(Request, string) (string, error)
	ExtractBody(Request, string) (any, error)
}

type (
	Request     = *http.Request
	Destination = any
)

func Extract[T Destination](
	r Request,
	units Units,
	updaters ...conf.Updater[Configuration],
) (T, error) {
	configuration := defaultConfiguration.Append(updaters...).Build()

	var result T

	buffer := make(map[string]any)

	for _, unit := range units {
		var value any
		var err error

		switch unit.From {
		case "url":
			value, err = configuration.Extractor.ExtractUrl(r, unit.Name)
		case "query":
			value, err = configuration.Extractor.ExtractQuery(r, unit.Name)
		case "body":
			value, err = configuration.Extractor.ExtractBody(r, unit.Name)
		}

		if err != nil {
			if unit.Required {
				return result, errors.New(fmt.Sprintf("required unit not found: %v", unit))
			}

			value = unit.DefaultValue
		}

		buffer[unit.Name] = value
	}

	return conv.Struct[T](buffer)
}
