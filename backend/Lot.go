package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"appengine"
)

type Location struct {
	LocationName string `json:"location_name"`
	Parking      []Lot  `json:"parking_list"`
}
type Lot struct {
	Title    string  `json:"title"`
	Source   string  `json:"source"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
	//	BookingURL string  `json:"booking_url"`
}

func GetLotsFromParkingPanda(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	lots, err := fetchParkingPandaLots(c)
	if err != nil {
		fmt.Errorf("fetchParkingPandaLots: %v", err)
	}
	/*io.WriteString(w, "<div>-----------------------------</div>")
	for i := 0; i < len(lots); i++ {
		io.WriteString(w,
			fmt.Sprintf("<center><div><img src='%v'>%v from %v, has price %v</div></center>",
				lots[i].ImageURL, lots[i].Title, lots[i].Source, lots[i].Price))
	}*/

	enc := json.NewEncoder(w)
	err = enc.Encode(lots)
	if err != nil {
		fmt.Errorf("Encoding: %v", err)
	}
}
