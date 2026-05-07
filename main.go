package main  
import (
    "fmt"  
    "os"
    "strconv"   
    "encoding/json" 
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
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    command := args[1]

    switch command {

    // ✅ 2. READ
    case "list":
        for _, t := range tasks {
            fmt.Println(t.Name)
        }

    // ✅ 3. MODIFY
    case "add":
        newTask := Task{
            ID: len(tasks) + 1,
            Name: args[2],
            Done: false,
        }

        tasks = append(tasks, newTask)

        // ✅ 4. SAVE
        err = saveTasks(tasks)
        if err != nil {
            fmt.Println("Error saving:", err)
            return
        }

    default:
        fmt.Println("Unknown command")
    }
}






type Task struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Done bool   `json:"done"`
}





func saveTasks(tasks []Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile("tasks.json", data, 0644)
}


func loadTasks() ([]Task, error) {
    data, err := os.ReadFile("tasks.json")
    if err != nil {
        if os.IsNotExist(err) {
            return []Task{}, nil
        }
        return nil, err
    }

    var tasks []Task
    err = json.Unmarshal(data, &tasks)
    return tasks, err
}


