package main

import "fmt"

// handleEmptyUnitElmt  returns an empty UnitValues if the unit/value is not found
func handleEmptyUnitElmt(u []UnitElmt) UnitValue {
	if len(u) > 0 {
		return u[0].toUnitValue()
	}
	return UnitValue{}
}

// GenerateCpusFromNodes generates CPU elements for a server from XML Nodes
func GenerateCpusFromNodes(n []Node) []Cpu {
	var p []Node
	var c []Cpu

	FindNodesByClass(n, "processor", &p)
	// p = FilterDisabledNodes(n)
	for _, node := range p {
		tempCpu := Cpu{Version: node.Version, Size: handleEmptyUnitElmt(node.Size), Clock: handleEmptyUnitElmt(node.Clock)}
		c = append(c, tempCpu)
	}
	return c
}

// GenerateBankFromXml generates Memory.banks from XML Nodes
func GenerateBankFromXml(n []Node) []Bank {
	// var b []Node
	// b = FilterDisabledNodes(n)
	var banks []Bank
	for _, node := range n {
		banks = append(banks, Bank{Description: node.Description, Size: handleEmptyUnitElmt(node.Size)})
	}
	return banks
}

// GenerateMemoryFromNodes generates memory elements from XML Nodes
func GenerateMemoryFromNodes(n []Node) []Memory {
	var m []Node
	FindNodesByClass(n, "memory", &m)
	var mem []Memory
	m = FilterMatchingNodes(m, "memory")
	for _, node := range m {
		tempMemory := Memory{TotalSize: handleEmptyUnitElmt(node.Size), Banks: make([]Bank, 0)}
		if len(node.ChildNodes) > 0 {
			banks := GenerateBankFromXml(node.ChildNodes)
			tempMemory.Banks = banks
		}
		mem = append(mem, tempMemory)
	}
	return mem
}

// GenerateDiskFromNodes generates disks from XML Nodes
func GenerateDiskFromNodes(n []Node) []Disk {
	var s []Node
	FindNodesByClass(n, "disk", &s)
	var d []Disk
	s = FilterMatchingNodes(s, "disk")
	for _, node := range s {
		tempDisk := Disk{
			Description: node.Description,
			Size:        handleEmptyUnitElmt(node.Size),
			Product:     node.Product,
			Vendor:      node.Vendor,
			Serial:      node.Serial,
			BusInfo:     node.BusInfo,
		}

		d = append(d, tempDisk)
	}
	return d
}

// GenerateDisplayFromNodes generates the display elements for the given XMLNode
func GenerateDisplayFromNodes(n []Node) []Display {
	var d []Node
	var disp []Display
	FindNodesByClass(n, "display", &d)
	for _, node := range d {
		tempDisplay := Display{
			Description: node.Description,
			Product:     node.Product,
		}
		disp = append(disp, tempDisplay)
	}
	return disp
}

// GenerateServerFromNodes uses other generators to generate a full server conf from XML nodes
func GenerateServerFromNodes(name string, n []Node) Server {
	newServer := Server{
		Name:     name,
		Cpus:     GenerateCpusFromNodes(n),
		Memories: GenerateMemoryFromNodes(n),
		Displays: GenerateDisplayFromNodes(n),
		Disks:    GenerateDiskFromNodes(n),
	}
	fmt.Println(newServer)
	return newServer
}
