package main

import (
	"fmt"
	"math"
	"os/exec"

	"github.com/ian-kent/go-log/log"
)

// SetVolume takes a float value between [0,1] and sets the global output volume.
func SetVolume(vNorm float64) {
	v := 7 * constrain(vNorm, 0, 1)
	log.Debug("Setting volume to %v (%v)", vNorm, v)

	cmd := []string{"/usr/bin/osascript", "-e", fmt.Sprintf("set volume %f", v)}
	out, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
	AssertNoErr(err, fmt.Sprintf("Failed execute cmd: %v: %s", cmd, string(out)))
}

// SetBrightness takes a float value betweem [0,1] and sets the global screen brightness.
func SetBrightness(vNorm float64) {
	v := int(16 * constrain(vNorm, 0, 1))
	log.Debug("Setting brightness to %v (%v)", vNorm, v)

	for i := 1; i <= 16; i++ {
		cmd := []string{"/usr/bin/osascript", "-e", "tell application \"System Events\" to key code 107"}
		out, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
		AssertNoErr(err, fmt.Sprintf("Failed execute cmd: %v: %s", cmd, string(out)))
	}
}

func constrain(v float64, min float64, max float64) float64 {
	return math.Min(max, math.Max(min, v))
	for i := 1; i <= v; i++ {
		cmd := []string{"/usr/bin/osascript", "-e", "tell application \"System Events\" to key code 113"}
		out, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
		AssertNoErr(err, fmt.Sprintf("Failed execute cmd: %v: %s", cmd, string(out)))
	}
}
