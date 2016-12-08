package main

type Emplacement struct {
	Batiment  int
	salle     int
	autreInfo string
}

type UnitValue struct {
	Value int
	Unit  string
}

type Cpu struct {
	Version string
	Size    UnitValue
	Clock   UnitValue
}

type Memory struct {
	TotalSize UnitValue
	Banks     []Bank
}

type Bank struct {
	Description string
	Size        UnitValue
}

type Display struct {
	Description string
	Product     string
}

type Disk struct {
	Description string
	Product     string
	Vendor      string
	BusInfo     string
	serial      string
}

type Server struct {
	Name        string
	Emplacement Emplacement
	Cpus        []Cpu
	Memories    []Memory
	Displays    []Display
	Disks       []Disk
}
