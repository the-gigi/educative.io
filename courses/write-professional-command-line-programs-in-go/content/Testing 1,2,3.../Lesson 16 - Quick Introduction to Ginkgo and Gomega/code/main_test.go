package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ginkgo and Gomega demo", func () {
	It("Should succeed", func () {
		value := 5
		Ω(value).Should(Equal(5))
		Expect(value).To(Equal(5))
	})

	It("Should also succeed", func () {
		today := "Sunday"
		Ω(today).ShouldNot(Equal("Monday"))
		Expect(today).ToNot(Equal("Monday"))
	})
})
