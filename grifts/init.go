package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/rishavs/integra/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
