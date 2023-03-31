package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/fatih/color"
)

func main() {
	compact := flag.Bool("c", false, "compact instead of pretty-printed output")
	nullInput := flag.Bool("n", false, "use 'null' as the single input value")
	exitStatus := flag.Bool("e", false, "set the exit status code based on the output")
	colorize := flag.Bool("C", false, "colorize JSON")
	monochrome := flag.Bool("M", false, "monochrome (don't colorize JSON)")
	flag.Parse()

	var input []byte
	var err error

	if *nullInput {
		input = []byte("null")
	} else if flag.NArg() > 0 {
		input, err = ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			exitWithError(err)
		}
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			exitWithError(err)
		}
	}

	var v interface{}
	err = json.Unmarshal(input, &v)
	if err != nil {
		exitWithError(err)
	}

	var output []byte
	if *compact {
		output, err = json.Marshal(v)
	} else {
		output, err = json.MarshalIndent(v, "", "    ")
	}
	if err != nil {
		exitWithError(err)
	}

	if *colorize {
		output, err = colorizeJSON(output)
		if err != nil {
			exitWithError(err)
		}
	}

	if *monochrome {
		output = removeColorCodes(output)
	}

	fmt.Println(string(output))

	if *exitStatus {
		os.Exit(1)
	}
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func colorizeJSON(jsonBytes []byte) ([]byte, error) {
	var obj interface{}
	err := json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		return nil, err
	}

	colorized, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return nil, err
	}

	outputString := string(colorized)
	outputString = color.GreenString(outputString)
	outputString = color.BlueString(outputString)

	return []byte(outputString), nil
}

func removeColorCodes(jsonBytes []byte) []byte {
	re := regexp.MustCompile("\x1b\\[[0-9;]*[m|K]")
	outputString := re.ReplaceAllString(string(jsonBytes), "")
	return []byte(outputString)
}

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
