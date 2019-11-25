package main

import "fmt"

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

func main() {
	hilux := truck{
		fourWheel: true,
		vehicle: vehicle{
			color: "balck",
			doors: 4,
		},
	}

	car := sedan{
		luxury: false,
		vehicle: vehicle{
			color: "OxBlood",
			doors: 4,
		},
	}

	fmt.Println(hilux)
	fmt.Println(car)
	fmt.Println(hilux.fourWheel)
	fmt.Println(car.luxury)
}
