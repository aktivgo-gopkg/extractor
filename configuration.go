package extractor

import (
	"github.com/semichkin-gopkg/conf"
)

type Configuration struct {
	Extractor Extractor
}

func WithExtractor(extractor Extractor) conf.Updater[Configuration] {
	return func(c *Configuration) {
		c.Extractor = extractor
	}
}
