Running the code

```
docker run -it educative-go-course bash
  git clone https://github.com/the-gigi/multi-git.git
  cd multi-git
  git checkout v0.5  
  go mod download
  cd pkg/repo_manager

```
$ go test
Running Suite: RepoManager Suite
================================
Random Seed: 1575175770
Will run 7 of 7 specs

•••••••
Ran 7 of 7 Specs in 3.956 seconds
SUCCESS! -- 7 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
ok  	github.com/the-gigi/multi-git/pkg/repo_manager	3.983s
```

The `go test` command has several other command-line arguments that can be useful. Check out the full list here: https://golang.org/cmd/go/#hdr-Testing_flags

However, when running Ginkgo tests I recommend using the ginkgo command itself.

## Running Ginkgo tests with ginkgo

First, make sure that ginkgo is installed. If it's not follow the instructions here:
https://onsi.github.io/ginkgo/#getting-ginkgo

Now, let's run the tests with `ginkgo -v`.
```
$ ginkgo -v
Running Suite: RepoManager Suite
================================
Random Seed: 1575179047
Will run 7 of 7 specs

Repo manager tests Tests for failure cases
  Should fail with invalid base dir
  /Users/gigi.sayfan/git/multi-git/pkg/repo_manager/repo_manager_test.go:33
•
------------------------------
Repo manager tests Tests for failure cases
  Should fail with empty repo list
  /Users/gigi.sayfan/git/multi-git/pkg/repo_manager/repo_manager_test.go:38
•
------------------------------
Repo manager tests Tests for success cases
  Should get repo list successfully
  /Users/gigi.sayfan/git/multi-git/pkg/repo_manager/repo_manager_test.go:45
•
------------------------------
Repo manager tests Tests for success cases
  Should get repo list successfully with non-git directories
  /Users/gigi.sayfan/git/multi-git/pkg/repo_manager/repo_manager_test.go:56
•
------------------------------
Repo manager tests Tests for success cases
  Should get repo list successfully with non-git directories
  /Users/gigi.sayfan/git/multi-git/pkg/repo_manager/repo_manager_test.go:69
•
------------------------------
Repo manager tests Tests for success cases
  Should create branches successfully
  /Users/gigi.sayfan/git/multi-git/pkg/repo_manager/repo_manager_test.go:82
•
------------------------------
Repo manager tests Tests for success cases
  Should commit files successfully
  /Users/gigi.sayfan/git/multi-git/pkg/repo_manager/repo_manager_test.go:96
•
Ran 7 of 7 Specs in 2.053 seconds
SUCCESS! -- 7 Passed | 0 Failed | 0 Pending | 0 Skipped
```

As you can see the context and the test name are displayed for each test.

## Checking test coverage

We can check the coverage with ginkgo as well:

```
$ ginkgo -cover
Running Suite: RepoManager Suite
================================
Random Seed: 1575179204
Will run 7 of 7 specs

•••••••
Ran 7 of 7 Specs in 2.263 seconds
SUCCESS! -- 7 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
coverage: 78.0% of statements

Ginkgo ran 1 suite in 10.383398546s
Test Suite Passed
```

OK. We have 78% coverage, but which parts of the code are not covered by tests? Let's generate a coverage profile that we can later review:

```
$ ginkgo -coverprofile=coverage.out
```

The coverage.out file is a simple text file with lines that are not covered:

```
$ cat coverage.out
mode: set
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:16.114,18.16 2 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:25.2,25.36 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:29.2,29.25 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:34.2,37.30 2 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:42.2,42.8 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:18.16,19.25 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:22.3,22.9 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:19.25,21.4 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:25.36,27.3 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:29.25,32.3 2 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:37.30,40.3 2 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:45.43,47.2 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:49.78,53.52 4 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:74.2,78.28 4 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:92.2,92.8 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:53.52,54.41 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:59.3,59.25 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:70.3,70.45 1 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:54.41,56.12 2 0
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:59.25,60.43 1 0
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:65.4,67.26 3 0
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:60.43,62.13 2 0
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:78.28,88.36 4 1
github.com/the-gigi/multi-git/pkg/repo_manager/repo_manager.go:88.36,90.4 1 0
```

This is not so user-friendly. But, we can use the `go cover` tool to visualize the results:

```
$ go tool cover -html=coverage.out
```

That will read the coverage.out file and open a browser window where the covered lines are displayed in green and the uncovered lines are displayed in red:
