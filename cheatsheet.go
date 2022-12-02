
// ==== Terminal
// go run file.go
// go build main.go
// go test // Run the test file

// ==== Declare variables (camelcase naming convention)
var newVar string                // Declare without initializing
var newVar string = "something"
newVar2 := "something else"      // Can only be used inside a function. Inferred type "String"
ExportVar := "Capitalize the 1st letter means the var is accessible outside of the package"


// ==== Declare a constant (same naming convention with variable's)
const newConst string = "truth"


// ==== Slices (projection of an underlying array)
cards := []string {"ace of spade", "king of heart"}
cards = append(cards, "Six of Spades") // Append an item to a slice


// ==== For loop
cards := []string {"Queen", "King"}    // A slice of string
for i, card := range cards {           // cards is a slice of type string
	fmt.Println(i, card)
}

// Print from 0 to 4
for i := 0; i < 5; i++ {
	fmt.Println(i)
}


// ==== Declare a FUNCTION
func funcName(arg argDataType) string {  //replace string with the data type the function returns
	return "Amazing"
}

// Call an outside function in main
func outsideFunc(myFunc func(int, int) int) {
	fmt.Println(myFunc(2, 3))
}

func main() {
	add := func(x, y int) int {
		return x + y
	}
	outsideFunc(add)
}



// ==== Receiver
//create a new dataType deck, which is a slice of type string
type deck []string 

// Every variable type deck (a receiver of a deck) now gets access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// (deck, deck) is the data type of the 2 values the function returns
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


// ==== Convert data type
[]byte("Hi there") // Convert string "Hi there" into byte


// ==== Struct (container of data having various types)
type structName struct {
	field1 dataType
	field2 dataType
}


// ==== Pointer & Dereference
TLDR:
	&x => memory location/ pointer/ address of x
	*y => dereference a given address and give back the value
	*{dataType} => pointer to a data type (see third example)

// Pointer - First example
func main() {
	x := 7
	fmt.Println(x, &x)   // &x = pointer of x, a memory location like 0xc00010a000	

	y := &x // storing the pointer in y | y is pointing to x or referencing x
	*y = 8  // dereference the pointer (location) and assign it to a new value
	fmt.Println(x, y) // the value of x is now 8
}

// Pointer - Second example
func changeValue(str *string) {
	*str := "changevalue func"
}

func changeValue2(str string) {
	str := "changeValue2 func"
}

func main() {
	toChange = "original value"
	changeValue(&toChange)
	fmt.Println(toChange) // This now becomes "changeValue func"

	toChange2 = "original value 2"
	changeValue2(toChange2)
	fmt.Println(toChange2) // This will stay as "original value 2"

}

// Third example
func main() {
	// ms is of type pointer-to-myStruct
	// ms must be a pointer (&var) to a myStruct instance. 
	var ms *myStruct{}  
	ms = &myStruct{foo: "test"}   // assign the value to pointer ms
	fmt.Println(ms) // &{test}
}

type myStruct struct {
	foo : string
}

// Fourth example
func main() {
	x := "change"
	var pointer *string = &x
	fmt.Println(x, &x, pointer) // eg: "change" 0xc00010a210 0xc00010a210
	fmt.Println(x, *pointer, pointer) // eg: "change" "change" 0xc00010a210
}


// ==== Operators
if guess >= 1 && guess <= 100 // And
if guess < 1 || guess > 100 // OR
if !true // if not true


// ==== Switch Statement
func main() {
	switch 9 {
	case 1, 5, 10:
		fmt.Println("two, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four, or six")
	default:
		fmt.Println("another number")
	}
}

// Condition can overlap
func main() {
	i := 9
	switch i {
	case i < 10:
		fmt.Println("smaller than 10")
	case i < 20:
		fmt.Println("greater than 20")
	default:
		fmt.Println("equal to 10")
	}
}

// Type switch
func main() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is another type")
	}
}


// ==== Print
var a int = 3

fmt.Printf("The type of a is: %T \n", a) // This will print int


// ==== PACKAGES https://pkg.go.dev/ 
func Exit(code int) // https://pkg.go.dev/os#Exit
os.Exit(1) // Example: exit if error 

func Split(s, sep string) []string // https://pkg.go.dev/strings#Split
strings.Split("Ace, King, Queen", ",") // Split a string by separator ","

// ==== go test code snippet
func Test_SomeTest(t *testing.T) {

	t.Run("Test name", func(t *testing.T) {
		t.Parallel()
		assert.NoErr(t, err)
	})
}
