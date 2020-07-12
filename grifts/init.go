package grifts

import (
	"github.com/S-KYUCHAN/coke/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
