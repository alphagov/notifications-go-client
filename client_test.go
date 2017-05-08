package notify

import (
	"fmt"
	"net/http"
	"net/url"

	httpmock "gopkg.in/jarcoal/httpmock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	Context("creating New() client instance", func() {
		var (
			client *Client
		)

		BeforeEach(func() {
			httpmock.Activate()
		})

		AfterEach(func() {
			httpmock.DeactivateAndReset()
		})

		It("should create New() client", func() {
			var err error

			u, _ := url.Parse("https://example.com")

			config := Configuration{
				APIKey: []byte(`#5K+･ｼミew{ｦ住ｳ(跼Tﾉ(ｩ┫ﾒP.ｿﾓ燾辻G�感ﾃwb="=.!r.Oﾀﾍ奎gﾐ｣`),
				BaseURL:   u,
				ServiceID: "test",
			}

			client, err = New(config)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(client).NotTo(BeNil())
		})

		It("should create New() client with default Congifuration.BaseURL", func() {
			config := Configuration{
				APIKey: []byte(`#5K+･ｼミew{ｦ住ｳ(跼Tﾉ(ｩ┫ﾒP.ｿﾓ燾辻G�感ﾃwb="=.!r.Oﾀﾍ奎gﾐ｣`),
				ServiceID: "test",
			}

			c, err := New(config)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(c).NotTo(BeNil())
			Expect(c.Configuration.BaseURL.String()).To(Equal(BaseURLProduction))
		})

		It("should be able to buildHeaders()", func() {
			httpmock.RegisterResponder("GET", "https://example.com",
				httpmock.NewStringResponder(http.StatusOK, ``))

			req, _ := http.NewRequest("GET", "https://example.com", nil)
			err := client.buildHeaders(req)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(req.Header.Get("User-agent")).To(Equal(fmt.Sprintf("NOTIFY-API-GO-CLIENT/%s", Version)))
		})

		It("should be able to httpCall()", func() {
			httpmock.RegisterResponder("GET", "https://example.com",
				httpmock.NewStringResponder(http.StatusOK, ``))

			res, err := client.httpCall("GET", "https://example.com", nil)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})

		It("should be fail the httpCall()", func() {
			res, err := client.httpCall("GET", "%gh&%ij", nil)

			Expect(err).Should(HaveOccurred())
			Expect(res).To(BeNil())
		})

		It("should be able to handleInvalidResponse()", func() {
			httpmock.RegisterResponder("GET", "https://example.com",
				httpmock.NewStringResponder(http.StatusUnauthorized, `[{"":"","":""}]`))

			res, _ := client.httpCall("GET", "https://example.com", nil)
			err := client.handleInvalidResponse(res)

			Expect(err).Should(HaveOccurred())
		})

		It("should be able to httpGet()", func() {
			httpmock.RegisterResponder("GET", "https://example.com/test",
				httpmock.NewStringResponder(http.StatusOK, ``))

			res, err := client.httpGet("/test", nil)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})

		It("should be able to httpPost()", func() {
			httpmock.RegisterResponder("POST", "https://example.com/test",
				httpmock.NewStringResponder(http.StatusAccepted, ``))

			res, err := client.httpPost("/test", nil)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusAccepted))
		})

		It("should allow to GetNotification() by ID", func() {
			httpmock.RegisterResponder("GET", "https://example.com/v2/notifications/n0t1-1234567890",
				httpmock.NewStringResponder(http.StatusOK, `{"id":"n0t1-1234567890"}`))

			notification, err := client.GetNotification("n0t1-1234567890")

			Expect(err).ShouldNot(HaveOccurred())
			Expect(notification.ID).To(Equal("n0t1-1234567890"))
		})

		It("should fallout if the GetNotification() fails with not found", func() {
			httpmock.RegisterResponder("GET", "https://example.com/v2/notifications/n0t1-1234567890",
				httpmock.NewStringResponder(http.StatusNotFound, `[{"error": "NoResultFound","message": "No result found"}]`))

			notification, err := client.GetNotification("n0t1-1234567890")

			Expect(err).Should(HaveOccurred())
			Expect(notification).To(BeNil())
		})

		It("should allow to ListNotifications()", func() {
			httpmock.RegisterResponder("GET", "https://example.com/v2/notifications",
				httpmock.NewStringResponder(http.StatusOK, `{"notifications":[{"id":"n0t1-1234567890"}, {"id":"n0t1-0123456789"}],"links":{"current":"/v2/notifications?status=delivered","next":"/v2/notifications?older_than=n0t1-0123456789&status=delivered"}}`))

			list, err := client.ListNotifications(Filters{Status: "delivered"})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(list.Notifications[0].ID).To(Equal("n0t1-1234567890"))
			Expect(list.Notifications[1].ID).To(Equal("n0t1-0123456789"))
			Expect(list.Links.Next).NotTo(BeEmpty())
		})

		It("should allow to SendEmail()", func() {
			httpmock.RegisterResponder("POST", "https://example.com/v2/notifications/email",
				httpmock.NewStringResponder(http.StatusCreated, `{"id":"df10a23e-2c6d-4ea5-87fb-82e520cbf93a"}`))

			res, err := client.SendEmail("test@example.com", "123456qwerty", templateData{}, "")

			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.ID).To(Equal("df10a23e-2c6d-4ea5-87fb-82e520cbf93a"))
		})

		It("should allow to SendLetter()", func() {
			httpmock.RegisterResponder("POST", "https://example.com/v2/notifications/letter",
				httpmock.NewStringResponder(http.StatusCreated, `{"id":"df10a23e-2c6d-4ea5-87fb-82e520cbf93a"}`))

			res, err := client.SendLetter("xxx", "123456qwerty", templateData{}, "")

			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.ID).To(Equal("df10a23e-2c6d-4ea5-87fb-82e520cbf93a"))
		})

		It("should allow to SendSms()", func() {
			httpmock.RegisterResponder("POST", "https://example.com/v2/notifications/sms",
				httpmock.NewStringResponder(http.StatusCreated, `{"id":"df10a23e-2c6d-4ea5-87fb-82e520cbf93a"}`))

			res, err := client.SendSms("00000000000", "123456qwerty", templateData{}, "")

			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.ID).To(Equal("df10a23e-2c6d-4ea5-87fb-82e520cbf93a"))
		})
	})
})
