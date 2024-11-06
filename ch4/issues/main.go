// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	LESS_THAN_A_MONTH = "less_than_a_month"
	LESS_THAN_A_YEAR  = "less_than_a_year"
	MORE_THAN_A_YEAR  = "more_than_a_year"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	categories := make(map[string][]*github.Issue)

	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()
	for _, item := range result.Items {
		switch{
		case item.CreatedAt.After(now.Add(-1 * 60 * 24 * 30 * time.Minute)):
			categories[LESS_THAN_A_MONTH] = append(categories[LESS_THAN_A_MONTH], item)
		case item.CreatedAt.After(now.Add(-1 * 60 * 24 * 365 * time.Minute)):
			categories[LESS_THAN_A_YEAR] = append(categories[LESS_THAN_A_YEAR], item)
		case item.CreatedAt.Before(now.Add(-1 * 60 * 24 * 365 * time.Minute)):
			categories[MORE_THAN_A_YEAR] = append(categories[MORE_THAN_A_YEAR], item)
		}
		// fmt.Printf("#%-5d %9.9s %.55s\n",
		// 	item.Number, item.User.Login, item.Title)
	}

	for category, issues := range categories {
		fmt.Printf("%s issues:\n", category)
		for _, issue := range issues {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				issue.Number, issue.User.Login, issue.Title)
		}
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
