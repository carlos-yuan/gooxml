// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/qifengzhang007/gooxml/schema/soo/ofc/math"
)

func TestCT_CharConstructor(t *testing.T) {
	v := math.NewCT_Char()
	if v == nil {
		t.Errorf("math.NewCT_Char must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Char should validate: %s", err)
	}
}

func TestCT_CharMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Char()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Char()
	xml.Unmarshal(buf, v2)
}
