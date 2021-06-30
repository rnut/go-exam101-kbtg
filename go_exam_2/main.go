package main

import (
	"fmt"
	"go-exam101-kbtg/go_exam_2/function"
)

func main() {
	if err := function.ValidateThailandCitizenID("1100500625945"); err != nil {
		fmt.Println("error validate thai citizen id : ", err)
		return
	}
	fmt.Println("validate thai citizen id success")
}
