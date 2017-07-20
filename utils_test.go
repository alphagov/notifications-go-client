package notify

import (
	"bytes"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testingStruct struct {
	Bool   bool
	Int    int
	String string
	Other  string `json:"the_other_field"`
}

type testingStructMap struct {
	Int     int64
	Uint    uint64
	Float32 float32
	Float64 float64
	Byte    []byte
	String  string
}

var _ = Describe("Utils", func() {
	It("should be able to parse jsonResponse()", func() {
		s := `{"bool":true,"int":123,"string":"testing","the_other_field":"Hazza!"}`
		obj := testingStruct{}
		err := jsonResponse(ioutil.NopCloser(bytes.NewReader([]byte(s))), &obj)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(obj.Bool).To(BeTrue())
		Expect(obj.Int).To(Equal(123))
		Expect(obj.String).NotTo(BeEmpty())
		Expect(obj.Other).NotTo(BeEmpty())
	})

	It("should return an error on jsonResponse() with wrong content", func() {
		s := `status: error`
		obj := testingStruct{}
		err := jsonResponse(ioutil.NopCloser(bytes.NewReader([]byte(s))), &obj)

		Expect(err).Should(HaveOccurred())
		Expect(obj.Bool).To(BeFalse())
		Expect(obj.Int).To(Equal(0))
		Expect(obj.String).To(BeEmpty())
		Expect(obj.Other).To(BeEmpty())
	})
})
