package main

import (
   "fmt"
   "log"
   "net/http"
   "io/ioutil"
   "encoding/json"
   "math/rand"
   "strconv"
)

type NasaViewer struct {
	Photos []struct {
		ID     int `json:"id"`
		Sol    int `json:"sol"`
		Camera struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			RoverID  int    `json:"rover_id"`
			FullName string `json:"full_name"`
		} `json:"camera"`
		ImgSrc    string `json:"img_src"`
		EarthDate string `json:"earth_date"`
		Rover     struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			LandingDate string `json:"landing_date"`
			LaunchDate  string `json:"launch_date"`
			Status      string `json:"status"`
			MaxSol      int    `json:"max_sol"`
			MaxDate     string `json:"max_date"`
			TotalPhotos int    `json:"total_photos"`
			Cameras     []struct {
				Name     string `json:"name"`
				FullName string `json:"full_name"`
			} `json:"cameras"`
		} `json:"rover"`
	} `json:"photos"`
}

type NasaPayload struct {
  Img string
}

func handler(w http.ResponseWriter, r *http.Request) {
  var page = strconv.Itoa(rand.Intn(5))
  response, _ := http.Get("https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&page=" + page + "&api_key=")
  catalog, _ := ioutil.ReadAll(response.Body)

  var nasaViewer NasaViewer
  err := json.Unmarshal(catalog, &nasaViewer)
  if err != nil {
    fmt.Fprintf(w, "Error during Unmarshal of data")
  }

  var randArrValue = rand.Intn(len(nasaViewer.Photos))
  nasaPayload := NasaPayload{ nasaViewer.Photos[randArrValue].ImgSrc }
  js, err := json.Marshal(nasaPayload)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(js)
}

func main() {
   http.HandleFunc("/rover", handler)
   log.Fatal(http.ListenAndServe(":8000", nil))
}
