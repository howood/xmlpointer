package parser

import (
	"testing"
)

var xmlDataTest = `
<?xml version="1.0" encoding="UTF-8"?>
<sar:ObjectDirection xmlns:sar="http://earth.esa.int/sar" xmlns:eop="http://earth.esa.int/eop" xmlns:gml="http://www.opengis.net/gml" xmlns:schemaLocation="http://earth.esa.int/sar" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" version="1.2.1">
    <gml:metaDataProperty>
        <eop:ObjectDirectionMetaData>
            <eop:offnadirAngle>34.30</eop:offnadirAngle>
            <eop:orbitDirection>ASCENDING</eop:orbitDirection>
        </eop:ObjectDirectionMetaData>
    </gml:metaDataProperty>
    <gml:validTime>
        <gml:TimePeriod>
            <gml:beginPosition>2010-08-17T13:23:50Z</gml:beginPosition>
            <gml:endPosition>2010-08-17T13:24:07Z</gml:endPosition>
        </gml:TimePeriod>
    </gml:validTime>
    <gml:using>
        <eop:ObjectDirectionEquipment>
            <eop:platform>
                <eop:shortName>StriX</eop:shortName>
                <eop:serialIdentifier>alpha</eop:serialIdentifier>
                <eop:orbitType>LEO</eop:orbitType>
            </eop:platform>
            <eop:instrument>
                <eop:shortName>SAR</eop:shortName>
            </eop:instrument>
            <eop:sensor>
                <eop:sensorType>RADAR</eop:sensorType>
                <eop:operationalMode>Stripmap</eop:operationalMode>
            </eop:sensor>
            <eop:acquisitionParameters>
                <sar:Acquisition>
                    <sar:polarisationChannels>VV</sar:polarisationChannels>
                    <sar:antennaLookDirection>RIGHT</sar:antennaLookDirection>
                </sar:Acquisition>
            </eop:acquisitionParameters>
        </eop:ObjectDirectionEquipment>
    </gml:using>
</sar:ObjectDirection>
`

func Test_XMLParser(t *testing.T) {
	xmlobj, err := ByteToXMLMap([]byte(xmlDataTest))
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log(xmlobj)

	parsedxml, err := NewXParsedXML([]byte(xmlDataTest))
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log(parsedxml.Name)
	t.Log(parsedxml.Attr)
	xmlbytes, err := parsedxml.ToXML()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log(string(xmlbytes))
	xmlbytes2, err := XMLToByte(parsedxml)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log(string(xmlbytes2))
}
