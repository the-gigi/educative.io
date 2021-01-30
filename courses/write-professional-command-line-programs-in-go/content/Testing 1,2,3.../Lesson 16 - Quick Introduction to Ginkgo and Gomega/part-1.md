# What is BDD?

BDD stands for behavior-driven development. The basic idea is to define your systems behavior at a higher level of abstraction. Often less technical stakeholders can participate and validate those specification, which increase the likelihood that the developers actually build a system that meets the requirements. The key point is to connect the behavioral spec to actual tests. Ginkgo is a BDD framework for Go that can help writing behavior-driven tests.

# Quick introduction to Ginkgo

Ginkgo is a framework that builds on top on the standard Go testing package and tooling. You define your BDD tests, but under the cover Ginkgo creates standard Go tests.

This is powerful for several reasons. It's easy to migrate to Ginkgo tests because they can co-exist side by side with normal Go tests. You can also take advantage of many other libraries tools that integrate with the standard Go testing support. This collaborative ecosystem effect is a testament to the wisdom of the Go designers that included testing as a core capability.

Ginkgo lets you specify your tests in hierarchical and higher level constructs that are close to natural language. You use building blocks like: `Describe`, `Context` and `It` to structure your tests. You provide setup and teardown at different levels using `BeforeEach`, `AfterEach`, `BeforeSuite` and `AfterSuite`. You can run asynchronous tests too using Ginkgo.

Here is an example of what Ginkgo tests look like:

```
var _ = Describe("Repo manager tests", func() {
    BeforeEach(func() {
        ...
    })
    AfterEach(func() {
        ...
    })
    
    Context("Tests for failure cases", func() {
        It("Should fail with invalid base dir", func() {
            ...
        })

        It("Should fail with empty repo list", func() {
            ...        
        })
    })

    Context("Tests for success cases", func() {
        It("Should get repo list successfully", func() {
            ...        
        })

        It("Should get repo list successfully with non-git directories", func() {
            ...        
        })

        It("Should get repo list successfully with non-git directories", func() {
            ...        
        })

        It("Should create branches successfully", func() {
            ...        
        })

        It("Should commit files successfully", func() {

        })
    })
})
```

# Quick introduction to Gomega

Gomega is a general-purpose assertion library that can be used with any Go test framework (including standard Go tests). However, it is a great complement to Ginkgo. It is not required, but it is a very popular option. Gomega is used inside the test functions and provides a slew of useful assertions. Gomega provides two alternative and equivalent forms of assertions. You can use the `Expect()` syntax or the `Ω()` syntax (which gives Gomega its name) interchangeably.

Here are a couple of examples in both forms that are functionally identical.

The first example checks that a value is equals to 5 using both forms:

```
Ω(value).Should(Equal(5))
Expect(value).To(Equal(5))
```

The second example checks that today is not Monday using both forms:

```
Ω(today).ShouldNot(Equal("Monday"))
Expect(today).ToNot(Equal("Monday"))
```

The Gomega syntax is a little more concise and IMO more cool :-). The `Expect` syntax may be more familiar and easier to understand.

On the right side of the expression a Gomega matcher is expected. There is a large number of matchers that cover many areas such as numbers, collections, errors, files and channels. Asynchronous assertions poll their values periodically until the expected value is set or until the timeout expires (defaults to one second), in which case the assertion fails.

Check out the complete list here: https://onsi.github.io/gomega/#provided-matchers

In addition you can write custom matchers that suit your test needs. Your custom matcher must implement the following interface:

```
type GomegaMatcher interface {
    Match(actual interface{}) (success bool, err error)
    FailureMessage(actual interface{}) (message string)
    NegatedFailureMessage(actual interface{}) (message string)
}
```

Now, that we got acquainted with Ginkgo and Gomega let's get busy and write some tests to ensure multi-git works as expected.


