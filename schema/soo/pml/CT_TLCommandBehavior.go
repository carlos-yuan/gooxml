// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/qifengzhang007/gooxml"
)

type CT_TLCommandBehavior struct {
	// Command Type
	TypeAttr ST_TLCommandType
	// Command
	CmdAttr *string
	CBhvr   *CT_TLCommonBehaviorData
}

func NewCT_TLCommandBehavior() *CT_TLCommandBehavior {
	ret := &CT_TLCommandBehavior{}
	ret.CBhvr = NewCT_TLCommonBehaviorData()
	return ret
}

func (m *CT_TLCommandBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_TLCommandTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CmdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cmd"},
			Value: fmt.Sprintf("%v", *m.CmdAttr)})
	}
	e.EncodeToken(start)
	secBhvr := xml.StartElement{Name: xml.Name{Local: "p:cBhvr"}}
	e.EncodeElement(m.CBhvr, secBhvr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLCommandBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CBhvr = NewCT_TLCommonBehaviorData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "cmd" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CmdAttr = &parsed
			continue
		}
	}
lCT_TLCommandBehavior:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cBhvr"}:
				if err := d.DecodeElement(m.CBhvr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TLCommandBehavior %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLCommandBehavior
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLCommandBehavior and its children
func (m *CT_TLCommandBehavior) Validate() error {
	return m.ValidateWithPath("CT_TLCommandBehavior")
}

// ValidateWithPath validates the CT_TLCommandBehavior and its children, prefixing error messages with path
func (m *CT_TLCommandBehavior) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.CBhvr.ValidateWithPath(path + "/CBhvr"); err != nil {
		return err
	}
	return nil
}
