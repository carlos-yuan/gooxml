// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/qifengzhang007/gooxml/schema/soo/pml"
)

func TestEG_ShowTypeConstructor(t *testing.T) {
	v := pml.NewEG_ShowType()
	if v == nil {
		t.Errorf("pml.NewEG_ShowType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.EG_ShowType should validate: %s", err)
	}
}

func TestEG_ShowTypeMarshalUnmarshal(t *testing.T) {
	v := pml.NewEG_ShowType()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewEG_ShowType()
	xml.Unmarshal(buf, v2)
}
