package main
import "fmt"
var global *int
func f() {                      
    var x int                      
    x = 1                           
    global = &x                 
}
func g() {
y := new(int)
*y = 1
}
func main(){
	fmt.Println(global)
	g()
	f()
	fmt.Println(*global)
}