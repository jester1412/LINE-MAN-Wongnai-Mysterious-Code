package main

import (
	"fmt"
	"encoding/base64"
)

func main() {
	
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"


	// Decode secret msg
	sd, _ := base64.StdEncoding.DecodeString(secret)
	
	// Print sd []byte as string
	fmt.Println(string(sd))
	

	// Reverse sd string
	// Get range of sd then loop through every elements 
	// Set whatIsIt with reverse elements of sd
	for i,_ := range sd {
		sdLen := len(sd)
		whatIsIt = whatIsIt + string(sd[sdLen - i -1])
	}
	
	fmt.Println(whatIsIt)
}
