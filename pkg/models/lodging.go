package models

import "time"

// Lodging : contains all needful infromation about lodging
type Lodging struct {
	ID            int          `json:"id"`
	Country       string       `json:"country"`
	City          string       `json:"city"`
	Street        string       `json:"street"`
	HouseNumber   int          `json:"house_number"`
	Type          string       `json:"type"`
	MaxPersons    int          `json:"max_persons"`
	CostPerPerson float32      `json:"cost_per_person"`
	User          *User        `json:"user"`
	Comments      []*Comment   `json:"comments"`
	Description   *Description `json:"description"`
	Statistic     *Statistic   `json:"statistic"`
}

// Comment : contains all needful infromation about comment of lodging
type Comment struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	Created  time.Time `json:"created"`
	Rating   int       `json:"rating"`
	Username string    `json:"username"`
	Lodging  *Lodging  `json:"lodging"`
	Date     time.Time `json:"date"`
}

// Statistic : contains all needful infromation about statistics of lodging
type Statistic struct {
	ID                int      `json:"id"`
	StarsSum          int      `json:"stars_sum"`
	VotedAmount       int      `json:"voted_amount"`
	RentAmount        int      `json:"rent_amount"`
	Comfort           int      `json:"comfort"`
	Cleanness         int      `json:"cleanness"`
	PriceQualityRatio int      `json:"price_quality_ratio"`
	Staff             int      `json:"staff"`
	Location          int      `json:"location"`
	FreeWifi          int      `json:"free_wifi"`
	Lodging           *Lodging `json:"lodging"`
}

// Description : contains all needful infromation about description of lodging
type Description struct {
	ID             int    `json:"id"`
	RoomsAmount    int    `json:"rooms_amount"`
	SleepingPlaces int    `json:"sleeping_places"`
	Wifi           bool   `json:"wifi"`
	BanquetPlaces  int    `json:"banquet_places"`
	SwimmingPool   bool   `json:"swimming_pool"`
	Spa            bool   `json:"spa"`
	Braizer        bool   `json:"braizer"`
	Billiard       bool   `json:"bool"`
	Hookah         bool   `json:"hookah"`
	AdditionalInfo string `json:"additional_info"`
	ParkingPlaces  int    `json:"parking_places"`
	Lodging        *Lodging
}

// BookingTime : contains all needful infromation about booking time lodging
type BookingTime struct {
	ID       int       `json:"id"`
	FromDate time.Time `json:"from_date"`
	ToDate   time.Time `json:"to_date"`
	Lodging  *Lodging  `json:"lodging"`
}
