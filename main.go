package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	searcher := ShakespeareSearch{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch(searcher))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type ShakespeareSearch struct {
	Works []Searcher
}

type Searcher struct {
	WorkTitle string
	CompleteWork string
	SuffixArray   *suffixarray.Index
}

type WorkParser struct {
	worksMap map[string]bool
}

type SearchResponse struct {
	Results []ResponseUnit `json:"results"`
}

type ResponseUnit struct {
	Work string `json:"work"`
	Match string `json:"match"`
}

func handleSearch(searcher ShakespeareSearch) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("missing search query in URL params"))
			return
		}
		response := searcher.Search(query[0])
		jData, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("json marshal failure"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	}
}

func (s *ShakespeareSearch) Load(fn string) (err error) {
	works := WorkParser{}
	works.Load()

	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	// https://stackoverflow.com/a/41741702/2295672
	reader := bufio.NewReader(file)
	var line string
	var currentWorkText string
	var currentWorkName string
	contentStarted := false
	contentEnd := false
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		// actual content is finished
		if line == "  FINIS\r\n" {
			contentEnd = true
		}

		// Process the line here.
		if works.Search(line) || contentEnd {
			// past work had started or content has ended
			if (contentStarted || contentEnd) {
				searcher := Searcher{}
				searcher.CompleteWork = currentWorkText
				dataLowerCased := strings.ToLower(currentWorkText)
				searcher.SuffixArray = suffixarray.New([]byte(dataLowerCased))
				searcher.WorkTitle = currentWorkName
				s.Works = append(s.Works, searcher)
			}
			// if contentEnd, close
			if contentEnd {
				break
			}
			// start new work
			contentStarted = true
			currentWorkText = ""
			currentWorkName = strings.TrimSuffix(line, "\r\n")
		} else if contentStarted {
			currentWorkText += line
		}

		if err != nil {
			break
		}
	}
	if !contentEnd && err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	return
}

func (s *ShakespeareSearch) Search(query string) SearchResponse {
	queryLowerCased := strings.ToLower(query)
	queryRegex := regexp.MustCompile("(?i)(" + query + ")")
	response := SearchResponse{}
	for _, work := range s.Works {
		idxs := work.SuffixArray.Lookup([]byte(queryLowerCased), -1)
		lastCoveredIdxMax := 0
		sort.Ints(idxs)
		for _, idx := range idxs {
			if idx < lastCoveredIdxMax {
				// if idxMax already covered next match, ignore it and go to next
				continue
			}
			idxMin := idx-250
			if idxMin < 0 {
				idxMin = 0
			}
			idxMax := idx+250
			if idxMax > len(work.CompleteWork) {
				idxMax = len(work.CompleteWork)
			}
			lastCoveredIdxMax = idxMax
			result := work.CompleteWork[idxMin:idxMax]
			resultHighlighted := queryRegex.ReplaceAllString(result, "<mark>${1}</mark>")
			responseUnit := ResponseUnit{}
			responseUnit.Match = resultHighlighted
			responseUnit.Work = work.WorkTitle
			response.Results = append(response.Results, responseUnit)
		}
	}
	return response
}

func (w *WorkParser) Search(val string) bool {
	if _, ok := w.worksMap[val]; ok {
		return true
	}
	return false
}

func (w *WorkParser) Load() {
	works := []string{"THE SONNETS",
		"ALL’S WELL THAT ENDS WELL",
		"ANTONY AND CLEOPATRA", // modified
		"AS YOU LIKE IT",
		"THE COMEDY OF ERRORS",
		"THE TRAGEDY OF CORIOLANUS",
		"CYMBELINE",
		"THE TRAGEDY OF HAMLET, PRINCE OF DENMARK",
		"THE FIRST PART OF KING HENRY THE FOURTH",
		"THE SECOND PART OF KING HENRY THE FOURTH",
		"THE LIFE OF KING HENRY V", // modified
		"THE FIRST PART OF HENRY THE SIXTH",
		"THE SECOND PART OF KING HENRY THE SIXTH",
		"THE THIRD PART OF KING HENRY THE SIXTH",
		"KING HENRY THE EIGHTH",
		"KING JOHN",
		"THE TRAGEDY OF JULIUS CAESAR",
		"THE TRAGEDY OF KING LEAR",
		"LOVE’S LABOUR’S LOST",
		"MACBETH", // modified
		"MEASURE FOR MEASURE",
		"THE MERCHANT OF VENICE",
		"THE MERRY WIVES OF WINDSOR",
		"A MIDSUMMER NIGHT’S DREAM",
		"MUCH ADO ABOUT NOTHING",  // modified in text
		"OTHELLO, THE MOOR OF VENICE", // modified
		"PERICLES, PRINCE OF TYRE",
		"KING RICHARD THE SECOND",
		"KING RICHARD THE THIRD",
		"THE TRAGEDY OF ROMEO AND JULIET",
		"THE TAMING OF THE SHREW",
		"THE TEMPEST",
		"THE LIFE OF TIMON OF ATHENS",
		"THE TRAGEDY OF TITUS ANDRONICUS",
		"THE HISTORY OF TROILUS AND CRESSIDA",
		"TWELFTH NIGHT: OR, WHAT YOU WILL", // modified
		"THE TWO GENTLEMEN OF VERONA",
		"THE TWO NOBLE KINSMEN", // modified in text
		"THE WINTER’S TALE",
		"A LOVER’S COMPLAINT",
		"THE PASSIONATE PILGRIM",
		"THE PHOENIX AND THE TURTLE",
		"THE RAPE OF LUCRECE",
		"VENUS AND ADONIS"} // modified in text

	w.worksMap = map[string]bool{}
	for _, v := range works {
		w.worksMap[v+"\r\n"] = true
	}
}
