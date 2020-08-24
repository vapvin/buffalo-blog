package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/vapvin/vins_blog/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
