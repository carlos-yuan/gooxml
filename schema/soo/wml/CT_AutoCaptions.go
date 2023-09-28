// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/qifengzhang007/gooxml"
)

type CT_AutoCaptions struct {
	// Single Automatic Captioning Setting
	AutoCaption []*CT_AutoCaption
}

func NewCT_AutoCaptions() *CT_AutoCaptions {
	ret := &CT_AutoCaptions{}
	return ret
}

func (m *CT_AutoCaptions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seautoCaption := xml.StartElement{Name: xml.Name{Local: "w:autoCaption"}}
	for _, c := range m.AutoCaption {
		e.EncodeElement(c, seautoCaption)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AutoCaptions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AutoCaptions:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "autoCaption"}:
				tmp := NewCT_AutoCaption()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AutoCaption = append(m.AutoCaption, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_AutoCaptions %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AutoCaptions
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AutoCaptions and its children
func (m *CT_AutoCaptions) Validate() error {
	return m.ValidateWithPath("CT_AutoCaptions")
}

// ValidateWithPath validates the CT_AutoCaptions and its children, prefixing error messages with path
func (m *CT_AutoCaptions) ValidateWithPath(path string) error {
	for i, v := range m.AutoCaption {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AutoCaption[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
