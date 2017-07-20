package notify

import "net/url"

// Filters for the notifications to be looked at.
type Filters struct {
	OlderThan    string `json:"older_than"`
	Reference    string `json:"reference"`
	Status       string `json:"status"`
	TemplateType string `json:"template_type"`
}

// ToURLValues will convert the struct into the url.Values.
func (f *Filters) ToURLValues() url.Values {
	m := url.Values{}

	f.addIfNotEmpty(&m, "older_than", f.OlderThan)
	f.addIfNotEmpty(&m, "reference", f.Reference)
	f.addIfNotEmpty(&m, "status", f.Status)
	f.addIfNotEmpty(&m, "template_type", f.TemplateType)

	return m
}

func (f *Filters) addIfNotEmpty(m *url.Values, key, value string) {
	if value != "" {
		m.Add(key, value)
	}
}
