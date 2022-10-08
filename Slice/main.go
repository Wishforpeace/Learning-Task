package main

import "fmt"

type Image struct {
	ID     string
	Avatar string
}
type User struct {
	ID   string
	Name string
	Im   Image
}

func main() {

	var u1 = User{

		Name: "hh",
	}

	var u2 = User{
		Name: "jj",
	}

	var image1 = Image{
		ID:     "1",
		Avatar: "11",
	}
	var image2 = Image{
		ID:     "2",
		Avatar: "22",
	}

	U := []User{
		u1, u2,
	}

	I := []Image{
		image1, image2,
	}

	for i, _ := range U {
		U[i].Im.Avatar = I[i].Avatar
	}
	fmt.Println(U)

	u1.Im = image2
	fmt.Println(u1)
}
