package main

import (
	"os"
	"tdd/mocking"
	"time"
)

func main() {
	sleeper := mocking.MakeSleeper(2*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, &sleeper)
}
