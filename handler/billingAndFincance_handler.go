package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// "os"

	"github.com/labstack/echo/v4"
	// paypal "github.com/logpacker/PayPal-Go-SDK"
)

func AcceptPayment(c echo.Context) error {
	url := "https://api.chapa.co/v1/transaction/initialize"
	method := "POST"
	payload := strings.NewReader(`{
		"amount":"100",
		"currency":"ETB",
		"email":"meles.zawdie@gmail.com",
		"first_name":"Meles",
		"last_name":"Zawude",
		"phone_number":"0920227833",
		"tx_ref":"chewatatest-6669",
		"callback_url":"https://webhook.site/077164d6-29cb-40df-ba29-8a00e59a7e60",
		"return_url": "https://www.google.com/",
		"customization[title]": "Payment for my favourite merchant",
		"customization[description]": "I love online payments"
	}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer CHASECK_TEST-MK9OJzrQdg39hUtFwNqjun7zNQtPjctR")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response := map[string]interface{}{
		"resp_body": string(body),
	}

	return c.JSON(http.StatusOK, response)
}

func VerifyPayment(c echo.Context) error {
	url := "https://api.chapa.co/v1/transaction/verify/chewatatest-6669"
	method := "GET"
	payload := strings.NewReader(``)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer CHASECK_TEST-MK9OJzrQdg39hUtFwNqjun7zNQtPjctR")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response := map[string]interface{}{
		"request_body": string(body),
	}
	return c.JSON(http.StatusOK, response)
}

func SplitPayment(c echo.Context) error {
	url := "https://api.chapa.co/v1/subaccount"
	method := "POST"

	payload := strings.NewReader(`{
		"business_name": "Abebe Souq", 
		"account_name": "Abebe Bikila ",
		"bank_code": "971bd28c-ff80-420b-a0db-0a1a4be6ee8b", 
		"account_number": "0123456789", 
		"split_value": 0.2,
		"split_type": "percentage"
		}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer CHASECK_TEST-MK9OJzrQdg39hUtFwNqjun7zNQtPjctR")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response := map[string]interface{}{
		"response_body": string(body),
	}
	return c.JSON(http.StatusOK, response)
}

func GetBank(c echo.Context) error {
	url := "https://api.chapa.co/v1/banks"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer CHASECK_TEST-MK9OJzrQdg39hUtFwNqjun7zNQtPjctR")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response := map[string]interface{}{
		"message_body": string(body),
	}
	return c.JSON(http.StatusOK, response)
}

func TransferBill(c echo.Context) error{
	url := "https://api.chapa.co/v1/transfers"
    method := "POST"
	payload:=strings.NewReader(`{
		"account_name":"Meles Zewude",
		"account_number":"415781533",
        "amount":"1",
        "currency":"ETB",
        "beneficiary_name":"Israel Goytom",
        "reference":"3241342142sfdd",
        "bank_code":"96e41186-29ba-4e30-b013-2ca36d7e7025"
	}`)
	client:=&http.Client{}
	req, err:=http.NewRequest(method, url, payload)
	if err !=nil{
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer CHASECK_TEST-MK9OJzrQdg39hUtFwNqjun7zNQtPjctR")
	req.Header.Add("Content-Type", "application/json")
	 res, err:=client.Do(req)
	 if err !=nil{
		fmt.Println(err)
		return err
	 }
	 defer res.Body.Close()
	 body, err:=ioutil.ReadAll(res.Body)
	 if err !=nil{
		fmt.Println(err)
		return err
	 }
	 response:=map[string]interface{}{
		"message_body":string(body),
	 }
	 return c.JSON(http.StatusOK, response)
}

// type PaymentMethod interface {
// 	Pay(amount float64) string
// }

// type CreditCard struct {
// 	name, cardNumber string
// }

// func (c *CreditCard) Pay(amount float64) string {
// 	return fmt.Sprintf("Paid %.2f using Credit Card (%s)", amount, c.cardNumber)
// }

// type PayPal struct {
// 	email string
// }

// func (p *PayPal) Pay(amount float64) string {
// 	return fmt.Sprintf("Paid %.2f using PayPal (%s)", amount, p.email)
// }

// type Cryptocurrency struct {
// 	walletAddress string
// }

// func (c *Cryptocurrency) Pay(amount float64) string {
// 	return fmt.Sprintf("Paid %.2f using Cryptocurrency (%s)", amount, c.walletAddress)
// }

// type Item struct {
// 	name  string
// 	price float64
// }

// type ShoppingCart struct {
// 	items         []Item
// 	paymentMethod PaymentMethod
// }

// func (s *ShoppingCart) SetPaymentMethod(paymentMethod PaymentMethod) {
// 	s.paymentMethod = paymentMethod
// }

// func (s *ShoppingCart) Checkout() string {
// 	var total float64
// 	for _, item := range s.items {
// 		total += item.price
// 	}
// 	return s.paymentMethod.Pay(total)
// }

// func PayBilling(c echo.Context) error {
// 	creditCard := &CreditCard{"Meles, M.Zawudie", "4111-1111-1111-1111"}
// 	paypal := &PayPal{"meles.zawdie@gmail.com"}
// 	cryptocurrency := &Cryptocurrency{"0xAbcDe1234FghIjKlMnOp"}

// 	shoppingCart := &ShoppingCart{
// 		items: []Item{
// 			{"Laptop", 1500},
// 			{"Smartphone", 1000},
// 		},
// 	}

// 	shoppingCart.SetPaymentMethod(creditCard)
// 	resultC := shoppingCart.Checkout()

// 	shoppingCart.SetPaymentMethod(paypal)
// 	resultP := shoppingCart.Checkout()

// 	shoppingCart.SetPaymentMethod(cryptocurrency)
// 	resultCR := shoppingCart.Checkout()

// 	// Concatenate the results into a single string and return it
// 	response := resultC + "\n" + resultP + "\n" + resultCR
// 	return c.String(http.StatusOK, response)
// }
