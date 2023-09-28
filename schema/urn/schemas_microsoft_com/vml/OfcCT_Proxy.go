// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/qifengzhang007/gooxml/schema/soo/ofc/sharedTypes"
)

type OfcCT_Proxy struct {
	StartAttr      sharedTypes.ST_TrueFalseBlank
	EndAttr        sharedTypes.ST_TrueFalseBlank
	IdrefAttr      *string
	ConnectlocAttr *int32
}

func NewOfcCT_Proxy() *OfcCT_Proxy {
	ret := &OfcCT_Proxy{}
	return ret
}

func (m *OfcCT_Proxy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.StartAttr != sharedTypes.ST_TrueFalseBlankUnset {
		attr, err := m.StartAttr.MarshalXMLAttr(xml.Name{Local: "start"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.EndAttr != sharedTypes.ST_TrueFalseBlankUnset {
		attr, err := m.EndAttr.MarshalXMLAttr(xml.Name{Local: "end"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idref"},
			Value: fmt.Sprintf("%v", *m.IdrefAttr)})
	}
	if m.ConnectlocAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "connectloc"},
			Value: fmt.Sprintf("%v", *m.ConnectlocAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Proxy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "start" {
			m.StartAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "end" {
			m.EndAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "idref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdrefAttr = &parsed
			continue
		}
		if attr.Name.Local == "connectloc" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ConnectlocAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_Proxy: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_Proxy and its children
func (m *OfcCT_Proxy) Validate() error {
	return m.ValidateWithPath("OfcCT_Proxy")
}

// ValidateWithPath validates the OfcCT_Proxy and its children, prefixing error messages with path
func (m *OfcCT_Proxy) ValidateWithPath(path string) error {
	if err := m.StartAttr.ValidateWithPath(path + "/StartAttr"); err != nil {
		return err
	}
	if err := m.EndAttr.ValidateWithPath(path + "/EndAttr"); err != nil {
		return err
	}
	return nil
}
