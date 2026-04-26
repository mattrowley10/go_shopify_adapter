package customers

import "fmt"

type CustomerQueryReq struct {
	CustomerID string
}

type CustomerQueryResp struct {
	Customer *Customer `json:"customer"`
}

type Customer struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func Query(req *CustomerQueryReq) string {
	return fmt.Sprintf(`
		query {
			customer(id: "%s") {
				id
				email
			}
		}
	`, req.CustomerID)
}
