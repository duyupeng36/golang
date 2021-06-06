package main

import (
	"fmt"
)

var m *Manager

func GetInstance() *Manager {
	if m == nil {
		m = &Manager{}
	}
	return m
}

type Manager struct {
}

func (m Manager) Manage() {
	fmt.Println("manage...")
}

//func main() {
//	m1 := GetInstance()
//	m2 := GetInstance()
//
//	fmt.Println(m1==m2)  // true
//}
