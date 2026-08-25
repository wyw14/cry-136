package journal

import "github.com/wyw14/cry-136/internal/model"

func Replay(events []model.Event) map[string]string {
	state := map[string]string{}
	for _, event := range events {
		state[event.Kind] = event.Payload
	}
	return state
}

func Last(events []model.Event, kind string) (model.Event, bool) {
	for index := len(events) - 1; index >= 0; index-- {
		if events[index].Kind == kind { return events[index], true }
	}
	return model.Event{}, false
}
