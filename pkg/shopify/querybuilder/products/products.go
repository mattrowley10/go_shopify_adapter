package products

import "fmt"

type ProductQueryReq struct {
	ProductID string
}

type VariantQueryReq struct {
	VariantID string
}

type ProductQueryResp struct {
	Product *Product `json:"product"`
}

type Product struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Handle   string    `json:"handle"`
	Variants []Variant `json:"variants"`
	Images   []Image   `json:"images"`
}

type Variant struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Price     string `json:"price"`
	Available bool   `json:"available"`
}

type Image struct {
	URL string `json:"url"`
}

func Query(req *ProductQueryReq) string {
	return fmt.Sprintf(`
		query {
			product(id: "%s") {
				id
				title
				handle
				variants(first: 10) {
					edges {
						node {
							id
							title
							price
							available
						}
					}
				}
				images(first: 10) {
					edges {
						node {
							url
						}
					}
				}
			}
		}
	`, req.ProductID)
}
