package schema

// Field represents a field in the schema
type Field struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description"`
	Ref         string       `json:"ref,omitempty"`
	Title       string       `json:"title,omitempty"`
	Required    *bool        `json:"required,omitempty"`
	Type        string       `json:"type,omitempty"`
	Default     *interface{} `json:"default,omitempty"`

	Properties map[string]Field `json:"properties,omitempty"`
}

// Map returns a map version of the field for gojsonschema
func (f *Field) Map() map[string]interface{} {
	out := map[string]interface{}{}
	includeStringIfSet(out, "id", f.ID)
	includeStringIfSet(out, "name", f.Name)
	includeStringIfSet(out, "description", f.Description)
	includeStringIfSet(out, "ref", f.Ref)
	includeStringIfSet(out, "title", f.Title)
	includeStringIfSet(out, "type", f.Type)

	if f.Required != nil {
		out["required"] = *f.Required
	}

	if f.Default != nil {
		out["default"] = *f.Default
	}

	return out
}

func includeStringIfSet(out map[string]interface{}, name string, value string) {
	if value != "" {
		out[name] = value
	}
}
