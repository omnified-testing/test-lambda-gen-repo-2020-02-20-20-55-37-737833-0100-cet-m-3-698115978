package main

	import (
		utility "github.com/omnified-testing/test-lambda-gen-repo-2020-02-20-20-55-37-737833-0100-cet-m-3-698115978/utility"
	)
	
	func main() {
	}
	
	// Call : The executable API of the function, just calls the subpackage.
	// Input and output needs to match that of utility subpackage's Call function
	func Call(s string) string {
		return utility.Call(s)
	}