package main

func main() {
	charset := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	myChan := make(chan byte, len(charset))

	// sending to channel
	for _, c := range charset {
		select {
		case myChan <- c:

		}
	}
	close(myChan)
	// receiving from channel
	for c := range myChan {
		println(string(c))
	}

}
