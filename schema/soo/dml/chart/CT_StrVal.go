// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/qifengzhang007/gooxml"
)

type CT_StrVal struct {
	IdxAttr uint32
	V       string
}

func NewCT_StrVal() *CT_StrVal {
	ret := &CT_StrVal{}
	return ret
}

func (m *CT_StrVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idx"},
		Value: fmt.Sprintf("%v", m.IdxAttr)})
	e.EncodeToken(start)
	sev := xml.StartElement{Name: xml.Name{Local: "c:v"}}
	gooxml.AddPreserveSpaceAttr(&sev, m.V)
	e.EncodeElement(m.V, sev)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_StrVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "idx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdxAttr = uint32(parsed)
			continue
		}
	}
lCT_StrVal:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "v"}:
				if err := d.DecodeElement(&m.V, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_StrVal %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_StrVal
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_StrVal and its children
func (m *CT_StrVal) Validate() error {
	return m.ValidateWithPath("CT_StrVal")
}

// ValidateWithPath validates the CT_StrVal and its children, prefixing error messages with path
func (m *CT_StrVal) ValidateWithPath(path string) error {
	return nil
}
