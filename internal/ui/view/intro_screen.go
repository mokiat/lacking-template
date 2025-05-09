package view

import (
	"time"

	"github.com/mokiat/gog/opt"
	"github.com/mokiat/lacking/ui"
	co "github.com/mokiat/lacking/ui/component"
	"github.com/mokiat/lacking/ui/layout"
	"github.com/mokiat/lacking/ui/std"

	"github.com/mokiat/lacking-template/internal/game/data"
	"github.com/mokiat/lacking-template/internal/ui/global"
	"github.com/mokiat/lacking-template/internal/ui/model"
)

var IntroScreen = co.Define(&introScreenComponent{})

type IntroScreenData struct {
	AppModel     *model.ApplicationModel
	ErrorModel   *model.ErrorModel
	LoadingModel *model.LoadingModel
	HomeModel    *model.HomeModel
}

type introScreenComponent struct {
	co.BaseComponent
}

func (c *introScreenComponent) OnCreate() {
	co.Window(c.Scope()).SetCursorVisible(false)

	globalState := co.TypedValue[global.State](c.Scope())
	engine := globalState.Engine
	resourceSet := globalState.ResourceSet

	componentData := co.GetData[IntroScreenData](c.Properties())
	appModel := componentData.AppModel
	errorModel := componentData.ErrorModel
	homeModel := componentData.HomeModel
	loadingModel := componentData.LoadingModel

	promise := model.NewLoadingPromise(
		co.Window(c.Scope()),
		data.LoadHomeData(engine, resourceSet),
		homeModel.SetData,
		errorModel.SetError,
	)
	loadingModel.SetState(model.LoadingState{
		Promise:         promise,
		SuccessViewName: model.ViewNameHome,
		ErrorViewName:   model.ViewNameError,
	})

	co.After(c.Scope(), time.Second, func() {
		appModel.SetActiveView(model.ViewNameLoading)
	})
}

func (c *introScreenComponent) OnDelete() {
	co.Window(c.Scope()).SetCursorVisible(true)
}

func (c *introScreenComponent) Render() co.Instance {
	return co.New(std.Container, func() {
		co.WithData(std.ContainerData{
			BackgroundColor: opt.V(ui.Black()),
			Layout:          layout.Anchor(),
		})

		co.WithChild("logo-picture", co.New(std.Picture, func() {
			co.WithLayoutData(layout.Data{
				Width:            opt.V(512),
				Height:           opt.V(128),
				HorizontalCenter: opt.V(0),
				VerticalCenter:   opt.V(0),
			})
			co.WithData(std.PictureData{
				BackgroundColor: opt.V(ui.Transparent()),
				Image:           co.OpenImage(c.Scope(), "ui/images/logo.png"),
				Mode:            std.ImageModeFit,
			})
		}))
	})
}
