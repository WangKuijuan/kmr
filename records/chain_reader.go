package records

// ChainReader Chain some reader to one
type ChainReader struct {
	RecordReader
	readers []RecordReader
	current int
}

// Peek Peek one record from ChainReader
func (cr *ChainReader) Peek() *Record {
	return cr.readers[cr.current].Peek()
}

// Pop Pop one record from ChainReader
func (cr *ChainReader) Pop() *Record {
	return cr.readers[cr.current].Pop()
}

// HasNext HasNext
func (cr *ChainReader) HasNext() bool {
	for ; cr.current < len(cr.readers); cr.current++ {
		if cr.readers[cr.current].HasNext() {
			return true
		}
	}
	return false
}

// Close Close all reader open by ChainReader
func (cr *ChainReader) Close() error {
	for _, reader := range cr.readers {
		err := reader.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func NewChainReader(readers []RecordReader) *ChainReader {
	return &ChainReader{
		readers: readers,
		current: 0,
	}
}
