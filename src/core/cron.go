package core

import (
	"sync"
	"time"
)

type Cron struct {
	once sync.Once
	Name string
	Mask []Day
	Repeat bool
	Serve func()
}


type Day struct {
	Days time.Weekday
	Hour int
	Minute int
	Second int
}