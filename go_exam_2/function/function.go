package function

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ValidateThailandCitizenID(idNo string) error {

	if len(idNo) != 13 {
		return errors.New("thai citizen no is not 13 digits")
	}

	_, err := strconv.Atoi(idNo)
	if err != nil {
		return errors.New("thai citizen no must be numeric only")
	}

	idNoSp := idNo[:12]
	//fmt.Println("idNoSp :", idNoSp)

	sum := 0

	for i := 0; i < 12; i++ {
		nd, err := strconv.Atoi(idNoSp[i : i+1])
		if err != nil {
			return err
		}
		//fmt.Println("num digit :", nd, " calculate by : ", nd*(13-i))
		sum += nd * (13 - i)
	}
	//fmt.Println("total summary of num :", sum)
	//fmt.Println("total summary of num with mod :", sum%11)

	ld := sum % 11
	ldStr := ""

	ld = 11 - ld

	if ld > 9 {
		ldStr = fmt.Sprintf("%d", ld)
		ldStr = ldStr[1:]
	} else {
		ldStr = fmt.Sprintf("%d", ld)
	}

	calIdno := fmt.Sprintf("%s%s", idNo[:12], ldStr)

	fmt.Println("calculate identification id :", calIdno)
	fmt.Println("    input identification id :", idNo)

	if strings.Compare(calIdno, idNo) != 0 {
		return errors.New("thai citizen id mismatch")
	}

	return nil
}
