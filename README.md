# Table of Contents

- [Hello Gogo](#hello-gogo)
- [Variables](#variables)
- [struct](#struct)
- [& (address-of operator) / * (dereference operator)](#)
- [exported (public) / unexported (private)](#exported-public--unexported-private)
- [Array / [n]T](#array)
- [Encapsulation](#encapsulation)
- [Composition](#composition)
- [Makefile for Database Setup and Migrations](#makefile-for-database-setup-and-migrations)
- [SQLC Configuration](#sqlc-configuration)
- [Final Notes](#final-notes)

---

### Hello Gogo

- Developer at Google in 2007
- Open-Source in 2009

- Why We need Go?
   - Multi-threading: Multiple tasks/things at once
      - Each thread process one task
      - you can run many of them parallel
         - e.g.
            - Google Drive: Downloading, uploading, changing file name
            - Youtube: watching video, commenting

   - Concurrency ≠ Exact same time: dealing lot of things at once
      - Multiple users editing Goolge doc/Miro/Figma etc
      - When process happen is same time
         - e.g. Multiple users try to book a ticket at same time
         - many languages have solutions
            - complex code
            - Expensive / Slow
         - **GO**
            - designed to run on mutiple code and support concurrency
               - cheap and easy



1. **Statically Typed Language**
   - Once a variable is declared with a specific type, it cannot be reassigned to a different type.
   - Example:
     ```go
     var number int = 1
     number = "one" // Compile-time error: cannot assign string to int
     ```

2. **Strongly Typed Language**
   - Operations between incompatible types (e.g., string + int) are not allowed.
   - Example:
     ```go
     a := 1
     b := "two"
     a + b // Compile-time error: mismatched types
     ```
   - Unlike JavaScript or PHP, Go does not implicitly convert types.

3. **Compiled Language**
   - Go code is compiled into machine code (binary files), enabling faster execution.
   - Example: Go is ~4x faster than Ruby for tasks like counting 100 million numbers (50ms vs 2s).
   - Ruby uses an interpreter (slower execution), while Go produces standalone binaries.

4. **Concurrency**
   - Go has built-in support for concurrency (e.g., goroutines and channels).

5. **Memory Management**
   - Go automatically frees unused memory (garbage collection).

6. **Packages and Modules**
   - A **package** is a collection of Go files.
   - A **module** is a collection of packages.
   - Initialize a module:
     ```bash
     go mod init hello_go
     ```
   - The `go.mod` file specifies the Go version and dependencies.

7. **Basic Program Structure**
   - Example:
     ```go
     package main // Entry point package
     import "fmt" // Import package for I/O

     func main() {
         fmt.Println("Hello Go.")
     }
     ```
   - Build and run:
     ```bash
     go build main.go // Produces a binary
     ./main           // Runs the binary
     go run main.go   // Compiles and runs in one step
     ```

---

### **Variables**

- When two or more consecutive variables share same type
   e.g.

   ```go
   x int, y int

   // we can shortened to

   x, y int
   ```

- `var` and short variable declaration (`:=`)
   - automatically initialized to their zero values
   ```go
   var c, python, java bool
   var i int
   var j, k int = 9, 99


	fmt.Println(i, j, k, c, python, java)
   // 0 9 99 false false false
   ```
   - Use `var` when you need zero-value initialization or package-level variables.
   - Use `:=` when inside a function and you have an initial value.

---
### **`struct`**

- A `struct` is a collection of fields.
   ```go
   type Student struct {
      Name string
      Number float32
   }
   // Named Fields: Student{Name: "GO", Number: 1} ✅
   // Positional Fields: Student{"GO", 1} ✅

   // WRONG Positional Fields: Student{1, "GO"} ❌

   student := Student{ Number: 1, Name: "GO"}
   student.Name // GO

   // student := Student{Number: 1, Name: 'GO'} // This will give an error! ❌
   // char := 'G' // A single rune (character) - can only assign a single character
   ```
      - A struct literal is simply an expression used to create and initialize a struct.

---

### **& (address-of operator) / * (dereference operator)**

- `&` is used to get the address of a variable.
- `*` is used to access or modify the value stored at the address a pointer is pointing to.
   ```go
      var name, city string = "Go", "Berlin"

      var namePointer = &name
      var cityPointer = &city

      namePointer // Memory address of Name
      *namePointer // Dereference to get value of Name

      *namePointer = "Golang"   // Change the value at the memory address pointed by namePointer
   ```

---

### **exported (public) / unexported (private)**

- Exported (Public): Starts with an uppercase letter and accessible outside the package.
  - FuncName, VariableName, StructName, etc.
- Unexported (Private): Starts with a lowercase letter and can only be called from within the package.
  - funcName, variableName, structName, etc.

```go
package pets

type Dog struct {
	Name  string
	Breed string

   lastSlept time.Time
}

```

[pets/dog.go](https://github.com/sajjadmurtaza/hello_go/blob/main/pets/dog.go)

---

### **array**

   ```go
   var information [2]string

	information[0] = "Go"
	information[1] = "Language"

	information // [Go Language]

	var houseNumbers = [5]int{1,2,3,4,5} // [1 2 3 4 5]

   // houseNumbers[startIndex:endIndex]
   // The start index is inclusive, but the end index is exclusive.
   houseNumbers[1:4]


   // using a make
   // we can create array with 0's
   var dynamicIntArray = make([]int, 2, 5) // [0 0]
   var dynamicIntArray = make([]int, 5) // [0 0 0 0 0 ]

 
   ```

---
### **Encapsulation**
- Encapsulation in Go is achieved through the use of exported and unexported identifiers

```go
   package main

	pet := pets.Dog{
		Name:  "Fast Doggy",
		Breed: "Strange",
	}

	fmt.Println(pet.Name)          // ✅ Allowed (exported)
	fmt.Println(pet.Breed)         // ✅ Allowed (exported)
	// fmt.Println(pet.lastSlept)  // ❌ ERROR: lastSlept is unexported
```


---
### **Composition**

- Inheritance (IS-A Relationship)
- Composition (HAS-A Relationship)

- Composition is a design principle a struct includes another struct to reuse its fields and methods, instead of using inheritance.

```go
type Dog struct {
	Name  string
	Breed string

	lastSlept time.Time // unexported

	Animal // composition
}
```

- The Dog struct "inherits" the fields and methods of Animal without explicitly extending it.
- Because Animal is embedded, Dog automatically gains access to Animal's methods.
   - Animal struct provides the lastAte field and Feed() method.


[pets/animal.go](https://github.com/sajjadmurtaza/hello_go/blob/main/pets/animal.go) |
[pets/dog.go](https://github.com/sajjadmurtaza/hello_go/blob/main/pets/dog.go)



---

### **Makefile for Database Setup and Migrations**

Below is a `Makefile` to simplify database setup, migrations, and other tasks for your project. It includes commands for creating/dropping databases, running PostgreSQL in Docker, and applying migrations.

```makefile

DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

create-db:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

db-migrate:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

```
---
- An interface defines method signatures but does not implement them.
---

### **SQLC Configuration**

- command line tool that generates code from SQL
- `sqlc init`

Below is the `sqlc` configuration file (`sqlc.yaml`) for generating Go code from SQL queries:

```yaml
version: "2"
sql:
  - schema: "db/migrations/"
    queries: "db/queries/"
    engine: "postgresql"
    gen:
      go:
        package: "db"
        out: "app"
```

#### **How to Use SQLC**
1. Install `sqlc` if you haven’t already:  
   ```bash
   brew install sqlc
   ```

2. Generate Go code from SQL queries:  
   ```bash
   sqlc generate
   ```

---

### **Final Notes**
- The `Makefile` simplifies common tasks like database setup and migrations, making it easier for developers to get started with your project.
- The `sqlc` configuration ensures that your SQL queries are type-safe and integrated into your Go codebase.


---

