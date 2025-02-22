package pets

import (
	"fmt"
	"strings"
	"time"
)

type Dog struct {
	// variable are exported and accessible outside the package.
	Name  string
	Breed string

	// encapsulation
	// unexported and only accessible within the package.
	lastSlept time.Time
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
// ======== encapsulation =========
// ============= #### =============
// needSleep and sleep are unexported methods (start with lowercase letters)
// and can only be called from within the pets package.

func (d *Dog) needSleep() bool {
	fmt.Printf("== Last Slept: %s\n", d.lastSlept)
	fmt.Println("== Time now:", time.Now())

	return time.Now().Sub(d.lastSlept) > 4*time.Hour
}

func (d *Dog) sleep() {
	d.lastSlept = time.Now()
}
