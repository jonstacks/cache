# Cache

Various caching implementations for go.

## TimedTextCache

The TimedTextCache is particularly useful for storing a `map[string]string`
where the keys should expire after a given duration. It also provides a nice
benefit of allowing for efficiently appending to a keys existing text.

## Testing

You can test by running `make test`.

## Dependencies

So far, there are no install dependencies outside of the standard lib. There are
currently testing dependencies to reduce the amount of testing code.
