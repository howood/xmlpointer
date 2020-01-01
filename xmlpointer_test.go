package xmlpointer

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/howood/xmlpointer/internal/parser"
)

type xmlTestData struct {
	Key          string
	CheckDataXML string
	CheckData    interface{}
	ResultHasErr bool
}

var xmlDataTest = `
<?xml version="1.0" encoding="UTF-8"?>
<ear:ObjDirectory xmlns:ear="http://moon.deo.net/ear" xmlns:aop="http://moon.deo.net/aop" xmlns:gml="http://www.opengis.net/gml" xmlns:schemaLocation="http://moon.deo.net/ear" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" version="1.2.1">
    <gml:metaDataProperty>
        <aop:ObjDirectoryMetaData>
            <aop:offnadirAngle>34.30</aop:offnadirAngle>
            <aop:orbitDirection>ASCENDING</aop:orbitDirection>
        </aop:ObjDirectoryMetaData>
    </gml:metaDataProperty>
    <gml:validTime>
        <gml:TimePeriod>
            <gml:beginPoint>2010-08-17T13:23:50Z</gml:beginPoint>
            <gml:endPoint>2010-08-17T13:24:07Z</gml:endPoint>
        </gml:TimePeriod>
    </gml:validTime>
    <gml:using>
        <aop:ObjDirectoryEquipment>
            <aop:platform>
                <aop:shortName>StriX</aop:shortName>
                <aop:serialIdentifier>alpha</aop:serialIdentifier>
                <aop:orbitType>LEO</aop:orbitType>
            </aop:platform>
            <aop:instrument>
                <aop:shortName>SAR</aop:shortName>
            </aop:instrument>
            <aop:sensor>
                <aop:seminorType>RADAR</aop:seminorType>
                <aop:operationalMode>Stripmap</aop:operationalMode>
            </aop:sensor>
            <aop:aquariumParameters>
                <ear:Aquarium>
                    <ear:polarisationChannels>VV</ear:polarisationChannels>
                    <ear:antennaLookDirection>RIGHT</ear:antennaLookDirection>
                </ear:Aquarium>
            </aop:aquariumParameters>
        </aop:ObjDirectoryEquipment>
	</gml:using>
	<arrayitem>
		<item>aaa</item>
		<item>bbb</item>
		<item>ccc</item>
		<item>ddd</item>
	</arrayitem>
</ear:ObjDirectory>
`

var timecheck, _ = time.Parse(time.RFC3339, "2010-08-17T13:23:50Z")

var xmlDataCheck = map[string]xmlTestData{
	"test1": {
		Key:          ".",
		CheckDataXML: xmlDataTest,
		ResultHasErr: false,
	},
	"test2": {
		Key: "ObjDirectory.metaDataProperty",
		CheckDataXML: `
<aop:ObjDirectoryMetaData>
	<aop:offnadirAngle>34.30</aop:offnadirAngle>
	<aop:orbitDirection>ASCENDING</aop:orbitDirection>
</aop:ObjDirectoryMetaData>
`,
		ResultHasErr: false,
	},
	"test3": {
		Key: "ObjDirectory.validTime",
		CheckDataXML: `
<gml:TimePeriod>
	<gml:beginPoint>2010-08-17T13:23:50Z</gml:beginPoint>
	<gml:endPoint>2010-08-17T13:24:07Z</gml:endPoint>
</gml:TimePeriod>

`,
		ResultHasErr: false,
	},
	"test4": {
		Key: "ObjDirectory.using",
		CheckDataXML: `
	<aop:ObjDirectoryEquipment>
		<aop:platform>
			<aop:shortName>StriX</aop:shortName>
			<aop:serialIdentifier>alpha</aop:serialIdentifier>
			<aop:orbitType>LEO</aop:orbitType>
		</aop:platform>
		<aop:instrument>
			<aop:shortName>SAR</aop:shortName>
		</aop:instrument>
		<aop:sensor>
			<aop:seminorType>RADAR</aop:seminorType>
			<aop:operationalMode>Stripmap</aop:operationalMode>
		</aop:sensor>
		<aop:aquariumParameters>
			<ear:Aquarium>
				<ear:polarisationChannels>VV</ear:polarisationChannels>
				<ear:antennaLookDirection>RIGHT</ear:antennaLookDirection>
			</ear:Aquarium>
		</aop:aquariumParameters>
	</aop:ObjDirectoryEquipment>

`,
		ResultHasErr: false,
	},
	"test5": {
		Key:          "ObjDirectory.validTime.TimePeriod.beginPoint",
		CheckData:    timecheck,
		ResultHasErr: false,
	},
	"test6": {
		Key:          "ObjDirectory.validTime.beginPoint",
		CheckData:    timecheck,
		ResultHasErr: true,
	},
	"test7": {
		Key:          "ObjDirectory.validTime.[*].beginPoint",
		CheckData:    timecheck,
		ResultHasErr: true,
	},
	"test8": {
		Key:          "ObjDirectory.arrayitem.item",
		CheckData:    []interface{}{"aaa", "bbb", "ccc", "ddd"},
		ResultHasErr: false,
	},
}

func Test_XMLPointer(t *testing.T) {
	var err error
	if _, err := NewXMLPointer(""); err == nil {
		t.Fatal("failed test string", err)
	} else {
		t.Logf("failed test %#v", err)
	}

	if _, err := NewXMLPointer([]byte("s")); err == nil {
		t.Fatal("failed test bytes", err)
	} else {
		t.Logf("failed test %#v", err)
	}
	_, err = NewXMLPointer([]byte(xmlDataTest))
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	xp, err := NewXMLPointer(xmlDataTest)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log(xp.Data)
	for k, v := range xmlDataCheck {
		xmldata, err := xp.Query(v.Key)
		if (err != nil) != v.ResultHasErr {
			t.Fatalf("failed test :%s %#v", k, err)
		} else {
			t.Logf("failed test :%s %#v", k, err)
		}
		if v.ResultHasErr == false {
			if v.CheckData != nil {
				switch v.CheckData.(type) {
				case time.Time:
					if reflect.DeepEqual(xmldata, v.CheckData) == false {
						t.Log(xmldata)
						t.Fatalf("not equal data %v, %v", xmldata, v.CheckData)
					}
				case []interface{}:
					if reflect.DeepEqual(xmldata, v.CheckData) == false {
						t.Log(xmldata)
						t.Log(reflect.ValueOf(xmldata).Type())
						t.Fatalf("not equal data %v, %v", xmldata, v.CheckData)
					}
				default:
					if fmt.Sprintf("%v", xmldata) != v.CheckData {
						t.Log(xmldata)
						t.Fatalf("not equal data %v, %s", xmldata, v.CheckData)
					}
				}
			} else {
				if checkEqualJSONByte(xmldata.(map[string]interface{}), []byte(v.CheckDataXML), t) == false {
					t.Log(xmldata)
					t.Fatalf("not equal data %v, %s", xmldata, v.CheckDataXML)
				}
			}
			t.Log(xmldata)
		}
	}

	t.Log("success JsonData")
}

func checkEqualJSONByte(input1 map[string]interface{}, input2 []byte, t *testing.T) bool {
	inputmap2, err := parser.ByteToXMLMap(input2)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log(inputmap2)
	return reflect.DeepEqual(input1, inputmap2)
}
