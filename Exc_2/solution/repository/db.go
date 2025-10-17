package repository

import "ordersystem/model"

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	drinks := []model.Drink{
		{
			Description: "Coca Cola",
			ID:          1,
			Name:        "Coke",
			Price:       1.50,
		},
		{
			Description: "Freshly brewed coffee",
			ID:          2,
			Name:        "Coffee",
			Price:       2.00,
		},
		{
			Description: "Orange juice from concentrate",
			ID:          3,
			Name:        "OJ",
			Price:       1.75,
		},
	}

	orders := []model.Order{
		{
			Amount:    1,
			CreatedAt: "2024-01-01T10:00:00Z",
			DrinkID:   1,
		},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += order.Amount
	}
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice
}
