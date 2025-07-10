package data

import (
	"github.com/mokiat/lacking/game"
	"github.com/mokiat/lacking/util/async"
)

func LoadPlayData(engine *game.Engine, resourceSet *game.ResourceSet) async.Promise[*PlayData] {
	var data PlayData
	return async.InjectionPromise(async.JoinOperations(
		resourceSet.FetchResource("play-screen.dat", &data.Scene),
		resourceSet.FetchResource("board.dat", &data.Board),
		resourceSet.FetchResource("ball.dat", &data.Ball),
	), &data)
}

type PlayData struct {
	Scene *game.ModelTemplate
	Board *game.ModelTemplate
	Ball  *game.ModelTemplate
}
