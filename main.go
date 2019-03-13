package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/cjiq/syn/data"
	"github.com/fatih/color"
)

type context struct {
	n int
	h bool
	e bool
	i bool
}

var con = &context{}

var apiURL = "https://tuna.thesaurus.com/pageData/"
var title = "Thesaurus v1.0"

func init() {
	flag.BoolVar(&con.h, "h", false, "Show help")
	flag.BoolVar(&con.h, "help", false, "Show help")
	flag.BoolVar(&con.e, "e", false, "Show example or definition")
	flag.BoolVar(&con.e, "example", false, "Show example or definition")
	flag.BoolVar(&con.i, "I", false, "Interactive mode")
	flag.BoolVar(&con.i, "interactive", false, "Interactive mode")
	flag.Parse()
	if con.h {
		fmt.Println(title + " - Find synonyms")
		showExample()
		fmt.Println("Help: ")
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {
	args := flag.Args()
	searchTerm := ""
	if con.i == false {
		if len(args) == 0 {
			showArgumentErr()
			os.Exit(1)
		}
		searchTerm := args[0]
		if searchTerm == "" {
			showArgumentErr()
			os.Exit(1)
		}
	}
	d := color.New(color.FgWhite, color.Bold)

	for {
		if con.i {
			d.Println(" " + title)
			searchTerm = prompt("> ")
		}

		resp, err := http.Get(apiURL + url.QueryEscape(searchTerm))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		res, err := ioutil.ReadAll(resp.Body)

		var result data.Result
		err = json.Unmarshal(res, &result)
		if err != nil {
			fmt.Printf("Failed to parse json: %s", err)
			os.Exit(1)
		}

		d.Printf("Synonyms for: %s\n", searchTerm)

		defs := result.ResultSet.DefinitionData.Definitions
		for _, def := range defs {
			sort.Slice(def.Synonyms, func(i, j int) bool {
				return def.Synonyms[i].Similarity > def.Synonyms[j].Similarity
			})
			for j, sym := range def.Synonyms {
				if j < 5 && sym.Similarity > 70 {
					fmt.Printf("  %s\n", trim(sym.Term))
				}
			}
		}
		fmt.Println()

		// Break loop if not interactive mode
		if !con.i {
			break
		}
	}

}

func showArgumentErr() {
	fmt.Println("Error: You have to enter a search term.")
	showExample()
}

func showExample() {
	fmt.Println("Usage: [-n..] syn <search-term>")
	fmt.Printf("Example: syn feel\n\n")
}

func trim(input string) string {
	return strings.Replace(input, "^M", "", -1)
}

func prompt(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return text
}
