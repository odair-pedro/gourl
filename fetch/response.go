package fetch

type Response struct {
	*Metrics
	length int64
	err    error
}

func (r Response) Length() int64 {
	return r.length
}

func (r Response) Err() error {
	return r.err
}
