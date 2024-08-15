//go:build js

package main

import (
	"fmt"

	jsapp "github.com/mokiat/lacking-js/app"
	jsui "github.com/mokiat/lacking-js/ui"
	gameui "github.com/mokiat/lacking-template/internal/ui"
	"github.com/mokiat/lacking-template/resources"
	"github.com/mokiat/lacking/game/asset"
	"github.com/mokiat/lacking/ui"
	"github.com/mokiat/lacking/util/resource"
)

func runApplication() error {
	registryStorage, err := asset.NewWebStorage(".")
	if err != nil {
		return fmt.Errorf("failed to initialize storage: %w", err)
	}

	registryFormatter := asset.NewBlobFormatter()

	registry, err := asset.NewRegistry(registryStorage, registryFormatter)
	if err != nil {
		return fmt.Errorf("failed to initialize registry: %w", err)
	}

	locator := ui.WrappedLocator(resource.NewFSLocator(resources.UI))

	uiController := ui.NewController(locator, jsui.NewShaderCollection(), func(w *ui.Window) {
		gameui.BootstrapApplication(w, registry)
	})

	cfg := jsapp.NewConfig("screen")
	cfg.AddGLExtension("EXT_color_buffer_float")
	cfg.SetFullscreen(false)
	cfg.SetAudioEnabled(false)
	return jsapp.Run(cfg, uiController)
}
