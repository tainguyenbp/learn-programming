package main

import "fmt"

func main() {
	rawstr :=
		`First line
some new lines
more new lines
"double quotes"
`
	fmt.Print(rawstr)

	rawstr1 :=
		`one line
two line
three line
"fours line"
`

	fmt.Print(rawstr1)

	rawstr2 :=
		`Line with tab	here
Line with newline
and more lines
`

	fmt.Print(rawstr2)

	rawstr3 :=
		`This is a raw string with backticks: 
	It spans multiple lines and includes "quotes" and \backslashes\.
`

	fmt.Print(rawstr3)

}
