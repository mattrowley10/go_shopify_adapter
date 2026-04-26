package checkout

import (
	"fmt"
)

type CheckoutMutateReq struct {
	CartID string
}

type CheckoutMutateResp struct {
	CartCheckoutURLCreate *CartCheckoutURLCreate `json:"cartCheckoutUrlCreate,omitempty"`
}

type CartCheckoutURLCreate struct {
	Checkout *Checkout `json:"checkout"`
}

type Checkout struct {
	URL string `json:"url"`
}

func Mutate(req *CheckoutMutateReq) string {
	return fmt.Sprintf(`
		mutation {
			cartCheckoutUrlCreate(cartId: "%s") {
				checkoutUrl
			}
		}
	`, req.CartID)
}
