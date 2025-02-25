# Table of Contents

- [Hello Gogo](#hello-gogo)
- [Why We Need Go](#why-we-need-go)
- [Basic Program Structure](#basic-program-structure)
- [Variables](#variables)
- [Struct](#struct)
- [& (Address-of Operator) / * (Dereference Operator)](#-address-of-operator--dereference-operator)
- [Exported (Public) / Unexported (Private)](#exported-public--unexported-private)
- [Array](#array)
- [Encapsulation](#encapsulation)
- [Composition](#composition)
- [Concurrency](#concurrency)
- [Memory Management](#memory-management)
- [Packages and Modules](#packages-and-modules)
- [Makefile for Database Setup and Migrations](#makefile-for-database-setup-and-migrations)
- [Database Migrations. CLI](#migrations)
- [SQLC Configuration](#sqlc-configuration)
- [Final Notes](#final-notes)

---

### Hello Gogo

- Developer at Google in 2007
- Open-Source in 2009

---

### Why We Need Go

- **Multi-threading**: Multiple tasks/things at once
  - Each thread processes one task
  - You can run many of them in parallel
    - e.g.
      - Google Drive: Downloading, uploading, changing file name
      - YouTube: Watching video, commenting

- **Concurrency ≠ Exact Same Time**: Dealing with many things at once
  - Multiple users editing Google Docs/Miro/Figma, etc.
  - When processes happen at the same time
    - e.g. Multiple users trying to book a ticket simultaneously
  - Many languages have solutions
    - Complex code
    - Expensive / Slow
  - **GO**
    - Designed to run multiple codes and support concurrency
      - Cheap and easy

---

### Basic Program Structure

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

### Variables

- When two or more consecutive variables share the same type:
  ```go
  x int, y int

  // We can shorten to
  x, y int
  ```

- `var` and short variable declaration (`:=`)
  - `var` Automatically initialized to their zero values

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

### Struct

- A `struct` is a collection of fields.
  ```go
  type Student struct {
      Name   string
      Number float32
  }
  // Named Fields: Student{Name: "GO", Number: 1} ✅
  // Positional Fields: Student{"GO", 1} ✅

  // WRONG Positional Fields: Student{1, "GO"} ❌

  student := Student{Number: 1, Name: "GO"}
  student.Name // GO
  ```

---

### & (Address-of Operator) / * (Dereference Operator)

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

### Exported (Public) / Unexported (Private)

- Exported (Public): Starts with an uppercase letter and is accessible outside the package.
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

---

### Array

```go
var information [2]string

information[0] = "Go"
information[1] = "Language"

information // [Go Language]

var houseNumbers = [5]int{1, 2, 3, 4, 5} // [1 2 3 4 5]

// houseNumbers[startIndex:endIndex]
// The start index is inclusive, but the end index is exclusive.
houseNumbers[1:4]

// Using make to create an array with 0's
var dynamicIntArray = make([]int, 2, 5) // [0 0]
var dynamicIntArray = make([]int, 5) // [0 0 0 0 0]
```

---

### Encapsulation

- Encapsulation in Go is achieved through the use of exported and unexported identifiers.

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

### Composition

- Inheritance (IS-A Relationship)
- Composition (HAS-A Relationship)

- Composition is a design principle where a struct includes another struct to reuse its fields and methods instead of using inheritance.

```go
type Dog struct {
    Name      string
    Breed     string
    lastSlept time.Time // unexported
    Animal    // composition
}
```

- The Dog struct "inherits" the fields and methods of Animal without explicitly extending it.
- Because Animal is embedded, Dog automatically gains access to Animal's methods.

---

### Concurrency

- Go has built-in support for concurrency (e.g., goroutines and channels).

- `sync.WaitGroup`
  - Used to wait for multiple goroutines to finish before the main function exits.
  - `wg.Add(1)` increments the WaitGroup counter by 1, indicating that a new goroutine is being started.
  - `go func() { ... }()` launches an anonymous goroutine.
  - `defer wg.Done()` indicates completion and decreases the counter.
  - `wg.Wait()` ensures all goroutines complete before moving on.

- Channels
  - Send and receive values with the channel operator, `<-`
  - `ch := make(chan int)`

---

### Memory Management

- Go automatically frees unused memory (garbage collection).

---

### Packages and Modules

- A **package** is a collection of Go files.
- A **module** is a collection of packages.
- Initialize a module:
  ```bash
  go mod init hello_go
  ```
- The `go.mod` file specifies the Go version and dependencies.

- Use `go get` to install something:
  ```bash
  go get github.com/lib/pq
  ```
- Use `go tidy` to install the dependencies.

---

### Makefile for Database Setup and Migrations

Below is a `Makefile` to simplify database setup, migrations, and other tasks for your project. It includes commands for creating/dropping databases, running PostgreSQL in Docker, and applying migrations.

```makefile
DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

create-db:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

db-migrate:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up
```

---

### Migrations

- Database migrations. CLI
- `migrate create -ext sql -dir db/migrations -seq init_schema`

---

### SQLC Configuration

- Command line tool that generates code from SQL.
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

#### How to Use SQLC
1. Install `sqlc` if you haven't already:  
   ```bash
   brew install sqlc
   ```

2. Generate Go code from SQL queries:
   ```bash
   sqlc generate
   ```

---

### Final Notes

- The `Makefile` simplifies common tasks like database setup and migrations, making it easier for developers to get started with your project.
- The `sqlc` configuration ensures that your SQL queries are type-safe and integrated into your Go codebase.


---

