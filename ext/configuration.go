package ext

import (
	"github.com/semichkin-gopkg/conf"
)

var defaultConfiguration = conf.NewBuilder[Configuration]().Append(
	WithExtractor(NewMuxExtractor()),
)

type Configuration struct {
	Extractor Extractor
}

func WithExtractor(extractor Extractor) conf.Updater[Configuration] {
	return func(c *Configuration) {
		c.Extractor = extractor
	}
}
