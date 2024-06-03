package response

import "xtt/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
