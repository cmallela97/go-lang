package main

import "fmt"

type Astro struct{
	name string
	age  int
	lastmission string
	isneeded bool
}

func main(){
	a1 := Astro{"Armstrong", 65, "Moon", false}
	a2 := Astro{"Kalpana", 45, "SpaceShuttle", false}
	a3 := Astro{"Jeff Bozos", 52, "Explorer", true}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	p := []Astro{a1, a2,a3}
	fmt.Println(p[2].lastmission)
}

