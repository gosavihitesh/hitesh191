package Exercise16
import (
	"strconv"
	"fmt"
)

func dogYears(age int) int {
	return 7*age
}

func isOld(age int) string {
	if age > 25 {
		return "is old"
	}else {
		return "is not old"
	}
}

func sentence(name string, age int) string {
	return name+" is "+strconv.Itoa(dogYears(age))+" in dog years and "+isOld(age)
}

func main() {
	fmt.Println(sentence("bholu",20))
	fmt.Println(sentence("petu",40))
}
