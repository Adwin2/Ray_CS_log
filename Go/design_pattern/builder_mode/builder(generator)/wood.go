package main

type wood House

func Newwood() *wood {
	return &wood{}
}

func (i *wood) SetFrame() {
	i.Frame = "woodFrame"
}

func (i *wood) SetStyle() {
	i.Style = "white"
}

func (i *wood) SetDoor() {
	i.Door = "woodDoor"
}

func (i *wood) SetBed() {
	i.Bed = "woodBed"
}

func (i *wood) BuildHouse() *House {
	return &House{
		Frame: i.Frame,
		Style: i.Style,
		Door:  i.Door,
		Bed:   i.Bed,
	}
}
