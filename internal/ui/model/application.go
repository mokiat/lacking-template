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

func (a *ApplicationModel) ActiveView() ViewName {
	return a.activeView
}

func (a *ApplicationModel) SetActiveView(view ViewName) {
	a.activeView = view
	a.eventBus.Notify(ApplicationActiveViewChangedEvent{
		ActiveView: view,
	})
}

type ApplicationActiveViewChangedEvent struct {
	ActiveView ViewName
}
