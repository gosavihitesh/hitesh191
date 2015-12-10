package Exercise15
import "fmt"

func main() {
	name1 := []string{"hitesh","gosavi"}
	name2 := []string{"amruta","gosavi"}
	name := append(name1,name2...)
	fmt.Println(name)
}