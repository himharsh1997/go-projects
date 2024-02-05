package main

import (
	"fmt"
	"os"
)

func main(){
    var jsonString string = os.Args[1];

	var parsedJSON map[string]interface{};
 
	var lenOfJSONString = len(jsonString);
	for  i:= 0; i < lenOfJSONString; i++{
      if(string(jsonString[0]) == "{" || jsonString[lenOfJSONString - 1] == '}'){
         
	  }
	}


   fmt.Println(parsedJSON);
}