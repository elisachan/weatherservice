package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type server struct {
	r *httprouter.Router
}

var (
	openWeatherAPIKEY    = os.Getenv("OPEN_WEATHER_API_KEY")
	openWeatherMapDomain = "api.openweathermap.org"
	//cache will hold city and weather info for 5 min (assumes weather does not change much in that time)
	//used to prevent unnecessary repeat/duplicate calls to api
	miniWeatherCache = make(map[string][]byte)
	cacheTTL = 5 * time.Minute
)

func main() {
	router := httprouter.New()
	// todo build FE page
	// router.Handle("GET", "/", getWeatherByCity)
	router.GET("/getweather", getWeatherByCity)

	// setup cacheTTL 
	go clearCache()
	log.Fatal(http.ListenAndServe(":8082", &server{router}))
}


// returns weather info based off city in query string
// if city was called within 5 min prior, will return last set info; 
// otherwise will fetch new weather data from api and load cache
func getWeatherByCity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	keys, ok := r.URL.Query()["city"]
	if !ok || len(keys[0]) < 1 {
		log.Printf("ERROR: getWeatherByCity(): param city not provided")
		return
	}
	city := keys[0]
	if val, ok := miniWeatherCache[city]; ok {
		log.Printf("INFO: getWeatherByCity(): city %s previously fetched, returning last given info", city)
		w.Write(val)
	} else {
		url := fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s", openWeatherMapDomain, city, openWeatherAPIKEY)
		log.Printf("INFO: getWeatherByCity(): fetching info from api %s", url)
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
		// add to cache
		miniWeatherCache[city] = weather
		w.Write(weather)
	}
}

func clearCache(){
	for {
		// assign cache to new map 
		miniWeatherCache = make(map[string][]byte)
		log.Printf("cache cleared and reset")
		time.Sleep(cacheTTL)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	s.r.ServeHTTP(w, r)
}
