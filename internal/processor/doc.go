// Package processor provides the core pipeline that connects a line reader,
// an optional filter chain, and a buffered writer to process log files
// efficiently.
//
// Basic usage:
//
//	p, err := processor.New(processor.Config{
//		Input:  inputFile,
//		Output: outputFile,
//		Filter: myFilter,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	stats, err := p.Run(context.Background())
package processor
