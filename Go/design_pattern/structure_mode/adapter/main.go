package main

func main() {
	c := &Client{}

	m := &Linux{}
	c.InsertUSBtoPC(m)

	other := &Other{}
	adapter := &Adapter{Othertmp: other}
	c.InsertUSBtoPC(adapter)
}
