# go-kv

This library provides a key-value field structure `kv.Field{}` that plays well with the `"fmt"` package. It gives an easy means of appending key-value fields to log entries, in a more performant manner that also happens to look nice! (it's not far removed from using a `map[string]interface{}`).

The formatting for these key-value fields is handled by the `"fmt"` package by default. If you set the `kvformat` build tag then it will use a custom formatting library found under `format/`.
