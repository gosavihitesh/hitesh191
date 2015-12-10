package Exercise15
import "fmt"

type customer struct {
	name string
	id int
}

func (c customer) printCustomer(){
	fmt.Println("Name:",c.name)
	fmt.Println("ID:",c.id)
}

func main() {
	cust1 := customer{"hitesh",100}
	cust2 := customer{"gosavi",200}
	cust1.printCustomer()
	cust2.printCustomer()
	cust2.id = 300
	cust2.printCustomer()
}
