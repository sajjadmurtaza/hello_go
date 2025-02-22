package pets

import "time"

type Cat struct {
	Name  string
	Breed string

	Animal
}

func (d *Cat) GivenAttention() string {
	return "Your Cat is Active!"
}

func AngryCat(name, breed string) Pet {
	return &Cat{
		Name:  name,
		Breed: breed,

		Animal: Animal{lastAte: time.Now()},
	}
}
