package main

import (
	"github.com/apache/beam/sdks/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/go/pkg/beam"
)

func main(){

	p, s := beam.NewPipelineWithRoot()
	lines := textio.Read(s, "gs://some/inputData.txt")

}