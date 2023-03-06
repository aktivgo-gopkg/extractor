package ext

type (
	Unit struct {
		From         string
		Name         string
		DefaultValue any
		Required     bool
	}

	Units = []Unit
)
