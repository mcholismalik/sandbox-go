package main

import "sandbox-go/worker"

func main() {
	// worker.RunWgWait()
	worker.RunWorkerPool()
	// worker.RunWorkerPoolBatch()
	// worker.RunWorkerPoolRace()
}
