package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const worldTimeAPI = "https://worldtimeapi.org/api/timezone/America/Toronto"

type TimeInfo struct {
	Datetime string `json:"datetime"`
}

func getTorontoTime() (string, error) {
	resp, err := http.Get(worldTimeAPI)

	if err != nil {
		return "Error retriving data", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error readinging data", err
	}
	var TimeInfo TimeInfo
	err = json.Unmarshal(body, &TimeInfo)
	if err != nil {
		return "Error parsing data", err
	}

	return TimeInfo.Datetime, nil
}

func TorontoTimeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime, err := getTorontoTime()
	if err != nil {
		http.Error(w, "error Fetchin Toronto time ", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Toronto time is %s\n", torontoTime)
	resp := map[string]string{"current_Time_Toronto": torontoTime}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func main() {

	http.HandleFunc("/api/torontotime", TorontoTimeHandler)

	fmt.Println("Server Started  at port 8014...")
	log.Fatal(http.ListenAndServe(":8014", nil))

}
