package main

import (
	"log"
	"github.com/Feresey/monutils/util"
	"time"
)

func main() {
	outputs := make(map[string]bool)
	current := ""
	currOutputs, err := util.GetOutputs()
	if err != nil {
		log.Panicln(err)
	}
	for _, val := range currOutputs {
		outputs[val.Name] = val.Connected
		if val.Connected {
			current = val.Name
		}
		log.Printf(`output: %q, connected: %t`, val.Name, val.Connected)
	}
	defaultOutput := current
	log.Printf(`current output: %q`, current)

	for range time.Tick(5 * time.Second) {
		currOutputs, err := util.GetOutputs()
		if err != nil {
			log.Panicln(err)
		}
		for _, val := range currOutputs {
			if val.Connected == outputs[val.Name] {
				continue
			}
			switch val.Connected {
			case true:
				log.Printf(`%q connected`, val.Name)
				if err := util.SwitchOutputs(current, val.Name); err != nil {
					log.Panicf("Failed to switch outputs: %q", err)
				}
				current = val.Name
			case false:
				log.Printf(`%q disconnected`, val.Name)
				if val.Name == current {
					if err := util.SwitchOutputs(val.Name, defaultOutput); err != nil {
						log.Panicf("Failed to switch outputs: %q", err)
					}
					current = defaultOutput
				}
			}
			outputs[val.Name] = val.Connected
		}
	}
}
