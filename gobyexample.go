package main

import (
	"fmt"
	"math"
	"time"
)

func demoValues() {
	/*
		Go has various types including strings, integers, floats, booleans, etc.
		Strings can be added/concatenated together with `+`.
	*/

	fmt.Println("go" + "lang")
	
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func demoVariables(){
	/* 
		In Go, variables are explicitely declared and used by the compiler 
		to e.g. check type-correctness of function calls. 
	*/ 

	// `var` declares 1 or more variables.
	var a string = "initial"
	fmt.Println(a)
	
	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	// For example, the zero value for an `int` is 0.
	var e int
	fmt.Println(e)

	// The `:=` syntaxt is shorthand for declaring and initializing a variable.
	// e.g. for var f string = "short" we could use:
	f := "short"
	fmt.Println(f)

}

// `const` declares a constant value
const s string = "constant"
func demoConstants(){
	/*
		Go supports `constants` of character, string, boolean and numeric values.
	*/
	
	fmt.Println(s)

	// A `const` statement can appear anywhere a var statement can.
	const n = 400000

	// Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20/n
	fmt.Println(d)

	// A numeric constant has no type until it's given one, such by an explicit cast.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call.
	// For example, here `math.Sin` expects a `float64`.
	fmt.Println(math.Sin(n))

}

func demoFor(){
	/*
		`for` is Go's only looping construct. Here are three basic
		types of `for` loops.
	*/

	// The most basic type is with a single condition:
	i := 1
	for i <= 3{
		fmt.Println(i)
		i = i + 1
	}

	// A classic intial/condition/after `for` loop:
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly until you
	// `break` out of the loop or `return` from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of the loop.
	for n := 0; n <= 11; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

func demoIfElse(){
	/*
		Branching with `if` and `else` in Go is straigh-forward.
	*/

	// Here is a basic example:
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an `if` statement without an `else`:
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables
	// declared in this statement are available in all branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}

func demoSwitch(){
	/*
		`switch` statements express conditionals across many branches.
	*/

	// Here's a basic `switch`:
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions in 
	// the same `case` statement. We use the optional `default` 
	// case in this example as well.
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend!")
		default:
			fmt.Println("It's a weekday...")
	}

	// `switch` without an expression is an alternate way to
	// express if/else logic. Here we also show how the `case` expressions
	// can be non-constants.
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("It's before noon!")
		default:
			fmt.Println("It's after noon!!")
	}

	// A type `switch` compares types instead of values. You can
	// use this to discover the type of an interface value. In this
	// example, the variable t will have the type corresponding to its clause.
	whatIAm := func(i interface{}){
		switch t := i.(type) {
			case bool:
				fmt.Println("I'm a bool.")
			case int: 
				fmt.Println("I'm an int.")
			default:
				fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatIAm(true)
	whatIAm(12)
	whatIAm("hey")

}

func demoArrays(){
	/*
		In Go, an `array` is a numbered sequence of elements of
		a specific length.
	*/

	// Here we create an `array` a that will hold exactly 5 `int`s. 
	// The type of elements and length are both of the array's
	// type. By default, an array is zero-valued, which for `int`s
	// means 0s.
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the array[index] = value syntax.
	// We can get the value with array[index]
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:",a[4])

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array in one line.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// Here is an example with a 3-dim array
	var threeD [2][3][4]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				threeD[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println("3d:", threeD)

	// Note that we will see `slices` much more often than arrays in typical Go.

}

func demoSlices(){
	/*
		`slices` are a key data type in Go, giving a more powerful
		interface to sequences than arrays.

		Unlike arrays, slices are typed only by the elements they
		contain (not the number of elements). To create an empty
		slice with non-zero length, use the builtin type `make`.
	*/

	// Here we make a slice of strings of length 3 (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support
	// several more that make them richer than arrays.
	// One is the builtin `append`, which returns a slice containing one
	// or more new values. Note that we need to accept a return
	// value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("app:", s)

	// Slices can also be `copy`'d. Here we can create an
	// empty slice `c` of the same length as `s` and copy intoc from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3] and s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) s[2]:
	l = s[2:]
	fmt.Println("sl3:", l)	

	// We can declare and initialize a variable for slice
	// in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can vary,
	// unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}

	}
	fmt.Println("2d", twoD)

}

func demoMaps(){
	/*
		Maps are Go's built-in associative data type.
		This is sometimes called hashes or dicts in other
		languages.
	*/

	// To create an empty map, use the builtin `make`:
	// >>> make(map[key-type]val-type)
	m := make(map[string]int)

	// Set key/value pairs using typical `name[key] = val` syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. `Println` will show
	// all of its key/value pairs
	fmt.Println("map:", m)

	// Get a value from a key with `name[key]`
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// The builtin `len` returns the number of
	// key/value pairs when called on a map
	fmt.Println("len:", len(m))

	// The builtin `delete` removes key/value pairs.
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional secod return value when getting
	// a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys 
	// with zero values like 0 or "". Here we didn't need the value
	// itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	var some_key string = "missing_key"
	fmt.Println("Testing key:", some_key)
	value, key_exists := m[some_key]
	if key_exists {
		fmt.Println("Key " + some_key + " exists.")
		fmt.Println("Value is:", value)
	} else {
		fmt.Println("Key " + some_key + " does NOT exist.")
		fmt.Println("Value is:", value)
	}

	some_key = "k1"
	fmt.Println("Testing key:", some_key)
	value, key_exists = m[some_key]
	if key_exists {
		fmt.Println("Key " + some_key + " exists.")
		fmt.Println("Value is:", value)
	} else {
		fmt.Println("Key " + some_key + " does NOT exist.")
		fmt.Println("Value is:", value)
	}

	// You can also delcare and initialize a new map
	// in the same line with the following syntax:
	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map:", n)

}

func demoRange(){
	/*
		`range` iterates over elements in a variety of
		data structures. Let's see how to use `range` with some
		of the data structures we've already learned.
	*/

	// Here we use `range` to sum the numbers in a sclice.
	// Arrays work like this too.
	nums := []int{2,3,4}
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` on arrays and slices provides both the index and
	// value for each entry. Above we didn't need the index, 
	// so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` on map iterates over key/value pairs
	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` on strings iterate over Unicode code points.
	// The first value is the starting byte index of the `rune`
	// and the second the `rune` itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}

func demoFunctions(){
	/*
		Funcstion are central in Go. We'll learn about
		functions with a few different examples.
	*/

	// Here's a function that takes two ints and returns
	// their sum as an int. Go requires explicit returns,
	// i.e. it won't automatically return the value of the
	// last expression.
	plus := func(a int, b int) int {
		return a + b
	}

	// When you have multiple consecutive parameters of the
	// same type, you may omit the type name for the like-typed
	// parameters up to the final parameter that declares the type.
	plusPlus := func(a, b, c int) int {
		return a + b + c
	}

	// Call a function just as you'd expect, with name(args)
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)

	// Go has built-in support for multiple return values.
	// This feature is used often in idiomatic Go, for
	// example to return both result and error values from a function.

	// The (int, int) in this function signature shows that the
	// function rturns 2 ints.
	vals := func() (int, int) {
		return 3, 7
	}

	// Here we use the 2 different return values from the call with 
	// multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values, use
	// the blank identifier _.
	_, c := vals()
	fmt.Println(c)

}

func demoVariadicFunctions(){
	/*
		Variadic functions can be called with any number of
		trailing arguments. For example, `fmt.Println` is a 
		common variadic function.
	*/

	// Here's a function that will take an arbitrary number
	// of ints as arguments.
	sum := func(nums ...int) {
		fmt.Print(nums, " ")
		var total int = 0
		for _, num := range nums {
			total += num
		} 
		fmt.Println("Total:", total)
	}

	// Variadic functions can be called in the usual way with 
	// individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// appy them to a variadic function using func(slice...) like this.
	nums := []int{1,2,3,4,5,6,7}
	sum(nums...)

}

func demoClosures(){
	/*
		Go supports anonymous functions, which can form
		closures. Anonymous functions are useful when you 
		want to define a function inline without having to name it.
	*/

	// This function `intSeq` return another function, which
	// we define anonymously in the body of `intSeq`. The
	// returned function closes over the variable i to form
	// a closure.
	intSeq := func(init_value int) func() int {
		return func() int {
			init_value += 1
			return init_value
		}
	}

	// We call `intSeq`, assigning the result (a function) to
	// `nextInt`. This functions value captures its own i value,
	// which will be updated each time we call `nextInt`.
	nextInt := intSeq(10)

	// See the effect of the closure by calling `nextInt` a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular
	// function, create and test a new one.
	newInts := intSeq(100)
	fmt.Println(newInts())

}


// This `fact` function calls itself until it reaches the base case
// of `fact(0)`
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
func demoRecursion(){
	/*
		Go supports recursive functions.
	*/
	fmt.Println(fact(7))

}

func demoPointers(){
	/*
		Go supports pointers, allowing you to pass references
		to values and records within your program.
	*/
	
	// We'll show how pointers work in contrast to values
	// with 2 functions: zeroval and zeroptr.
	// zeroval has an int parameters, so arguments will be passed to it by value.
	// zeroval will get a copy of ival distinct from the one in the calling function.
	zeroval := func(ival int){
		ival = 0
	}

	// zeroptr in contrast has an *int parameters, meaning that it takes an int
	// pointer. The *iptr code in the function body then dereferences the pointer
	// from its memory address to the current value at that address. Assigning
	// a value to a dereferenced pointer changes the value at the referenced address.
	zeroptr := func(iptr *int) {
		*iptr = 0
	}

	var i int = 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The `&i` syntax gives the memory address of `i`, i.e. a pointer to `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("pointer:", &i)

	// `zeroval` doesn't change the i in `main`, but `zeroptr` does because
	// it has a reference to the memory address for that variable.

}

func demoStructs(){
	/*
		Go's `structs` are typed collections of fields. They're useful
		for grouping data together to form records.
	*/

	// This `person` struct type has `name` and `age` fields.
	type person struct {
		name string
		age int
	}

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct
	fmt.Println(
		person{
			name: "Alice",
			age: 30
		}
	)

	// Omitted fields will be zero-valued.
	fmt.Println(
		person{
			name: "Fred"
		}
	)

	// An `&` prefix yields a pointer to the struct
	fmt.Println(
		&person{
			name: "Ann",
			age: 40
		}
	)

	
}


func main(){
	fmt.Println("你好!")
	//demoValues()
	//demoVariables()
	//demoConstants()
	//demoFor()
	//demoIfElse()
	//demoSwitch()
	//demoArrays()
	//demoSlices()
	//demoMaps()
	//demoRange()
	//demoFunctions()
	//demoVariadicFunctions()
	//demoClosures()
	//demoRecursion()
	//demoPointers()
}
