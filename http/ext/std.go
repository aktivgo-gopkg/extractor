package ext

import (
	"github.com/aktivgo-gopkg/extractor/http/def"
	"github.com/mitchellh/mapstructure"
	"github.com/semichkin-gopkg/conv"
	"gitlab.collabox.dev/go/errors"
	"io"
)

type StandardExtractor struct {
}

func NewStandardExtractor() *StandardExtractor {
	return &StandardExtractor{}
}

func (s StandardExtractor) ExtractUrlVars(
	_ def.Request,
	_ def.Destination,
) error {
	return nil
}

func (s StandardExtractor) ExtractUrlVarsUnits(
	_ def.Request,
	_ def.Destination,
	_ def.UnitCollection,
) error {
	return nil
}

func (s StandardExtractor) ExtractQueryParams(
	r def.Request,
	dest def.Destination,
) error {
	return mapstructure.Decode(r.URL.Query(), &dest)
}

func (s StandardExtractor) ExtractQueryParamsUnits(
	r def.Request,
	dest def.Destination,
	units def.UnitCollection,
) error {
	extractedUnits := make(map[string]any)

	for _, unit := range units {
		value := r.URL.Query().Get(unit.Name)
		if value == "" {
			if unit.Required {
				return errors.ErrBadRequest.WithMessage("query params unit not found")
			}

			extractedUnits[unit.Name] = unit.DefaultValue
			continue
		}

		extractedUnits[unit.Name] = value
	}

	return mapstructure.Decode(extractedUnits, &dest)
}

func (s StandardExtractor) ExtractBody(
	r def.Request,
	dest def.Destination,
) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.ErrInternal.WithMessage("body reading error")
	}

	return mapstructure.Decode(body, &dest)
}

func (s StandardExtractor) ExtractBodyUnits(
	r def.Request,
	dest def.Destination,
	units def.UnitCollection,
) error {
	extractedUnits := make(map[string]any)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.ErrInternal.WithMessage("body reading error")
	}

	extractedBody, err := conv.FromJSON[map[string]any](body)
	if err != nil {
		return errors.ErrBadRequest.WithMessage("body is invalid")
	}

	for _, unit := range units {
		value, ok := extractedBody[unit.Name]
		if !ok {
			if unit.Required {
				return errors.ErrBadRequest.WithMessage("body unit not found")
			}

			value = unit.DefaultValue
		}

		extractedUnits[unit.Name] = value
	}

	return mapstructure.Decode(extractedUnits, &dest)
}
