package stat

import (
	"api-main/pkg/event"
	"log"
)

type StatServiseDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatServise struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatServise(deps *StatServiseDeps) *StatServise {
	return &StatServise{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatServise) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Bad EventLinkVisited Data:", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}
