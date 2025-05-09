package model

import "github.com/mokiat/lacking/ui/mvc"

const (
	ViewNameIntro    ViewName = "intro"
	ViewNameError    ViewName = "error"
	ViewNameLoading  ViewName = "loading"
	ViewNameLicenses ViewName = "licenses"
	ViewNameHome     ViewName = "home"
	ViewNamePlay     ViewName = "play"
)

type ViewName = string

func NewApplicationModel(eventBus *mvc.EventBus) *ApplicationModel {
	return &ApplicationModel{
		eventBus:   eventBus,
		activeView: ViewNameIntro,
	}
}

type ApplicationModel struct {
	eventBus   *mvc.EventBus
	activeView ViewName
}

func (m *ApplicationModel) ActiveView() ViewName {
	return m.activeView
}

func (m *ApplicationModel) SetActiveView(view ViewName) {
	m.activeView = view
	m.eventBus.Notify(ApplicationActiveViewChangedEvent{
		ActiveView: view,
	})
}

type ApplicationActiveViewChangedEvent struct {
	ActiveView ViewName
}
