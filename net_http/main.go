// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
// )

// func main() {
//     // Serve static files from the "static" directory
//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         path := "static/" + r.URL.Path[1:]
//         if r.URL.Path == "/" {
//             path = "static/index.html" // default to index.html
//         }
//         http.ServeFile(w, r, path)
//     })

//     // Still support dynamic route
//     http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "Hi")
//     })

//     log.Println("Serving on http://localhost:8081")
//     log.Fatal(http.ListenAndServe(":8081", nil))
// }



// package main

// import (
//     "log"
//     "net/http"
// )

// func main() {
//     // Serve everything from the "static" folder
//     http.Handle("/", http.FileServer(http.Dir("./static")))

//     log.Println("Serving at http://localhost:8081")
//     log.Fatal(http.ListenAndServe(":8081", nil))
// }


package main

import (
    "log"
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./static")))

    log.Println("Serving over HTTPS at https://localhost:8081")
    err := http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
    if err != nil {
        log.Fatal("ListenAndServeTLS error: ", err)
    }
}
