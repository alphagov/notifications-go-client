package notify

import (
	"net/http"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

var _ = Describe("Notification", func() {
	Context("List", func() {
		var (
			client *Client
		)

		BeforeEach(func() {
			httpmock.Activate()

			u, _ := url.Parse("https://example.com")

			config := Configuration{
				APIKey: []byte(`#5K+･ｼミew{ｦ住ｳ(跼Tﾉ(ｩ┫ﾒP.ｿﾓ燾辻G�感ﾃwb="=.!r.Oﾀﾍ奎gﾐ｣`),
				BaseURL:   u,
				ServiceID: "test",
			}

			client, _ = New(config)
		})

		AfterEach(func() {
			httpmock.DeactivateAndReset()
		})

		It("should loead the Next() page", func() {
			httpmock.RegisterResponder("GET", "https://example.com/v2/notifications",
				httpmock.NewStringResponder(http.StatusOK, `{"notifications":[{"id":"n0t1-1234567890"}, {"id":"n0t1-0123456789"}],"links":{"current":"/v2/notifications?status=delivered","next":"/v2/notifications?older_than=n0t1-0123456789&status=delivered"}}`))
			httpmock.RegisterResponder("GET", "https://example.com/v2/notifications?older_than=n0t1-0123456789&status=delivered",
				httpmock.NewStringResponder(http.StatusOK, `{"notifications":[{"id":"n0t1-2345678901"}, {"id":"n0t1-9012345678"}],"links":{"current":"/v2/notifications?status=delivered","previous":"/v2/notifications?status=delivered"}}`))

			list, err := client.ListNotifications(Filters{Status: "delivered"})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(list.Notifications[0].ID).To(Equal("n0t1-1234567890"))
			Expect(list.Notifications[1].ID).To(Equal("n0t1-0123456789"))
			Expect(list.Links.Next).To(Equal("/v2/notifications?older_than=n0t1-0123456789&status=delivered"))

			err = list.Next()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(list.Notifications[0].ID).To(Equal("n0t1-2345678901"))
			Expect(list.Notifications[1].ID).To(Equal("n0t1-9012345678"))
		})

		It("should loead the Previous() page", func() {
			httpmock.RegisterResponder("GET", "https://example.com/v2/notifications?older_than=n0t1-0123456789&status=delivered",
				httpmock.NewStringResponder(http.StatusOK, `{"notifications":[{"id":"n0t1-2345678901"}, {"id":"n0t1-9012345678"}],"links":{"current":"/v2/notifications?status=delivered","previous":"/v2/notifications?status=delivered"}}`))
			httpmock.RegisterResponder("GET", "https://example.com/v2/notifications",
				httpmock.NewStringResponder(http.StatusOK, `{"notifications":[{"id":"n0t1-1234567890"}, {"id":"n0t1-0123456789"}],"links":{"current":"/v2/notifications?status=delivered","next":"/v2/notifications?older_than=n0t1-0123456789&status=delivered"}}`))

			list, err := client.ListNotifications(Filters{OlderThan: "n0t1-0123456789", Status: "delivered"})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(list.Notifications[0].ID).To(Equal("n0t1-2345678901"))
			Expect(list.Notifications[1].ID).To(Equal("n0t1-9012345678"))
			Expect(list.Links.Previous).To(Equal("/v2/notifications?status=delivered"))

			err = list.Previous()

			Expect(err).ShouldNot(HaveOccurred())
			Expect(list.Notifications[0].ID).To(Equal("n0t1-1234567890"))
			Expect(list.Notifications[1].ID).To(Equal("n0t1-0123456789"))
		})

		It("should fail to load the Next() page", func() {
			httpmock.RegisterResponder("GET", "https://example.com/v2/notifications?status=delivered",
				httpmock.NewStringResponder(http.StatusOK, `{"notifications":[{"id":"n0t1-2345678901"}, {"id":"n0t1-9012345678"}],"links":{"current":"/v2/notifications?status=delivered","previous":"/v2/notifications?status=delivered"}}`))

			list, err := client.ListNotifications(Filters{Status: "delivered"})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(list.Links.Next).To(BeEmpty())

			err = list.Next()

			Expect(err).Should(HaveOccurred())
		})

		It("should fail to load Previous() page", func() {
			httpmock.RegisterResponder("GET", "https://example.com/v2/notifications?status=delivered",
				httpmock.NewStringResponder(http.StatusOK, `{"notifications":[{"id":"n0t1-1234567890"}, {"id":"n0t1-0123456789"}],"links":{"current":"/v2/notifications?status=delivered","next":"/v2/notifications?older_than=n0t1-0123456789&status=delivered"}}`))

			list, err := client.ListNotifications(Filters{Status: "delivered"})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(list.Links.Previous).To(BeEmpty())

			err = list.Previous()

			Expect(err).Should(HaveOccurred())
		})
	})
})
