package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func main() {
  args := os.Args
  inputfile := args[1]
  dayPart, err := strconv.Atoi(args[2])
  file, err := os.Open(inputfile)
  if err != nil {
    log.Fatal(err)
  }
  //close the file at the end of this function
  defer file.Close()

  // reading the input file
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  input_text := scanner.Text()

  // spliting the input into individual digits
  digits := strings.Split(input_text,"")
  num_digits := len(digits)

  // control how which numbers we compare numbers in part1 vs part2
  // use strconv.Atoi to convert string to integer
  step_len := 1
  if dayPart == 1 {
    step_len = 1
  } else dayPart == 2{
    step_len = num_digits / 2
  }
  sum := 0
  for ii:= 0; ii < num_digits; ii++ {
    num1,err := strconv.Atoi(digits[ii])
    num2,err := strconv.Atoi(digits[(ii+step_len) % num_digits])
    if err != nil {
      log.Fatal(err)
    }
    if num1 == num2 {
      sum += num1
    }
  }
  fmt.Printf("Captcha is %d\n", sum)

}
