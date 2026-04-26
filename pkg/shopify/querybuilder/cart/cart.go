package cart

import (
	"fmt"
)

type CartQueryReq struct {
	CartID string
}

type CartQueryResp struct {
	Cart *Cart `json:"cart"`
}

type Cart struct {
	ID            string     `json:"id"`
	Lines         []CartLine `json:"lines"`
	EstimatedCost CartCost   `json:"estimatedCost"`
}

type CartLine struct {
	ID          string      `json:"id"`
	Quantity    int         `json:"quantity"`
	Merchandise Merchandise `json:"merchandise"`
}

type Merchandise struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
}

type CartCost struct {
	TotalAmount CartPrice `json:"totalAmount"`
}

type CartPrice struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
}

type CartMutateReq struct {
	CartID string
	Lines  []CartLineInput
}

type CartLineInput struct {
	VariantID string
	Quantity  int
}

type CartMutateResp struct {
	CartLinesAdd    *CartLinesAddPayload    `json:"cartLinesAdd,omitempty"`
	CartLinesUpdate *CartLinesUpdatePayload `json:"cartLinesUpdate,omitempty"`
	CartLinesRemove *CartLinesRemovePayload `json:"cartLinesRemove,omitempty"`
}

type CartLinesAddPayload struct {
	Cart *Cart `json:"cart"`
}

type CartLinesUpdatePayload struct {
	Cart *Cart `json:"cart"`
}

type CartLinesRemovePayload struct {
	Cart *Cart `json:"cart"`
}

func Query(req *CartQueryReq) string {
	return fmt.Sprintf(`
		query {
			cart(id: "%s") {
				id
				lines(first: 10) {
					edges {
						node {
							id
							quantity
							merchandise {
								... on ProductVariant {
									id
									title
									price {
										amount
									}
								}
							}
						}
					}
				}
				estimatedCost {
					totalAmount {
						amount
						currencyCode
					}
				}
			}
		}
	`, req.CartID)
}

func Mutate(req *CartMutateReq) string {
	linesStr := ""
	for _, line := range req.Lines {
		linesStr += fmt.Sprintf(`{merchandiseId: "%s", quantity: %d}`, line.VariantID, line.Quantity)
	}

	return fmt.Sprintf(`
		mutation {
			cartLinesAdd(cartId: "%s", lines: [%s]) {
				cart {
					id
					lines(first: 10) {
						edges {
							node {
								id
								quantity
								merchandise {
									... on ProductVariant {
										id
										title
										price {
											amount
										}
									}
								}
							}
						}
					}
					estimatedCost {
						totalAmount {
							amount
							currencyCode
						}
					}
				}
			}
		}
	`, req.CartID, linesStr)
}
