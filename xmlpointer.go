package xmlpointer

import (
	"fmt"
	"strings"

	"github.com/howood/xmlpointer/internal/parser"
)

// XMLPointer represents XMLPointer entity
type XMLPointer struct {
	// Data is XML Data
	Data interface{}
}

// NewXMLPointer create XMLPointer pointer from []byte / string / interface{}
func NewXMLPointer(inputdata interface{}) (*XMLPointer, error) {
	switch converteddata := inputdata.(type) {
	case []byte:
		data, err := parser.ByteToXMLMap(converteddata)
		if err != nil {
			return nil, err
		}
		return &XMLPointer{data}, nil
	case string:
		data, err := parser.ByteToXMLMap([]byte(converteddata))
		if err != nil {
			return nil, err
		}
		return &XMLPointer{data}, nil
	case map[string]interface{}:
		return &XMLPointer{Data: converteddata}, nil
	default:
		return &XMLPointer{Data: inputdata}, nil
	}
}

// Query is extract data from XMLData with item key
func (xp *XMLPointer) Query(key string) (interface{}, error) {
	if key == "." {
		return xp.Data, nil
	}
	keyList := strings.Split(key, ".")
	var context interface{} = xp.Data
	for _, xmlKey := range keyList {
		var err error
		context, err = xp.searchKey(context, xmlKey)
		if err != nil {
			return nil, err
		}
	}
	return context, nil
}

// searchKey returns data from xmldata with xmlKey
func (xp *XMLPointer) searchKey(xmldata interface{}, xmlKey string) (interface{}, error) {
	if xmlKey == "[*]" {
		switch v := xmldata.(type) {
		case []interface{}:
			xmldata = v
		default:
			return nil, fmt.Errorf("not array: %s", xmlKey)
		}
	}
	switch v := xmldata.(type) {
	case map[string]interface{}:
		if val, ok := v[xmlKey]; ok {
			return val, nil
		} else {
			return nil, fmt.Errorf("not exist: %s", xmlKey)
		}
	case []interface{}:
		ret := make([]interface{}, 0)
		for _, vv := range v {
			vvv, err := xp.searchKey(vv, xmlKey)
			if err != nil {
				return nil, err
			}
			ret = append(ret, vvv)
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("not exist: %s", xmlKey)
	}
}
