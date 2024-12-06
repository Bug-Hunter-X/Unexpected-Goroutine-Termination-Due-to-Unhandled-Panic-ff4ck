func main() {


        fmt.Println("Before the panic")
        defer func() {
                if r := recover(); r != nil {
                        fmt.Println("Recovered from panic:", r)
                }
        }()
        panic("Something went wrong")
        fmt.Println("After the panic")
}
