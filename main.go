package main

import (
  "fmt"

  "github.com/ua-parser/uap-go/uaparser"
)

func main() {
  uagent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"

  parser := uaparser.NewFromSaved()
  
  

  client := parser.Parse(uagent)

  fmt.Println(client.UserAgent.Family)  // "Amazon Silk"
  fmt.Println(client.UserAgent.Major)   // "1"
  fmt.Println(client.UserAgent.Minor)   // "1"
  fmt.Println(client.UserAgent.Patch)   // "0-80"
  fmt.Println(client.Os.Family)         // "Android"
  fmt.Println(client.Os.Major)          // ""
  fmt.Println(client.Os.Minor)          // ""
  fmt.Println(client.Os.Patch)          // ""
  fmt.Println(client.Os.PatchMinor)     // ""
  fmt.Println(client.Device.Family)     // "Kindle Fire"
}
