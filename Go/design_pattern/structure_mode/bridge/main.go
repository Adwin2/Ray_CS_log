package main

func main() {

	foo := &Monitor01{}
	bar := &Monitor02{}

	m := &PC01{str: "mac-pro"}
	w := &PC02{str: "dell"}

	m.SetMonitor(foo)
	m.Show()

	m.SetMonitor(bar)
	m.Show()

	w.SetMonitor(foo)
	w.Show()

	w.SetMonitor(bar)
	w.Show()
}
