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

type CT_Constraint struct {
	OpAttr         ST_BoolOperator
	ValAttr        *float64
	FactAttr       *float64
	ExtLst         *dml.CT_OfficeArtExtensionList
	TypeAttr       ST_ConstraintType
	ForAttr        ST_ConstraintRelationship
	ForNameAttr    *string
	PtTypeAttr     ST_ElementType
	RefTypeAttr    ST_ConstraintType
	RefForAttr     ST_ConstraintRelationship
	RefForNameAttr *string
	RefPtTypeAttr  ST_ElementType
}

func NewCT_Constraint() *CT_Constraint {
	ret := &CT_Constraint{}
	return ret
}

func (m *CT_Constraint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OpAttr != ST_BoolOperatorUnset {
		attr, err := m.OpAttr.MarshalXMLAttr(xml.Name{Local: "op"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	if m.FactAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fact"},
			Value: fmt.Sprintf("%v", *m.FactAttr)})
	}
	if m.TypeAttr != ST_ConstraintTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ForAttr != ST_ConstraintRelationshipUnset {
		attr, err := m.ForAttr.MarshalXMLAttr(xml.Name{Local: "for"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ForNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "forName"},
			Value: fmt.Sprintf("%v", *m.ForNameAttr)})
	}
	if m.PtTypeAttr != ST_ElementTypeUnset {
		attr, err := m.PtTypeAttr.MarshalXMLAttr(xml.Name{Local: "ptType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RefTypeAttr != ST_ConstraintTypeUnset {
		attr, err := m.RefTypeAttr.MarshalXMLAttr(xml.Name{Local: "refType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RefForAttr != ST_ConstraintRelationshipUnset {
		attr, err := m.RefForAttr.MarshalXMLAttr(xml.Name{Local: "refFor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RefForNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refForName"},
			Value: fmt.Sprintf("%v", *m.RefForNameAttr)})
	}
	if m.RefPtTypeAttr != ST_ElementTypeUnset {
		attr, err := m.RefPtTypeAttr.MarshalXMLAttr(xml.Name{Local: "refPtType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Constraint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "op" {
			m.OpAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
		if attr.Name.Local == "fact" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.FactAttr = &parsed
			continue
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "for" {
			m.ForAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "forName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ForNameAttr = &parsed
			continue
		}
		if attr.Name.Local == "ptType" {
			m.PtTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "refType" {
			m.RefTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "refFor" {
			m.RefForAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "refForName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefForNameAttr = &parsed
			continue
		}
		if attr.Name.Local == "refPtType" {
			m.RefPtTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Constraint:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Constraint %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Constraint
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Constraint and its children
func (m *CT_Constraint) Validate() error {
	return m.ValidateWithPath("CT_Constraint")
}

// ValidateWithPath validates the CT_Constraint and its children, prefixing error messages with path
func (m *CT_Constraint) ValidateWithPath(path string) error {
	if err := m.OpAttr.ValidateWithPath(path + "/OpAttr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.ForAttr.ValidateWithPath(path + "/ForAttr"); err != nil {
		return err
	}
	if err := m.PtTypeAttr.ValidateWithPath(path + "/PtTypeAttr"); err != nil {
		return err
	}
	if err := m.RefTypeAttr.ValidateWithPath(path + "/RefTypeAttr"); err != nil {
		return err
	}
	if err := m.RefForAttr.ValidateWithPath(path + "/RefForAttr"); err != nil {
		return err
	}
	if err := m.RefPtTypeAttr.ValidateWithPath(path + "/RefPtTypeAttr"); err != nil {
		return err
	}
	return nil
}
