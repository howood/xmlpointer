package parser

import (
	"bytes"
	"encoding/xml"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	// marshalPrefix is prefix for indented line beginning with
	marshalPrefix = ""
	// marshalPrefix is indented according to the indentation nesting
	marshalIndent = "    "
)

var datetimeFormat = []string{
	"2006/01/02 15:04:05",
	"2006-01-02 15:04:05",
	"2006年01月02日 15時04分05秒",
	"2006-01-02T15:04:05",
	"2006-01-02T15:04:05Z07:00",
	"2006-01-02T15:04:05-07:00",
	"2006-01-02T15:04:05-0700",
	"2006-01-02T15:04:05.999999999Z07:00",
}

// ParsedXML represents Parsed XML data entity
type ParsedXML struct {
	Name          xml.Name
	Attr          []xml.Attr
	ChildNodes    []interface{}
	ChildNodesMap map[string]interface{}
}

// MarshalXML Marchal XML for ParsedXML struct
func (p *ParsedXML) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = p.Name
	start.Attr = p.Attr
	e.EncodeToken(start)
	for _, v := range p.ChildNodes {
		switch v.(type) {
		case *ParsedXML:
			child := v.(*ParsedXML)
			if err := e.Encode(child); err != nil {
				return err
			}
		case xml.CharData:
			e.EncodeToken(v.(xml.CharData))
		case xml.Comment:
			e.EncodeToken(v.(xml.Comment))
		}
	}
	e.EncodeToken(start.End())
	return nil
}

// UnmarshalXML Unmarshal XML for ParsedXML struct
func (p *ParsedXML) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	p.Name = start.Name
	p.Attr = start.Attr
	if p.ChildNodesMap == nil || reflect.DeepEqual(p.ChildNodesMap, map[string]interface{}{}) == true {
		p.ChildNodesMap = make(map[string]interface{}, 0)
	}
	for {
		token, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.EndElement:
			childmap := make(map[string]interface{}, 0)
			if v, ok := p.ChildNodesMap[start.Name.Local]; ok {
				childmap[start.Name.Local] = v
			} else {
				childmap[start.Name.Local] = p.ChildNodesMap
			}
			p.ChildNodesMap = childmap
		case xml.StartElement:
			tok := token.(xml.StartElement)
			var data *ParsedXML
			if err := d.DecodeElement(&data, &tok); err != nil {
				return err
			}
			p.ChildNodes = append(p.ChildNodes, data)
			if v, ok := data.ChildNodesMap[data.Name.Local]; ok {
				p.ChildNodesMap[data.Name.Local] = v
			} else {
				p.ChildNodesMap[data.Name.Local] = data.ChildNodesMap
			}
		case xml.CharData:
			tokendata := string(p.trimData(token.(xml.CharData).Copy()))
			if tokendata != "" {
				p.ChildNodes = append(p.ChildNodes, tokendata)
				p.ChildNodesMap[start.Name.Local] = p.cast(tokendata)
			}
		case xml.Comment:
			tokendata := string(p.trimData(token.(xml.Comment).Copy()))
			if tokendata != "" {
				p.ChildNodes = append(p.ChildNodes, tokendata)
				p.ChildNodesMap[start.Name.Local] = p.cast(tokendata)
			}
		}
	}
}

// cast returns casted data
func (p *ParsedXML) cast(v string) interface{} {
	for _, format := range datetimeFormat {
		if f, err := time.Parse(format, v); err == nil {
			return f
		}
	}
	if f, err := strconv.ParseBool(v); err == nil {
		return f
	}
	if f, err := strconv.ParseFloat(v, 64); err == nil {
		return f
	}
	if f, err := strconv.ParseInt(v, 10, 64); err == nil {
		return f
	}
	return v
}

// trimData trims space and line feed
func (p *ParsedXML) trimData(v []byte) []byte {
	str := string(v)
	str = strings.TrimSpace(strings.NewReplacer(
		"\r\n", "",
		"\r", "",
		"\n", "",
	).Replace(str))
	return []byte(str)
}

// ToXML generate XML bytes for ParsedXML struct
func (p *ParsedXML) ToXML() ([]byte, error) {
	w := &bytes.Buffer{}
	w.Write([]byte(xml.Header))
	enc := xml.NewEncoder(w)
	enc.Indent(marshalPrefix, marshalIndent)
	err := enc.Encode(p)
	return w.Bytes(), err
}

// NewXParsedXML create ParsedXML pounter
func NewXParsedXML(xmlbyte []byte) (*ParsedXML, error) {
	xmldata := &ParsedXML{}
	err := xml.NewDecoder(bytes.NewBuffer(xmlbyte)).Decode(&xmldata)
	return xmldata, err
}

// ByteToXMLMap is convert bytes to xml map[string]interface{}
func ByteToXMLMap(xmlbyte []byte) (map[string]interface{}, error) {
	xmldata, err := NewXParsedXML(xmlbyte)
	return xmldata.ChildNodesMap, err
}

// XMLToByte is convert json struct to bytes
func XMLToByte(xmldata interface{}) ([]byte, error) {
	return xml.MarshalIndent(xmldata, marshalPrefix, marshalIndent)
}
