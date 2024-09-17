package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type DaySummaryResponse struct {
	Coord      coordinates `json:"coord,omitempty"`
	Weather    []weather   `json:"weather,omitempty"`
	Base       string      `json:"base,omitempty"`
	Main       maininfo    `json:"main"`
	Visibility uint16      `json:"visibility,omitempty"`
	Wind       wind        `json:"wind,omitempty"`
	Clouds     clouds      `json:"clouds,omitempty"`
	DT         uint64      `json:"dt,omitempty"`
	Sys        sys         `json:"sys,omitempty"`
	Timezone   uint16      `json:"timezone,omitempty"`
	ID         uint        `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	StatusCode string      `json:"status_code,omitempty"`
}

type coordinates struct {
	Lon float32 `json:"lon,omitempty"`
	Lat float32 `json:"lat,omitempty"`
}
type weather struct {
	ID          uint   `json:"id,omitempty"`
	Main        string `json:"main,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
}
type maininfo struct {
	Temp       float32 `json:"temp,omitempty"`
	Feels_like float32 `json:"feels___like,omitempty"`
	Temp_min   float32 `json:"temp___min,omitempty"`
	Temp_max   float32 `json:"temp___max,omitempty"`
	Pressure   uint16  `json:"pressure,omitempty"`
	Humidity   uint8   `json:"humidity,omitempty"`
	Sea_level  uint32  `json:"sea___level,omitempty"`
	Grnd_level uint32  `json:"grnd___level,omitempty"`
}

type wind struct {
	Speed float32 `json:"speed,omitempty"`
	Deg   int16   `json:"deg,omitempty"`
	Gust  float32 `json:"gust,omitempty"`
}

type clouds struct {
	All int8 `json:"all,omitempty"`
}

type sys struct {
	Type    int8   `json:"type,omitempty"`
	ID      int    `json:"id,omitempty"`
	Country string `json:"country,omitempty"`
	Sunrise int    `json:"sunrise,omitempty"`
	Sunset  int    `json:"sunset,omitempty"`
}

// https://api.openweathermap.org/data/3.0/onecall/timemachine?lat={lat}&lon={lon}&dt={time}&appid={API key}
func main() {

	output := DaySummaryResponse{}

	resp, err := getContent(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=54.972638&lon=-7.857291&appid=%s", os.Getenv("OWM_API_KEY")))
	if err != nil {
		log.Println(err)
	}

	unmErr := json.Unmarshal(resp, &output)

	if unmErr != nil {
		log.Println(unmErr)
	}

	fmt.Println(output.Coord)

}

func getContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}

	return data, nil
}
