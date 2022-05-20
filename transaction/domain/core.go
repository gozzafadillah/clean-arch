package transactionDomain

import "time"

type Transaction struct {
	ID               int
	Code             string
	User_Id          int
	Total_Qty        int
	Total_Price      float64
	Shipping_Name    string
	Shipping_Package string
	Shipping_Price   float64
	Etd              string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Checkout struct {
	ID              int
	ProductID       int
	TransactionCode string
	Qty             int
	Price           float64
	Weight          float64
	Destination     string
	Courier         string
	Package         string
	Etd             string
	Shipping_Price  float64
	Status          bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CityRO struct {
	Rajaongkir struct {
		Query struct {
			Province string `json:"province"`
			ID       string `json:"id"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

type Ongkir struct {
	Rajaongkir struct {
		Query struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Weight      int    `json:"weight"`
			Courier     string `json:"courier"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		OriginDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"origin_details"`
		DestinationDetails struct {
			CityID     string `json:"city_id"`
			ProvinceID string `json:"province_id"`
			Province   string `json:"province"`
			Type       string `json:"type"`
			CityName   string `json:"city_name"`
			PostalCode string `json:"postal_code"`
		} `json:"destination_details"`
		Results []struct {
			Code  string `json:"code"`
			Name  string `json:"name"`
			Costs []struct {
				Service     string `json:"service"`
				Description string `json:"description"`
				Cost        []struct {
					Value int    `json:"value"`
					Etd   string `json:"etd"`
					Note  string `json:"note"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
}
