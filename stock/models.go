package stock

type Stock struct {
	OrderID int `json:"orderid""`
	Time string `json:"time"`
	Name string `json:"name"`
	Action string `json:"action"`
	Quantity int `json:"qty"`
	Price float32 `json:"price"`
}

type Profile struct {
	Name string
	TotalStocksPurchased int
	Prices []int
}
