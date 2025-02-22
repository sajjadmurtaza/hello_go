### **Hello Gogo**

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

### **SQLC Configuration**

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



- Encapsulation in Go is achieved through the use of exported and unexported identifiers


- Inheritance (IS-A Relationship)
- Inheritance is a mechanism where a child (subclass) inherits properties and behavior from a parent (superclass).

- Composition (HAS-A Relationship)
- Composition is a design principle where one struct contains another struct as a field instead of inheriting from it.

- interface is collection of method signatures