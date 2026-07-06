//go:build !js

package main

import (
	"fmt"

	nativeapp "github.com/mokiat/lacking-native/app"
	nativegame "github.com/mokiat/lacking-native/game"
	nativeui "github.com/mokiat/lacking-native/ui"
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
	storage, err := resource.NewFileStore("./assets")
	if err != nil {
		return fmt.Errorf("failed to initialize storage: %w", err)
	}

	locator := resource.OneOfLocator(
		resource.SchemaLocator("ui", resource.NewFSStore(icons.FS)),
		resource.SchemaLocator("ui", resource.NewFSStore(fonts.FS)),
		resource.NewFSStore(resources.UI),
	)

	gameController := game.NewController(storage, nativegame.NewShaderCollection(), nativegame.NewShaderBuilder())
	uiController := ui.NewController(locator, nativeui.NewShaderCollection(), func(w *ui.Window) {
		gameui.BootstrapApplication(w, gameController, view.Application)
	})

	cfg := nativeapp.NewConfig("Game", 1280, 800)
	cfg.SetFullscreen(false)
	cfg.SetMaximized(false)
	cfg.SetMinSize(1024, 576)
	cfg.SetVSync(true)
	cfg.SetIcon("ui/images/icon.png")
	cfg.SetLocator(locator)
	cfg.SetAudioEnabled(false)
	return nativeapp.Run(cfg, app.NewLayeredController(gameController, uiController))
}
