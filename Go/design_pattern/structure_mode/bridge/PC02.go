package main

type PC02 struct {
	str     string
	monitor Monitor
}

func (p *PC02) SetMonitor(m Monitor) {
	p.monitor = m
}

func (p *PC02) Show() {
	p.monitor.Show(p.str)
}
