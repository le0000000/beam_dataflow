package main

import (
	"context"
	"flag"
	"reflect"
	"strconv"

	"github.com/apache/beam/sdks/go/pkg/beam"
	"github.com/apache/beam/sdks/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/go/pkg/beam/log"
	"github.com/apache/beam/sdks/go/pkg/beam/transforms/stats"
	"github.com/apache/beam/sdks/go/pkg/beam/x/beamx"

	// The following packages are required to read files from GCS or local.
	_ "github.com/apache/beam/sdks/go/pkg/beam/io/filesystem/gcs"
	_ "github.com/apache/beam/sdks/go/pkg/beam/io/filesystem/local"
)

func init() {
	beam.RegisterType(reflect.TypeOf((*parseLineFn)(nil)).Elem())
	beam.RegisterFunction(formatResultFn)
}

var (
	inputURI  = flag.String("input_uri", "", "Input file with a number on each line.")
	outputURI = flag.String("output_uri", "", "Output file.")
)

func main() {
	flag.Parse()
	beam.Init()

	ctx := context.Background()
	pipeline := beam.NewPipeline()

	scope := pipeline.Root()
	lines := textio.ReadSdf(scope, *inputURI)
	reshuffledLines := beam.Reshuffle(scope, lines)

	values := beam.ParDo(scope, &parseLineFn{}, reshuffledLines)
	sum := stats.Sum(scope, values)

	formatted := beam.ParDo(scope, formatResultFn, sum)
	textio.Write(scope, *outputURI, formatted)

	if err := beamx.Run(ctx, pipeline); err != nil {
		log.Exitf(ctx, "Failed to execute job: %s", err)
	}
}

type parseLineFn struct {
	lineCounter beam.Counter
}

func (fn *parseLineFn) Setup() {
	fn.lineCounter = beam.NewCounter("simple_example", "parse-line-count")
}

func (fn *parseLineFn) ProcessElement(ctx context.Context, line string, emit func(uint64)) error {
	// value, err := strconv.ParseUint(line, 10, 64)
	// if err != nil {
	// 	return err
	// }
	// emit(value)
	emit(1)
	fn.lineCounter.Inc(ctx, 1)
	return nil
}

func formatResultFn(sum uint64, emit func(string)) error {
	emit(strconv.FormatUint(sum, 10))
	return nil
}
