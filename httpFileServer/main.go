package main

//https://blog.csdn.net/idwtwt/article/details/81180460
import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Args)
	http.Handle("/", http.FileServer(http.Dir(".")))//to wrapperHTMLFile
	http.Handle("/api", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}
