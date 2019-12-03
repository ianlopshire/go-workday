package wd

import (
	"encoding/xml"
	"io"
	"sort"
	"strings"
)

const (
	// DescriptorKey is the key used to store the descriptor in a reference object.
	DescriptorKey = "__Descriptor"
)

// Reference is a map based representation of a Workday reference object.
type Reference map[string]string

// Find returns the first type/id pair presents in the provided type list.
//
// If no types are given or no type in the provided list exists, t and id are both returned as empty strings.
func (r Reference) Find(types ...string) (t, id string) {
	for _, t := range types {
		if v, ok := r[t]; ok {
			return t, v
		}
	}
	return "", ""
}

// Any returns the first type/id pair it finds.
//
// If no type/id pairs exist, t and id are both returned as empty strings.
//
// Reference is a map type causing this to be non-deterministic. Repeated calls are not guaranteed to return the same
// type/id pair.
func (r Reference) Any() (t, id string) {
	for t, id = range r {
		if t != DescriptorKey {
			return t, id
		}
	}
	return "", ""
}

// Is returns true if the two reference objects share any ID. If one or either is nil, false is returned.
func (r Reference) Is(ref Reference) bool {
	if r == nil || len(r) < 1 || ref == nil || len(ref) < 1 {
		return false
	}

	for t, id1 := range r {
		if t == DescriptorKey {
			continue
		}

		if id2, ok := ref[t]; ok && id1 == id2 {
			return true
		}
	}

	return false
}

// Equal returns true if the two reference objects share any ID or if both reference objects are empty.
func (r Reference) Equal(ref Reference) bool {
	return r.Is(ref) || len(ref) == 0 && len(r) == 0
}

// String builds a string representation of a Reference.
//
// The string representation follows the template: "{Reference} [{type}:{value},{type}:{value},...]"
func (r Reference) String() string {
	s := r[DescriptorKey]

	var ids []string
	for k, v := range r {
		if k != DescriptorKey {
			ids = append(ids, k+":"+v)
		}
	}

	if len(ids) > 0 {
		if s != "" {
			s += " "
		}
		sort.Strings(ids)
		s += "[" + strings.Join(ids, ",") + "]"
	}

	return s
}

// UnmarshalXML implements the xml.Unmarshaler interface. UnmarshalXML assumes the reference is in the format used by
// Workday RaaS.
//
// Example:
//
//     <Reference Descriptor="John Doe (9999)">
//         <ID type="WID">7ae0e6d8b46e108b602c1540015dca87</ID>
//         <ID type="Employee_ID">9999</ID>
//     </Reference>
//
func (r *Reference) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	ref := Reference{}

	for i := range start.Attr {
		if start.Attr[i].Name.Local == "Descriptor" {
			ref[DescriptorKey] = start.Attr[i].Value
			break
		}
	}

	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if se, ok := token.(xml.StartElement); ok {
			var value string
			err = d.DecodeElement(&value, &se)
			if err != nil {
				return err
			}

			for i := range se.Attr {
				if se.Attr[i].Name.Local == "type" {
					ref[se.Attr[i].Value] = value
					break
				}
			}
		}
	}

	*r = ref
	return nil
}

// UnmarshalXML implements the xml.Marshaler interface.
func (r Reference) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if desc, ok := r[DescriptorKey]; ok {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "Descriptor"},
			Value: desc,
		})
	}

	tokens := []xml.Token{start}
	for key, value := range r {
		if key == DescriptorKey {
			continue
		}

		t := xml.StartElement{
			Name: xml.Name{Local: "ID"},
			Attr: []xml.Attr{{
				Name:  xml.Name{Local: "type"},
				Value: key,
			}},
		}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}
	tokens = append(tokens, xml.EndElement{start.Name})

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	// flush to ensure tokens are written
	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}
