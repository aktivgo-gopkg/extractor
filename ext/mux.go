package ext

import (
	"errors"
	"github.com/gorilla/mux"
)

type MuxExtractor struct {
	*StandardExtractor
}

func NewMuxExtractor() *MuxExtractor {
	return &MuxExtractor{
		StandardExtractor: NewStandardExtractor(),
	}
}

func (s MuxExtractor) ExtractUrl(
	r Request,
	name string,
) (string, error) {
	value, ok := mux.Vars(r)[name]
	if !ok {
		return "", errors.New("name not found")
	}

	return value, nil
}
