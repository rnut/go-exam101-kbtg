package main

import (
	"fmt"
	"go-basic-exam/go_exam_2/function"
)

func main() {
	if err := function.ValidateThailandCitizenID("1516712728928"); err != nil {
		fmt.Println("error validate thai citizen id : ", err)
		return
	}
	fmt.Println("validate thai citizen id success")
}
