package Exercise14
import "fmt"

func main() {
	name := "My name is Hitesh Gosavi"
	fmt.Println("Initials:",string(name[11]),string(name[17]))
	fmt.Println("First Name:",name[11:16])
}
