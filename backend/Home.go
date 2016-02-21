package backend

import "net/http"

// LocationsAPI defines all the endpoints
//type LocationsAPI struct{}

//

func init() {
	//http.HandleFunc("/api/yelp", yelp_lots)
	http.HandleFunc("/api/park_whiz", GetParkWhizLot)
	http.HandleFunc("/api/park_panda", GetParkingPandaLots)
	http.HandleFunc("/api/lots", GetLotsFromParkingPanda)

}
