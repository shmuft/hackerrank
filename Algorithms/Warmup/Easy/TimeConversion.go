package main
import "fmt"

func main() {
    var h, m, s int
    var st string
    fmt.Scanf("%2d:%2d:%2d%2s", &h, &m, &s, &st)
    if st=="PM" {
        if (h!=12){
          h+=12
        }
    }
    if st=="AM" {
        if (h==12){
            h = 0
        }
    }

    fmt.Printf("%02d:%02d:%02d", h, m, s)    
}
