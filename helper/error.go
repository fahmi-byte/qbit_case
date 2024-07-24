package helper

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Println(err, "error apa")
		panic(err)
	}
}
