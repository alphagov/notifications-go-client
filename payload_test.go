package notify

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Payload", func() {
	var (
		payload     *Payload
		phoneNumber string
	)

	It("should addIfNotEmpty()", func() {
		m := url.Values{}
		p := Payload{}

		p.addIfNotEmpty(&m, "success", "true")
		p.addIfNotEmpty(&m, "failure", "")

		Expect(m["success"][0]).To(Equal("true"))
		_, ok := m["failure"]
		Expect(ok).To(BeFalse())
	})

	It("should be able to run NewPayload() for sms service", func() {
		phoneNumber = "00000000000"
		templateID := "12345"
		personalisation := map[string]string{}
		reference := "123456qwerty"

		payload = NewPayload("sms", phoneNumber, templateID, personalisation, reference)

		Expect(payload.PhoneNumber).To(Equal(phoneNumber))
		Expect(payload.Reference).To(Equal(reference))
		Expect(payload.EmailAddress).To(BeEmpty())
		Expect(payload.Letter).To(BeEmpty())
	})

	It("should be able to run NewPayload() for email service", func() {
		emailAddress := "test@example.com"
		templateID := "12345"
		personalisation := map[string]string{}
		reference := "123456qwerty"

		p := NewPayload("email", emailAddress, templateID, personalisation, reference)

		Expect(p.EmailAddress).To(Equal(emailAddress))
		Expect(p.Reference).To(Equal(reference))
		Expect(p.PhoneNumber).To(BeEmpty())
		Expect(p.Letter).To(BeEmpty())
	})

	It("should be able to run NewPayload() for letter service", func() {
		letter := "xxx"
		templateID := "12345"
		personalisation := map[string]string{}
		reference := "123456qwerty"

		p := NewPayload("letter", letter, templateID, personalisation, reference)

		Expect(p.Letter).To(Equal(letter))
		Expect(p.Reference).To(Equal(reference))
		Expect(p.EmailAddress).To(BeEmpty())
		Expect(p.PhoneNumber).To(BeEmpty())
	})
})
