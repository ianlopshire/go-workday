package raas

import (
	"errors"
	"strings"

	"github.com/ianlopshire/go-workday/pkg/wd"
)

// ReferenceParam constructs a report parameter for a given parameter key and reference object.
func ReferenceParam(key string, ref wd.Reference) (k, v string) {
	t, id := ref.Any()
	return key + "!" + t, id
}

// ReferenceListParam constructs a report parameter for a given parameter key and a slice of reference objects.
//
// The slice of reference objects must share at least 1 id type. If no id type is shared, an error is returned.
func ReferenceListParam(key string, refs []wd.Reference) (k, v string, err error) {
	if refs == nil || len(refs) == 0 {
		k, v := ReferenceParam(key, nil)
		return k, v, nil
	}

	if len(refs) == 1 {
		k, v := ReferenceParam(key, refs[0])
		return k, v, nil
	}

	candidates := make([]string, 0, len(refs[0]))
	for t := range refs[0] {
		if t == wd.DescriptorKey {
			continue
		}
		candidates = append(candidates, t)
	}

	z := refs[0]
	refs = refs[1:]

	for _, ref := range refs {
		nc := make([]string, 0, len(candidates))
		for _, c := range candidates {
			if _, ok := ref[c]; ok {
				nc = append(nc, c)
			}

			if len(nc) < 1 {
				return "", "", errors.New("all references must share at least one id type")
			}
		}
		candidates = nc
	}

	t := candidates[0]
	ids := make([]string, 0, len(refs)+1)
	ids = append(ids, z[t])
	for _, ref := range refs {
		ids = append(ids, ref[t])
	}

	return key + "!" + t, strings.Join(ids, "!"), nil
}
