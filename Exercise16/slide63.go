package Exercise16
import "fmt"

func printFriends(names ...string) {
	for _,val := range names{
		fmt.Println(val)
	}
}

func main() {
	friends := []string{"dev","saurab","Ajay","vinit","nikhil"}
	printFriends(friends...)
}
