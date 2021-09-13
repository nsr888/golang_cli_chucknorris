package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)


// Joke is struct for response from api.chucknorris.io.
type Joke struct {
	Value string `json:"value"`
}

// get JSON data from given URL.
func getJSON(url string) ([]byte, error) {
	const timeoutSeconds = 10

	client := http.Client{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		return []byte{}, getErr
	}

	dataJSON, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return []byte{}, readErr
	}

	res.Body.Close()

	return dataJSON, nil
}

// get string with joke from JSON.
func getStringFromJSON(dataJSON []byte) (string, error) {
	joke1 := Joke{Value: ""}
	jsonErr := json.Unmarshal(dataJSON, &joke1)

	if jsonErr != nil {
		return "", jsonErr
	}

	return joke1.Value, nil
}

// print to stdout random joke.
func getRandomJoke() error {
	dataJSON, err := getJSON("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return err
	}

	joke, err := getStringFromJSON(dataJSON)
	if err != nil {
		return err
	}

	fmt.Println(joke)

	return nil
}

// get available categories from API.
func getCategories() ([]string, error) {
	var arr []string

	dataJSON, err := getJSON("https://api.chucknorris.io/jokes/categories")
	if err != nil {
		return []string{}, err
	}

	jsonErr := json.Unmarshal(dataJSON, &arr)

	if jsonErr != nil {
		return []string{}, fmt.Errorf("failed to Unmarshal dataJSON: %w", jsonErr)
	}
	return arr, nil
}

// save n jokes to separate <categories> files.
func saveDump(n int) error {
	arr, err := getCategories()
	if err != nil {
		return err
	}

	for _, v := range arr {
		f, err := os.OpenFile(v+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600)
		if err != nil {
			return fmt.Errorf("failed to OpenFile %v: %w", v+".txt", err)
		}

		for i := 0; i < n; i++ {
			dataJSON, err := getJSON("https://api.chucknorris.io/jokes/random?category=" + v)
			if err != nil {
				return err
			}

			text, err := getStringFromJSON(dataJSON)
			if err != nil {
				return err
			}

			if _, err = f.WriteString(text + "\n"); err != nil {
				return err
			}
		}
		f.Close()
	}
	return nil
}

// print usage of programm.
func printUsage() {
	fmt.Println("Usage:\n\n\tjoker random\n\tjoker dump\n\tjoker dump -n <number>")
}

func main() {
	// if only one argument case.
	if len(os.Args) == 2 {
		const defaultNumbersOfJokes = 5

		switch os.Args[1] {
		case "random":
			err := getRandomJoke()
			if err != nil {
				log.Fatal(err)
			}
		case "dump":
			err := saveDump(defaultNumbersOfJokes)
			if err != nil {
				log.Fatal(err)
			}
		}

		return
	}

	// if number of jokes flag provided.
	if len(os.Args) == 4 && os.Args[2] == "-n" {
		numbersOfJokes, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}

		if numbersOfJokes > 0 {
			err := saveDump(numbersOfJokes)
			if err != nil {
				log.Fatal(err)
			}

			return
		}
	}

	// default behaviour
	printUsage()
}
