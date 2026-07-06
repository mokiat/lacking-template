//go:build js

package main

import (
	jsapp "github.com/mokiat/lacking-js/app"
	jsgame "github.com/mokiat/lacking-js/game"
	jsui "github.com/mokiat/lacking-js/ui"
	gameui "github.com/mokiat/lacking-template/internal/ui"
	"github.com/mokiat/lacking-template/internal/ui/view"
	"github.com/mokiat/lacking-template/resources"
	"github.com/mokiat/lacking/app"
	"github.com/mokiat/lacking/core/resource"
	"github.com/mokiat/lacking/game"
	"github.com/mokiat/lacking/ui"
	"github.com/mokiat/lacking/ui/resources/fonts"
	"github.com/mokiat/lacking/ui/resources/icons"
)

func runApplication() error {
	store := resource.NewWebStore(".")

	locator := resource.OneOfLocator(
		resource.SchemaLocator("ui", resource.NewFSStore(icons.FS)),
		resource.SchemaLocator("ui", resource.NewFSStore(fonts.FS)),
		resource.NewFSStore(resources.UI),
	)

	gameController := game.NewController(store, jsgame.NewShaderCollection(), jsgame.NewShaderBuilder())
	uiController := ui.NewController(locator, jsui.NewShaderCollection(), func(w *ui.Window) {
		gameui.BootstrapApplication(w, gameController, view.Application)
	})

	cfg := jsapp.NewConfig("screen")
	cfg.AddGLExtension("EXT_color_buffer_float")
	cfg.SetFullscreen(false)
	cfg.SetAudioEnabled(false)
	return jsapp.Run(cfg, app.NewLayeredController(gameController, uiController))
}
