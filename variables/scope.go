package main
import "fmt"
func main() {
    {
        v := 1
        {
            fmt.Println(v)
        }
        fmt.Println(v)
    }
    // “undefined: v” compilation error
    // fmt.Println(v)
}