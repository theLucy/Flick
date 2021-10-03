package main

import "github.com/mitchellh/go-finger"
import (
  "fmt"
  "bufio"
  "os"
  "io/ioutil"
  "io"
  "context"
)

func main() {

  profile_folder := "profiles/"

  go finger.Serve(finger.HandlerFunc(func(ctx context.Context, w io.Writer, q *finger.Query) {
    text := ""
    filename := ""

    if q.Username == "" {
      filename = "root"
    } else {
      filename = profile_folder + q.Username
    }

    file, err := os.Open(filename)
    if err != nil {
      text = "ERR!\nNo such user!"
    }
    defer func() {
      if err = file.Close(); err != nil {
      }
    }()

    if err == nil {
      b, err := ioutil.ReadAll(file)
      if err != nil {
        text = "ERR!\nNo such user!"
      } else {
        text = "OK.\n" + string(b)
      }
    }

    w.Write([]byte(fmt.Sprintln(text)))
    fmt.Printf("Requested username: %s\n", q.Username)
  }))
  fmt.Print("Press 'Enter' to exit...\n")
  bufio.NewReader(os.Stdin).ReadBytes('\n') 
}
