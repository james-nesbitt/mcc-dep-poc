package main

type Components []Component

type Component interface {
	Name() string
}
