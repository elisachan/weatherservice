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
	port = os.Getenv("PORT") 
)

func main() {
	router := httprouter.New()
	router.GET("/api/getweather", getWeatherByCity)

	// serve FE files
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("dist")))
	mux.Handle("/api/", router)
	// quick handle for heroku port binding
	if port == "" {
		port = "8080"
	}
	// setup cacheTTL 
	go clearCache()
	log.Fatal(http.ListenAndServe(":" + port, mux))
}


// returns weather info based off city in query string
// if city was called within 5 min prior, will return last set info; 
// otherwise will fetch new weather data from api and load cache
func getWeatherByCity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	keys, ok := r.URL.Query()["city"]
	if !ok || len(keys[0]) < 1 {
		log.Printf("ERROR: getWeatherByCity(): param city not provided")
		w.Write([]byte("Sorry you'll need to enter in a city"))
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
			w.Write([]byte("Sorry, there was a problem fetching the weather info. Please try again"))
		}
		defer resp.Body.Close()
		weather, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("ERROR: getWeatherByCity() error reading from for url %s", url)
			w.Write([]byte("sorry, there was a problem fetching the weather info. Please try again"))
		}
		// add to cache
		miniWeatherCache[city] = weather
		w.Write(weather)
	}
}

// set to clear cache based on cachettl set
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
