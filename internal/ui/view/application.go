package view

import (
	"github.com/mokiat/lacking-template/internal/ui/model"
	co "github.com/mokiat/lacking/ui/component"
	"github.com/mokiat/lacking/ui/mvc"
	"github.com/mokiat/lacking/ui/std"
)

var Application = mvc.EventListener(co.Define(&applicationComponent{}))

type applicationComponent struct {
	co.BaseComponent

	appModel     *model.Application
	errorModel   *model.ErrorModel
	loadingModel *model.Loading
	homeModel    *model.Home
}

func (c *applicationComponent) OnCreate() {
	eventBus := co.TypedValue[*mvc.EventBus](c.Scope())
	c.appModel = model.NewApplicationModel(eventBus)
	c.errorModel = model.NewErrorModel()
	c.loadingModel = model.NewLoadingModel()
	c.homeModel = model.NewHomeModel()
}

func (c *applicationComponent) Render() co.Instance {
	return co.New(std.Switch, func() {
		co.WithData(std.SwitchData{
			ChildKey: c.appModel.ActiveView(),
		})

		co.WithChild(model.ViewNameIntro, co.New(IntroScreen, func() {
			co.WithData(IntroScreenData{
				AppModel:     c.appModel,
				ErrorModel:   c.errorModel,
				LoadingModel: c.loadingModel,
				HomeModel:    c.homeModel,
			})
		}))
		// TODO: Add error screen
		co.WithChild(model.ViewNameLoading, co.New(LoadingScreen, func() {
			co.WithData(LoadingScreenData{
				AppModel:     c.appModel,
				LoadingModel: c.loadingModel,
			})
		}))
		co.WithChild(model.ViewNameLicenses, co.New(LicensesScreen, func() {
			co.WithData(LicensesScreenData{
				AppModel: c.appModel,
			})
		}))
		co.WithChild(model.ViewNameHome, co.New(HomeScreen, func() {
			co.WithData(HomeScreenData{
				AppModel:     c.appModel,
				LoadingModel: c.loadingModel,
				HomeModel:    c.homeModel,
			})
		}))
		co.WithChild(model.ViewNamePlay, co.New(PlayScreen, func() {
			co.WithData(PlayScreenData{
				AppModel: c.appModel,
			})
		}))
	})
}

func (c *applicationComponent) OnEvent(event mvc.Event) {
	switch event.(type) {
	case model.ApplicationActiveViewChangedEvent:
		c.Invalidate()
	}
}
