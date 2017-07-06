package main

import "fmt"

/*https://www.hackerrank.com/challenges/lisa-workbook*/
func main() {
	var numberChapters, problemsPerPage, numberProblems int
	fmt.Scanf("%d", &numberChapters)
	fmt.Scanf("%d", &problemsPerPage)
	var amountSpecialPages, currentPage int = 0, 0
	for i := 0; i < numberChapters; i++ {
		fmt.Scanf("%v", &numberProblems)
		countProblems := 0
		for numberProblems >= problemsPerPage {
			currentPage++
			countProblems += problemsPerPage
			if countProblems-problemsPerPage < currentPage && currentPage <= countProblems {
				amountSpecialPages++
			}
			numberProblems -= problemsPerPage
		}
		if numberProblems != 0 {
			currentPage++
			countProblems += numberProblems
			if countProblems-numberProblems < currentPage && currentPage <= countProblems {
				amountSpecialPages++
			}
		}
	}
	fmt.Println(amountSpecialPages)

}
