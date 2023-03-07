# gjq

### gjq is a Go implementation of the popular command-line JSON processor jq.

# Usage
```bash
gjq [options] <jq filter> [file...]
gjq [options] --args <jq filter> [strings...]
gjq [options] --jsonargs <jq filter> [JSON_TEXTS...]
```


`gjq` is a tool for processing JSON inputs, applying the given filter to
its JSON text inputs and producing the filter's results as JSON on
standard output.

The simplest filter is ., which copies gjq's input to its output
unmodified (except for formatting, but note that `IEEE754` is used
for number representation internally, with all that that implies).


# Options

Some of the available options include:

    -c: Compact instead of pretty-printed output
    -n: Use null as the single input value
    -e: Set the exit status code based on the output
    -s: Read (slurp) all inputs into an array; apply filter to it
    -r: Output raw strings, not JSON texts
    -R: Read raw strings, not JSON texts
    -C: Colorize JSON
    -M: Monochrome (don't colorize JSON)
    -S: Sort keys of objects on output
    --tab: Use tabs for indentation
    --arg a v: Set variable $a to value <v>
    --argjson a v: Set variable $a to JSON value <v>
    --slurpfile a f: Set variable $a to an array of JSON texts read from <f>
    --rawfile a f: Set variable $a to a string consisting of the contents of <f>
    --args: Remaining arguments are string arguments, not files
    --jsonargs: Remaining arguments are JSON arguments, not files
    --: Terminates argument processing

Named arguments are also available as $ARGS.named[], while
positional arguments are available as $ARGS.positional[].

# Example

```bash
echo '{"foo": 0}' | gjq .
```

```json
{
    "foo": 0
}
```

License

gjq is licensed under the MIT License. See LICENSE for more information.
