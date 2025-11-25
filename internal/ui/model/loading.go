package model

func NewLoadingModel() *LoadingModel {
	return &LoadingModel{}
}

type LoadingModel struct {
	state LoadingState
}

func (m *LoadingModel) State() LoadingState {
	return m.state
}

func (m *LoadingModel) SetState(state LoadingState) {
	m.state = state
}

type LoadingState struct {
	Promise         LoadingPromise
	SuccessViewName ViewName
	ErrorViewName   ViewName
}
