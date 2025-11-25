package model

import (
	"github.com/mokiat/lacking/game"

	"github.com/mokiat/lacking-template/internal/game/data"
)

func NewPlayModel() *PlayModel {
	return &PlayModel{}
}

type PlayModel struct {
	sceneData *data.PlayData
	scene     *PlayScene
}

func (m *PlayModel) Data() *data.PlayData {
	return m.sceneData
}

func (m *PlayModel) SetData(data *data.PlayData) {
	m.sceneData = data
}

func (m *PlayModel) Scene() *PlayScene {
	return m.scene
}

func (m *PlayModel) SetScene(scene *PlayScene) {
	m.scene = scene
}

type PlayScene struct {
	Scene *game.Scene
}
