package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"appengine"
	"appengine/datastore"
)

// LocationsAPI defines all the endpoints
//type LocationsAPI struct{}

//

func init() {
	//http.HandleFunc("/api/yelp", yelp_lots)
	http.HandleFunc("/api/parkingpanda", GetParkingPandaLots)
	http.HandleFunc("/api/parkwhiz", GetParkWhizLots)
	http.HandleFunc("/api/lots", ServeLots)
	http.HandleFunc("/api/locations", ServeLocations)
	http.HandleFunc("/do/save/parkingpanda", SaveAllFromParkingPanda)
	http.HandleFunc("/do/save/parkwhiz", SaveAllFromParkWhiz)
	http.HandleFunc("/api/singlelot", SingleLot)
}

func ServeLocations(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Lot")
	var lots []Lot
	if _, err := q.GetAll(c, &lots); err != nil {
		fmt.Errorf("Encoding: %v", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	locs := make([]Location, len(lots))
	for i := 0; i < len(lots); i++ { // store all Locations names
		if lots[i].Source == "www.ParkingPanda.com" { // this is the location
			locs[i].LocationName = lots[i].Title
			locs[i].Key = lots[i].Key
		}
	}
	var rootKey *datastore.Key
	for i := 0; i < len(locs); i++ {
		rootKey = datastore.NewKey(c, "Lot", locs[i].Key, 0, nil)
		q := datastore.NewQuery("Lot").Ancestor(rootKey)
		if _, err := q.GetAll(c, &locs[i].Listing); err != nil {
			fmt.Errorf("query get all: %v", err)
		}
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(locs)
	if err != nil {
		fmt.Errorf("Encoding: %v", err)
	}

}

func SingleLot(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	var l Lot
	marketStKey := datastore.NewKey(c, "Lot", "101marketst", 0, nil)
	err := datastore.Get(c, marketStKey, &l)
	if err != nil {
		fmt.Errorf("datastore get: %v", err)
	}

	q := datastore.NewQuery("Lot").Ancestor(marketStKey)
	var lots []Lot
	if _, err := q.GetAll(c, &lots); err != nil {
		fmt.Errorf("query get all: %v", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	err = enc.Encode(lots)
	if err != nil {
		fmt.Errorf("Encoding: %v", err)
	}
}
