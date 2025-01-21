package model

func NewLoadingModel() *Loading {
	return &Loading{}
}

type Loading struct {
	state LoadingState
}

func (l *Loading) State() LoadingState {
	return l.state
}

func (l *Loading) SetState(state LoadingState) {
	l.state = state
}

type LoadingState struct {
	Promise         LoadingPromise
	SuccessViewName ViewName
	ErrorViewName   ViewName
}
