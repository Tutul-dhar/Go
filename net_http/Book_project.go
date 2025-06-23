package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/google/uuid"
)


type Book struct {
    UUID        string   `json:"uuid"`
    Name        string   `json:"name"`
    AuthorList  []string `json:"authorList"`
    PublishDate string   `json:"publishDate"`
    ISBN        string   `json:"isbn"`
}

var bookStore = make(map[string]Book)


func createBook(w http.ResponseWriter, r *http.Request) {
    var book Book

    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte(fmt.Sprintf("Invalid request: %v", err)))
        return
    }

    book.UUID = uuid.NewString()

    bookStore[book.UUID] = book

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}


func getBook(w http.ResponseWriter, r *http.Request) {
    bookId := chi.URLParam(r, "id")

    book, found := bookStore[bookId]

    if !found {
        w.WriteHeader(404)

        w.Write([]byte("Book not Found"))
        return
    }

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(book)
}

func listBooks(w http.ResponseWriter, r *http.Request) {
    
    var books []Book
    for _, book := range bookStore {
        books = append(books, book)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    bookID := chi.URLParam(r, "id")

    _, found := bookStore[bookID]
    if !found {
        w.WriteHeader(404)
        w.Write([]byte("Book not found"))
        return
    }

    var updatedBook Book
    err := json.NewDecoder(r.Body).Decode(&updatedBook)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte(fmt.Sprintf("Invalid request: %v", err)))
        return
    }

    updatedBook.UUID = bookID

    bookStore[bookID] = updatedBook

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    bookID := chi.URLParam(r, "id")

    book, found := bookStore[bookID]
    if !found {
        w.WriteHeader(404)
        w.Write([]byte("Book not found"))
        return
    }

    delete(bookStore, bookID)


    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}


// func main() {
//     r := chi.NewRouter()

//     r.Use(middleware.Logger)

//     r.Post("/api/v1/books", createBook)

//     r.Get("/api/v1/books/{id}", getBook)

//     r.Get("/api/v1/books", listBooks)

//     r.Put("/api/v1/books/{id}", updateBook)

//     r.Delete("/api/v1/books/{id}", deleteBook)

//     http.ListenAndServe(":8080", r)
// }

// using mount ... ... 
func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Mount("/api/v1/books", bookRoutes())

    http.ListenAndServe(":8080", r)
}

// bookRoutes returns the subrouter mounted at /api/v1/books
func bookRoutes() http.Handler {
    r := chi.NewRouter()

    r.Post("/", createBook)         // POST    /api/v1/books
    r.Get("/", listBooks)           // GET     /api/v1/books
    r.Get("/{id}", getBook)         // GET     /api/v1/books/{id}
    r.Put("/{id}", updateBook)      // PUT     /api/v1/books/{id}
    r.Delete("/{id}", deleteBook)   // DELETE  /api/v1/books/{id}

    return r
}








-------------------------------------------------------------------------------------------------
Adding cli/flag, authentucation...

package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "net/http"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/jwtauth/v5"
    "github.com/google/uuid"
)


type Book struct {
    UUID        string   `json:"uuid"`
    Name        string   `json:"name"`
    AuthorList  []string `json:"authorList"`
    PublishDate string   `json:"publishDate"`
    ISBN        string   `json:"isbn"`
}

var (
    tokenAuth *jwtauth.JWTAuth
    adminUser = "AdminUser"
    adminPass = "AdminPassword"
    bookStore = make(map[string]Book)
)


func createBook(w http.ResponseWriter, r *http.Request) {
    var book Book

    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte(fmt.Sprintf("Invalid request: %v", err)))
        return
    }

    book.UUID = uuid.NewString()

    bookStore[book.UUID] = book

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}


func getBook(w http.ResponseWriter, r *http.Request) {
    bookId := chi.URLParam(r, "id")

    book, found := bookStore[bookId]

    if !found {
        w.WriteHeader(404)

        w.Write([]byte("Book not Found"))
        return
    }

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(book)
}

func listBooks(w http.ResponseWriter, r *http.Request) {
    
    var books []Book
    for _, book := range bookStore {
        books = append(books, book)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    bookID := chi.URLParam(r, "id")

    _, found := bookStore[bookID]
    if !found {
        w.WriteHeader(404)
        w.Write([]byte("Book not found"))
        return
    }

    var updatedBook Book
    err := json.NewDecoder(r.Body).Decode(&updatedBook)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte(fmt.Sprintf("Invalid request: %v", err)))
        return
    }

    updatedBook.UUID = bookID

    bookStore[bookID] = updatedBook

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    bookID := chi.URLParam(r, "id")

    book, found := bookStore[bookID]
    if !found {
        w.WriteHeader(404)
        w.Write([]byte("Book not found"))
        return
    }

    delete(bookStore, bookID)


    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}

func init() {
    tokenAuth = jwtauth.New("HS256", []byte("your-secret-key"), nil)
}

// jwt token
func getTokenHandler(w http.ResponseWriter, r *http.Request) {
    user, pass, ok := r.BasicAuth()
    if !ok || user != adminUser || pass != adminPass {
        w.Header().Set("WWW-Authenticate", Basic realm="Restricted")
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    //exp := jwtauth.ExpireIn(100000 * 60) // returns int64 ...not work
    exp := time.Now().Add(100 * time.Minute).Unix()

    _, tokenString, err := tokenAuth.Encode(map[string]interface{}{
        "user_id":  108,
        "username": adminUser,
        "exp":      exp,
    })
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": tokenString,
    })
}


// using mount ... ... 
func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Mount("/api/v1/books", bookRoutes())
    r.Get("/api/v1/get-token", getTokenHandler)
    
    var port string
    flag.StringVar(&port, "port", "8080", "Port to run the book server")
    flag.Parse()
    addr := fmt.Sprintf(":%s", port)
    fmt.Println("Starting server or port", port)

    http.ListenAndServe(addr, r)
    // http.ListenAndServe(":8080", r)
}

// bookRoutes returns the subrouter mounted at /api/v1/books
func bookRoutes() http.Handler {
    r := chi.NewRouter()

    r.Post("/", createBook)         // POST    /api/v1/books
    r.Get("/", listBooks)           // GET     /api/v1/books
    r.Get("/{id}", getBook)         // GET     /api/v1/books/{id}
    r.Put("/{id}", updateBook)      // PUT     /api/v1/books/{id}
    r.Delete("/{id}", deleteBook)   // DELETE  /api/v1/books/{id}

    return r
}
