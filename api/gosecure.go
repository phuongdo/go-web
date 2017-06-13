package main
import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}
/**
* openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
*/
func main() {
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)
	http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
}