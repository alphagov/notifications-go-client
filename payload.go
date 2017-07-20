package notify

import "net/url"

// Payload that will be send with different set of requests by the client.
type Payload struct {
	EmailAddress    string            `json:"email_address"`
	Letter          string            `json:"letter"` // TODO Establish what this actually is and prepare for usage.
	Personalisation map[string]string `json:"personalisation"`
	PhoneNumber     string            `json:"phone_number"`
	Reference       string            `json:"reference"`
	TemplateID      string            `json:"template_id"`
}

func (p *Payload) addIfNotEmpty(m *url.Values, key, value string) {
	if value != "" {
		m.Add(key, value)
	}
}

// NewPayload is a function that takes different parameters and initialises the
// Payload struct, to be used in the calls.
func NewPayload(service, recipient, templateID string, personalisation templateData, reference string) *Payload {
	p := Payload{
		Personalisation: personalisation,
		Reference:       reference,
		TemplateID:      templateID,
	}

	switch service {
	case "sms":
		p.PhoneNumber = recipient
	case "email":
		p.EmailAddress = recipient
	case "letter":
		p.Letter = recipient
	}

	return &p
}
