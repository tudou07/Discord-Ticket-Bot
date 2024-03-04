package main

import (
	"bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

const (
	RepoOwner 	= "Rai-Sahil"
	RepoName 	= "Machine-Learning"
	Token		= "TOKEN_HERE"
)

type Ticket struct {
    Title       string   `json:"title"`
    Description string   `json:"body"`
    Labels      []string `json:"labels"`
}

func Welcome() string {
	var input string

	fmt.Println("Welcome!!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Make a new ticket")
	fmt.Println("2. Close")
	fmt.Println("")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "err"
	}

	return input
}

func createTicket(repoOwner string, repoName string, token string, ticket Ticket) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", repoOwner, repoName)

	body, err := json.Marshal(ticket)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        return err
    }

	req.Header.Set("Authorization", "token " + token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("failed to create ticket, status code: %d", resp.StatusCode)
    }

    fmt.Println("Ticket created successfully")
    return nil
}

func main() {

    ticket := Ticket{
    	Title: "Test 2",
    	Description: "This is first API testing from Golang Codebase",
    	Labels: []string {"bugs"},
    }

	// API - /ticket GET
	// Creates the Issue
	http.HandleFunc("/ticket", func (w http.ResponseWriter, r *http.Request) {
		err := createTicket(RepoOwner, RepoName, Token, ticket)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			fmt.Fprintf(w, `{"result": "Failed", "Error": err}`)
			return
		}

		fmt.Fprintf(w, `{"result": "Success"}`)
	})

	fmt.Println("Server listening on port 8080")
    http.ListenAndServe(":8080", nil)
}
