package main

import (
    "fmt"
    "os"
    "os/user"
    "github.com/Sawansunar56/repl"
)

func main() {
  user, err := user.Current() 

  if err != nil {
    panic(err)
  }

  fmt.Printf("Hello %s! This is the Monkey Programming Language \n", user.Username)
  fmt.Printf("Feel free to type anything")

  repl.Start(os.Stdin, os.Stdout)
}
