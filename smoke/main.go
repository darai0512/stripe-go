package main

import (
	"fmt"
	stripe "github.com/stripe/stripe-go/v73"
	subscription "github.com/stripe/stripe-go/v73/subscription"
	customer "github.com/stripe/stripe-go/v73/customer"
)

func main() {
	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	status := "active"
	params := &stripe.SubscriptionListParams{
		Status: &status,
	}
	params.Filters.AddFilter("limit", "", "3")
	// params.Filters.AddFilter("offset", "", "0")
	i := subscription.List(params)
	for i.Next() {
	    s := i.Subscription()
	    fmt.Println(s.Customer)
	    sc, _ := subscription.Cancel(s.ID, nil)
	    fmt.Println(sc)
		fmt.Println(sc.Customer)
	    c, err := customer.Del(s.Customer.ID, nil)
	    fmt.Println(err)
		fmt.Println(c.ID)
	    fmt.Println(s.Customer)

	    break
	}
}
