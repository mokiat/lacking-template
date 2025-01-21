package model

import (
	"github.com/mokiat/lacking/game"

	"github.com/mokiat/lacking-template/internal/game/data"
)

func NewHomeModel() *Home {
	return &Home{}
}

type Home struct {
	sceneData *data.HomeData
	scene     *HomeScene
}

func (h *Home) Data() *data.HomeData {
	return h.sceneData
}

func (h *Home) SetData(data *data.HomeData) {
	h.sceneData = data
}

func (h *Home) Scene() *HomeScene {
	return h.scene
}

func (h *Home) SetScene(scene *HomeScene) {
	h.scene = scene
}

type HomeScene struct {
	Scene *game.Scene
}
