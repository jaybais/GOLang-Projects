package main  // Declaring the main package, which is the entry point of the Go program
import (
    "fmt"  // Importing the "fmt" package for formatted I/O operations
    "os"
    "strconv"   // Importing the "strconv" package to convert strings to numbers
)

type Customer struct {     // Defining a struct named "Customer" to hold customer information
    Name         string    // Field to store the customer's name
    Address      string
    Services     string
    Monthly_Rate float32
}
func (self Customer) Print() {  // Method to print the details of a Customer instance
    fmt.Printf("Name: %s\nAddress: %s\nServices: %s\nMonthly Rate: $%.2f\n\n",   // Formatting the output to display customer details
        self.Name, self.Address, self.Services, self.Monthly_Rate)   // Using fmt.Printf to format the output with the customer's name, address, services, and monthly rate
}
func main() {    // The main function is the entry point of the program
    args := os.Args  // Accessing command-line arguments passed to the program
    if len(args) < 2 {   // Checking if the number of arguments is less than 2 (the first argument is the program name)
    fmt.Println("Usage: go run main.go [list | add]") // Printing usage instructions if no valid command is provided
    return
}

fmt.Println()

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


    command := args[1]  // Storing the second command-line argument (the command) in a variable for later use

    switch command {  // Using a switch statement to handle different commands based on the user input

    case "list": // If the command is "list", iterate through the customers slice and print each customer's details
        for _, c := range customers {  
            c.Print() 
        }

    case "add":  // If the command is "add", check if there are enough arguments to add a new customer
        if len(args) < 6 {
            fmt.Println(`Usage: add "Name" "Address" "Services" Rate`)
            return
        }

        name := args[2]   // Storing the third command-line argument as the customer's name
        address := args[3]
        services := args[4]
        rateStr := args[5]  // Storing the sixth command-line argument as a string to be converted to a float

        rateFloat, err := strconv.ParseFloat(rateStr, 32) // Converting the rate string to a float32 value, and checking for errors during conversion
        if err != nil {
            fmt.Println("Invalid rate.")
            return
        }

        newCustomer := Customer{  // Creating a new Customer struct with the provided details
            Name:         name,
            Address:      address,
            Services:     services,
            Monthly_Rate: float32(rateFloat),
        }

        customers = append(customers, newCustomer) // Adding the new customer to the customers slice

        fmt.Println("Customer added!\n") // Printing a confirmation message that the customer has been added

        for _, c := range customers {  // Iterating through the customers slice again to print all customers, including the newly added one
            c.Print()
        }

    default:  // If the command does not match "list" or "add", print an error message indicating an unknown command
        fmt.Println("Unknown command")
    }
}

    

