package core

import (
	"gitlab.com/gelleson/trader/matching-engine-go/src/trading"
	"time"
)

type Match struct {
	Buy   trading.Order
	Sell  trading.Order
	Date  time.Time
	Size  int64
	Price int64
}

type MatchBook struct {

}
