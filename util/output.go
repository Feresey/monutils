package util

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var r = regexp.MustCompile(`(?i)^(.*)\s(.*connected)`)

type Output struct {
	Name      string
	Connected bool
}

func GetOutputs() (res []Output, err error) {
	out, err := exec.Command("xrandr").Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(out), "\n")
	for _, val := range lines {
		matches := r.FindStringSubmatch(val)
		if len(matches) < 2 {
			continue
		}
		res = append(res, Output{matches[1], matches[2] == "connected"})
	}
	return
}

func SwitchOutputs(from, to string) error {
	off := exec.Command("xrandr", "--output", from, "--off")
	off.Stdout = os.Stdout
	if err := off.Run(); err != nil {
		return err
	}
	connect := exec.Command("xrandr", "--output", to, "--auto", "--primary")
	connect.Stdout = os.Stdout
	if err := connect.Run(); err != nil {
		return err
	}

	return nil
}
