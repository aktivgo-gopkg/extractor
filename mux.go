package extractor

import (
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"log"
)

type MuxExtractor struct {
	*StandardExtractor
}

func NewMuxExtractor() *MuxExtractor {
	return &MuxExtractor{
		StandardExtractor: NewStandardExtractor(),
	}
}

func (s MuxExtractor) extractUrlVar(
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

	return decoder.Decode(mux.Vars(r))
}
