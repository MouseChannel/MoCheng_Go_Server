package timer

import (
	"time"
)

type MyTimer struct {
	ticker     time.Ticker
	interval   int //单位：纳秒
	callback   func()
	needDelete chan bool
}

func (myTimer *MyTimer) UpdateTask() {
	defer myTimer.ticker.Stop()
	ticks := myTimer.ticker.C

	for range ticks {
		select {
		case <-myTimer.needDelete:
			return
		default:		 
			myTimer.callback()
		}

	}

}

func (myTimer *MyTimer) DeleteTick() {
	myTimer.needDelete <- true
	close(myTimer.needDelete)
}

func (myTimer *MyTimer) AddTickTimerTask(interval int, callback func()) {
	myTimer.ticker = *time.NewTicker(time.Millisecond * time.Duration(interval))
	myTimer.callback = callback
	myTimer.needDelete = make(chan bool)
 
	go myTimer.UpdateTask()

}
