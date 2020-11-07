package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)
type server struct {
	r *httprouter.Router
}

var (
	openWeatherAPIKEY = os.Getenv("OPEN_WEATHER_API_KEY")
	openWeatherMapDomain = "api.openweathermap.org"
)

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
        log.Printf("ERROR: getWeatherByCity(): param city not provided")
        return
	}	
	city := keys[0]
	log.Printf("city: %s", city)
	url := fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s",openWeatherMapDomain, city, openWeatherAPIKEY)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: getWeatherByCity() error fetching weather for url %s : %v", url, err)
		return
	}
	defer resp.Body.Close()
	weather, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERROR: getWeatherByCity() error reading from for url %s", url)
		return
	}
	w.Write(weather)
	return 
}




func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	s.r.ServeHTTP(w, r)
}