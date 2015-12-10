package Exercise16
import "fmt"

func hello() func() {
	return func() { fmt.Println("Hello") }
}

func main() {
	sayHello := hello()
	sayHello()
}
