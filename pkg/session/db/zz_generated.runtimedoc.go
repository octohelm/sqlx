/*
Package db GENERATED BY gengo:runtimedoc 
DON'T EDIT THIS FILE
*/
package db

// nolint:deadcode,unused
func runtimeDoc(v any, names ...string) ([]string, bool) {
	if c, ok := v.(interface {
		RuntimeDoc(names ...string) ([]string, bool)
	}); ok {
		return c.RuntimeDoc(names...)
	}
	return nil, false
}

func (v Database) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Endpoint":
			return []string{
				"Endpoint of database",
			}, true
		case "NameOverwrite":
			return []string{
				"Overwrite dbname when not empty",
			}, true
		case "UsernameOverwrite":
			return []string{
				"Overwrite username when not empty",
			}, true
		case "PasswordOverwrite":
			return []string{
				"Overwrite password when not empty",
			}, true
		case "EnableMigrate":
			return []string{
				"auto migrate before run",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v Endpoint) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Scheme":
			return []string{}, true
		case "Hostname":
			return []string{}, true
		case "Port":
			return []string{}, true
		case "Path":
			return []string{}, true
		case "Username":
			return []string{}, true
		case "Password":
			return []string{}, true
		case "Extra":
			return []string{}, true

		}

		return nil, false
	}
	return []string{
		"openapi:strfmt endpoint",
	}, true
}
