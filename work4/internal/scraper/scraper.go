package scraper

import "sync"

func Run(conf *Config, input <-chan string) <-chan *Result {
	output := make(chan *Result)

	var wg sync.WaitGroup

	for range conf.ThreadCount {
		wg.Add(1)

		go func() {
			NewProcessor(conf).Do(input, output)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
