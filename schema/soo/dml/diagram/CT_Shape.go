// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/qifengzhang007/gooxml"
	"github.com/qifengzhang007/gooxml/schema/soo/dml"
)

type CT_Shape struct {
	RotAttr       *float64
	TypeAttr      *ST_LayoutShapeType
	BlipAttr      *string
	ZOrderOffAttr *int32
	HideGeomAttr  *bool
	LkTxEntryAttr *bool
	BlipPhldrAttr *bool
	AdjLst        *CT_AdjLst
	ExtLst        *dml.CT_OfficeArtExtensionList
}

func NewCT_Shape() *CT_Shape {
	ret := &CT_Shape{}
	return ret
}

func (m *CT_Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rot"},
			Value: fmt.Sprintf("%v", *m.RotAttr)})
	}
	if m.TypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"},
			Value: fmt.Sprintf("%v", *m.TypeAttr)})
	}
	if m.BlipAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:blip"},
			Value: fmt.Sprintf("%v", *m.BlipAttr)})
	}
	if m.ZOrderOffAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zOrderOff"},
			Value: fmt.Sprintf("%v", *m.ZOrderOffAttr)})
	}
	if m.HideGeomAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hideGeom"},
			Value: fmt.Sprintf("%d", b2i(*m.HideGeomAttr))})
	}
	if m.LkTxEntryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lkTxEntry"},
			Value: fmt.Sprintf("%d", b2i(*m.LkTxEntryAttr))})
	}
	if m.BlipPhldrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "blipPhldr"},
			Value: fmt.Sprintf("%d", b2i(*m.BlipPhldrAttr))})
	}
	e.EncodeToken(start)
	if m.AdjLst != nil {
		seadjLst := xml.StartElement{Name: xml.Name{Local: "adjLst"}}
		e.EncodeElement(m.AdjLst, seadjLst)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Shape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "blip" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BlipAttr = &parsed
			continue
		}
		if attr.Name.Local == "rot" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.RotAttr = &parsed
			continue
		}
		if attr.Name.Local == "type" {
			parsed, err := ParseUnionST_LayoutShapeType(attr.Value)
			if err != nil {
				return err
			}
			m.TypeAttr = &parsed
			continue
		}
		if attr.Name.Local == "zOrderOff" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ZOrderOffAttr = &pt
			continue
		}
		if attr.Name.Local == "hideGeom" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HideGeomAttr = &parsed
			continue
		}
		if attr.Name.Local == "lkTxEntry" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LkTxEntryAttr = &parsed
			continue
		}
		if attr.Name.Local == "blipPhldr" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BlipPhldrAttr = &parsed
			continue
		}
	}
lCT_Shape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "adjLst"}:
				m.AdjLst = NewCT_AdjLst()
				if err := d.DecodeElement(m.AdjLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Shape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Shape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Shape and its children
func (m *CT_Shape) Validate() error {
	return m.ValidateWithPath("CT_Shape")
}

// ValidateWithPath validates the CT_Shape and its children, prefixing error messages with path
func (m *CT_Shape) ValidateWithPath(path string) error {
	if m.TypeAttr != nil {
		if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
			return err
		}
	}
	if m.AdjLst != nil {
		if err := m.AdjLst.ValidateWithPath(path + "/AdjLst"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
