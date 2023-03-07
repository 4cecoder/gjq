package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"github.com/fatih/color"
	"regexp"
	"os"
)

func main() {
	compact := flag.Bool("c", false, "compact instead of pretty-printed output")
	nullInput := flag.Bool("n", false, "use 'null' as the single input value")
	exitStatus := flag.Bool("e", false, "set the exit status code based on the output")
	slurp := flag.Bool("s", false, "read (slurp) all inputs into an array; apply filter to it")
	rawStrings := flag.Bool("r", false, "output raw strings, not JSON texts")
	rawInput := flag.Bool("R", false, "read raw strings, not JSON texts")
	colorize := flag.Bool("C", false, "colorize JSON")
	monochrome := flag.Bool("M", false, "monochrome (don't colorize JSON)")
	sortKeys := flag.Bool("S", false, "sort keys of objects on output")
	useTabs := flag.Bool("tab", false, "use tabs for indentation")
	argValues := flag.String("arg", "", "set variable $a to value <v>")
	argJSONValues := flag.String("argjson", "", "set variable $a to JSON value <v>")
	slurpFile := flag.String("slurpfile", "", "set variable $a to an array of JSON texts read from <f>")
	rawFile := flag.String("rawfile", "", "set variable $a to a string consisting of the contents of <f>")
	args := flag.Bool("args", false, "remaining arguments are string arguments, not files")
	jsonArgs := flag.Bool("jsonargs", false, "remaining arguments are JSON arguments, not files")
	flag.Parse()

	var input []byte
	var err error

	if *nullInput {
		input = []byte("null")
	} else if *args {
		input = []byte(flag.Arg(0))
	} else if *jsonArgs {
		input = []byte(flag.Arg(0))
	} else if flag.NArg() > 0 {
		input, err = ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	var v interface{}
	if *rawInput {
		v = string(input)
	} else {
		err = json.Unmarshal(input, &v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	var filter string
	if *args || *jsonArgs {
		filter = flag.Arg(0)
	} else {
		filter = flag.Arg(1)
	}

	output, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *compact {
		output, _ = json.Marshal(v)
	}

	if *rawStrings {
		output = []byte(fmt.Sprintf("%v", v))
	}

	if *colorize {
	// Colorize JSON
	var obj interface{}
	err = json.Unmarshal(output, &obj)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	colorized, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Add color to keys and strings
	outputString := string(colorized)
	outputString = color.GreenString(outputString)
	outputString = color.BlueString(outputString)

	fmt.Println(outputString)
	}
	
	if *monochrome {
	// Remove color codes from output
	re := regexp.MustCompile("\x1b\\[[0-9;]*[m|K]")
	outputString := re.ReplaceAllString(string(output), "")
	fmt.Println(outputString)
	}

	fmt.Println(string(output))

	if *exitStatus {
		os.Exit(1)
	}
}
