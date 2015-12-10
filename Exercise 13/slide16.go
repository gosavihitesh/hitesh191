package slide16
import "fmt"

func printStrings(a ...string){
	for _,val := range a{
		fmt.Println(val)
	}
}

func main() {
	a:="hitesh"
	b:="gosavi"
	c:= []string{"ajay","nikhil","dev"}
	printStrings(a,b)
	printStrings(c...)
}
