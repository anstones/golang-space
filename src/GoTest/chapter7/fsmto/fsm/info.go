package fsm

type StateInfo struct {
	name string
}

func (s *StateInfo) Name() string {
	return s.name
}

func (s *StateInfo) SetName(name string)  {
	s.name = name
}

func (s *StateInfo) EnableSanmeTransit() bool {
	return false
}

func (s *StateInfo) OnBegin()  {

}

func (s *StateInfo) OnEnd()  {

}

func (s *StateInfo) CanTransitTo(name string) bool {
	return true
}