package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rwirdemann/marketsync/application/domain"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Marketplace struct {
	accessToken string
}

func (m *Marketplace) CreateOrUpdateProduct(product domain.Product) error {
	products := []domain.Product{
		product,
	}

	b, err := json.Marshal(products)
	if err != nil {
		return err
	}
	apiUrl := os.Getenv("API_URL")
	uri := fmt.Sprintf("%s/v3/products", apiUrl)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.accessToken))
	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	b, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Printf("POST /v3/products %v status Code: %d body %s", product, res.StatusCode, b)
	return nil
}

func (m *Marketplace) Login() error {
	apiUrl := os.Getenv("API_URL")
	uri := fmt.Sprintf("%s/v1/token", apiUrl)
	data := url.Values{}
	data.Set("username", os.Getenv("USERNAME"))
	data.Set("password", os.Getenv("PASSWORD"))
	data.Set("grant_type", "password")
	data.Set("client_id", os.Getenv("CLIENT_ID"))

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, uri, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		log.Printf("POST %s failed", apiUrl)
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var response = struct {
		AccessToken string `json:"access_token"`
	}{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("POST %s failed", apiUrl)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("POST %s failed", apiUrl)
		return errors.New("Login failed")
	}
	m.accessToken = response.AccessToken
	log.Printf("POST %s success", apiUrl)
	return nil
}
