package Exercise15
import "fmt"

func main() {
	name := []string{"dev","saurab","Ajay","vinit","nikhil"}
	fmt.Println(name)
	name = append(name[:2],name[3:]...)
	fmt.Println(name)
}
