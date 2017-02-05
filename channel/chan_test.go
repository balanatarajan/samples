package main

import "testing"
import "github.com/pkg/profile"

func TestChan(t *testing.T) {
	defer profile.Start(profile.CPUProfile).Stop()
	work()
}
