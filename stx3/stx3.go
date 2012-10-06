package stx3

import (
	"encoding/xml"
)

type Stx3 struct {
	XMLName     xml.Name `xml:"stx3"`
	VehicleData struct {
		XMLName  xml.Name `xml:"vehicle_data"`
		Vehicles struct {
			XMLName xml.Name  `xml:"vehicles"`
			Vehicle []Vehicle `xml:"vehicle"`
		}
	}
}

type Vehicle struct {
	XMLName         xml.Name `xml:"vehicle"`
	DealerId        string   `xml:"dealer_id"`
	OwnersvehicleId string   `xml:"ownersvehicle_id"`
	Type            string   `xml:"type"`
	Category        string   `xml:"category"`
	Body            string   `xml:"body"`
	Brand           string   `xml:"brand"`
	Model           string   `xml:"model"`
	BodyColor       string   `xml:"body_color"`
	BodyColorgroup  string   `xml:"body_colorgroup"`
	Doors           string   `xml:"doors"`
	GearType        string   `xml:"gear_type"`
	Gears           string   `xml:"gears"`
	FuelType        string   `xml:"fuel_type"`
	Kilowatt        string   `xml:"kilowatt"`
	Cylinder        string   `xml:"cylinder"`
	Capacity        string   `xml:"capacity"`
	Consumption     struct {
		XMLName xml.Name `xml:"consumption"`
		Liquid  struct {
			XMLName    xml.Name `xml:"liquid"`
			Urban      string   `xml:"urban"`
			ExtraUrban string   `xml:"extra_urban"`
			Combined   string   `xml:"combined"`
		}
		Electric struct {
			XMLName    xml.Name `xml:"electric"`
			Urban      string   `xml:"urban"`
			ExtraUrban string   `xml:"extra_urban"`
			Combined   string   `xml:"combined"`
		}
	}
	Mileage             string `xml:"mileage"`
	InitialRegistration string `xml:"initial_registration"`
	Notes               string `xml:"notes"`
	LicenseNumber       string `xml:"license_number"`
	Prices              struct {
		XMLName xml.Name `xml:"prices"`
		Price   []Price  `xml:"price"`
	}
	Media struct {
		XMLName xml.Name `xml:"media"`
		Images  struct {
			XMLName xml.Name `xml:"images"`
			Image   []Image  `xml:"image"`
		}
	}
	KerbWeight string `xml:"kerb_weight"`
	Seats      string `xml:"seats"`
	Equipments struct {
		XMLName   xml.Name    `xml:"equipments"`
		Equipment []Equipment `xml:"equipment"`
	}
}

type Price struct {
	XMLName  xml.Name `xml:"price"`
	Type     string   `xml:"type"`
	Currency string   `xml:"currency"`
	Value    string   `xml:"value"`
}

type Equipment struct {
	XMLName xml.Name `xml:"equipment"`
	Text    string   `xml:"text"`
}

type Image struct {
	XMLName xml.Name `xml:"image"`
	Local   string   `xml:"local"`
}

func (stx3 *Stx3) Marshal() ([]byte, error) {
	result, err := xml.MarshalIndent(stx3, "  ", "    ")

	fullXml := xml.Header + string(result)

	return []byte(fullXml), err
}
