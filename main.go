package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go"
	"github.com/jmespath/go-jmespath"
	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("My new project!"))
}

