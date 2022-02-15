package mystring

import (
	"fmt"
	"regexp"
)

func aa() {

	r, _ := regexp.Compile("202108.*/*history.dat")

	fmt.Println(r.MatchString("/home/scada/data/plc-sync/20210818/history/dat/TYSFCB_149_2021_06_history.dat"))

}
