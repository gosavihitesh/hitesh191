package Exercise16
import "fmt"

func run(f func(string), names ...string) {
	for _, val := range names{
		f(val)
	}
}

func main() {
	names := []string{"dev","saurab","Ajay","vinit","nikhil"}
	greet := func(name string){fmt.Println("Hello,",name)}
	run(greet,names...)
}
