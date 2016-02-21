package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"appengine"
	"appengine/datastore"
)

type Location struct {
	LocationName string `json:"location_name"`
	Parking      []Lot  `json:"parking_list"`
}
type Lot struct {
	Title          string  `json:"title"`
	Source         string  `json:"source"`
	ImageURL       string  `json:"image_url"`
	Price          float64 `json:"price"`
	Lat            float64 `json:"lat"`
	Lng            float64 `json:"lng"`
	Key            string
	DisplayAddress string `json:"display_address"`
	//	BookingURL string  `json:"booking_url"`
}

func GetLotsFromParkingPanda(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	lots, err := fetchParkingPandaLots(c)
	if err != nil {
		fmt.Errorf("fetchParkingPandaLots: %v", err)
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(lots)
	if err != nil {
		fmt.Errorf("Encoding: %v", err)
	}
}

func SaveAllFromParkingPanda(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	lots, err := fetchParkingPandaLots(c)
	if err != nil {
		fmt.Errorf("fetchParkingPandaLots: %v", err)
	}

	for i := 0; i < len(lots); i++ {
		k := datastore.NewKey(c, "Lot", lots[i].Key, 0, nil)
		_, err := datastore.Put(c, k, &lots[i])
		if err != nil {
			fmt.Errorf("datastore.Put: %v", err)
		}
		io.WriteString(w, fmt.Sprintf("<p>Saved: %v; ERR = %v</p>", lots[i], err))
	}

}
