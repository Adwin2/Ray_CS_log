package main

type PC01 struct {
	str     string
	monitor Monitor
}

func (p *PC01) SetMonitor(m Monitor) {
	p.monitor = m
}

func (p *PC01) Show() {
	p.monitor.Show(p.str)
}
