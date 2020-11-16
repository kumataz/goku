package main

import (
  "flag"
  "fmt"
)



type Config struct{
    name string
    age string
    bornDate string
    interest string
}


// Define INFO in here or Next Step
var cfg Config = Config{
    name: "kawavi",
    age: "20",
    bornDate: "20000101",
    interest: "game",
}



func init(){
    flag.StringVar(&cfg.name,     "n", "kawayi", "Name input.")
    flag.StringVar(&cfg.age,      "a", "20", "Age input.")
    flag.StringVar(&cfg.bornDate, "b", "20200101", "bornDate input.")
    flag.StringVar(&cfg.interest, "i", "game", "interest input.")
}


func main(){
    // TODO parse from command
    flag.Parse()
    fmt.Printf("Name: %s\nAge: %s\nDate: %s\nInterest: %s\n", cfg.name, cfg.age, cfg.bornDate, cfg.interest)
}
