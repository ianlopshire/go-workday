package wd

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestReference_String(t *testing.T) {
	for _, tt := range []struct {
		name string
		ref  Reference
		want string
	}{
		{
			name: "Empty",
			ref:  Reference{},
			want: "",
		},
		{
			name: "Descriptor and IDs",
			ref: Reference{
				DescriptorKey: "John Doe (9999)",
				"Employee_ID": "9999",
				"WID":         "7ae0e6d8b46e108b602c1540015dca87",
			},
			want: "John Doe (9999) [Employee_ID:9999,WID:7ae0e6d8b46e108b602c1540015dca87]",
		},
		{
			name: "Only Descriptor",
			ref: Reference{
				DescriptorKey: "John Doe (9999)",
			},
			want: "John Doe (9999)",
		},
		{
			name: "Only IDs",
			ref: Reference{
				"Employee_ID": "9999",
				"WID":         "7ae0e6d8b46e108b602c1540015dca87",
			},
			want: "[Employee_ID:9999,WID:7ae0e6d8b46e108b602c1540015dca87]",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			have := tt.ref.String()
			if have != tt.want {
				t.Errorf("String() have %v, want %v", have, tt.want)
			}
		})
	}
}

func TestReference_Is(t *testing.T) {
	for _, tt := range []struct {
		name string
		ref1 Reference
		ref2 Reference
		want bool
	}{
		{
			name: "both empty",
			ref1: Reference{},
			ref2: Reference{},
			want: false,
		},
		{
			name: "both nil",
			ref1: nil,
			ref2: nil,
			want: false,
		},
		{
			name: "empty and nil",
			ref1: Reference{},
			ref2: nil,
			want: false,
		},
		{
			name: "ref1 is ref2",
			ref1: Reference{"Employee_ID": "9999", "OtherID1": "111"},
			ref2: Reference{"Employee_ID": "9999", "OtherID2": "222"},
			want: true,
		},
		{
			name: "ref1 is not ref2",
			ref1: Reference{"Employee_ID": "8888", "OtherID1": "111"},
			ref2: Reference{"Employee_ID": "9999", "OtherID2": "222"},
			want: false,
		},
		{
			name: "ref1 is not ref2 (same descriptor)",
			ref1: Reference{"Employee_ID": "8888", "OtherID1": "111", DescriptorKey: "DescriptorKey"},
			ref2: Reference{"Employee_ID": "9999", "OtherID2": "222", DescriptorKey: "DescriptorKey"},
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			have := tt.ref1.Is(tt.ref2)
			if have != tt.want {
				t.Errorf("ref1.Is(ref2) have %v, want %v", have, tt.want)
			}

			have = tt.ref2.Is(tt.ref1)
			if have != tt.want {
				t.Errorf("ref2.Is(ref1) have %v, want %v", have, tt.want)
			}
		})
	}
}

func TestReference_UnmarshalXML(t *testing.T) {
	for _, tt := range []struct {
		name      string
		data      []byte
		want      Reference
		shouldErr bool
	}{
		{
			name: "id with descriptor",
			data: []byte(`
				<Reference Descriptor="John Doe (9999)">
    				<ID type="Employee_ID">9999</ID>
				</Reference>
			`),
			want: Reference{
				DescriptorKey: "John Doe (9999)",
				"Employee_ID": "9999",
			},
			shouldErr: false,
		},
		{
			name: "no ids no descriptor",
			data: []byte(`
				<Reference>
				</Reference>
			`),
			want:      Reference{},
			shouldErr: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			have := Reference{}
			err := xml.Unmarshal(tt.data, &have)

			if (err != nil) != tt.shouldErr {
				t.Errorf("unexpected error value %v, want %v (%v)", err != nil, tt.shouldErr, err)
			}

			if !reflect.DeepEqual(have, tt.want) {
				t.Errorf("UnmarshalXML() have %v, want %v", have, tt.want)
			}
		})
	}
}

func TestReference_MarshalXML(t *testing.T) {
	for _, tt := range []struct {
		name      string
		reference Reference
		want      []byte
		shouldErr bool
	}{
		{
			name: "ids with descriptor",
			reference: Reference{
				DescriptorKey: "John Doe (9999)",
				"WID":         "7ae0e6d8b46e108b602c1540015dca87",
				"Employee_ID": "9999",
			},
			want:      []byte(`<Reference Descriptor="John Doe (9999)"><ID type="WID">7ae0e6d8b46e108b602c1540015dca87</ID><ID type="Employee_ID">9999</ID></Reference>`),
			shouldErr: false,
		},
		{
			name:      "no ids no descriptor",
			reference: Reference{},
			want:      []byte(`<Reference></Reference>`),
			shouldErr: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			have, err := xml.Marshal(tt.reference)

			if (err != nil) != tt.shouldErr {
				t.Errorf("unexpected error value %v, want %v (%v)", err != nil, tt.shouldErr, err)
			}

			have = bytes.TrimSpace(have)
			if !bytes.Equal(have, tt.want) {
				t.Errorf("MarshalXML() have %s, want %s", have, tt.want)
			}
		})
	}
}
