package main  // Declaring the main package, which is the entry point of the Go program
import (
    "fmt"  // Importing the "fmt" package for formatted I/O operations
    "os"   // Importing the "os" package to access command-line arguments
)

type Customer struct {     // Defining a struct named "Customer" to hold customer information
    Name         string
    Address      string
    Services     string
    Monthly_Rate float32
}
func (c Customer) Print() {  // Method to print the details of a Customer instance
    fmt.Printf("Name: %s\nAddress: %s\nServices: %s\nMonthly Rate: $%.2f\n\n",   // Formatting the output to display customer details
        c.Name, c.Address, c.Services, c.Monthly_Rate)   // Using fmt.Printf to format the output with the customer's name, address, services, and monthly rate
}
func main() {    // The main function is the entry point of the program
    args := os.Args  // Accessing command-line arguments passed to the program
    if len(args) < 2 {   // Checking if the number of arguments is less than 2 (the first argument is the program name)
    fmt.Println("Usage: go run main.go [list | add]") // Printing usage instructions if no valid command is provided
    return
}

    customers := []Customer{   // Creating a slice of Customer structs to hold multiple customer records
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
    for _, c := range customers {   // Iterating over the slice of customers and calling the Print method for each customer to display their details
        c.Print()
    }
}

