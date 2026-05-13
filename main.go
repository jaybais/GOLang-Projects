package main  
import (   
    "fmt"  
    "os" 
    "encoding/json" 
    "strconv"
)


type Customer struct {      
    Name         string    
    Address      string
    Services     string
    Monthly_Rate float32
}
func (self Customer) Print() {     
    fmt.Printf("Name: %s\nAddress: %s\nServices: %s\nMonthly Rate: $%.2f\n\n",   
        self.Name, self.Address, self.Services, self.Monthly_Rate)   
}


func main() {    
    args := os.Args

    if len(args) < 2 {  
        fmt.Println("Usage: go run main.go [list | add]")
        return
    }

    // ✅ 1. LOAD DATA    
    customers, err := loadCustomers()
    if err != nil {
        fmt.Println("Error loading customers:", err)
        return
    }

    command := args[1]    

    switch command {    

    // ✅ 2. READ
    case "list":    
        for _, c := range customers {
            c.Print()
        }

    // ✅ 3. ADD NEW CUSTOMER
    case "add":    

        if len(args) < 6 {    
            fmt.Println("Usage:")
            fmt.Println(`go run main.go add "Name" "Address" "Service" MonthlyRate`)
            return
        }

        rate64, err := strconv.ParseFloat(args[5], 32)    
        if err != nil {
            fmt.Println("Invalid monthly rate")
            return
        }

        newCustomer := Customer{    
            Name: args[2],
            Address: args[3],
            Services: args[4],
            Monthly_Rate: float32(rate64),
        }

        customers = append(customers, newCustomer)    

        // ✅ 4. SAVE
        err = saveCustomers(customers)    
    if err != nil {
        fmt.Println("Error saving customers:", err)
        return
    }

    fmt.Println("Customer added successfully!") 
    }
}


func saveCustomers(customers []Customer) error {    
    data, err := json.MarshalIndent(customers, "", "  ")    
    if err != nil {    
        return err    
    }
    err = os.WriteFile("customers.json", data, 0644)    
    if err != nil {
        return err
    }
    return nil      
}


func loadCustomers() ([]Customer, error) {    

    data, err := os.ReadFile("customers.json")    
    if err != nil {

        if os.IsNotExist(err) {    
            return []Customer{}, nil
        }

        return nil, err
    }

    var customers []Customer    

    err = json.Unmarshal(data, &customers)
    if err != nil {
        return nil, err
    }

    return customers, nil
}

