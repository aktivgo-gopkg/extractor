package ext

import (
	"errors"
	"github.com/semichkin-gopkg/conv"
	"io"
)

type StandardExtractor struct {
}

func NewStandardExtractor() *StandardExtractor {
	return &StandardExtractor{}
}

func (s StandardExtractor) ExtractUrl(
	_ Request,
	_ string,
) (string, error) {
	return "", nil
}

func (s StandardExtractor) ExtractQuery(
	r Request,
	name string,
) (string, error) {
	value := r.URL.Query().Get(name)
	if value == "" {
		return "", errors.New("name not found")
	}

	return value, nil
}

func (s StandardExtractor) ExtractBody(
	r Request,
	name string,
) (any, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", errors.New("body reading error")
	}

	encBody, err := conv.FromJSON[map[string]any](body)
	if err != nil {
		return "", err
	}

	value, ok := encBody[name]
	if !ok {
		return "", errors.New("name not found")
	}

	return value, nil
}
