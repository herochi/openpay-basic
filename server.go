package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	privateKey string = "llave privada"
	merchanID  string = "id del comercio"
	clientID   string = "client id  del cliente creado"
	cardID     string = "card id del cliente"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/create-client", createClient)
	http.HandleFunc("/create-card", createCard)
	http.HandleFunc("/pay", pay)

	addr := "localhost:4242"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func createClient(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:4242/register.html"
	err := prepareRequestClient(r)

	if err != nil {
		fmt.Println("error: ", err)
		http.Redirect(w, r, domain, http.StatusBadRequest)
	}
	http.Redirect(w, r, domain, http.StatusOK)
}

func prepareRequestClient(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	nombre := r.Form.Get("nombre")
	email := r.Form.Get("email")

	postBody, _ := json.Marshal(map[string]string{
		"name":             nombre,
		"email":            email,
		"requires_account": "false",
	})

	url := fmt.Sprintf("https://sandbox-api.openpay.mx/v1/%s/customers", merchanID)

	err = doRequest(url, postBody)
	if err != nil {
		return err
	}
	return nil
}

func doRequest(url string, body []byte) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.SetBasicAuth(privateKey, "")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("el error fue: ", err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	fmt.Println("RESPONSE: ", s)

	if resp.Status >= "400 Bad Request" {
		fmt.Println("error: ", s)
		return errors.New("an error ocurred")
	}

	return nil
}

func createCard(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:4242/card.html"
	err := prepareRequestCard(r)

	if err != nil {
		fmt.Println("error: ", err)
		http.Redirect(w, r, domain, http.StatusBadRequest)
	}
	http.Redirect(w, r, domain, http.StatusOK)
}

func prepareRequestCard(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	card1 := r.Form.Get("card1")
	card2 := r.Form.Get("card2")
	card3 := r.Form.Get("card3")
	card4 := r.Form.Get("card4")
	card := card1 + card2 + card3 + card4
	holder := r.Form.Get("holder")
	expMonth := r.Form.Get("expiration-month")
	expYear := r.Form.Get("expiration-year")
	cvv := r.Form.Get("cvv")

	postBody, _ := json.Marshal(map[string]string{
		"card_number":      card,
		"holder_name":      holder,
		"expiration_year":  expYear[len(expYear)-2:],
		"expiration_month": expMonth,
		"cvv2":             cvv,
	})

	url := fmt.Sprintf("https://sandbox-api.openpay.mx/v1/%s/customers/%s/cards", merchanID, clientID)

	err = doRequest(url, postBody)
	if err != nil {
		return err
	}
	return nil
}

func pay(w http.ResponseWriter, r *http.Request) {
	domain := "http://localhost:4242/pay.html"

	err := prepareRequestPay(r)

	if err != nil {
		fmt.Println("error: ", err)
		http.Redirect(w, r, domain, http.StatusBadRequest)
	}
	http.Redirect(w, r, domain, http.StatusOK)
}

func prepareRequestPay(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	postBody, _ := json.Marshal(map[string]string{
		"source_id":         cardID,
		"method":            "card",
		"amount":            "200",
		"currency":          "MXN",
		"description":       "Cargo inicial a mi cuenta",
		"order_id":          strconv.Itoa(randomNum()),
		"device_session_id": "AJSANDB1212KE1N",
	})

	url := fmt.Sprintf("https://sandbox-api.openpay.mx/v1/%s/customers/%s/charges", merchanID, clientID)

	err = doRequest(url, postBody)
	if err != nil {
		return err
	}
	return nil
}

func randomNum() int {
	//semilla a nivel nanosegundo en unix time
	rand.Seed(time.Now().UTC().UnixNano())

	return rand.Intn(1000) // valor entre 0 - 1000
}
