package main

import "fmt"

type Artist interface {
	fulfill(fantancy interface{})
	say()
}

type Van struct {
	name string
}

func (van *Van) say() {
	fmt.Println("my name's ", van.name)
}
func (van *Van) fulfill(fantancy interface{}) {
	van.say()
	fmt.Printf("fulfill your %v,deep dark  fantacy\n", fantancy)
}

type PerformanArtist struct {
	name string
}

func (your PerformanArtist) say() {
	fmt.Println("my name's ", your.name)
}
func (you PerformanArtist) fulfill(fantancy interface{}) {
	you.say()
	fmt.Printf("I can't fulfill your  %v, \n", fantancy)
}

func main() {
	van := Van{"van"}
	van.fulfill("fantacy")

	van2 := Van{"van2"}
	van2.fulfill("fantacy")

	var your PerformanArtist
	your.name = "哲学家"
	your.fulfill("fantacy")
}
