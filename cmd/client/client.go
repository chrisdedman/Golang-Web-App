package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func MakeAuthenticatedRequest() {
	url := "http://localhost:4000/"
	token := "admin"
	email := "admin@example.com"

	// Prepare request body
	reqBody := []byte("")

	// Create a new request
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("AuthToken", token)
	req.Header.Set("AuthEmail", email)

	// Create a client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print response
	fmt.Println("Response Status:", resp.Status)
}
