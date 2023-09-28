// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"

	"github.com/qifengzhang007/gooxml"
)

type CT_SSup struct {
	SSupPr *CT_SSupPr
	E      *CT_OMathArg
	Sup    *CT_OMathArg
}

func NewCT_SSup() *CT_SSup {
	ret := &CT_SSup{}
	ret.E = NewCT_OMathArg()
	ret.Sup = NewCT_OMathArg()
	return ret
}

func (m *CT_SSup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SSupPr != nil {
		sesSupPr := xml.StartElement{Name: xml.Name{Local: "m:sSupPr"}}
		e.EncodeElement(m.SSupPr, sesSupPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	sesup := xml.StartElement{Name: xml.Name{Local: "m:sup"}}
	e.EncodeElement(m.Sup, sesup)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SSup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
	m.Sup = NewCT_OMathArg()
lCT_SSup:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSupPr"}:
				m.SSupPr = NewCT_SSupPr()
				if err := d.DecodeElement(m.SSupPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"}:
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sup"}:
				if err := d.DecodeElement(m.Sup, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_SSup %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SSup
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SSup and its children
func (m *CT_SSup) Validate() error {
	return m.ValidateWithPath("CT_SSup")
}

// ValidateWithPath validates the CT_SSup and its children, prefixing error messages with path
func (m *CT_SSup) ValidateWithPath(path string) error {
	if m.SSupPr != nil {
		if err := m.SSupPr.ValidateWithPath(path + "/SSupPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	if err := m.Sup.ValidateWithPath(path + "/Sup"); err != nil {
		return err
	}
	return nil
}
