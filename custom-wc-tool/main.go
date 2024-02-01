package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
     var fileName string;
	 var option string;

	if(len(os.Args) == 2){
		fileName = os.Args[1];
		option = "";
	} else {
		fileName = os.Args[2];
		option = os.Args[1];
	}
    
	content, err := ioutil.ReadFile(fileName);
    if(err != nil){
       fmt.Println("Error while reading file: ", err);
	   return;
	}
	
	if(option == "-c"){
     fmt.Println(len(content));
	} else if(option == "-l"){
	 fmt.Println(len(strings.Split(string(content), "\n")));
	} else if(option == "-w"){

	} else if(option == "-m"){

	} else if(option == "-L"){
	 var lines = strings.Split(string(content), "\n")
     max:= 0
	 for i:=0; i< len(lines); i++{
      if(len(lines[i]) > max){
       max = len(lines[i])
	  }
	}
	fmt.Println(max);
	}


		// Print default <lines> <words> <characters> <filename>

}