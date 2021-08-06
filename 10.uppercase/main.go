
package main 
  
import ( 
    "fmt"
    "strings"
) 
  

func main() { 
  
   
    s1 := "anonna ferdaus"
    s2 := "tareq ibna rahman"
    s3 := " GOLANG"
    s4 := "uppercase conversion"
  
  
    fmt.Println("Strings (before):") 
    fmt.Println("String 1: ", s1) 
    fmt.Println("String 2:", s2) 
    fmt.Println("String 3:", s3) 
    fmt.Println("String 4:", s4) 
  
    // Converting all the string into uppercase 
    // Using ToUpper() function 
    r1 := strings.ToUpper(s1) 
    r2 := strings.ToUpper(s2) 
    r3 := strings.ToUpper(s3) 
    r4 := strings.ToUpper(s4) 
  
    // Displaying the results 
    fmt.Println("\nStrings (after):") 
    fmt.Println("Result 1: ", r1) 
    fmt.Println("Result 2:", r2) 
    fmt.Println("Result 3:", r3) 
    fmt.Println("Result 4:", r4) 
} 