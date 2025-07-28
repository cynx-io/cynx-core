package helper

func SetIfNotNil[T any](target *T, source *T) {
	if source != nil {
		*target = *source
	}
}
