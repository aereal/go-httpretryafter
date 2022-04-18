package httpretryafter

import "fmt"

func ExampleParse_httpDate() {
	parsedTime, _ := Parse("Fri, 22 Mar 2019 15:04:05 JST")
	fmt.Printf("%s\n", parsedTime)
	// Output:
	// 2019-03-22 06:04:05 +0000 UTC
}
