package main

import (
  "path/filepath"
  "fmt"
  "os"

  "github.com/neut0ne/todo/db"
  "github.com/neut0ne/todo/cmd"
  homedir "github.com/mitchellh/go-homedir"
  )

func main (){
  home, _ := homedir.Dir()
  dbPath := filepath.Join(home, "tasks.db")
  must(db.Init(dbPath))
  must(cmd.RootCmd.Execute())
}

func must(err error) {
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }
}
