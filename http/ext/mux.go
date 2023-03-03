package ext

import (
	"github.com/aktivgo-gopkg/extractor/http/def"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"gitlab.collabox.dev/go/errors"
)

type MuxExtractor struct {
	*StandardExtractor
}

func NewMuxExtractor() *MuxExtractor {
	return &MuxExtractor{
		StandardExtractor: NewStandardExtractor(),
	}
}

func (s MuxExtractor) ExtractUrlVarsUnits(
	r def.Request,
	dest def.Destination,
	units def.UnitCollection,
) error {
	extractedUnits := make(map[string]any)

	for _, unit := range units {
		value, ok := mux.Vars(r)[unit.Name]
		if !ok {
			if unit.Required {
				return errors.ErrBadRequest.WithMessage(unit.Name).
					WithWrappedMessage("requirement url unit not found")
			}

			extractedUnits[unit.Name] = unit.DefaultValue
			continue
		}

		extractedUnits[unit.Name] = value
	}

	return mapstructure.Decode(extractedUnits, &dest)
}
