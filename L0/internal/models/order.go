package models

import "encoding/json"

type Order struct {
	OrderUID        string   `json:"OrderUID"`
	TrackNumber     string   `json:"TrackNumber"`
	Entry           string   `json:"Entry"`
	Delivery        Delivery `json:"Delivery"`
	Payment         Payment  `json:"Payment"`
	Items           []Item   `json:"Items"`
	Locale          string   `json:"Locale"`
	InternalSig     string   `json:"InternalSig"`
	CustomerID      string   `json:"CustomerID"`
	DeliveryService string   `json:"DeliveryService"`
	ShardKey        string   `json:"ShardKey"`
	SmID            int      `json:"SmID"`
	DateCreated     string   `json:"DateCreated"`
	OofShard        string   `json:"OofShard"`
}

type Delivery struct {
	Name    string `json:"Name"`
	Phone   string `json:"Phone"`
	Zip     string `json:"Zip"`
	City    string `json:"City"`
	Address string `json:"Address"`
	Region  string `json:"Region"`
	Email   string `json:"Email"`
}

type Payment struct {
	Transaction  string `json:"Transaction"`
	RequestID    string `json:"RequestID"`
	Currency     string `json:"Currency"`
	Provider     string `json:"Provider"`
	Amount       int    `json:"Amount"`
	PaymentDt    int64  `json:"PaymentDt"`
	Bank         string `json:"Bank"`
	DeliveryCost int    `json:"DeliveryCost"`
	GoodsTotal   int    `json:"GoodsTotal"`
	CustomFee    int    `json:"CustomFee"`
}

type Item struct {
	ChrtID      int    `json:"ChrtID"`
	TrackNumber string `json:"TrackNumber"`
	Price       int    `json:"Price"`
	Rid         string `json:"Rid"`
	Name        string `json:"Name"`
	Sale        int    `json:"Sale"`
	Size        string `json:"Size"`
	TotalPrice  int    `json:"TotalPrice"`
	NmID        int    `json:"NmID"`
	Brand       string `json:"Brand"`
	Status      int    `json:"Status"`
}

// MarshalOrder превращает Order в JSON
func MarshalOrder(order *Order) ([]byte, error) {
	return json.Marshal(order)
}

// UnmarshalOrder распаковывает JSON в Order
func UnmarshalOrder(data []byte, order *Order) error {
	return json.Unmarshal(data, order)
}
