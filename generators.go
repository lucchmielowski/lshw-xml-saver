package main

import (
	"fmt"
)

func GenerateCpusFromNodes(n []Node) {
	var p []Node
	//c := make([]Cpu, 0)

	FindNodesByClass(n, "processor", &p)
	p = FilterDisabledNodes(n)
	for _, node := range p {
		tempCpu := Cpu{Version: node.Version, Size: node.Size[0].toUnitValue(), Clock: node.Clock[0].toUnitValue()}
		fmt.Println(tempCpu)
	}
}

func GenerateBankFromXml(n []Node) []Bank {
	var b []Node
	b = FilterDisabledNodes(n)
	var banks []Bank
	for _, node := range b {
		if len(node.Size) > 0 {
			banks = append(banks, Bank{Description: node.Description, Size: node.Size[0].toUnitValue()})
		} else {
			banks = append(banks, Bank{Description: node.Description, Size: UnitValue{}})
		}
	}
	return banks
}

func GenerateMemoryFromNodes(n []Node) {
	var m []Node
	FindNodesByClass(n, "memory", &m)
	m = FilterDisabledNodes(FilterMatchingNodes(m, "memory"))
	for _, node := range m {
		tempMemory := Memory{TotalSize: node.Size[0].toUnitValue(), Banks: make([]Bank, 0)}
		if len(node.ChildNodes) > 0 {
			banks := GenerateBankFromXml(node.ChildNodes)
			tempMemory.Banks = banks
		}
		fmt.Println(tempMemory)
	}
}

func GenerateDiskFromNodes(n []Node) {
	var s []Node
	FindNodesByClass(n, "storage", &s)
	s = FilterDisabledNodes(n)
}

func GenerateDisplayFromNodes(n []Node) {}

func GenerateServerFromNodes(n []Node) {}
