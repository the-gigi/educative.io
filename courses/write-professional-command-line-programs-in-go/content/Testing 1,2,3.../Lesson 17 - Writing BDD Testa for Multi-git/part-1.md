# Writing BDD tests for multi-git

There are several steps when writing Ginkgo tests for a package:
- Creating a test suite
- Planning the test hierarchy
- Writing the actual tests

Let's tackle these steps one by one

## Bootstrapping a test suite

The ginkgo bootstrap command creates a test suite file that registers the Ginkgo fail handler and runs all your Ginkgo tests as nested tests of the single test method. This is the integration of Ginkgo with the standard testing package of Go. Here is the result for the repo_manager package of multi-git:

```
package repo_manager_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestRepoManager(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "RepoManager Suite")
}
```

## Building the hierarchy

The top-level of the hierarchy is a `Describe` block that describes what the group of test cases in this group does. You could have multiple `Describe` blocks if you need to divide your tests even further. For example, you may want to put performance tests and benchmarks in a separate Describe block. You can nest Describe block too.

For the repo_manager package there is just one Describe block.

```
var _ = Describe("Repo manager tests", func() {
    ...
})
```

The second level, which is optional is the Context. You can have as many nested Context containers as you want. The idea is that the tests grouped under a Context block have some shared context. Here I created two contexts. One for testing failure conditions and one for testing success conditions.

```
var _ = Describe("Repo manager tests", func() {
    Context("Tests for failure cases", func() {
    })

    Context("Tests for success cases", func() {
    })
})
```

Often context will have shared initialization and cleanup blocks (ForEach, AfterEach).

Finally the tests themselves are identified by `It` blocks.

Here is a sample test that verifies the GetRepos() method works correctly.

```
It("Should get repo list successfully", func() {
    rm, err := NewRepoManager(baseDir, repoList, true)
    Ω(err).Should(BeNil())

    repos := rm.GetRepos()
    Ω(repos).Should(HaveLen(1))
    Ω(repos[0] == path.Join(baseDir, repoList[0])).Should(BeTrue())
})
```

## Writing the tests

Let's look at the repo_manager_test.go file that contains the tests for the repo_manager package. It starts with several imports. Note that Ginkgo and Gomega use the dot import notation. That means that we can use the Ginkgo blocks and Gomega assertion without qualifying them with the package name.

```
package repo_manager

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/the-gigi/multi-git/pkg/helpers"
	"os"
	"path"
	"strings"
)
```

We also import a helpers package from multi-git itself using the same dot notation. The helpers package provides some convenience functions useful for writing multi-git tests, such as creating a directory that is a git repo, adding files and committing files to a git repository. Check out the full source code here:
https://github.com/the-gigi/multi-git/blob/v0.3/pkg/helpers/helpers.go

OK. With the imports out of the way we can do a little initialization and define the base directory where the test will create sub-directories to work with as well as a variable for the list of repositories.

```
const baseDir = "tmp/test-multi-git"

var repoList = []string{}
```

Inside the Describe block we define an `err` variable we can reuse as well as a `removeAll()` to clean up the base directory. This is important to make sure we don't leave directories and files around between tests, so subsequent tests don't get confused.

```
var _ = Describe("Repo manager tests", func() {
    var err error

    removeAll := func() {
        err = os.RemoveAll(baseDir)
        Ω(err).Should(BeNil())
    }
```

The next piece of the puzzle is to write the `BeforeEach()` and `AfterEach()` blocks. Those run before and after each test. They both use the `removeAll()` function we just defined. The `BeforeEach()` function calls `removeAll()` and then uses the `CreateDir()` function from the helpers package to create a directory called "dir-1" under the base directory and initialize it as a git repository.

```
    BeforeEach(func() {
        removeAll()
        err = CreateDir(baseDir, "dir-1", true)
        Ω(err).Should(BeNil())
        repoList = []string{"dir-1"}
    })
```

The AfterEach() function just calls `removeAll()` to clean up everything. Note how concise the syntax is.
```
    AfterEach(removeAll)
``` 

It is not strictly necessary to call removeAll() after each test because the BeforeEach() of the next test will clean it up anyway. We could use AfterSuite() instead, which runs once after all tests have finished.

With all the preliminaries out of the way we can write some tests. The first test is in the context of "Tests for failure cases". It attempts to instantiate a RepoManager class by calling the NewRepoManager() function with a non-existent base directory. The test then verifies that the call returned an error using the Gomega assertion `Ω(err)ShouldNot(BeNil())`.

```
    Context("Tests for failure cases", func() {
        It("Should fail with invalid base dir", func() {
            _, err := NewRepoManager("/no-such-dir", repoList, true)
            Ω(err).ShouldNot(BeNil())
        })
```

The next test verifies that if the repository list is empty then NewRepoManager() returns an error:

```
        It("Should fail with empty repo list", func() {
            _, err := NewRepoManager(baseDir, []string{}, true)
            Ω(err).ShouldNot(BeNil())
        })
    })
```

Let's look at a success test case. This is the most complicated test case so far. It instantiates a RepoManager object. It then creates a brunch called "test-branch" and makes sure that the output is correct. Then it uses the AddFiles() helper function to add some dummy files to the repository and commit them. Finally, it runs the `git log --oneline` command and verifies that indeed the files were committed properly.

```
It("Should commit files successfully", func() {
    rm, err := NewRepoManager(baseDir, repoList, true)
    Ω(err).Should(BeNil())

    output, err := rm.Exec("checkout -b test-branch")
    Ω(err).Should(BeNil())

    for _, out := range output {
        Ω(out).Should(Equal("Switched to a new branch 'test-branch'\n"))
    }

    AddFiles(baseDir, repoList[0], true, "file_1.txt", "file_2.txt")

    // Restore working directory after executing the command
    wd, _ := os.Getwd()
    defer os.Chdir(wd)

    dir := path.Join(baseDir, repoList[0])
    err = os.Chdir(dir)
    Ω(err).Should(BeNil())

    output, err = rm.Exec("log --oneline")
    Ω(err).Should(BeNil())

    ok := strings.HasSuffix(output[dir], "added some files...\n")
    Ω(ok).Should(BeTrue())
})
```

There are more test cases you can review here: [repo_manager_test.go](https://github.com/the-gigi/multi-git/blob/v0.3/pkg/repo_manager/repo_manager_test.go)

Feel free to browse and look around in the live terminal below

```
cd /tmp
git clone --depth 1 --branch v0.3 https://github.com/the-gigi/multi-git.git
cd multi-git/pkg/repo_manager
<<< Terminal >>>
```

In the next lesson we will actually run our tests.
