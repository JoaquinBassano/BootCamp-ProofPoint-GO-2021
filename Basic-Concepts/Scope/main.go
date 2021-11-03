package main

import (
	"myapp/packageone"
)

var myVar = "I am myVar"

func main() {
	blockVar := "I am blockVar"

	packageone.PrintMe(myVar)
	packageone.PrintMe(blockVar)
	packageone.PrintMe(packageone.PackageVar)
}
