::go
::header
package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "math"
    "fmt"
)

::code
// Insert your SimpsonsIndex() function here, along with any subroutines that you need.
func SimpsonsIndex(sample map[string]int) float64 {

}

::footer

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

func main() {
    //input
    scanner := bufio.NewScanner(os.Stdin)

    //make the map storing the sample
    sample := make(map[string]int)

    for scanner.Scan() {
        currentLine := scanner.Text()
        arrayString := strings.TrimSpace(currentLine) // converts to string
        words := strings.Fields(arrayString) // converts string to array of strings

        //now we need to convert each element of the map
        if len(words) != 2 {
            panic("Error: some line of sample does not have appropriate length.")
        }

        key := words[0]
        val, err := strconv.Atoi(words[1])

        if err != nil {
            panic("Issue with converting element of array.")
        }

        //ready to set current element of sample
        sample[key] = val
    }

    if err := scanner.Err(); err != nil {
        panic("Error reading from stdin")
    }

    //function call and output
    s := SimpsonsIndex(sample)
    fmt.Println(roundFloat(s,4))
}


    
