package data

import (
	"github.com/mokiat/lacking/game"
	"github.com/mokiat/lacking/util/async"
)

func LoadHomeData(engine *game.Engine, resourceSet *game.ResourceSet) async.Promise[*HomeData] {
	var data HomeData
	return async.InjectionPromise(async.JoinOperations(
		resourceSet.FetchResource("home-screen.dat", &data.Scene),
	), &data)
}

type HomeData struct {
	Scene *game.ModelTemplate
}
