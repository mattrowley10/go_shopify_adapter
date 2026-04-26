package collections

import (
	"fmt"
)

type CollectionQueryReq struct {
	CollectionID string
}

type CollectionQueryResp struct {
	Collection *Collection `json:"collection"`
}

type Collection struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Handle string `json:"handle"`
}

func Query(req *CollectionQueryReq) string {
	return fmt.Sprintf(`
		query {
			collection(id: "%s") {
				id
				title
				handle
			}
		}
	`, req.CollectionID)
}
