package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// LocationsAPI defines all the endpoints
//type LocationsAPI struct{}

//

func init() {
	//http.HandleFunc("/api/yelp", yelp_lots)
	http.HandleFunc("/api/parkingpanda", GetParkingPandaLots)
	http.HandleFunc("/api/parkwhiz", GetParkWhizLots)
	http.HandleFunc("/api/lots", ServeLots)
	http.HandleFunc("/do/save/parkingpanda", SaveAllFromParkingPanda)
	http.HandleFunc("/do/save/parkwhiz", SaveAllFromParkWhiz)
	http.HandleFunc("/api/singlelot", SingleLot)
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
		fmt.Errorf("Encoding: %v", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	err = enc.Encode(lots)
	if err != nil {
		fmt.Errorf("Encoding: %v", err)
	}
}
