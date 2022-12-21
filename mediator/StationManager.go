package main

type StationManager struct {
	isFree bool
	queue  []ITrain
}

func newStationManger() *StationManager {
	return &StationManager{
		isFree: true,
	}
}

func (s *StationManager) canArrive(t ITrain) bool {
	if s.isFree {
		s.isFree = false
		return true
	}
	s.queue = append(s.queue, t)
	return false
}
func (s *StationManager) notifyAboutDeparture() {
	s.isFree = true

	firstOfQueue := s.queue[0]
	firstOfQueue.permitArrival()
	s.queue = s.queue[1:]
}
