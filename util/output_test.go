package util

import (
	"reflect"
	"testing"
)

func TestGetOutputs(t *testing.T) {
	gotRes, err := GetOutputs()
	if err != nil {
		t.Errorf("error = %v", err)
		return
	}
	res := []Output{Output{Name: "eDP-1", Connected: true}, Output{Name: "HDMI-1"}}
	if !reflect.DeepEqual(gotRes, res) {
		t.Errorf("GetOutputs() = %v, want %v", gotRes, res)
	}
}
