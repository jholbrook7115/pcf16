package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"appengine"

	"appengine/urlfetch"
)

type ParkWhizList struct {
	ParkingListings []struct {
		LocationName    string  `json:"location_name"`
		LocationID      int     `json:"location_id"`
		ListingID       int     `json:"listing_id"`
		Start           int     `json:"start"`
		End             int     `json:"end"`
		ParkwhizURL     string  `json:"parkwhiz_url"`
		APIURL          string  `json:"api_url"`
		Address         string  `json:"address"`
		City            string  `json:"city"`
		State           string  `json:"state"`
		Zip             string  `json:"zip"`
		Lat             float64 `json:"lat"`
		Lng             float64 `json:"lng"`
		Distance        int     `json:"distance"`
		Recommendations int     `json:"recommendations"`
		Reservation     int     `json:"reservation"`
		Eticket         int     `json:"eticket"`
		Valet           int     `json:"valet"`
		Indoor          int     `json:"indoor"`
		Shuttle         int     `json:"shuttle"`
		Tailgate        int     `json:"tailgate"`
		Security        int     `json:"security"`
		Restroom        int     `json:"restroom"`
		Attended        int     `json:"attended"`
		Rv              int     `json:"rv"`
		AvailableSpots  int     `json:"available_spots"`
		Price           int     `json:"price"`
		PriceFormatted  string  `json:"price_formatted"`
		WwwReserveURL   string  `json:"www_reserve_url"`
		APIReserveURL   string  `json:"api_reserve_url"`
	} `json:"parking_listings"`
}

func GetParkWhizLot(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	res, err := client.Get(urlParkWhiz)
	if err != nil {
		fmt.Errorf("get: %v", err)
	}

	var listing ParkWhizList

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&listing)
	if err != nil {
		fmt.Errorf("decode: %v", err)
	}

	io.WriteString(w, "<div>-----------------------------</div>")
	for i := 0; i < len(listing.ParkingListings); i++ {
		io.WriteString(w, fmt.Sprintf("<div>%v has a price of %v<p>%v</p></div>",
			listing.ParkingListings[i].LocationName,
			listing.ParkingListings[i].Price,
			listing.ParkingListings[i].Address,
		))
	}
}

//func (p *ParkWhizList) Save() error {}

//func fetchParkWhizLots(c appengine.Context) ([]Lot, error) {}
