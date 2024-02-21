package slice

func Filter[T any](s []T, predicate func(T) (ok bool)) (result []T) {

	out := make([]chan FilterResult[T], len(s))

	for i := 0; i < len(s); i++ {
		out[i] = make(chan FilterResult[T], 1)
	}

	for i := 0; i < len(s); i++ {
		idx := i
		go func() {
			out[idx] <- FilterResult[T]{
				Ok:   predicate(s[idx]),
				Item: s[idx],
			}
		}()
	}

	for i := 0; i < len(s); i++ {

		outChan := out[i]
		filterResult := <-outChan

		if filterResult.Ok {
			result = append(result, filterResult.Item)
		}

	}

	return result
}

type FilterResult[T any] struct {
	Ok   bool
	Item T
}
