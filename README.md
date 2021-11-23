# Cache

[![CI](https://github.com/jonstacks/cache/actions/workflows/ci.yml/badge.svg)](https://github.com/jonstacks/cache/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/jonstacks/cache?status.png)](https://godoc.org/github.com/jonstacks/cache)
[![codecov](https://codecov.io/gh/jonstacks/cache/branch/master/graph/badge.svg)](https://codecov.io/gh/jonstacks/cache)

Various caching implementations for go.

## TimedText

The cache.TimedText is particularly useful for storing a `map[string]string`
where the keys should expire after a given duration. It also provides a nice
benefit of allowing for efficiently appending to a keys existing text.

## Testing

You can test by running `make test`.

## Dependencies

So far, there are no install dependencies outside of the standard lib. There are
currently testing dependencies to reduce the amount of testing code.
