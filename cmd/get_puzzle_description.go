package main

import (
	"fmt"
	"log"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/mahalde/advent-of-code/utils/conv"
	"github.com/mahalde/advent-of-code/utils/req"
)

func main() {
	intArgs := conv.ToIntSlice(os.Args[1:])
	year, day := intArgs[0], intArgs[1]

	resBody := req.GetPuzzleDescription(year, day)

	defer resBody.Close()

	doc, err := goquery.NewDocumentFromReader(resBody)

	if err != nil {
		log.Fatal(err)
	}

	converter := md.NewConverter("", true, nil)

	doc.Find(".day-desc").Each(func(i int, s *goquery.Selection) {
		markdown := converter.Convert(s)
		fmt.Println(markdown)
	})
}
