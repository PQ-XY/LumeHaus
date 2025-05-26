package internal

import (
   "fmt"

   "estore/util"

   "github.com/stripe/stripe-go/v74"
   "github.com/stripe/stripe-go/v74/price"
   "github.com/stripe/stripe-go/v74/product"
   "github.com/stripe/stripe-go/v74/checkout/session"
)

func CreateProductWithPrice(config *util.StripeInfo, appTitle string, appDescription string, appPrice int64) (productID, priceID string, err error) {
   stripe.Key = config.APIKey
   product_params := &stripe.ProductParams{
       Name:        &appTitle,
       Description: &appDescription,
   }
   newProduct, err := product.New(product_params)
   if err != nil {
       fmt.Println("Failed to create product:" + err.Error())
       return "", "", err
   }

   price_params := &stripe.PriceParams{
       Currency:   stripe.String(string(stripe.CurrencyUSD)),
       Product:    stripe.String(newProduct.ID),
       UnitAmount: &appPrice,
   }
   newPrice, err := price.New(price_params)
   if err != nil {
       fmt.Println("Failed to create price:" + err.Error())
       return "", "", err
   }

   fmt.Println("Success! Here is your product id: " + newProduct.ID)
   fmt.Println("Success! Here is your price id: " + newPrice.ID)

   return newProduct.ID, newPrice.ID, nil
}

func CreateCheckoutSession(config *util.StripeInfo, domain string, priceID string) (*stripe.CheckoutSession, error) {
   stripe.Key = config.APIKey
   params := &stripe.CheckoutSessionParams{
       LineItems: []*stripe.CheckoutSessionLineItemParams{
          {
               Price:    &priceID,
               Quantity: stripe.Int64(1),
           },
       },
       Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
       SuccessURL: stripe.String(domain + "?success=true"),
       CancelURL:  stripe.String(domain + "?canceled=true"),
   }

   s, err := session.New(params)

   if err != nil {
       fmt.Printf("session.New: %v", err)
       return nil, err
   }
   return s, err
}

