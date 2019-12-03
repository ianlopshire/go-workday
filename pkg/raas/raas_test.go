package raas

import (
	"testing"

	"github.com/ianlopshire/go-workday/pkg/wd"
)

func TestReferenceParam(t *testing.T) {
	type param struct {
		k, v string
	}

	for _, tt := range []struct {
		name      string
		key       string
		ref       wd.Reference
		wantOneOf []param
	}{
		{
			name: "one id",
			key:  "id",
			ref:  wd.Reference{"Employee_ID": "99999"},
			wantOneOf: []param{
				{k: "id!Employee_ID", v: "99999"},
			},
		},
		{
			name: "multiple id",
			key:  "id",
			ref: wd.Reference{
				"WID":         "99999999999999",
				"Employee_ID": "99999",
			},
			wantOneOf: []param{
				{k: "id!Employee_ID", v: "99999"},
				{k: "id!WID", v: "99999999999999"},
			},
		},
		{
			name: "nil ref",
			key:  "id",
			ref:  nil,
			wantOneOf: []param{
				{k: "id!", v: ""},
			},
		},
		{
			name: "empty ref",
			key:  "id",
			ref:  wd.Reference{},
			wantOneOf: []param{
				{k: "id!", v: ""},
			},
		},
		{
			name: "descriptor only",
			key:  "id",
			ref:  wd.Reference{wd.DescriptorKey: "descriptor value"},
			wantOneOf: []param{
				{k: "id!", v: ""},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			k, v := ReferenceParam(tt.key, tt.ref)

			for _, parm := range tt.wantOneOf {
				if k == parm.k && v == parm.v {
					return
				}
			}

			t.Errorf("ReferenceParam() have %+v, want one of %+v", param{k: k, v: v}, tt.wantOneOf)
		})
	}
}

func TestReferenceListParam(t *testing.T) {
	type param struct {
		k, v string
	}

	for _, tt := range []struct {
		name      string
		key       string
		refs      []wd.Reference
		shouldErr bool
		wantOneOf []param
	}{
		{
			name: "one ref",
			key:  "id",
			refs: []wd.Reference{
				{"Employee_ID": "999"},
			},
			shouldErr: false,
			wantOneOf: []param{
				{k: "id!Employee_ID", v: "999"},
			},
		},
		{
			name: "multiple refs",
			key:  "id",
			refs: []wd.Reference{
				{"Employee_ID": "999"},
				{"Employee_ID": "888"},
				{"Employee_ID": "777"},
			},
			shouldErr: false,
			wantOneOf: []param{
				{k: "id!Employee_ID", v: "999!888!777"},
			},
		},
		{
			name: "no shared id type",
			key:  "id",
			refs: []wd.Reference{
				{"Employee_ID": "999"},
				{"Worker_ID": "888"},
				{"employee_ID": "777"},
			},
			shouldErr: true,
			wantOneOf: nil,
		},
		{
			name: "multiple refs multiple ids",
			key:  "id",
			refs: []wd.Reference{
				{"Employee_ID": "999", "Worker_ID": "99", "A": "9"},
				{"Employee_ID": "888", "Worker_ID": "88", "B": "8", "D": "8"},
				{"Employee_ID": "777", "Worker_ID": "77", "C": "7", "D": "7", "E": "7"},
			},
			shouldErr: false,
			wantOneOf: []param{
				{k: "id!Employee_ID", v: "999!888!777"},
				{k: "id!Worker_ID", v: "99!88!77"},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			k, v, err := ReferenceListParam(tt.key, tt.refs)

			if (err != nil) != tt.shouldErr {
				t.Errorf("ReferenceListParam() unexpected error value (%v)", err)
			}
			if tt.shouldErr {
				return
			}

			for _, param := range tt.wantOneOf {
				if k == param.k && v == param.v {
					return
				}
			}

			t.Errorf("ReferenceParam() have %+v, want one of %+v", param{k: k, v: v}, tt.wantOneOf)
		})
	}
}
