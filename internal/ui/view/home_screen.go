package view

import (
	"github.com/mokiat/gog/opt"
	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/lacking/game"
	"github.com/mokiat/lacking/game/graphics"
	"github.com/mokiat/lacking/ui"
	co "github.com/mokiat/lacking/ui/component"
	"github.com/mokiat/lacking/ui/layout"
	"github.com/mokiat/lacking/ui/std"

	"github.com/mokiat/lacking-template/internal/game/data"
	"github.com/mokiat/lacking-template/internal/ui/global"
	"github.com/mokiat/lacking-template/internal/ui/model"
	"github.com/mokiat/lacking-template/internal/ui/widget"
)

var HomeScreen = co.Define[*homeScreenComponent]()

type HomeScreenData struct {
	AppModel     *model.ApplicationModel
	ErrorModel   *model.ErrorModel
	LoadingModel *model.LoadingModel
	HomeModel    *model.HomeModel
	PlayModel    *model.PlayModel
}

type homeScreenComponent struct {
	co.BaseComponent

	engine      *game.Engine
	resourceSet *game.ResourceSet

	appModel     *model.ApplicationModel
	errorModel   *model.ErrorModel
	loadingModel *model.LoadingModel
	homeModel    *model.HomeModel
	playModel    *model.PlayModel
}

func (c *homeScreenComponent) OnCreate() {
	globalState := co.TypedValue[global.State](c.Scope())
	c.engine = globalState.Engine
	c.resourceSet = globalState.ResourceSet

	componentData := co.GetData[HomeScreenData](c.Properties())
	c.appModel = componentData.AppModel
	c.errorModel = componentData.ErrorModel
	c.loadingModel = componentData.LoadingModel
	c.homeModel = componentData.HomeModel
	c.playModel = componentData.PlayModel

	homeScene := c.homeModel.Scene()
	if homeScene == nil {
		homeScene = c.createScene()
		c.homeModel.SetScene(homeScene)
	}
	c.engine.SetActiveScene(homeScene.Scene)
	c.engine.ResetDeltaTime()
}

func (c *homeScreenComponent) OnDelete() {
	c.engine.SetActiveScene(nil)
}

func (c *homeScreenComponent) Render() co.Instance {
	return co.New(std.Element, func() {
		co.WithData(std.ElementData{
			Layout: layout.Anchor(),
		})

		co.WithChild("pane", co.New(std.Container, func() {
			co.WithLayoutData(layout.Data{
				Top:    opt.V(0),
				Bottom: opt.V(0),
				Left:   opt.V(0),
				Width:  opt.V(320),
			})
			co.WithData(std.ContainerData{
				BackgroundColor: opt.V(ui.RGBA(0, 0, 0, 192)),
				Layout:          layout.Anchor(),
			})

			co.WithChild("holder", co.New(std.Element, func() {
				co.WithLayoutData(layout.Data{
					Left:           opt.V(75),
					VerticalCenter: opt.V(0),
				})
				co.WithData(std.ElementData{
					Layout: layout.Vertical(layout.VerticalSettings{
						ContentAlignment: layout.HorizontalAlignmentLeft,
						ContentSpacing:   15,
					}),
				})

				co.WithChild("play-button", co.New(widget.Button, func() {
					co.WithData(widget.ButtonData{
						Text: "Play",
					})
					co.WithCallbackData(widget.ButtonCallbackData{
						OnClick: c.onPlayClicked,
					})
				}))

				co.WithChild("licenses-button", co.New(widget.Button, func() {
					co.WithData(widget.ButtonData{
						Text: "Licenses",
					})
					co.WithCallbackData(widget.ButtonCallbackData{
						OnClick: c.onLicensesClicked,
					})
				}))

				co.WithChild("exit-button", co.New(widget.Button, func() {
					co.WithData(widget.ButtonData{
						Text: "Exit",
					})
					co.WithCallbackData(widget.ButtonCallbackData{
						OnClick: c.onExitClicked,
					})
				}))
			}))
		}))
	})
}

func (c *homeScreenComponent) createScene() *model.HomeScene {
	sceneData := c.homeModel.Data()

	scene := c.engine.CreateScene(game.SceneInfo{
		IncludePhysics: opt.V(false),
		IncludeECS:     opt.V(false),
	})

	sceneModel := scene.InstantiateModel(game.ModelInfo{
		Template:  sceneData.Scene,
		Name:      opt.V("Scene"),
		IsDynamic: false,
	})

	camera := c.createCamera(scene.Graphics())
	scene.Graphics().SetActiveCamera(camera)

	if cameraNode := sceneModel.FindNode("Camera"); !cameraNode.IsNil() {
		scene.CameraBindingSet().Bind(cameraNode, camera)
	}

	const animationName = "CameraRotation"
	if recording := sceneModel.FindRecording(animationName); recording != nil {
		playback := recording.Playback(true)
		player := sceneModel.BindAnimation(playback)
		scene.PlayAnimation(player)
	}

	return &model.HomeScene{
		Scene: scene,
	}
}

func (c *homeScreenComponent) createCamera(scene *graphics.Scene) *graphics.Camera {
	result := scene.CreateCamera()
	result.SetFoVMode(graphics.FoVModeHorizontalPlus)
	result.SetFoV(sprec.Degrees(30))
	result.SetAutoExposure(false)
	result.SetExposure(1.0)
	result.SetAutoFocus(false)
	result.SetAutoExposureSpeed(0.1)
	result.SetCascadeDistances([]float32{32.0})
	return result
}

func (c *homeScreenComponent) onPlayClicked() {
	promise := model.NewLoadingPromise(
		co.Window(c.Scope()),
		data.LoadPlayData(c.engine, c.resourceSet),
		c.playModel.SetData,
		c.errorModel.SetError,
	)
	c.loadingModel.SetState(model.LoadingState{
		Promise:         promise,
		SuccessViewName: model.ViewNamePlay,
		ErrorViewName:   model.ViewNameError,
	})
	c.appModel.SetActiveView(model.ViewNameLoading)
}

func (c *homeScreenComponent) onLicensesClicked() {
	c.appModel.SetActiveView(model.ViewNameLicenses)
}

func (c *homeScreenComponent) onExitClicked() {
	co.Window(c.Scope()).Close()
}
