package main

import "fmt"

func main() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	pickup := truck{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		fourWheel: true,
	}

	camry := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "gold",
		},
		luxury: false,
	}

	fmt.Println(camry)
	fmt.Println(pickup)
	fmt.Println(pickup.color)
	fmt.Println(camry.luxury)
}
