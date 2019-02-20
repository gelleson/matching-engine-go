package core

import "gitlab.com/gelleson/trader/matching-engine-go/src/trading"

type Queue struct {
	Name string
	Order chan trading.Order
}

func NewQueue(name string) *Queue {
	return &Queue{Name: name, Order: make(chan trading.Order, 1)}
}

func (q *Queue) Publish(order trading.Order)  <- chan trading.Order {
	q.Order <- order
	return q.Order
}

func (q *Queue) Subscribe() chan <- trading.Order  {
	return q.Order
}