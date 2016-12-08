package main

import (
	"fmt"
)

type Emplacement struct {
	Batiment  int
	salle     int
	autreInfo string
}

type UnitValue struct {
	Value int
	Unit  string
}

func (u UnitValue) String() string {
	return fmt.Sprintf("%d|%s", u.Value, u.Unit)
}

type Cpu struct {
	Version string
	Size    UnitValue
	Clock   UnitValue
}

func (c Cpu) String() string {
	return fmt.Sprintf("Version: %s, Size: %s, Clock: %s", c.Version, c.Size, c.Clock)
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
