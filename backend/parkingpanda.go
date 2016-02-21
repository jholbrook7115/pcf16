package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"appengine"

	"appengine/urlfetch"
)

type ParkingPandaList struct {
	Data struct {
		IsLive    bool        `json:"isLive"`
		Venue     interface{} `json:"venue"`
		Locations []struct {
			AffiliateURL              interface{}   `json:"affiliateUrl"`
			AppURL                    string        `json:"appUrl"`
			ArriveParts               []string      `json:"arriveParts"`
			CloseTime                 string        `json:"closeTime"`
			ConvenienceFee            float64       `json:"convenienceFee"`
			DailyAmounts              []float64     `json:"dailyAmounts"`
			DepartParts               []string      `json:"departParts"`
			Distance                  float64       `json:"distance"`
			DistanceFormatted         interface{}   `json:"distanceFormatted"`
			DistanceType              string        `json:"distanceType"`
			DriveUpPrice              interface{}   `json:"driveUpPrice"`
			DurationDescription       string        `json:"durationDescription"`
			EndDate                   interface{}   `json:"endDate"`
			EndDateAndTime            string        `json:"endDateAndTime"`
			EndDateAndTimeFormatted   string        `json:"endDateAndTimeFormatted"`
			EndTime                   string        `json:"endTime"`
			HoursAfterEvent           interface{}   `json:"hoursAfterEvent"`
			IsAvailable               bool          `json:"isAvailable"`
			IsCheapest                bool          `json:"isCheapest"`
			IsEventOnlyParking        bool          `json:"isEventOnlyParking"`
			IsLive                    bool          `json:"isLive"`
			IsSoldOut                 bool          `json:"isSoldOut"`
			IsSuggested               bool          `json:"isSuggested"`
			IsTimeExtended            bool          `json:"isTimeExtended"`
			MinimumStayMinutes        interface{}   `json:"minimumStayMinutes"`
			OpenTime                  string        `json:"openTime"`
			ParkingTypes              interface{}   `json:"parkingTypes"`
			NumberOfGamesInPackage    int           `json:"numberOfGamesInPackage"`
			PandaDiscount             float64       `json:"pandaDiscount"`
			PandaRating               float64       `json:"pandaRating"`
			Price                     float64       `json:"price"`
			ServiceFee                float64       `json:"serviceFee"`
			SiteDiscountAmount        interface{}   `json:"siteDiscountAmount"`
			Slug                      string        `json:"slug"`
			StartDate                 interface{}   `json:"startDate"`
			StartDateAndTime          string        `json:"startDateAndTime"`
			StartDateAndTimeFormatted string        `json:"startDateAndTimeFormatted"`
			StartTime                 string        `json:"startTime"`
			TimeAmount                interface{}   `json:"timeAmount"`
			Upsells                   []interface{} `json:"upsells"`
			URL                       string        `json:"url"`
			WalkingTime               interface{}   `json:"walkingTime"`
			ID                        int           `json:"id"`
			IDUser                    int           `json:"idUser"`
			Address1                  string        `json:"address1"`
			AllowOvernight            bool          `json:"allowOvernight"`
			City                      string        `json:"city"`
			Country                   string        `json:"country"`
			Currency                  int           `json:"currency"`
			CurrencyFormatted         string        `json:"currencyFormatted"`
			CurrencyVerbose           string        `json:"currencyVerbose"`
			CustomAgreement           string        `json:"customAgreement"`
			EntranceAddress           interface{}   `json:"entranceAddress"`
			EntranceExits             interface{}   `json:"entranceExits"`
			EntranceLatitude          interface{}   `json:"entranceLatitude"`
			EntranceLongitude         interface{}   `json:"entranceLongitude"`
			Guarantee                 string        `json:"guarantee"`
			IntegrationType           int           `json:"integrationType"`
			IsAutomated               bool          `json:"isAutomated"`
			IsGateEnabled             bool          `json:"isGateEnabled"`
			IsOpen247                 bool          `json:"isOpen247"`
			Label                     struct {
				BackgroundColor string `json:"backgroundColor"`
				HoverText       string `json:"hoverText"`
				Text            string `json:"text"`
				TextColor       string `json:"textColor"`
			} `json:"label"`
			Latitude                   float64 `json:"latitude"`
			Longitude                  float64 `json:"longitude"`
			MaxVehicleHeight           string  `json:"maxVehicleHeight"`
			MustDisplayConfirmation    bool    `json:"mustDisplayConfirmation"`
			MustPrintConfirmation      bool    `json:"mustPrintConfirmation"`
			OneTimeFee                 float64 `json:"oneTimeFee"`
			OverrideArriveText         string  `json:"overrideArriveText"`
			OverrideDepartText         string  `json:"overrideDepartText"`
			Postal                     string  `json:"postal"`
			PrintSetting               int     `json:"printSetting"`
			CustomerReceiptPrintMode   int     `json:"customerReceiptPrintMode"`
			MerchantReceiptPrintMode   int     `json:"merchantReceiptPrintMode"`
			VoidReceiptPrintMode       int     `json:"voidReceiptPrintMode"`
			RequiresLicensePlate       bool    `json:"requiresLicensePlate"`
			RequiresVehicleDescription bool    `json:"requiresVehicleDescription"`
			ShowPriceOnConfirmation    bool    `json:"showPriceOnConfirmation"`
			State                      string  `json:"state"`
			ThumbnailURL               string  `json:"thumbnailUrl"`
			TowDescription             string  `json:"towDescription"`
			Amenities                  []struct {
				ID          int    `json:"id"`
				Value       int    `json:"value"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"amenities"`
			Gates            []interface{} `json:"gates"`
			HoursOfOperation []interface{} `json:"hoursOfOperation"`
			Images           []struct {
				ID               int         `json:"id"`
				ImageDescription interface{} `json:"imageDescription"`
				ImagePath        string      `json:"imagePath"`
				ImagePathMedium  string      `json:"imagePathMedium"`
				ImagePathSmall   string      `json:"imagePathSmall"`
				IsDefault        bool        `json:"isDefault"`
			} `json:"images"`
			AvailableSpaces        int         `json:"availableSpaces"`
			IDCompany              int         `json:"idCompany"`
			IDMarket               int         `json:"idMarket"`
			Description            string      `json:"description"`
			Directions             string      `json:"directions"`
			IgnoresDaylightSavings bool        `json:"ignoresDaylightSavings"`
			Instructions           string      `json:"instructions"`
			IsGarage               bool        `json:"isGarage"`
			IsMonthly              bool        `json:"isMonthly"`
			IsNonRefundable        bool        `json:"isNonRefundable"`
			PublicName             string      `json:"publicName"`
			TimeZoneOffsetFromUtc  int         `json:"timeZoneOffsetFromUtc"`
			CityStateAndPostal     string      `json:"cityStateAndPostal"`
			DisplayAddress         string      `json:"displayAddress"`
			DisplayName            string      `json:"displayName"`
			DistanceString         string      `json:"distanceString"`
			IsWebSpecial           bool        `json:"isWebSpecial"`
			Entrances              interface{} `json:"entrances"`
			PrimaryImage           struct {
				ID               int         `json:"id"`
				ImageDescription interface{} `json:"imageDescription"`
				ImagePath        string      `json:"imagePath"`
				ImagePathMedium  string      `json:"imagePathMedium"`
				ImagePathSmall   string      `json:"imagePathSmall"`
				IsDefault        bool        `json:"isDefault"`
			} `json:"primaryImage"`
		} `json:"locations"`
		Promotions       []interface{} `json:"promotions"`
		ResultsCount     int           `json:"resultsCount"`
		HasLiveLocations bool          `json:"hasLiveLocations"`
		Search           struct {
			AmenityIds    interface{} `json:"amenityIds"`
			AmenityNames  interface{} `json:"amenityNames"`
			Address1      interface{} `json:"address1"`
			City          interface{} `json:"city"`
			CompanyIds    interface{} `json:"companyIds"`
			CompanyNames  interface{} `json:"companyNames"`
			Country       interface{} `json:"country"`
			Daily         bool        `json:"daily"`
			DisplayText   string      `json:"displayText"`
			Duration      string      `json:"duration"`
			DurationShort string      `json:"durationShort"`
			EndDate       string      `json:"endDate"`
			EndTime       string      `json:"endTime"`
			EventID       interface{} `json:"eventId"`
			Garage        bool        `json:"garage"`
			Latitude      float64     `json:"latitude"`
			LocationID    interface{} `json:"locationId"`
			LocationsLive []struct {
				City      string  `json:"city"`
				State     string  `json:"state"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Distance  float64 `json:"distance"`
				IsLive    bool    `json:"isLive"`
			} `json:"locationsLive"`
			Longitude     float64     `json:"longitude"`
			Miles         float64     `json:"miles"`
			Km            interface{} `json:"km"`
			Monthly       bool        `json:"monthly"`
			OnlyAvailable bool        `json:"onlyAvailable"`
			Peer          bool        `json:"peer"`
			Postal        interface{} `json:"postal"`
			Query         string      `json:"query"`
			SeoPageURL    interface{} `json:"seoPageUrl"`
			StartDate     string      `json:"startDate"`
			StartTime     string      `json:"startTime"`
			State         interface{} `json:"state"`
			VenueID       interface{} `json:"venueId"`
		} `json:"search"`
		ShowSoldOutModal bool        `json:"showSoldOutModal"`
		Performance      interface{} `json:"performance"`
	} `json:"data"`
	Error   interface{} `json:"error"`
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

func GetParkingPandaLots(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	res, err := client.Get(urlParkingPanda)
	if err != nil {
		fmt.Errorf("get: %v", err)
	}

	var listing ParkingPandaList

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&listing)
	if err != nil {
		fmt.Errorf("decode: %v", err)
	}

	io.WriteString(w, "<div>-----------------------------</div>")
	for i := 0; i < len(listing.Data.Locations); i++ {
		io.WriteString(w, fmt.Sprintf("<div>%v has a price of %v</div>",
			listing.Data.Locations[i].DisplayAddress,
			listing.Data.Locations[i].Price))
	}
}

// fetch the prefered readable title
func fetchTitle(displayAddress, displayName string) string {
	if displayName == "" {
		return clean(displayAddress)
	} else {
		return clean(displayName)
	}
}

func fetchParkingPandaLots(c appengine.Context) ([]Lot, error) {
	client := urlfetch.Client(c)
	res, err := client.Get(urlParkingPanda)
	if err != nil {
		return nil, fmt.Errorf("get: %v", err)
	}

	var listing ParkingPandaList

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&listing)
	if err != nil {
		return nil, fmt.Errorf("decode: %v", err)
	}
	Result := make([]Lot, len(listing.Data.Locations))
	for i := 0; i < len(listing.Data.Locations); i++ {
		Result[i] = Lot{
			// work to clean(fetchTitle(listing.Data.Locations[i])),
			Title: fetchTitle(listing.Data.Locations[i].DisplayAddress,
				listing.Data.Locations[i].DisplayName),
			Key:            strings.Join(strings.Fields(clean(listing.Data.Locations[i].DisplayAddress)), ""),
			ImageURL:       listing.Data.Locations[i].PrimaryImage.ImagePath,
			Price:          listing.Data.Locations[i].Price,
			Lat:            listing.Data.Locations[i].Latitude,
			Lng:            listing.Data.Locations[i].Longitude,
			DisplayAddress: listing.Data.Locations[i].DisplayAddress,
			Descr:          listing.Data.Locations[i].Description,
			Direc:          listing.Data.Locations[i].Directions,
			Source:         "www.ParkingPanda.com",
		}
	}

	return Result, nil
}

//func storeAllParkingPandaLots(c appengine.Context) {}
