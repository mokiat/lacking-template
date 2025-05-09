package model

import (
	"github.com/mokiat/lacking/game"

	"github.com/mokiat/lacking-template/internal/game/data"
)

func NewHomeModel() *HomeModel {
	return &HomeModel{}
}

type HomeModel struct {
	sceneData *data.HomeData
	scene     *HomeScene
}

func (m *HomeModel) Data() *data.HomeData {
	return m.sceneData
}

func (m *HomeModel) SetData(data *data.HomeData) {
	m.sceneData = data
}

func (m *HomeModel) Scene() *HomeScene {
	return m.scene
}

func (m *HomeModel) SetScene(scene *HomeScene) {
	m.scene = scene
}

type HomeScene struct {
	Scene *game.Scene
}
