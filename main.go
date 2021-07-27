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


  go finger.Serve(finger.HandlerFunc(func(ctx context.Context, w io.Writer, q *finger.Query) {
    text := ""
    filename := ""

    if q.Username == "" {
      filename = "root"
    } else {
      filename = q.Username
    }

    file, err := os.Open(filename)
    if err != nil {
      text = "ERR: No such user!"
    }
    defer func() {
      if err = file.Close(); err != nil {
      }
    }()


    if err == nil {
      b, err := ioutil.ReadAll(file)
      if err != nil {
        text = "ERR: No such user!"
      } else {
        text = "MSG: " + string(b)
      }
    }

    w.Write([]byte(fmt.Sprintln(text)))
  }))
  fmt.Print("Press 'Enter' to exit...")
  bufio.NewReader(os.Stdin).ReadBytes('\n') 
}
