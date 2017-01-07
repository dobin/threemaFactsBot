package main

import (
	"crypto/rand"
//	"fmt"
	"log"
	"os"
	"bufio"
	"math/big"
)


func readFile() []string {
    var facts []string

    if file, err := os.Open("facts.txt"); err == nil {
    defer file.Close()

    // create a new scanner and read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	facts = append(facts, scanner.Text())
    }

    // check for errors
    if err = scanner.Err(); err != nil {
      log.Fatal(err)
    }

  } else {
    log.Fatal(err)
  }

	return facts
}

func getEntry(list []string) string {
	l := len(list)
	var ll int64
	ll = int64(l)

	nBig, err := rand.Int(rand.Reader, big.NewInt(ll))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()

	return list[n]
}


func readPrompts() []string {
	prompts := []string{
		"i love chuck norris",
		"chuck <3",
		"wow",
		"miau",
		"your birthday",
		"your mothers maiden name",
		"your password" }
/*	prompts := []string{
		"a", 
		"b"}
*/
	return prompts
}

/*
func main() {
	facts := readFile();

	x := getEntry(facts);
	fmt.Println(x)
}*/
