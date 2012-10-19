package stx3

import (
	"encoding/xml"
)

type Stx3 struct {
	XMLName     xml.Name `xml:"stx3"`
	VehicleData struct {
		Vehicles struct {
			Vehicle []Vehicle `xml:"vehicle"`
		} `xml:"vehicles"`
	} `xml:"vehicle_data"`
}

type Vehicle struct {
	DealerId        string `xml:"dealer_id" bson:"dealer_id"`
	OwnersvehicleId string `xml:"ownersvehicle_id" bson:"ownersvehicle_id"`
	Type            string `xml:"type"`
	Category        string `xml:"category"`
	Body            string `xml:"body"`
	Brand           string `xml:"brand"`
	Model           string `xml:"model"`
	Version         string `xml:"version"`
	BodyColor       string `xml:"body_color" bson:"body_color"`
	BodyColorgroup  string `xml:"body_colorgroup" bson:"body_colorgroup"`
	InteriorColor   string `xml:"interior_color" bson:"interior_color"`
	Covering        string `xml:"covering"`
	Doors           int    `xml:"doors,omitempty"`
	GearType        string `xml:"gear_type" bson:"gear_type"`
	Gears           int    `xml:"gears,omitempty"`
	FuelType        string `xml:"fuel_type" bson:"fuel_type"`
	Kilowatt        int    `xml:"kilowatt,omitempty"`
	Cylinder        int    `xml:"cylinder,omitempty"`
	Capacity        int    `xml:"capacity,omitempty"`
	Consumption     struct {
		Liquid   Consumption `xml:"liquid"`
		Electric Consumption `xml:"electric"`
	} `xml:"consumption"`
	Mileage             int    `xml:"mileage,omitempty"`
	InitialRegistration string `xml:"initial_registration" bson:"initial_registration"`
	Notes               string `xml:"notes"`
	LicenseNumber       string `xml:"license_number" bson"license_number"`
	Prices              struct {
		Price []Price `xml:"price"`
	} `xml:"prices"`
	Media struct {
		Images struct {
			Image []Image `xml:"image"`
		} `xml:"images"`
	} `xml:"media"`
	KerbWeight int `xml:"kerb_weight,omitempty" bson:"kerb_weight,omitempty"`
	Seats      int `xml:"seats,omitempty"`
	Equipments struct {
		Equipment []Equipment `xml:"equipment"`
	} `xml:"equipments"`
}

type Price struct {
	Type     string `xml:"type"`
	Currency string `xml:"currency"`
	Value    int    `xml:"value,omitempty"`
}

type Equipment struct {
	Text string `xml:"text"`
}

type Image struct {
	Local string `xml:"local"`
}

type Consumption struct {
	Urban      string `xml:"urban"`
	ExtraUrban string `xml:"extra_urban" bson:"extra_urban"`
	Combined   string `xml:"combined"`
}

func NewStx3(inputXml []byte) (result Stx3, err error) {
	if err = xml.Unmarshal(inputXml, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (stx3 *Stx3) Marshal() ([]byte, error) {
	result, err := xml.MarshalIndent(stx3, "  ", "    ")

	fullXml := xml.Header + string(result)

	return []byte(fullXml), err
}
