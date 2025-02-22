package pets

import (
	"fmt"
	"strings"
	"time"
)

// ============= #### =============
// ===== exported/unexported ======
// ============= #### =============
// variable start Capital: are exported and accessible outside the package.
// variable start camelCase: are unexported and only accessible within the package.

type Dog struct {
	Name  string
	Breed string

	lastSlept time.Time // unexported

	Animal // composition
}

func (d Dog) Feed(food string) string {
	return fmt.Sprintf("Dog with Name %s: is eating => %s ", d.Name, strings.ToUpper(food))
}

func (d *Dog) GivenAttention() string {
	if d.needSleep() {
		d.sleep()
		return "Your Dog iS Tired..."
	}
	return "Your Dog is Active!"
}

// ============= #### =============
// ======== Encapsulation =========
// ============= #### =============
// needSleep are unexported methods (start with lowercase letters)
// and can only be called from within the pets package.

func (d *Dog) needSleep() bool {
	fmt.Printf("\n\n== Last Slept: %s\n", d.lastSlept)
	fmt.Printf("== Time now: %s\n\n", time.Now())

	return time.Now().Sub(d.lastSlept) > 4*time.Hour
}

func (d *Dog) sleep() {
	d.lastSlept = time.Now()
}

// ============= #### =============
// ========= Composition ==========
// ============= #### =============

func GuardDog(name, breed string, lastSlept time.Time) Dog {
	return Dog{
		Name:  name,
		Breed: breed,

		lastSlept: lastSlept,

		Animal: Animal{lastAte: time.Now()},
	}
}
