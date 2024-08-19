package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

type DefaultSleeper struct{}

func (ds DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{duration: time.Second * 3, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
