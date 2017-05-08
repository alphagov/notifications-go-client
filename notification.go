package notify

import (
	"errors"
	"time"
)

// Template may be returned as part of Notification response.
type Template struct {
	ID      int64  `json:"id"`
	URI     string `json:"uri"`
	Version int64  `json:"version"`
}

type templateData map[string]string

// Pagination of the list that's returned as part of the JSON response.
type Pagination struct {
	Current  string `json:"current"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// Notification is the object build and returned by GOV.UK Notify.
type Notification struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	Subject   string    `json:"subject"`
	Reference string    `json:"reference"`
	Email     string    `json:"email_address"`
	Phone     string    `json:"phone_number"`
	Line1     string    `json:"line_1"`
	Line2     string    `json:"line_2"`
	Line3     string    `json:"line_3"`
	Line4     string    `json:"line_4"`
	Line5     string    `json:"line_5"`
	Line6     string    `json:"line_6"`
	Postcode  string    `json:"postcode"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	Template  Template  `json:"template"`
	CreatedAt time.Time `json:"created_at"`
	SentAt    time.Time `json:"sent_at"`
}

// NotificationEntry is the struct aroung the successful response from the API
// collected upon the creation of a new Notification.
type NotificationEntry struct {
	Content   map[string]string `json:"content"`
	ID        string            `json:"id"`
	Reference string            `json:"reference"`
	Template  Template          `json:"template"`
	URI       string            `json:"uri"`
}

// NotificationList is one the responses from GOV.UK Notify.
type NotificationList struct {
	Client *Client `json:"-"`

	Notifications []Notification `json:"notifications"`
	Links         Pagination     `json:"links"`
}

// Next page of the list should be loaded in place of the old one.
func (nl *NotificationList) Next() error {
	if nl.Links.Next == "" {
		return errors.New("pagination: already on last page")
	}

	res, err := nl.Client.httpGet(nl.Links.Next, nil)
	if err != nil {
		return err
	}

	return jsonResponse(res.Body, nl)
}

// Previous page of the list should be loaded in place of the old one.
func (nl *NotificationList) Previous() error {
	if nl.Links.Previous == "" {
		return errors.New("pagination: already on first page")
	}

	res, err := nl.Client.httpGet(nl.Links.Previous, nil)
	if err != nil {
		return err
	}

	return jsonResponse(res.Body, nl)
}
