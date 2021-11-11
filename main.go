package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/cors"
)

func main() {
	port := 7000
	dirStatic := filepath.Join("..", "www")
	nameHost, err := os.Hostname()
	if err != nil {
		log.Println(err.Error())
		return
	}
	/*
		str := strings.Split(os.Getenv("SESSION_MANAGER"), `/`)[1]
		nameHost := str[:len(str)-2]
	*/
	log.Println(nameHost)
	startHttpServer(port, dirStatic, nameHost)
}

func startHttpServer(port int, dirStatic string, nameHost string) {
	log.Println("Server started @ http://" + nameHost + ":" + strconv.Itoa(port))
	mux := http.NewServeMux()
	mux.Handle(`/`, http.FileServer(http.Dir(dirStatic)))
	//mux.HandleFunc("/", handlerRoot)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(`:`+strconv.Itoa(port), handler))
}

/*
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	pageContent, err := ioutil.ReadFile(`../aavarthi/index.html`)
	if err != nil {
		w.Write([]byte("File read Error"))
		return
	}
	w.Write([]byte(pageContent))
}
*/
