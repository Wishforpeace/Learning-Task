package handle

import "fmt"

func Error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
