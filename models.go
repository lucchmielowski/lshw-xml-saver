package main

import (
	"fmt"
)

// Emplacement is generated from JSON in server files
type Emplacement struct {
	Batiment  int
	Salle     int
	AutreInfo string
}

// UnitValue represents a unit/value tupple from xml
type UnitValue struct {
	Value int
	Unit  string
}

func (u UnitValue) String() string {
	return fmt.Sprintf("%d|%s", u.Value, u.Unit)
}

// Cpu type represents a cpu instance in XML
type Cpu struct {
	Version string
	Size    UnitValue
	Clock   UnitValue
}

func (c Cpu) String() string {
	return fmt.Sprintf("Version: %s, Size: %s, Clock: %s\n", c.Version, c.Size, c.Clock)
}

// Memory represents a RAM slot in a machine
type Memory struct {
	TotalSize UnitValue
	Banks     []Bank
}

func (m Memory) String() string {
	return fmt.Sprintf("TotalSize: %s, Banks: %s\n", m.TotalSize, m.Banks)
}

// Bank model represent a memory slot's bank
type Bank struct {
	Description string
	Size        UnitValue
}

func (b Bank) String() string {
	return fmt.Sprintf("Description: %s, Size: %s\n", b.Description, b.Size)
}

// Display model represents the display capabilities of a machine (VGA, etc..)
type Display struct {
	Description string
	Product     string
}

func (d Display) String() string {
	return fmt.Sprintf("Description: %s\n, Product: %s\n", d.Description, d.Product)
}

type Disk struct {
	Description string
	Product     string
	Vendor      string
	Size        UnitValue
	BusInfo     string
	Serial      string
}

func (d Disk) String() string {
	return fmt.Sprintf("Description: %s\n, Product: %s\n, Size: %s\n, Vendor: %s\n, BusInfo: %s\n, Serial: %s\n",
		d.Description,
		d.Product,
		d.Size,
		d.Vendor,
		d.BusInfo,
		d.Serial,
	)
}

type Server struct {
	Name string
	//Emplacement Emplacement
	Cpus     []Cpu
	Memories []Memory
	Displays []Display
	Disks    []Disk
}

func (s Server) String() string {
	return fmt.Sprintf("Name: %s,\n Cpus: %s,\n Memories: %s,\n Displays: %s,\n, Disks: %s,\n",
		s.Name,
		s.Cpus,
		s.Memories,
		s.Displays,
		s.Disks,
	)
}
