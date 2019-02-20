package trading

import (
	"sync"
	"time"
)

const (
	SELL = true
	BUY  = false
)

type Order struct {
	ID            int64
	TypeOperation bool
	Created       time.Time
	Instrument    int64
	Price         int64
	Size          int64
	Broker	 	  int64
	Invest		  int64
}

func (o *Order) setSequence(id int64)  {
	o.ID = id
}

func NewOrder(typeOperation bool, price, size,broker, invest, instrument int64) Order {
	date := time.Now()
	return Order{TypeOperation: typeOperation, Price: price, Size: size, Created: date, Instrument:instrument, Broker: broker, Invest: invest}
}

type OrderBook struct {
	seq     int64
	mutex   sync.Mutex
	Buy     []Order
	Sell    []Order
}

func NewOrderBook(seq int64) *OrderBook {
	return &OrderBook{seq: seq}
}

func (ob *OrderBook) AddOrder(order Order)  {
	ob.mutex.Lock()
	defer ob.mutex.Unlock()
	order.setSequence(ob.seq)
	if !order.TypeOperation  {
		(*ob).Buy = append((*ob).Buy, order)
	} else {
		(*ob).Sell = append((*ob).Buy, order)
	}

	(*ob).seq++
}

func (ob *OrderBook) SetSequence(seq int64)  {
	(*ob).seq = seq
}
