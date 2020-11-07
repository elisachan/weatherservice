package main


import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

)
type server struct {
	r *httprouter.Router
}

func main() {
	router := httprouter.New()
	// todo build FE page
	// router.Handle("GET", "/", getWeatherByCity)
	router.GET("/getweather", getWeatherByCity)
	log.Fatal(http.ListenAndServe(":8082", &server{router}))
}



func getWeatherByCity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	keys, ok := r.URL.Query()["city"]
    if !ok || len(keys[0]) < 1 {
        log.Println("ERROR: getWeatherByCity(): param city not provided")
        return
	}	
	city := keys[0]
	log.Printf("city: %s", city)
	return 
}




func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	s.r.ServeHTTP(w, r)
}