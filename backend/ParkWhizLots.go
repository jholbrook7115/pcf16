package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"appengine/datastore"

	"appengine/urlfetch"

	"appengine"
)

type ParkWhizLot struct {
	Listings []struct {
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

func fetchParkWhizLots(c appengine.Context) ([]Lot, error) {

	client := urlfetch.Client(c)
	res, err := client.Get(urlParkWhiz)
	if err != nil {
		fmt.Errorf("get: %v", err)
	}

	var listing ParkWhizLot

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&listing)
	if err != nil {
		fmt.Errorf("decode: %v", err)
	}
	Result := make([]Lot, len(listing.Listings))
	for i := 0; i < len(listing.Listings); i++ {
		Result[i] = Lot{
			Title:          listing.Listings[i].LocationName,
			Price:          float64(listing.Listings[i].Price),
			Lat:            listing.Listings[i].Lat,
			Lng:            listing.Listings[i].Lng,
			DisplayAddress: listing.Listings[i].Address,
			Source:         listing.Listings[i].ParkwhizURL,
		}
	}

	return Result, nil
}

func SaveAllFromParkWhiz(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	lots, err := fetchParkWhizLots(c)
	if err != nil {
		fmt.Errorf("fetchParkWhizLots: %v", err)
	}
	io.WriteString(w, "<html><body>")
	for i := 0; i < len(lots); i++ {
		// clean name
		name := strings.Join(strings.Fields(clean(lots[i].DisplayAddress)), "")
		// is the lot in the datastore alread
		var parentLot Lot

		io.WriteString(w, fmt.Sprintf("<h3>Name: %v</h3>", name))

		parentLotKey := datastore.NewKey(c, "Lot", name, 0, nil)
		err = datastore.Get(c, parentLotKey, &parentLot)
		if err != nil {
			fmt.Errorf("datastore get: %v", err)
		}
		if parentLot.Title == "" { // lot doesnt exist; create root
			_, err = datastore.Put(c, parentLotKey, &lots[i])
			if err != nil {
				fmt.Errorf("datastore put: %v", err)
			}
			io.WriteString(w,
				fmt.Sprintf("<p>Saved %v as root<br><%v></p>",
					lots[i], err))

		} else { // lot exists and we make whiz a child
			childLotKey := datastore.NewKey(c, "Lot", name, 0, parentLotKey)

			_, err = datastore.Put(c, childLotKey, &lots[i])
			if err != nil {
				fmt.Errorf("datastore put: %v", err)
			}

			io.WriteString(w,
				fmt.Sprintf("<p>Saved %v under %v<br><%v></p>",
					lots[i], parentLot, err))
		}
	}
	io.WriteString(w, "</body></html>")
}

func GetParkWhizLots(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	res, err := client.Get(urlParkWhiz)
	if err != nil {
		fmt.Errorf("get: %v", err)
	}

	var listing ParkWhizLot

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&listing)
	if err != nil {
		fmt.Errorf("decode: %v", err)
	}

	io.WriteString(w, "<div>-----------------------------</div>")
	for i := 0; i < len(listing.Listings); i++ {
		io.WriteString(w, fmt.Sprintf("<div>%v has a price of %v</div>",
			listing.Listings[i].LocationName,
			listing.Listings[i].Price))
	}
}
