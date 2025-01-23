package data

import (
	"errors"

	"github.com/mokiat/lacking/game"
	"github.com/mokiat/lacking/util/async"
)

func LoadPlayData(engine *game.Engine, resourceSet *game.ResourceSet) async.Promise[*PlayData] {
	scenePromise := resourceSet.OpenModelByName("PlayScreen")
	boardPromise := resourceSet.OpenModelByName("Board")
	ballPromise := resourceSet.OpenModelByName("Ball")

	promise := async.NewPromise[*PlayData]()
	go func() {
		var data PlayData
		err := errors.Join(
			scenePromise.Inject(&data.Scene),
			boardPromise.Inject(&data.Board),
			ballPromise.Inject(&data.Ball),
		)
		if err != nil {
			promise.Fail(err)
		} else {
			promise.Deliver(&data)
		}
	}()
	return promise
}

type PlayData struct {
	Scene *game.ModelDefinition
	Board *game.ModelDefinition
	Ball  *game.ModelDefinition
}
