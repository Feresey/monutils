package main

import (
	"github.com/Feresey/monutils/util"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		panic("Want two outputs")
	}

	outs, err := util.GetOutputs()
	if err != nil {
		panic(err)
	}
	var (
		from string
		to   string
	)

	for _, val := range outs {
		if strings.Contains(strings.ToLower(val.Name), strings.ToLower(os.Args[1])) {
			from = val.Name
		} else if strings.Contains(strings.ToLower(val.Name), strings.ToLower(os.Args[2])) {
			to = val.Name
		}
	}

	println("from", from, "to", to)

	if err = util.SwitchOutputs(from, to); err != nil {
		panic(err)
	}
}
