package packageone

import "fmt"

var privateVar = "I am private"

var PublicVar = "I am public (or exported)"

func notExported() {

}

func Exported(myVar, blockVar string) {
	fmt.Println(myVar, ",", blockVar, ",", PublicVar)
}
