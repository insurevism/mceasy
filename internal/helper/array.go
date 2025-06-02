package helper

func ArrayChunk[Data any](items []Data, size int) (chunks [][]Data) {

	var buffers []Data
	buffers = append(buffers, items...)

	for size < len(buffers) {
		chunks = append(chunks, buffers[0:size:size])
		buffers = buffers[size:]
	}
	return append(chunks, buffers)
}
