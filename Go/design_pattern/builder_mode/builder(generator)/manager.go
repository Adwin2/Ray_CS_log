package main

type Manager struct {
	Worker Build
}

func NewManager(b Build) *Manager {
	return &Manager{Worker: b}
}

func (m *Manager) SetWorker(b Build) {
	m.Worker = b
}

func (m *Manager) GetHouse() *House {
	m.Worker.SetFrame()

	m.Worker.SetStyle()

	m.Worker.SetDoor()

	m.Worker.SetBed()

	return m.Worker.BuildHouse()
}
