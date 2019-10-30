# Notes on learning Go

Go doc: `go doc {package}.{function}`     
    
</br>

======== (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ ========    
### Section I: Introduction
    
To compile, `go build {programName}.go`   
To run, `./{execName}`    
A command that combines compile and run, `go run {programName}.go` (this does not create an executable in the folder)    
    
      
first program: `HelloWorld.go`    
    
To import multiple packages:    
```
import "package1"
import "package2"
```
or 
```
import (
  "package1"
  "package2"
)
```
You can also create alias for a package, like this:   
```
import (
  p1 "package1"
  "package2"
)
```
and this will allow you to do something like `p1.func1()` instead of `package1.func1()`
    
Some fmt functions: 
* Println()    
`fmt.Println("Hello", "hello", "my dear")` will print `Hello hello my dear`, inserting a space with each argument    
`fmt.Println("ily")` will print on the next line    

* Print()    
`fmt.Print("Hello", "hello", "my dear")` will print `Hellohellomy dear`, not adding anything
`fmt.Print("ily")` will print directly after `dear` and will not start on a new line

* Printf()     
`animal1 := cat animal2 := dog  fmt.Printf("Are you a %v or a %v person?", animal1, animal2)` prints "Are you a cat or a dog person?"    
`num := 42  fmt.Printf("This value's type is %T.", num) ` prints "This value's type is int"    
`votingAge := 21 fmt.Printf("You must be %d years old to vote.", votingAge)` prints "You must be 21 years old to vote."    
`gpa := 3.9 fmt.Printf("You're averaging: %.2f.", gpa)` Prints "You're averaging 3.90."    

* Sprint()    
stored the string into a var    
```
grade := "100"
compliment := "Good job!"
teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)
fmt.Print(teacherSays)
```
* Sprintln()    
same as Sprint(), but adds a trailing space
```
quote = fmt.Sprintln("Look,", "space!")
fmt.Print(quote) 
```
* Sprintf()    
```
template := "I wish I had a %v."
pet := "puppy"
wish := fmt.Sprintf(template,pet)
fmt.Println(wish)
```
* User input .Scan()    
```
var food string
fmt.Scan(&food)
fmt.Printf("Let's have %v for lunch", food)
```
</br>

======== (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ ========    
### Section II: Values and Variables    
constant variable is declared and defined as: `const earthsGravity = 9.80665`    
varibles (use camelCase in Go):
```
var lengthOfSong uint16
var isMusicOver bool
var songRating float32
```
To assign values (2 ways):
```
var numOfFlavors int
numOfFlavors = 57
======== OR =========
var numOfFlavors int = 57
```
**Note**: if you don't assign values, ints & floats will be defaulted to **0**, strings will be defaulted to **""**, bools will be defaulted to **false**    
</br>
String concatenation: `fmt.Println("My favorite snack is "+favoriteSnack+".")`    
</br>
You can also _infer types_:
```
myMessage := "gergerger"
======== OR =========
var myMessage = "gergerger"
```
Multiple variables declaration:
```
amount, unit := 10, "item"
======== OR =========
var str1, str2 string
str1 = "hello"
str2 = "world"
```
</br>

======== (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ ========  
### Section III: Conditionals    
**Operators:**     
equal: `==`   
unequal: `!=`    
and: `&&`    
or: `||`    
not: `!`   

_code see conditionals.go_     (if else; switch; scoped short declaration)    

**Random:**    
_code see random.go_

</br>

======== (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ (๑˃̵ᴗ˂̵)ﻭ ʕ•ᴥ•ʔ ٩(♡ε♡ )۶ (•ө•)♡ ========  
### Section IV: Functions   
**Return Type:**
```
func isItLateInNewYork() string {  // return type is a string

}
```
**Function Parameters:**
```
func multiplier(x int32, y int32) int32 {  // or func multiplier(x, y int32) int32
  return x * y
}
```
**Multiple Return Values:**
```
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {  // returns a string and a float32

}
```
_code see functions.go_     
</br>

**Deferring Resolution:**
```
func calculateTaxes(revenue, deductions, credits float64) float64 {
  defer fmt.Println("Taxes Calculated!") // defer statement here

  taxRate := .06143
  fmt.Println("Calculating Taxes")

  if deductions == 0 || credits == 0 {
    return revenue * taxRate
  }

  taxValue := (revenue - (deductions * credits)) * taxRate
  if taxValue >= 0 {
    return taxValue
  } else {
    return 0
  }
}
```
when we call calculateTaxes(), we immediately defer the message "Taxes Calculated!", which does not print until the end of the function (after the taxes have been calculated and are about to be returned).    
Normally, we would just add fmt.Println("Taxes Calculated!") at _the end_ of calculateTaxes(); But here, we have multiple return statements in our code, so instead of adding a print statement right before each return, we use defer and it prints regardless of when our function ends.

_code see defer.go_     
</br>