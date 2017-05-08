package notify

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filters", func() {
	var f Filters

	BeforeEach(func() {
		f = Filters{
			"older_than",
			"reference",
			"status",
			"template_type",
		}
	})

	It("should addIfNotEmpty()", func() {
		m := url.Values{}

		f.addIfNotEmpty(&m, "success", "true")
		f.addIfNotEmpty(&m, "failure", "")

		Expect(m["success"][0]).To(Equal("true"))
		_, ok := m["failure"]
		Expect(ok).To(BeFalse())
	})

	It("should be handle converting ToURLValues()", func() {
		v := f.ToURLValues()

		Expect(v).NotTo(BeEmpty())
		Expect(v["older_than"][0]).To(Equal("older_than"))
		Expect(v["reference"][0]).To(Equal("reference"))
		Expect(v["status"][0]).To(Equal("status"))
		Expect(v["template_type"][0]).To(Equal("template_type"))
	})

	It("should ignore empty fields in the ToURLValues() conversion", func() {
		f.Status = ""
		f.TemplateType = ""

		v := f.ToURLValues()

		Expect(v).NotTo(BeEmpty())
		Expect(v["older_than"][0]).To(Equal("older_than"))
		Expect(v["reference"][0]).To(Equal("reference"))
		_, okStatus := v["status"]
		Expect(okStatus).To(BeFalse())
		_, okTemplateType := v["template_type"]
		Expect(okTemplateType).To(BeFalse())
	})
})
