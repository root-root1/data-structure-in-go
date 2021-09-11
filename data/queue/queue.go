package queue

type Queue []*Order

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func (order *Order) New(priority, quantity int, product, customerName string) {
	order.priority = priority
	order.product = product
	order.quantity = quantity
	order.customerName = customerName
}

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool
		appended = false
		for index, addedOrder := range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:index], append(Queue{order}, (*queue)[index:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}
