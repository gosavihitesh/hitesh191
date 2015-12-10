package Exercise15
import "fmt"

func main() {
	ages := map[string]int{
		"hitesh":26,
		"amruta":27,
		"harshad":24,
		"Babygirl":1,
		"bhavesh":23,
	}
	ages["arjun"]=14
	ages["hitesh"]=26
	if _, exists := ages["bhavesh"]; exists {
		delete(ages, "bhavesh")
	}
	for key,val := range ages{
		fmt.Println(key,"is",val,"years old.")
	}
	fmt.Println("Length of map is",len(ages))
}
