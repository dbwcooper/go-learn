package main

import "strings"

import "strconv"


func main() {

	str := "abcad"
	numic := "100"
	print("total a: ", strings.Count(str, "a"))
	print("\n ToLower:", strings.ToLower(str))

	print(" numic: ", numic)


	var b , _ = strconv.Atoi(numic)
	// var isEqual = numic == b
	print(b)
	// print(isEqual)

}