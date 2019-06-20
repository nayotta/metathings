package fx_helper

import "go.uber.org/fx"

type FxAppGetter struct {
	app **fx.App
}

func (g *FxAppGetter) Get() *fx.App {
	return *g.app
}

func NewFxAppGetter(app **fx.App) func() *FxAppGetter {
	return func() *FxAppGetter { return &FxAppGetter{app: app} }
}
