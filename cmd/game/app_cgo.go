//go:build !js

package main

import (
	"fmt"

	nativeapp "github.com/mokiat/lacking-native/app"
	nativeui "github.com/mokiat/lacking-native/ui"
	gameui "github.com/mokiat/lacking-template/internal/ui"
	"github.com/mokiat/lacking-template/resources"
	"github.com/mokiat/lacking/game/asset"
	"github.com/mokiat/lacking/ui"
	"github.com/mokiat/lacking/util/resource"
)

func runApplication() error {
	registryStorage, err := asset.NewFSStorage("./assets")
	if err != nil {
		return fmt.Errorf("failed to initialize storage: %w", err)
	}

	registryFormatter := asset.NewBlobFormatter()

	registry, err := asset.NewRegistry(registryStorage, registryFormatter)
	if err != nil {
		return fmt.Errorf("failed to initialize registry: %w", err)
	}

	locator := ui.WrappedLocator(resource.NewFSLocator(resources.UI))

	uiController := ui.NewController(locator, nativeui.NewShaderCollection(), func(w *ui.Window) {
		gameui.BootstrapApplication(w, registry)
	})

	cfg := nativeapp.NewConfig("Game", 1280, 800)
	cfg.SetFullscreen(false)
	cfg.SetMaximized(false)
	cfg.SetMinSize(1024, 576)
	cfg.SetVSync(true)
	cfg.SetIcon("ui/images/icon.png")
	cfg.SetLocator(locator)
	cfg.SetAudioEnabled(false)
	return nativeapp.Run(cfg, uiController)
}
