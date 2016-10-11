package main
import (
"net/http"
)
func main() {
mux := http.NewServeMux()
fs := http.FileServer(http.Dir("public"))
mux.Handle("/", fs)
http.ListenAndServe(":1908", mux)
}