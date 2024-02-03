package main

import (
  "fmt"
  "net/http"
  "encoding/json"
)


func handler(w http.ResponseWriter, r *http.Request){

  response := map[string]interface{}{
	"status": "success",
  };

  jsonData, err := json.Marshal(response);
  fmt.Print(err);

  w.Header().Set("Content-Type", "application/json");
  w.Header().Set("H-Key", "xyz");
  w.Write(jsonData);

}

func main(){
  http.HandleFunc("/", handler)

  fmt.Println("Server listing on port 8080")
  err := http.ListenAndServe(":8080", nil)
  if(err != nil){
   fmt.Println("Error:", err)
  }
}
