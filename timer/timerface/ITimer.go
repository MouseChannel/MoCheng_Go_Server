package timerface

type ITimer interface {
	AddTickTimerTask(interval int, callback func())
	UpdateTask()
	DeleteTick()
}
