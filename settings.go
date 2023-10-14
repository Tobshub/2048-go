package main

type Settings struct {
	open bool
}

func ToggleSettings(s *Settings) {
	s.open = true
}

func (s *Settings) Draw() {
	if !s.open {
		return
	}
}
