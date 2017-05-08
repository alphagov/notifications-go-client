package notify

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	It("should Authenticate() the JWT", func() {
		config := Configuration{
			APIKey: []byte(`#5K+･ｼミew{ｦ住ｳ(跼Tﾉ(ｩ┫ﾒP.ｿﾓ燾辻G�感ﾃwb="=.!r.Oﾀﾍ奎gﾐ｣`),
			ServiceID: "test",
		}

		token, err := config.Authenticate(config.APIKey)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(*token).NotTo(BeEmpty())
	})
})
