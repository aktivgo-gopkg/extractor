package ext

import (
	"github.com/aktivgo-gopkg/extractor/http/def"
	"github.com/semichkin-gopkg/conf"
)

var defaultConfiguration = conf.NewBuilder[Configuration]().Append(
	WithExtractor(NewMuxExtractor()),
)

type UnitsCollection struct {
	UrlVars     def.UnitCollection
	QueryParams def.UnitCollection
	Body        def.UnitCollection
}

func ExtractUnits[T def.Destination](
	r def.Request,
	units UnitsCollection,
	updaters ...conf.Updater[Configuration],
) (T, error) {
	configuration := defaultConfiguration.Append(updaters...).Build()

	var result T

	if err := configuration.Extractor.ExtractUrlVarsUnits(r, &result, units.UrlVars); err != nil {
		return result, err
	}

	if err := configuration.Extractor.ExtractQueryParamsUnits(r, &result, units.QueryParams); err != nil {
		return result, err
	}

	if err := configuration.Extractor.ExtractBodyUnits(r, &result, units.Body); err != nil {
		return result, err
	}

	return result, nil
}

func ExtractUrlVars[T def.Destination](
	r def.Request,
	updaters ...conf.Updater[Configuration],
) (T, error) {
	configuration := defaultConfiguration.Append(updaters...).Build()

	var result T

	if err := configuration.Extractor.ExtractUrlVars(r, &result); err != nil {
		return result, err
	}

	return result, nil
}

func ExtractUrlVarsUnits[T def.Destination](
	r def.Request,
	units def.UnitCollection,
	updaters ...conf.Updater[Configuration],
) (T, error) {
	configuration := defaultConfiguration.Append(updaters...).Build()

	var result T

	if err := configuration.Extractor.ExtractUrlVarsUnits(r, &result, units); err != nil {
		return result, err
	}

	return result, nil
}

func ExtractQueryParams[T def.Destination](
	r def.Request,
	updaters ...conf.Updater[Configuration],
) (T, error) {
	configuration := defaultConfiguration.Append(updaters...).Build()

	var result T

	if err := configuration.Extractor.ExtractQueryParams(r, &result); err != nil {
		return result, err
	}

	return result, nil
}

func ExtractQueryParamsUnits[T def.Destination](
	r def.Request,
	units def.UnitCollection,
	updaters ...conf.Updater[Configuration],
) (T, error) {
	configuration := defaultConfiguration.Append(updaters...).Build()

	var result T

	if err := configuration.Extractor.ExtractQueryParamsUnits(r, &result, units); err != nil {
		return result, err
	}

	return result, nil
}

func ExtractBody[T def.Destination](
	r def.Request,
	updaters ...conf.Updater[Configuration],
) (T, error) {
	configuration := defaultConfiguration.Append(updaters...).Build()

	var result T

	if err := configuration.Extractor.ExtractBody(r, &result); err != nil {
		return result, err
	}

	return result, nil
}

func ExtractBodyUnits[T def.Destination](
	r def.Request,
	units def.UnitCollection,
	updaters ...conf.Updater[Configuration],
) (T, error) {
	configuration := defaultConfiguration.Append(updaters...).Build()

	var result T

	if err := configuration.Extractor.ExtractBodyUnits(r, &result, units); err != nil {
		return result, err
	}

	return result, nil
}
