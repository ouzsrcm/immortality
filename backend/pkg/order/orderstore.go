package order

import "immortality/pkg/common"

type OrderStore struct {
	common.StoreBase
}

func NewOrderStore() *OrderStore {
	store := new(OrderStore)
	store.Connect()
	return store
}

type IOrderStore interface {
	GetOrderById(id uint) (*Order, error)
	GetOrderInfoById(id uint) (*OrderInfo, error)
	GetOrderItemById(id uint) (*OrderItem, error)
	GetOrderItemsById(id uint) (*[]OrderItem, error)

	GetBasketById(id uint) (*Basket, error)
	GetBasketInfoById(id uint) (*BasketInfo, error)
	GetBasketItemById(id uint) (*BasketItem, error)
	GetBasketItemsById(id uint) (*[]BasketItem, error)

	CreateOrder(order *Order) error
	CreateOrderInfo(orderInfo *OrderInfo) error
	CreateOrderItem(orderItem *OrderItem) error

	CreateBasket(basket *Basket) error
	CreateBasketInfo(basketInfo *BasketInfo) error
	CreateBasketItem(basketItem *BasketItem) error

	UpdateOrder(order *Order) error
	UpdateOrderInfo(orderInfo *OrderInfo) error
	UpdateOrderItem(orderItem *OrderItem) error

	UpdateBasket(basket *Basket) error
	UpdateBasketInfo(basketInfo *BasketInfo) error
	UpdateBasketItem(basketItem *BasketItem) error

	DeleteOrder(order *Order) error
	DeleteOrderInfo(orderInfo *OrderInfo) error
	DeleteOrderItem(orderItem *OrderItem) error

	DeleteBasket(basket *Basket) error
	DeleteBasketInfo(basketInfo *BasketInfo) error
	DeleteBasketItem(basketItem *BasketItem) error

	GetOrders() (*[]Order, error)
	GetOrderInfosByOrderId(orderId uint) (*[]OrderInfo, error)
	GetOrderItemsByOrderId(orderId uint) (*[]OrderItem, error)

	GetBaskets() (*[]Basket, error)
	GetBasketInfosByBasketId(basketId uint) (*[]BasketInfo, error)
	GetBasketItemsByBasketId(basketId uint) (*[]BasketItem, error)
}
