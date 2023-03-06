package extractor

import (
	"github.com/mitchellh/mapstructure"
	"github.com/semichkin-gopkg/conv"
	"gitlab.collabox.dev/go/errors"
	"io"
	"log"
)

type StandardExtractor struct {
}

func NewStandardExtractor() *StandardExtractor {
	return &StandardExtractor{}
}

func (s StandardExtractor) extractUrlVars(
	_ Request,
	_ Destination,
) error {
	return nil
}

func (s StandardExtractor) extractQueryParams(
	r Request,
	dest Destination,
) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		MatchName: func(mapKey, fieldName string) bool {
			log.Println(mapKey, fieldName)
			return true
		},
		Result: dest,
	})
	if err != nil {
		return err
	}

	return decoder.Decode(r.URL.Query())
}

func (s StandardExtractor) extractBody(
	r Request,
	dest Destination,
) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.ErrInternal.WithMessage(err).WithWrappedMessage("body reading error")
	}

	return mapstructure.Decode(body, &dest)
}

func (s StandardExtractor) extractBodyUnits(
	r Request,
	dest Destination,
	units UnitCollection,
) error {
	extractedUnits := make(map[string]any)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.ErrInternal.WithMessage(err).WithWrappedMessage("body reading error")
	}

	extractedBody, err := conv.FromJSON[map[string]any](body)
	if err != nil {
		return errors.ErrBadRequest.WithMessage(err).WithWrappedMessage("body is invalid")
	}

	for _, unit := range units {
		value, ok := extractedBody[unit.Name]
		if !ok {
			if unit.Required {
				return errors.ErrBadRequest.WithMessage(unit.Name).
					WithWrappedMessage("requirement body unit not found")
			}

			value = unit.DefaultValue
		}

		extractedUnits[unit.Name] = value
	}

	return mapstructure.Decode(extractedUnits, &dest)
}
