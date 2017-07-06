package main

import "fmt"

func main() {
	var h, m int
	fmt.Scanf("%d", &h)
	fmt.Scanf("%d", &m)

	dictionary := make(map[int]string)
	dictionary[0] = "zero"
	dictionary[1] = "one"
	dictionary[2] = "two"
	dictionary[3] = "three"
	dictionary[4] = "four"
	dictionary[5] = "five"
	dictionary[6] = "six"
	dictionary[7] = "seven"
	dictionary[8] = "eight"
	dictionary[9] = "nine"
	dictionary[10] = "ten"
	dictionary[11] = "eleven"
	dictionary[12] = "twelve"
	dictionary[13] = "thirteen"
	dictionary[14] = "fourteen"
	dictionary[15] = "fifteen"
	dictionary[16] = "sixteen"
	dictionary[17] = "seventeen"
	dictionary[18] = "eighteen"
	dictionary[19] = "nineteen"
	dictionary[20] = "twenty"
	dictionary[30] = "thirty"
	dictionary[40] = "forty"
	dictionary[50] = "fifty"
	dictionary[60] = "sixty"
	s := ""
	switch {
	case m == 0:
		s = dictionary[h] + " o' clock"
	case m == 1:
		s = "one minute past " + dictionary[h]
	case m == 10:
		s = dictionary[m] + " minutes past " + dictionary[h]
	case m == 15:
		s = "quarter past " + dictionary[h]
	case m == 30:
		s = "half past " + dictionary[h]
	case m == 40:
		s = dictionary[60-m] + " minutes to " + dictionary[h+1]
	case m == 45:
		s = "quarter to " + dictionary[h+1]
	case m > 30:
		if _, ok := dictionary[60-m]; !ok {
			s2 := dictionary[(60-m)%10]
			s1 := dictionary[(60-m)-(60-m)%10]
			s = s1 + " " + s2 + " minutes to " + dictionary[h+1]
		} else {
			s = dictionary[60-m] + " minutes to " + dictionary[h+1]
		}
	case m < 30:
		if _, ok := dictionary[60-m]; !ok {
			s2 := dictionary[m%10]
			s1 := dictionary[m-m%10]
			s = s1 + " " + s2 + " minutes past " + dictionary[h]
		} else {
			s = dictionary[60-m] + " minutes past " + dictionary[h]
		}
	}
	fmt.Println(s)
}
