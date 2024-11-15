/*
Package filter GENERATED BY gengo:runtimedoc 
DON'T EDIT THIS FILE
*/
package filter

// nolint:deadcode,unused
func runtimeDoc(v any, prefix string, names ...string) ([]string, bool) {
	if c, ok := v.(interface {
		RuntimeDoc(names ...string) ([]string, bool)
	}); ok {
		doc, ok := c.RuntimeDoc(names...)
		if ok {
			if prefix != "" && len(doc) > 0 {
				doc[0] = prefix + doc[0]
				return doc, true
			}

			return doc, true
		}
	}
	return nil, false
}

func (v Composed) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Filters":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v ErrInvalidFilter) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Filter":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v ErrInvalidFilterOp) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Op":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v ErrUnsupportedQLField) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "FieldName":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (Op) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}
