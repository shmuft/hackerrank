package main
import "fmt"

func main() {
    var num, i, j int
    fmt.Scanf("%v", &num)
    for i=0; i<num; i++{
        for j=0;j<num-i-1;j++{
            fmt.Print(" ")
        }
        for j=0;j<=i;j++{
            fmt.Print("#")
        }
        fmt.Print("\n")
    }
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}
