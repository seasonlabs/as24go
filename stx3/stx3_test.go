package stx3

import (
	"reflect"
	"testing"
)

const testInput = `
<?xml version="1.0" encoding="UTF-8"?>
<stx3>
  <vehicle_data>
    <vehicles>
      <vehicle>
        <dealer_id>1234</dealer_id>
        <ownersvehicle_id>123456AA</ownersvehicle_id>
        <status>active</status>
        <visibility>AS24_Público</visibility>
        <type>Car</type>
        <category>Used</category>
        <body>Sedán</body>
        <brand>Make</brand>
        <model>Model</model>
        <version>Version 1.0</version>
        <body_color>blanco</body_color>
        <body_colorgroup>Blanco</body_colorgroup>
        <interior_color/>
        <covering/>
        <doors></doors>
        <gear_type>Automático</gear_type>
        <fuel_type>Diésel</fuel_type>
        <fuel_types>
          <primary_fuel_type/>
        </fuel_types>
        <transmission/>
        <capacity>1995</capacity>
        <kilowatt>105</kilowatt>
        <cylinder>4</cylinder>
        <consumption>
          <liquid>
            <urban>7.7</urban>
            <extra_urban>5.4</extra_urban>
            <combined>6.2</combined>
          </liquid>
          <electric>
            <urban/>
            <extra_urban/>
            <combined/>
          </electric>
        </consumption>
        <emission>
          <co2_liquid>164</co2_liquid>
          <co2_electric/>
          <efficiency_class/>
        </emission>
        <mileage>22084</mileage>
        <service>
          <last_technical_service/>
          <last_change_cam_belt/>
        </service>
        <initial_registration>2011-11-01</initial_registration>
        <prices>
          <price>
            <type>public</type>
            <vat_type>1</vat_type>
            <currency>Euro</currency>
            <negotiable>0</negotiable>
            <value>27900</value>
          </price>
          <price>
            <type>dealer</type>
            <vat_type>1</vat_type>
            <currency>Euro</currency>
            <negotiable/>
            <value/>
          </price>
          <price>
            <type>former_new</type>
            <vat_type>1</vat_type>
            <currency>Euro</currency>
            <negotiable/>
            <value>39104</value>
          </price>
          <price>
            <type>public_euro</type>
            <vat_type>1</vat_type>
            <currency>EUR</currency>
            <negotiable/>
            <value>27900</value>
          </price>
          <price>
            <type>dealer_euro</type>
            <vat_type>1</vat_type>
            <currency>EUR</currency>
            <negotiable/>
            <value/>
          </price>
        </prices>
        <leasing>
          <initial_payment/>
          <declining_balance/>
          <unused_mileage_price/>
          <excess_mileage_price/>
          <contract_value/>
        </leasing>
        <warranty_duration>24</warranty_duration>
        <delivery>
          <delivery_time/>
          <delivery_date/>
        </delivery>
        <equipments>
          <equipment>
            <text>Aire Acondicionado</text>
          </equipment>
          <equipment>
            <text>Dirección asistida</text>
          </equipment>
          <equipment>
            <text>Faros antiniebla</text>
          </equipment>
          <equipment>
            <text>Climatizador</text>
          </equipment>
          <equipment>
            <text>Garantia</text>
          </equipment>
          <equipment>
            <text>Control de velocidad</text>
          </equipment>
          <equipment>
            <text>Sistema de ayuda de aparcamiento</text>
          </equipment>
        </equipments>
        <seals>
          <seal>
            <name>BMW Premium Selection</name>
            <status>1</status>
          </seal>
        </seals>
        <media>
          <images>
            <image>
              <local>0221316922_01.jpg</local>
              <position>1</position>
            </image>
            <image>
              <local>0221316922_02.jpg</local>
              <position>2</position>
            </image>
            <image>
              <local>0221316922_03.jpg</local>
              <position>3</position>
            </image>
          </images>
          <videos>
            <video>
              <uri/>
            </video>
          </videos>
        </media>
        <accident_free>1</accident_free>
        <alloy_wheels_size/>
      </vehicle>
    </vehicles>
  </vehicle_data>
</stx3>`

func TestDecoding(t *testing.T) {
	decoded, err := NewStx3([]byte(testInput))
	if err != nil {
		t.Fatalf("Not parsed: %v", err)
	}

	if reflect.TypeOf(decoded.VehicleData.Vehicles.Vehicle[0].Doors).Name() == "int" {
		t.Log("Doors is int and 5")
	}
}
