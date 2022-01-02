package main

type Animal int

//go:generate stringer -type=Animal
const (
	Dog Animal = iota
	Cat
	Hamster
	Tiger
	Lion
	Fish
)
