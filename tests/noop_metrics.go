package tests

type NoopObserver struct{}

func (o *NoopObserver) Record(source, operation string, failed bool) {}

func NewNoopObserver() *NoopObserver {
	return &NoopObserver{}
}
