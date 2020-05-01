package main

import (

  "fmt"
  "os"
  "bufio"

)

type maintenance struct {
     
  Name string
  Miles int

}

func Maintain(mileage int){

   // Create an array of 8 maintenance objects, each with a Name and a Miles Trait specific to the users vehicle model by reading 
   // in from a file that contains the users vehicles information. 

   //create array of objects
  var  maint [8] maintenance

  //read from file
  fileName, err := os.Open("ZZR600.txt")
  if err != nil {
        panic(err)
    }
  scannerName := bufio.NewScanner(fileName)
  count := 0
  i := 0 //location in array
  //read each line at a time
  for scannerName.Scan(){
    count = count+1
    if count %2 == 1{
          fmt.Println(scannerName.Text())
      maint [i].Name = scannerName.Text()
    } else{
          fmt.Println(scannerName.Text())
      maint [i].Miles = scannerName.Text()
      i=i+1
    }

    //save to an array
    
    //manipulate
  } 

  fileName.Close()

   //Prints message showing user the their maintenance schedule with how many miles until each one.

   for i:= 0; i <=8; i++{
    x := (mileage % maint[i].Miles)
    until := (maint[i].Miles - x)
   
    if until < 100 {
    fmt.Println()
    fmt.Println("[","You have",until,"Miles until", maint[i].Name,"Don't forget!","]")
    } else {

    fmt.Println()
    
    fmt.Println("[",until, "Miles until:", maint[i].Name,"]")
    } 
  }
}

func main() {

//assign a variable for mileage
  var mileage int

  fmt.Println("To check your upcoming maintenance please enter your current vehicle mileage.")
  
  //scan for mileage
  fmt.Scanln(&mileage)

  //call maintenance
  Maintain(mileage)

}
