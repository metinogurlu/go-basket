package basket

import "encoding/json"

type CustomerBasket struct {
	PromoCode         string        `json:"promoCode"`
	Note              string        `json:"note"`
	ShippingAddressId string        `json:"shippingAddressId"`
	BillingAddressId  string        `json:"billingAddressId"`
	Items             *[]BasketItem `json:"items"`
}

func (b CustomerBasket) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}
