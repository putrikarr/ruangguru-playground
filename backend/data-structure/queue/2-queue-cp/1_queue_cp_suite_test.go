package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test2QueueCp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2QueueCp Suite")
}
