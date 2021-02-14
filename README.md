# Overview

This repo is my attempt to manage the content for [educative.io](https://www.educative.io) courses
I author in a GitOps manner.

It's not easy because the educative.io lessons are built on a very cool
GUI-based CMS where snippets of Markdown are mixed with custom widgets.

The educative author guide is very useful:
https://www.educative.io/courses/author-guide

# Custom docker images

To use the code widgets or the terminal with custom dependencies build a Docker image that contains all the dependencies.

Check out [Docker on Educative](https://www.educative.io/courses/author-guide/N0jNJnZPPYN) for detailed instructions.

Use the Makefile preapre the tarball with the Dockerfile

```
make tarball
```

Upload the resulting tarball.tr.gz


# Using code samples in terminal widgets

The code samples for each course are in a separate gitlab repo. To use thos in a terminal widget, clone the repo.

For example, the code for the calc-app of the Go command can be used in a terminal widget by adding the following incantation to the start script of the terminal widget:

```
cd /
git clone https://gitlab.com/the-gigi/educative-go-command-line-course.git
cd educative-go-command-line-course/calc-app
clear
```
