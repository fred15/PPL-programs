package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func contains(s []int) bool {
    for _, a := range s {
      if a == 7 || a == 8 { 
        return true
      }
    }
    return false
}

func getRegion(){
  fmt.Scanf("%d", &inputRegion)
  if 12 < inputRegion && inputRegion < 1 {
    fmt.Printf("no such Geographical area , try again\n"+
    "\nGeographical area: ")
    getRegion()
  } 
}

func getfilename() {
  fmt.Printf("input valid file name: ")
  fmt.Scanf("%s", &fileName)
}
func getResponse() bool{
  var response string
  fmt.Scanf("%s", &response)
  if response == "yes" {
    return true
  }else if response == "no"{
    return false
  }else{
    fmt.Printf("wrong input, try again: ")
    return getResponse()
  }
}
func getYear(){
  fmt.Printf("\nyear from %d-%d: ", minYear, maxYear)
  fmt.Scanf("%d", &inputyear)

  if inputyear < minYear || inputyear > maxYear{
    fmt.Printf("wrong input, try again: ")
    getYear()
  }
}

func getInput() { 
  var stringIn string
  fmt.Printf("input can be 0-8 individually, All"+
              "(will do 1-6 for given year) or\n"+
              "lists such as 1,2,4 or 6,3,1 for given year too,"+
              "\nif 7 and 8 are chosen 1 Geographical area will be kept for both\n\n"+
              "input: ")
  fmt.Scanf("%s", &stringIn)
  if stringIn == "All" || stringIn == "all"{
    inputs = []int{1,2,3,4,5,6}
  } else {
    sArray := strings.Split( stringIn, ",")
    for  j := 0; j < len(sArray); j++{
      input, err := strconv.ParseInt( sArray[j], 10 , 64)
      if err != nil || input > 8 || input < 0{
        fmt.Printf("wrong input, try again\n")
        getInput()
        return
      }  
      intIn := int(input)
      inputs = append(inputs, intIn)
    }
  }
}

func getMaxIndex(slice []int) int { //gets max int index

  var index int

  for i:=0; i< len(slice); i++{
    if( i == 0){
      index = 0
    }
    if slice[i] > slice[index] {
      index = i
    }
  }

  return index
}

func getMinIndex(slice []int) int { // gets min int index

  var index int

  for i:=0; i< len(slice); i++{
    if( i == 0){
      index = 0
    }
    if slice[i] < slice[index] {
      index = i
    }
  }

  return index
}

func getMinMaxYears(slice []int){
  maxIndex := getMaxIndex(slice)
  minIndex := getMinIndex(slice)
  maxYear = slice[maxIndex]
  minYear = slice[minIndex]
  run++
  return
}

//variables 

var run int = 0
var maxYear int
var minYear int
var fileName string
var inputs []int
var inputyear int
var inputRegion int

func main() {


  ints := make([]int, 0)
  countries := make([]string, 0)
  var i int
  var end bool
  start := false 

  if run ==2 {

    fmt.Printf("\n* Input an interger to choose the funcion you would like to run *"+
    "\n\n  1. Top 5 exporting countries for a given year \n"+
    "  2. Worst 5 exporting countries for a given year \n"+
    "  3. 5 countries with the best balance of trade for a given year. \n"+
    "  4. 5 countries with the worst balance of trade for a given year. \n"+
    "  5. 5 countries with the best ratio of exports to balance of trade for a given year. \n"+
    "  6. 5 countries with the worst ratio of exports to balance of trade for a given year. \n"+
    "  7. 5 countries with the worst ratio of exports to balance of trade for a given year & region. \n"+
    "  8. 5 countries with the worst ratio of exports to balance of trade for a given year & region. \n"+
    "  0. Exit.\n\n")
    getInput()
    getYear()
    run += len(inputs)
    i = inputs[0]
    if contains(inputs){
      fmt.Printf("\ninput integer for Geograhical area index (1-12)\n"+
        "1 US + Canada\n"+
        "2 Mexico + Central Am. + Carabien\n"+
        "3 South Am\n"+
        "4 Australia, New Zealand\n"+
        "5 Asia\n"+
        "6 Africa\n"+
        "7 Mid-east\n"+
        "8 Western Europe\n"+
        "9 Former Soviet Block\n"+
        "10 other or unknown\n"+
        "11 European Union\n"+
        "12 Pacific Islands\n"+
        "\nGeographical area: ")
      getRegion()
    }
  }else if run > 2{
    i = inputs[0]
  }else{
    getfilename()
  }
  file, err := os.Open(fileName)
  for err != nil {
    fmt.Printf("\n*no such file, try again*\n\n")
    getfilename()
    file, err = os.Open(fileName)
  }

  scanner := bufio.NewScanner(file)


  if run == 0{
    for scanner.Scan(){
      str:= scanner.Text()
      if start == true{
        sArray := strings.Split( str, ",")
        year, _ := strconv.ParseInt( sArray[3], 10 , 64)
        intYear := int(year)
        ints = append(ints, intYear)
      }
      if str == "%Data"{
        start = true
      }
    }
    getMinMaxYears(ints)
  }

  for{

    if i == 1 || i == 2 {

      for scanner.Scan(){
        str:= scanner.Text()
        if start == true {
          sArray := strings.Split( str, ",")

          year, _ := strconv.ParseInt( sArray[3], 10 , 64)
          intYear := int(year)

          if inputyear == intYear {

            num, _ := strconv.ParseInt( sArray[1], 10 , 64)
            numInt := int(num)

            ints = append(ints, numInt)
            countries = append(countries, sArray[0])

          }
        }  
        if str == "%Data"{
          start = true
        }
      }// scanning putting in countries and ints into arrays
      fmt.Printf("\noutput ")
      if i == 1{
        fmt.Printf("for 1:\n\n")
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMaxIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMaxIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }  
      }else if i == 2{
        fmt.Printf("for 2:\n\n")
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMinIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMinIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }      
      }
      break
    }else if i == 3 || i == 4 {


      for scanner.Scan(){
        str:= scanner.Text()
        if start == true {
          sArray := strings.Split( str, ",")
          year, _ := strconv.ParseInt( sArray[3], 10 , 64)
          intYear := int(year)

          if inputyear == intYear {

            num, _ := strconv.ParseInt( sArray[2], 10 , 64)
            numInt := int(num)

            ints = append(ints, numInt)
            countries = append(countries, sArray[0])

          }
        }  
        if str == "%Data"{
          start = true
        }
      }// scanning putting in countries and ints into arrays

      fmt.Printf("\noutput ")

      if i == 3{
        fmt.Printf("for 3:\n\n")
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMaxIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMaxIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }  
      }else if i == 4{
        fmt.Printf("for 4:\n\n")
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMinIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMinIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }
      }    
      break
    }else if i == 5 || i == 6 {

      for scanner.Scan(){
        str:= scanner.Text()
        if start == true {
          sArray := strings.Split( str, ",")
          year, _ := strconv.ParseInt( sArray[3], 10 , 64)
          intYear := int(year)

          if inputyear == intYear {

            export, _ := strconv.ParseInt( sArray[1], 10 , 64)
            trade, _ := strconv.ParseInt( sArray[2], 10 , 64)
            ratio := export/trade
            numInt := int(ratio)

            ints = append(ints, numInt)
            countries = append(countries, sArray[0])

          }
        }  
        if str == "%Data"{
          start = true
        }
      }// scanning putting in countries and ints into arrays

      fmt.Printf("\noutput: ")

      if i == 5{
        fmt.Printf("for 5:\n\n")
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMaxIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMaxIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }  
      }else if i == 6{
        fmt.Printf("for 6:\n\n")
        if len(countries) >= 0 {
          if len(countries) >= 5{
            for j := 0; j < 5; j++{
              k := getMinIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }else{
            length := len(countries)
            for j := 0; j < length; j++{
              k := getMinIndex(ints)
              fmt.Printf("  %s   %d\n", countries[k], ints[k])
              ints  = append(ints[:k], ints[k+1:]...)
              countries  = append(countries[:k], countries[k+1:]...)
            }
          }
        }      
      }
      break
    }else if i == 7 || i == 8 {

      for scanner.Scan(){
        str:= scanner.Text()
        if start == true {
          sArray := strings.Split( str, ",")
          year, _ := strconv.ParseInt( sArray[3], 10 , 64)
          intYear := int(year)
          region, _ := strconv.ParseInt( sArray[5], 10 , 64)
          intRegion := int(region)

          if inputyear == intYear && inputRegion == intRegion {

            export, _ := strconv.ParseInt( sArray[1], 10 , 64)
            trade, _ := strconv.ParseInt( sArray[2], 10 , 64)
            ratio := export/trade
            numInt := int(ratio)

            ints = append(ints, numInt)
            countries = append(countries, sArray[0])

          }
        }  
        if str == "%Data"{
          start = true
        }
      }// scanning putting in countries and ints into arrays
      
      fmt.Printf("\noutput ")

        if i == 7{
        fmt.Printf("for 7:\n\n")
          if len(countries) >= 0 {
            if len(countries) >= 5{
              for j := 0; j < 5; j++{
                k := getMaxIndex(ints)
                fmt.Printf("%s   %d\n", countries[k], ints[k])
                ints  = append(ints[:k], ints[k+1:]...)
                countries  = append(countries[:k], countries[k+1:]...)
              }
            }else{
              length := len(countries)
              for j := 0; j < length; j++{
                k := getMaxIndex(ints)
                fmt.Printf("  %s   %d\n", countries[k], ints[k])
                ints  = append(ints[:k], ints[k+1:]...)
                countries  = append(countries[:k], countries[k+1:]...)
              }
            }
          }    
        }else if i == 8{
        fmt.Printf("for 8:\n\n")
          if len(countries) >= 0 {
            if len(countries) >= 5{
              for j := 0; j < 5; j++{
                k := getMinIndex(ints)
                fmt.Printf("  %s   %d\n", countries[k], ints[k])
                ints  = append(ints[:k], ints[k+1:]...)
                countries  = append(countries[:k], countries[k+1:]...)
              }
            }else{
              length := len(countries)
              for j := 0; j < length; j++{
                k := getMinIndex(ints)
                fmt.Printf("  %s   %d\n", countries[k], ints[k])
                ints  = append(ints[:k], ints[k+1:]...)
                countries  = append(countries[:k], countries[k+1:]...)
              }
            }
          }    
        }
        break

    }else if i == 0{
      break
    }else{
      fmt.Printf("wrong input, try again: ")
    }
    fmt.Scanf("%d", &i)
  }
  

  if run == 1{
    end = true
    run++
  }else if run >= 2{
    if len(inputs) > 1{
      inputs  = append(inputs[:0], inputs[1:]...)
      end = true
    }else{
      inputs = make([]int, 0)
      fmt.Printf("\nWant to run another process(yes or no)? ")
      end = getResponse()
      run = 2
      inputyear = 0
    }
  }
  if end {
    main()
  }
}