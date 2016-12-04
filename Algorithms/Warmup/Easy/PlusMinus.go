package main
import "fmt"

func main() {
    var rea, num, positive, negative, zeros int
    positive = 0
    negative = 0
    zeros = 0
    fmt.Scanf("%v", &num)
    for i:=0; i<num; i++{
        fmt.Scanf("%v", &rea);
        switch {
            case rea > 0:
                positive++
            case rea < 0:
                negative++
        }
    }
    zeros = num - positive - negative
    
    fmt.Printf("%.6f\n", float64(positive)/float64(num))
    fmt.Printf("%.6f\n", float64(negative)/float64(num))
    fmt.Printf("%.6f\n", float64(zeros)/float64(num))
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}
