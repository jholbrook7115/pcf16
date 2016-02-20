package backend

import (
	"fmt"
	"io"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/JustinBeckwith/go-yelp/yelp"
)

func init() {

	http.HandleFunc("/", lots)
}

func lots(w http.ResponseWriter, r *http.Request) {
	options, err := getOptions(w)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, fmt.Sprintf("ERROR: %v", err))
	}

	// google app engine requires it's own class for making http requests
	c := appengine.NewContext(r)
	httpClient := urlfetch.Client(c)

	// create a new yelp client with the auth keys and the custom http client
	client := yelp.New(options, httpClient)

	// make a simple query
	term := r.URL.Query().Get("term")
	location := r.URL.Query().Get("location")

	// call the yelp API
	results, err := client.DoSimpleSearch(term, location)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, fmt.Sprintf("ERROR: %v", err))
	}

	// print the results
	io.WriteString(w, fmt.Sprintf("<div>Found a total of %v results for \"%v\" in \"%v\".</div>", results.Total, term, location))
	io.WriteString(w, "<div>-----------------------------</div>")
	for i := 0; i < len(results.Businesses); i++ {
		io.WriteString(w, fmt.Sprintf("<div>%v, %v<br>%+v</div>", results.Businesses[i].Name, results.Businesses[i].Rating, results.Businesses[i]))
	}
}
