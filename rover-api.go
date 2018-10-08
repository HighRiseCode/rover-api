package main

import (
   "fmt"
   "log"
   "net/http"
   "io/ioutil"
   "encoding/json"
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

func handler(w http.ResponseWriter, r *http.Request) {
  /*response, _ := http.Get("https://api.nasa.gov/mars-photos/api/v1/rovers/curiosity/photos?sol=1000&api_key=DEMO_KEY")
  temp, _ := ioutil.ReadAll(response.Body)

  var nasaViewer NasaViewer
  err := json.Unmarshal(temp, &nasaViewer)
  if err != nil {
    fmt.Fprintf(w, "Error during Unmarshal of data")
  }

  fmt.Fprintf(w, "ID: ", nasaViewer.Photos)*/

  fmt.Fprintf(w, "<p>NASA IMAGE HERE</p>")
}

func main() {
   http.HandleFunc("/", handler)
   log.Fatal(http.ListenAndServe(":8080", nil))
}
