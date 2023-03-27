// All rights reserved, Viktor Boman 2023.

// # Better flags because why not?
//
// Reason for this package: I tried using the flag package during a project at 01.edu and it refused to work
// as it should. This package does not care which order the command line arguments come in. They can come before or after
// the flags.
//
// Improvements:
//
// Can use any type, so you only have to call "betterflags.Create(name, usage, defaultValue, print)" to make a flag of any type.
// It will parse flags in any position in the os.Args.
// It will work with complex numbers.
package betterflags

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Usage:
// Create your flags with betterflags.Create(). The function returns a pointer to the data of the flag.
// This data can be of any type and is defined by the chosen default value.
// If you set last argument as "true" then it will print the structure to the command-line.
// You can also use function "Lookup()" for that task.

// The flag structure. Contains name, usage message, defaultvalue and data.
type Flag struct {
	Name         string
	Usage        string
	DefaultValue any
	Data         any
}

var Flags []*Flag
var help, parsed bool

// Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
// If print is set to true it will call the "PrintParsed" method and print out the parsed flags to the standard output.
func Parse(print bool) error {
	parsed = true
	// Checks for data from the flags and puts it in the flag struct
	Arguments := os.Args[1:]
	for _, arg := range Arguments {
		if arg == "-h" || arg == "-help" {
			help = true
		}
	}
	for _, flag := range Flags {
		for _, arg := range Arguments {
			if arg == flag.Name {
				fmt.Printf("flag provided but not defined: %s\n", flag.Name)
				os.Exit(2)
			}
			if len(arg) > len(flag.Name) {
				if arg[0:len(flag.Name)] == flag.Name {
					flag.Data = arg[len(flag.Name)+1:]
					break
				}
			}
		}
		if flag.Data == nil || flag.Data == "" {
			flag.Data = flag.DefaultValue
		}
	}
	for _, flag := range Flags {
		var err error
		var value = flag.Data
		if value != nil && print {
			flag.PrintParsed()
		}
		switch flag.Data.(type) {
		case string:
			if value.(string) == "" {
				value = flag.DefaultValue
			}
			switch flag.DefaultValue.(type) {
			case uint8:
				if flag.Data, err = strconv.ParseUint(string(value.(string)), 10, 8); err != nil {
					log.Fatalf("Fails ParseUint conversion from string to uint8.")
				}
			case uint16:
				if flag.Data, err = strconv.ParseUint(string(value.(string)), 10, 16); err != nil {
					log.Fatalf("Fails ParseUint conversion from string to uint16.")
				}
			case uint32:
				if flag.Data, err = strconv.ParseUint(string(value.(string)), 10, 32); err != nil {
					log.Fatalf("Fails ParseUint conversion from string to uint32.")
				}
			case uint64:
				if flag.Data, err = strconv.ParseUint(string(value.(string)), 10, 64); err != nil {
					log.Fatalf("Fails ParseUint conversion from string to uint64.")
				}
			case int:
				if flag.Data, err = strconv.ParseInt(string(value.(string)), 10, 32); err != nil {
					log.Fatalf("Fails ParseInt conversion from string to int.")
				}
			case int8:
				if flag.Data, err = strconv.ParseInt(string(value.(string)), 10, 8); err != nil {
					log.Fatalf("Fails ParseInt conversion from string to int8.")
				}
			case int16:
				if flag.Data, err = strconv.ParseInt(string(value.(string)), 10, 16); err != nil {
					log.Fatalf("Fails ParseInt conversion from string to int16.")
				}
			case int32:
				if flag.Data, err = strconv.ParseInt(string(value.(string)), 10, 32); err != nil {
					log.Fatalf("Fails ParseInt conversion from string to int32.")
				}
			case int64:
				if flag.Data, err = strconv.ParseInt(string(value.(string)), 10, 64); err != nil {
					log.Fatalf("Fails ParseInt conversion from string to int64.")
				}
			case float32:
				if flag.Data, err = strconv.ParseFloat(string(value.(string)), 32); err != nil {
					log.Fatalf("Fails ParseFloat conversion from string to float32.")
				}
			case float64:
				if flag.Data, err = strconv.ParseFloat(string(value.(string)), 64); err != nil {
					log.Fatalf("Fails ParseFloat conversion from string to float64.")
				}
			case complex64:
				if flag.Data, err = strconv.ParseComplex(string(value.(string)), 64); err != nil {
					log.Fatalf("Fails ParseComplex conversion from string to Complex64.")
				}
			case complex128:
				if flag.Data, err = strconv.ParseComplex(string(value.(string)), 128); err != nil {
					log.Fatalf("Fails ParseComplex conversion from string to Complex128.")
				}
			case uint:
				if flag.Data, err = strconv.ParseUint(string(value.(string)), 10, 32); err != nil {
					log.Fatalf("Fails ParseUint conversion from string to uint.")
				}
			default:
				flag.Data = value
				continue
			}
		case interface{}:
			if value == nil {
				flag = nil
				break
			}
		}
	}

	for _, flag := range Flags {
		if flag.Data != nil && help {
			fmt.Println(flag.Usage)
		}
	}
	return nil
}

// Parsed reports whether the command-line flags have been parsed.
func Parsed() bool {
	return parsed
}

// Create just creates a flag and stores it in the global Flags slice.
//
// Use it by giving it a name (add '-' to make it a terminator flag), eventual help text, and the data you want to store in the flag.
// It will return the flag struct containing:
//
//	type Flag struct {
//		Name         string
//		Usage        string
//		DefaultValue any
//		Data         any
//	}
func Create(name, usage string, defaultValue any, print bool) *any {
	if name == "" {
		fmt.Println("Empty name in creation of flag. Returning nil")
		return nil
	}

	if name == "-h" || name == "-help" {
		fmt.Println("Helper name in creation of flag. Returning nil")
		return nil
	}

	if name[0] != '-' {
		name = "-" + name
	}

	var f = Flag{name, usage, defaultValue, nil}
	Flags = append(Flags, &f)
	if print {
		f.Print()
	}
	return &f.Data
}

// CreateVar is like Create but assigning value to already created variable.
// In order to use you must use the "any" keyword of the variable.
func CreateVar(p *any, name, usage string, defaultValue any, print bool) {
	if name == "" {
		fmt.Println("Empty name in creation of flag. Returning nil")
		return
	}

	if name == "-h" || name == "-help" {
		fmt.Println("Helper name in creation of flag.")
		return
	}

	if name[0] != '-' {
		name = "-" + name
	}

	var f = Flag{name, usage, defaultValue, nil}
	Flags = append(Flags, &f)
	if print {
		f.Print()
	}
	*p = &f.Data
}

// Visit visits the command-line flags in lexicographical order, calling fn for each. It visits only those flags that have been set.
func Visit(fn func(*Flag)) {
	sortedSlice := []string{}
	sortedFlags := []*Flag{}
	for _, flag := range Flags {
		if flag.Data != nil {
			sortedSlice = append(sortedSlice, flag.Name)
		}
	}
	sort.Strings(sortedSlice)
	for _, name := range sortedSlice {
		for _, flag := range Flags {
			if flag.Name == name {
				sortedFlags = append(sortedFlags, flag)
			}
		}
	}
	for _, flag := range sortedFlags {
		fn(flag)
	}
}

// VisitAll visits the command-line flags in lexicographical order, calling fn for each. It visits all flags, even those not set.
func VisitAll(fn func(*Flag)) {
	sortedSlice := []string{}
	sortedFlags := []*Flag{}
	for _, flag := range Flags {
		sortedSlice = append(sortedSlice, flag.Name)

	}
	sort.Strings(sortedSlice)
	for _, name := range sortedSlice {
		for _, flag := range Flags {
			if flag.Name == name {
				sortedFlags = append(sortedFlags, flag)
			}
		}
	}
	for _, flag := range sortedFlags {
		fn(flag)
	}
}

// Lookup returns the Flag structure of the named command-line flag, returning nil if none exists.
func Lookup(name string) *Flag {
	for _, flag := range Flags {
		if flag.Name[1:] == name {
			return flag
		}
	}
	return nil
}

// TODO: NArg
// NArg is the number of arguments remaining after flags have been processed.
// func NArg() int {
// 	return len(os.Args[1:]) - len(Flags)
// }

// NFlag is the number of created flags.
func NFlag() int {
	return len(Flags)
}

// Prints the creation of the flag. Can be toggled. Example:
//
//	"Flag Created!"
//	"Name: -name"
//	"Default: true"
//	"Type: bool"
func (f Flag) Print() {
	fmt.Printf("\tFlag Created!\n\tName: %s\n\tDefault: %v\n\tType: %T\n\n", f.Name, f.DefaultValue, f.DefaultValue)
}

// Prints a parsed flag. Example:
//
//	"Flag -name Parsed!"
//	"Data: true"
//	"Type: bool"
func (f Flag) PrintParsed() {
	fmt.Printf("\tFlag %s Parsed!\n\tData: %v\n\tType: %T\n\n", f.Name, f.Data, f.Data)
}

// Prints all default values and their types. Example:
//
//	"-name bool"
//			"usage msg (default true)"
func PrintDefaults() {
	for _, flag := range Flags {
		fmt.Printf("%s %T\n\t%s (default %v)\n\n", flag.Name, flag.DefaultValue, flag.Usage, flag.DefaultValue)
	}
}
