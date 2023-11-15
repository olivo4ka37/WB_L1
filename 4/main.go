package main

func main() {
	const workerPool = 5

	for w := 0; w < workerPool; w++ {
		go worker(w)
	}

}

func worker(id int) {

}
