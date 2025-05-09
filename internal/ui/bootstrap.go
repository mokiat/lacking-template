package ui

import (
	"github.com/mokiat/lacking-template/internal/ui/global"
	"github.com/mokiat/lacking/game"
	"github.com/mokiat/lacking/ui"
	co "github.com/mokiat/lacking/ui/component"
	"github.com/mokiat/lacking/ui/mvc"
)

func BootstrapApplication(window *ui.Window, gameController *game.Controller, component co.Component) {
	engine := gameController.Engine()
	eventBus := mvc.NewEventBus()

	scope := co.RootScope(window)
	scope = co.TypedValueScope(scope, eventBus)
	scope = co.TypedValueScope(scope, global.State{
		Engine:      engine,
		ResourceSet: engine.CreateResourceSet(),
	})
	co.Initialize(scope, co.New(component, nil))
}
