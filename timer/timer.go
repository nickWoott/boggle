package timer

import (
	"fmt"
	"time"
)

type GameTimer struct {
	timerDuration time.Duration
	stopChan      chan struct{}
}

func NewGameTimer(duration time.Duration) *GameTimer {
	return &GameTimer{
		timerDuration: duration,
		stopChan:      make(chan struct{}),
	}
}

func (t *GameTimer) StartCountdown() {
	expirationTime := time.Now().Add(t.timerDuration)
	timer := time.NewTimer(t.timerDuration)

	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Printf("\rTime's up! The game has ended.    \n")
				return
			case <-t.stopChan:
				return
			default:
				remaining := expirationTime.Sub(time.Now())

				if remaining <= 0 {
					fmt.Printf("\rTime's up! The game has ended.    \n")
					return
				}

				minutes := int(remaining.Minutes())
				seconds := int(remaining.Seconds()) % 60
				fmt.Printf("\rTime remaining: %02d:%02d", minutes, seconds)

				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}
