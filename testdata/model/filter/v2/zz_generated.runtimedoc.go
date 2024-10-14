/*
Package filter GENERATED BY gengo:runtimedoc 
DON'T EDIT THIS FILE
*/
package filter

// nolint:deadcode,unused
func runtimeDoc(v any, names ...string) ([]string, bool) {
	if c, ok := v.(interface {
		RuntimeDoc(names ...string) ([]string, bool)
	}); ok {
		return c.RuntimeDoc(names...)
	}
	return nil, false
}

func (v OrgByCreatedAt) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "CreatedAt":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v OrgByID) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v OrgByName) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Name":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v OrgUserByID) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v OrgUserByOrgID) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "OrgID":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v OrgUserByUserID) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "UserID":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v UserByAge) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Age":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v UserByCreatedAt) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "CreatedAt":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v UserByDeletedAt) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "DeletedAt":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v UserByID) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v UserByName) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Name":
			return []string{
				"按 姓名 筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v UserByNickname) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Nickname":
			return []string{
				"按  筛选",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}
