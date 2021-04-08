package main

import (
	"bankzhb/bankcore"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/prometheus/common/log"
)

var accounts = map[float64]*bankcore.Account{}

func main() {
	accounts[1001] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeless, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}

	accounts[1002] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "Mark",
			Address: " Haha",
			Phone:   "131 111 000",
		},
		Number: 1002,
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintln(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	numbers := strings.Split(numberqs, ",")
	if len(numbers) == 0 {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		a1, _ := strconv.ParseFloat(numbers[0], 64)
		a2, _ := strconv.ParseFloat(numbers[1], 64)
		account1, ok1 := accounts[a1]
		account2, ok2 := accounts[a2]
		if !ok1 && !ok2 && account1.Balance < amount {
			fmt.Fprintf(w, "Account with number %v can't be found!", numbers[0])
		} else {
			err := account1.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintln(w, account1.Statement())
			}
			err = account2.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account2.Statement())
			}
		}
	}
}
