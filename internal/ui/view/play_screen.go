package view

import (
	"time"

	"github.com/mokiat/gog/opt"
	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/lacking-template/internal/ui/global"
	"github.com/mokiat/lacking-template/internal/ui/model"
	"github.com/mokiat/lacking/core/spatial/placement3d"
	"github.com/mokiat/lacking/core/spatial/shape3d"
	"github.com/mokiat/lacking/debug/metric/metricui"
	"github.com/mokiat/lacking/game"
	"github.com/mokiat/lacking/game/graphics"
	"github.com/mokiat/lacking/game/hierarchy"
	"github.com/mokiat/lacking/game/physics"
	"github.com/mokiat/lacking/ui"
	co "github.com/mokiat/lacking/ui/component"
	"github.com/mokiat/lacking/ui/layout"
	"github.com/mokiat/lacking/ui/std"
)

var PlayScreen = co.Define[*playScreenComponent]()

type PlayScreenData struct {
	AppModel  *model.ApplicationModel
	PlayModel *model.PlayModel
}

type playScreenComponent struct {
	co.BaseComponent

	debugVisible bool

	engine      *game.Engine
	resourceSet *game.ResourceSet

	appModel  *model.ApplicationModel
	playModel *model.PlayModel
}

var _ ui.ElementKeyboardHandler = (*playScreenComponent)(nil)

func (c *playScreenComponent) OnCreate() {
	c.debugVisible = false

	globalState := co.TypedValue[global.State](c.Scope())
	c.engine = globalState.Engine
	c.resourceSet = globalState.ResourceSet

	componentData := co.GetData[PlayScreenData](c.Properties())
	c.appModel = componentData.AppModel
	c.playModel = componentData.PlayModel

	playScene := c.playModel.Scene()
	if playScene == nil {
		playScene = c.createScene()
		c.playModel.SetScene(playScene)
	}
	c.engine.SetActiveScene(playScene.Scene)
	c.engine.ResetDeltaTime()
}

func (c *playScreenComponent) OnDelete() {
	c.engine.SetActiveScene(nil)
}

func (c *playScreenComponent) OnKeyboardEvent(element *ui.Element, event ui.KeyboardEvent) bool {
	switch event.Code {

	case ui.KeyCodeEscape:
		co.Window(c.Scope()).Close()
		return true

	case ui.KeyCodeTab:
		if event.Action == ui.KeyboardActionDown {
			c.debugVisible = !c.debugVisible
			c.Invalidate()
		}
		return true

	default:
		return false
	}
}

func (c *playScreenComponent) Render() co.Instance {
	return co.New(std.Element, func() {
		co.WithData(std.ElementData{
			Essence:       c,
			CanAutoFocus:  opt.V(true),
			CreateFocused: true,
			Layout:        layout.Anchor(),
		})

		if c.debugVisible {
			co.WithChild("flamegraph", co.New(metricui.FlameGraph, func() {
				co.WithData(metricui.FlameGraphData{
					UpdateInterval: time.Second,
				})
				co.WithLayoutData(layout.Data{
					Top:   opt.V(0),
					Left:  opt.V(0),
					Right: opt.V(0),
				})
			}))
		}
	})
}

func (c *playScreenComponent) createScene() *model.PlayScene {
	sceneData := c.playModel.Data()

	scene := c.engine.CreateScene(game.SceneInfo{
		IncludeECS: opt.V(false),
	})

	scene.InstantiateModel(game.ModelInfo{
		Template:  sceneData.Scene,
		Name:      opt.V("Scene"),
		IsDynamic: false,
	})

	boardModel := scene.InstantiateModel(game.ModelInfo{
		Template:  sceneData.Board,
		Name:      opt.V("Board"),
		IsDynamic: false,
	})

	camera := c.createCamera(scene.Graphics())
	scene.Graphics().SetActiveCamera(camera)

	if cameraNode := boardModel.FindNode("Camera"); cameraNode != hierarchy.NilNodeID {
		scene.CameraBindingSet().Bind(cameraNode, camera)
	}

	ballModel := scene.InstantiateModel(game.ModelInfo{
		Template:  sceneData.Ball,
		Name:      opt.V("Ball"),
		Position:  opt.V(dprec.NewVec3(-1.0, 3.0, 2.0)),
		IsDynamic: true,
	})
	ballModelNode := scene.Hierarchy().Nodes().Handle(ballModel.Root())

	physicsScene := scene.Physics()
	ballBody := physicsScene.Bodies().CreateHandle(ballModelNode.Position(), ballModelNode.Rotation())
	ballBody.SetMass(1.0)
	ballBody.SetMomentOfInertia(physics.SolidSphereMomentOfInertia(1.0, 1.0))
	ballBody.AttachCollisionSphere(physics.CollisionSphere{
		Shape:                  shape3d.NewSphere(dprec.ZeroVec3(), 1.0),
		FrictionCoefficient:    0.5,
		RestitutionCoefficient: 0.5,
		Filtering: placement3d.FilterInfo{
			RejectGroup: physicsScene.NextCollisionRejectGroup(),
		},
	})

	ballBody.SetVelocity(dprec.NewVec3(0.0, 0.0, 3.0))
	scene.BodyBindingSet().Bind(ballModelNode.ID(), ballBody.ID())

	physicsScene.GlobalAccelerators().Create(physics.NewGravitySolver())

	return &model.PlayScene{
		Scene: scene,
	}
}

func (c *playScreenComponent) createCamera(scene *graphics.Scene) *graphics.Camera {
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
