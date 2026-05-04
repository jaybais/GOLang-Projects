
package main

import (
    "fmt"
    "os"
    "strconv"
)

type Customer struct {
    Name         string
    Address      string
    Services     string
    Monthly_Rate float32
}

func (c Customer) Print() {
    fmt.Printf("Name: %s\nAddress: %s\nServices: %s\nMonthly Rate: $%.2f\n\n",
        c.Name, c.Address, c.Services, c.Monthly_Rate)
}

func main() {
    args := os.Args

    if len(args) < 2 {
        fmt.Println("Usage: go run main.go [list | add]")
        return
    }

fmt.Println()

    customers := []Customer{
        {
            Name: "Thomas O'Connor",
            Address: "425 17th St. Google Dr. Pittsburgh, PA 16789",
            Services: "Ultimate TV, Gig Extra Internet, Xfinity Mobile",
            Monthly_Rate: 245.00,
        },
        {
            Name: "Patrick DeAngelo",
            Address: "25 13th St. Dairy Hill Lane, Small City, IN 20978",
            Services: "Premium TV, Gigabit Internet, Xfinity Mobile",
            Monthly_Rate: 265.00,
        },
        {
            Name: "Kathleen Davis",
            Address: "321 Country Road, West City, VA 32879",
            Services: "1.2 Gig Internet, Xfinity Mobile",
            Monthly_Rate: 205.00,
        },
    }

    command := args[1]

    switch command {

    case "list":
        for _, c := range customers {
            c.Print()
        }

    case "add":
        if len(args) < 6 {
            fmt.Println(`Usage: add "Name" "Address" "Services" Rate`)
            return
        }

        name := args[2]
        address := args[3]
        services := args[4]
        rateStr := args[5]

        rateFloat, err := strconv.ParseFloat(rateStr, 32)
        if err != nil {
            fmt.Println("Invalid rate.")
            return
        }

        newCustomer := Customer{
            Name:         name,
            Address:      address,
            Services:     services,
            Monthly_Rate: float32(rateFloat),
        }

        customers = append(customers, newCustomer)

        fmt.Println("Customer added!\n")

        for _, c := range customers {
            c.Print()
        }

    default: 
        fmt.Println("Unknown command")
    }
}