package main

type ice House

func Newice() *ice {
	return &ice{}
}

func (i *ice) SetFrame() {
	i.Frame = "iceFrame"
}

func (i *ice) SetStyle() {
	i.Style = "white"
}

func (i *ice) SetDoor() {
	i.Door = "iceDoor"
}

func (i *ice) SetBed() {
	i.Bed = "iceBed"
}

func (i *ice) BuildHouse() *House {
	return &House{
		Frame: i.Frame,
		Style: i.Style,
		Door:  i.Door,
		Bed:   i.Bed,
	}
}
