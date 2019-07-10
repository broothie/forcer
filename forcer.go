package forcer

import (
	"github.com/stretchr/testify/require"
)

//go:generate go run result_gen.go

type TestingT interface {
	require.TestingT
	Helper()
}

type Forcer func(...interface{}) Result

func New(t TestingT) Forcer {
	return func(returns ...interface{}) Result {
		t.Helper()
		rest, err := force(returns...)
		require.NoError(t, err)
		return Result(rest)
	}
}

func force(returns ...interface{}) ([]interface{}, error) {
	rest := returns[:len(returns)-1]
	last := returns[len(returns)-1]
	if err, isError := last.(error); isError && err != nil {
		return rest, err
	}

	return rest, nil
}
