package def

import "net/http"

type (
	Request     = *http.Request
	Destination = any

	Unit struct {
		Name         string
		DefaultValue any
		Required     bool
	}

	UnitCollection = []Unit
)

type Extractor interface {
	ExtractUrlVars(Request, Destination) error
	ExtractUrlVarsUnits(Request, Destination, UnitCollection) error
	ExtractQueryParams(Request, Destination) error
	ExtractQueryParamsUnits(Request, Destination, UnitCollection) error
	ExtractBody(Request, Destination) error
	ExtractBodyUnits(Request, Destination, UnitCollection) error
}
