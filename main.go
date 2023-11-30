package main

import (
	

	"github.com/Kesha005/bc/cmd"
	//import fmt and os package
)

// func main() {
//    directory, err := os.Getwd()    //get the current directory using the built-in function
//    if err != nil {
//       fmt.Println(err) //print the error if obtained
//    }
//    fmt.Println("Current working directory:", directory) //print the required directory
// }


func main(){
   cmd.Execute()
}