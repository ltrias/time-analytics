package api

import "time"

//TimeEventRepository repository for time events
type TimeEventRepository struct {
}

//LoadAllEvents loads all events
func (t *TimeEventRepository) LoadAllEvents() []TimeEvent {
	var result []TimeEvent

	result = append(result, TimeEvent{time.Now(), "1:1", "Trias", 60, "Programming", "Devel", false})

	return result
}
