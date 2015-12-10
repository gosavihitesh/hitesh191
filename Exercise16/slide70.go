package Exercise16
import "fmt"

func main() {
	fmt.Println("Hello, hitesh")
	defer func() { fmt.Println("Goodbye, hitesh") }()
	fmt.Println("Do some things")
	fmt.Println("Execute some code")
	fmt.Println("Finish some things up")
}
