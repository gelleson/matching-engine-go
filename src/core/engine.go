package core

import (
	"gitlab.com/gelleson/trader/matching-engine-go/src/trading"
)

var Sequence *trading.OrderBook

type Engine struct {
	queue bool
	started bool
	config Config
	cron []Cron
}

func (e *Engine) SetConfig(config Config)  {
	(*e).config = config
}

func (e *Engine) SetCron(cron ...Cron)  {
	(*e).cron = cron
}

func (e *Engine) Serve()  {
	queue := NewQueue("transaction")
	for {
		queue.Publish(trading.Order{})
	}
}