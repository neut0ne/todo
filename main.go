package main

import (
  "path/filepath"

  "github.com/neut0ne/todo/db"
  "github.com/neut0ne/todo/cmd"
  homedir "github.com/mitchellh/go-homedir"
  )

func main (){
  home, _ := homedir.Dir()
  dbPath := filepath.Join(home, "tasks.db")
  err := db.Init(dbpath)
  if err != nil {
    panic(err)
  }
  cmd.RootCmd.Execute()
}
