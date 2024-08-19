package main

import (
	"bytes"
	"fmt"
	"slices"
	"testing"
	"time"
)

const (
	write string = "write"
	sleep string = "sleep"
)

type SpySleeper struct {
	Calls int
}

type SpyTime struct {
	durationSlept time.Duration
}

func (st *SpyTime) Sleep(duration time.Duration) {
	st.durationSlept = duration
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationSpy{})

		got := buffer.String()
		want := `3
	2
	1
	Go!`

		if got != want {
			fmt.Printf("got: %q want: %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !slices.Equal(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := time.Second * 5

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{duration: sleepTime, sleep: spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
