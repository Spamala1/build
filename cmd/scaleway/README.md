<!-- Auto-generated by x/build/update-readmes.go -->

[![GoDoc](https://godoc.org/golang.org/x/build/cmd/scaleway?status.svg)](https://godoc.org/golang.org/x/build/cmd/scaleway)

# golang.org/x/build/cmd/scaleway

The scaleway command creates ARM servers on Scaleway.com.

The Makefile, Dockerfiles, and kubernetes deployment file
are for running the command in daemon mode for server repair.
This will check on Scaleway and restart down or wedged instances.
