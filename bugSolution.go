func main() {
        fmt.Println("Before the panic")
        defer func() {
                if r := recover(); r != nil {
                        fmt.Println("Recovered from panic:", r)
                        // Handle the panic appropriately here.
                        // This might involve logging the error, retrying the operation,
                        // or taking other corrective actions.
                }
        }()
        panic("Something went wrong")
        fmt.Println("After the panic")
}
