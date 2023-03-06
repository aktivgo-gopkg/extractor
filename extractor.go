package extractor

import (
	"github.com/semichkin-gopkg/conf"
	"log"
	"net/http"
	"reflect"
	"strings"
)

const (
	HttpExtTag = "httpext"
	HttpExtFrom
)

type (
	Request     = *http.Request
	Destination = any

	Unit struct {
		Name         string
		DefaultValue any
		Required     bool
	}

	UnitCollection = []Unit
)

type Extractor interface {
	extractUrlVar(Request, Destination) error
	extractQueryParams(Request, Destination) error
	extractBody(Request, Destination) error
}

var defaultConfiguration = conf.NewBuilder[Configuration]().Append(
	WithExtractor(NewMuxExtractor()),
)

//type UnitsCollection struct {
//	UrlVars     def.UnitCollection
//	QueryParams def.UnitCollection
//	Body        def.UnitCollection
//}

type unit struct {
	From     string
	Name     string
	Required bool
	Default  any
}

func Extract[T any](
	r Request,
	updaters ...conf.Updater[Configuration],
) (T, error) {
	configuration := defaultConfiguration.Append(updaters...).Build()

	var result T

parsing:
	for i := 0; i < reflect.TypeOf(result).NumField(); i++ {
		field := reflect.TypeOf(result).Field(i)

		tag := field.Tag.Get(HttpExtTag)
		if tag == "" {
			continue
		}

		log.Println(tag)

		unit := unit{
			From:     "",
			Name:     "",
			Required: false,
			Default:  nil,
		}

		for _, u := range strings.Split(tag, ",") {
			kv := strings.Split(u, ":")
			if len(kv) != 2 {
				continue parsing
			}

			key := kv[0]
			value := kv[1]

			switch key {
			case "from":
				if value == "" {
					continue parsing
				}
				unit.From = value
			case "name":
				if value == "" {
					continue parsing
				}
				unit.Name = value
			case "required":
				unit.Required = true
			case "default":
				if value == "" {
					continue parsing
				}
				unit.Default = value
			}
		}

		log.Fatalln(unit)

		switch unit.From {
		case "url":
			err := configuration.Extractor.extractUrlVar(r, &result)
			log.Println(err)
		case "query":
			_ = configuration.Extractor.extractQueryParams(r, &result)
		}

		log.Println(1)
	}

	return result, nil
}

//func ExtractUrlVars[T def.Destination](
//	r def.Request,
//	updaters ...conf.Updater[Configuration],
//) (T, error) {
//	configuration := defaultConfiguration.Append(updaters...).Build()
//
//	var result T
//
//	if err := configuration.Extractor.ExtractUrlVars(r, &result); err != nil {
//		return result, err
//	}
//
//	return result, nil
//}
//
//func ExtractUrlVarsUnits[T def.Destination](
//	r def.Request,
//	units def.UnitCollection,
//	updaters ...conf.Updater[Configuration],
//) (T, error) {
//	configuration := defaultConfiguration.Append(updaters...).Build()
//
//	var result T
//
//	if err := configuration.Extractor.ExtractUrlVarsUnits(r, &result, units); err != nil {
//		return result, err
//	}
//
//	return result, nil
//}
//
//func ExtractQueryParams[T def.Destination](
//	r def.Request,
//	updaters ...conf.Updater[Configuration],
//) (T, error) {
//	configuration := defaultConfiguration.Append(updaters...).Build()
//
//	var result T
//
//	if err := configuration.Extractor.ExtractQueryParams(r, &result); err != nil {
//		return result, err
//	}
//
//	return result, nil
//}
//
//func ExtractQueryParamsUnits[T def.Destination](
//	r def.Request,
//	units def.UnitCollection,
//	updaters ...conf.Updater[Configuration],
//) (T, error) {
//	configuration := defaultConfiguration.Append(updaters...).Build()
//
//	var result T
//
//	if err := configuration.Extractor.ExtractQueryParamsUnits(r, &result, units); err != nil {
//		return result, err
//	}
//
//	return result, nil
//}
//
//func ExtractBody[T def.Destination](
//	r def.Request,
//	updaters ...conf.Updater[Configuration],
//) (T, error) {
//	configuration := defaultConfiguration.Append(updaters...).Build()
//
//	var result T
//
//	if err := configuration.Extractor.ExtractBody(r, &result); err != nil {
//		return result, err
//	}
//
//	return result, nil
//}
//
//func ExtractBodyUnits[T def.Destination](
//	r def.Request,
//	units def.UnitCollection,
//	updaters ...conf.Updater[Configuration],
//) (T, error) {
//	configuration := defaultConfiguration.Append(updaters...).Build()
//
//	var result T
//
//	if err := configuration.Extractor.ExtractBodyUnits(r, &result, units); err != nil {
//		return result, err
//	}
//
//	return result, nil
//}
