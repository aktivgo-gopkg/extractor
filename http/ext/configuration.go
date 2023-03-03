package ext

import (
	"github.com/aktivgo-gopkg/extractor/http/def"
	"github.com/semichkin-gopkg/conf"
)

type Configuration struct {
	Extractor def.Extractor
}

func WithExtractor(extractor def.Extractor) conf.Updater[Configuration] {
	return func(c *Configuration) {
		c.Extractor = extractor
	}
}
