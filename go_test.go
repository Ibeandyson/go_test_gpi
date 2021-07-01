// I imported the fmt package
import "fmt"

type Cement struct {
	NoOfCement int
}

func (c *Cement) BuyCement(howmany int) {
	c.NoOfCement += howmany
}

// the o is not defined but c was never used. Now I used the c
func (c *Cement) SellCement(howmany int) {
	c.NoOfCement -= howmany
}

// the  fmt will undfined, becuse i was not was not import
func (c *Cement) String() string {
	return fmt.Sprintf("%v", c.NoOfCement)
}

// the cememet variables missed the = to assign the value of  Cement. Now i wil put the =

// the  fmt will undfined, becuse i was not was not import
func main() {
	var cement =  Cement
	cement.BuyCement(15)
	cement.SellCement(9)
	fmt.Println(cement)
}