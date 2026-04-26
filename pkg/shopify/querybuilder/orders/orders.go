package orders

import (
	"fmt"
)

type OrderQueryReq struct {
	OrderID string
}

type OrderQueryResp struct {
	Order *Order `json:"order"`
}

type Order struct {
	ID            string     `json:"id"`
	OrderNumber   string     `json:"orderNumber"`
	StatusPageURL string     `json:"statusPageUrl"`
	TotalPrice    OrderPrice `json:"totalPrice"`
	CreatedAt     string     `json:"createdAt"`
}

type OrderPrice struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
}

func Query(req *OrderQueryReq) string {
	return fmt.Sprintf(`
		query {
			order(id: "%s") {
				id
				orderNumber
				statusPageUrl
				totalPrice {
					amount
					currencyCode
				}
				createdAt
			}
		}
	`, req.OrderID)
}
