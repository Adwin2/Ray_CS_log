package main

type Build interface {
	SetFrame()
	SetStyle()
	SetDoor()
	SetBed()
	BuildHouse() *House
}

func GetBuild(buildType string) Build {
	if buildType == "ice" {
		return Newice()
	}
	if buildType == "wood" {
		return Newwood()
	}
	return nil
}
