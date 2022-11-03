package main

import "fmt"

func main() {
  var publisher, writer, artist, title string
  var year, pageNumber uint
  var grade float32

// Comic Book Collection - Mr. GoToSleep
  title = "Mr. GoToSleep"; writer = "Tracey Hatchet"; artist = "Jewel Tampson"; publisher = "DizzyBooks Publishing Inc."; year = 1997; pageNumber = 14; grade = 6.5

//Print Statement "Mr. GoToSleep"
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "printed on", year, "with a total number of pages", pageNumber, "at grade", grade)

// Comic Book Collection - Epic Vol. 1
  title = "Epic Vol. 1"; writer = "Ryan N. Shawn"; artist = "Phoebe PaperClips"; publisher = "DizzyBooks Publishing Inc."; year = 2013; pageNumber = 160; grade = 9.0

//Print Statement "Epic Vol. 1"
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "printed on", year, "with a total number of pages", pageNumber, "at grade", grade)

}