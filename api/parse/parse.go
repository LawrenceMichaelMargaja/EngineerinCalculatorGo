package parse

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	SectionDimensionEnglishName string `json:"section_dimension_english_name" db:"section_dimension_english_name"`
    SectionDimensionEnglishD    string `json:"section_dimension_english_d" db:"section_dimension_english_d"`
    SectionDimensionEnglishB    string `json:"section_dimension_english_b" db:"section_dimension_english_b"`
    SectionDimensionEnglishTW   string `json:"section_dimension_english_tw" db:"section_dimension_english_tw"`
    SectionDimensionEnglishBF   string `json:"section_dimension_english_bf" db:"section_dimension_english_bf"`
    SectionDimensionEnglishTF   string `json:"section_dimension_english_tf" db:"section_dimension_english_tf"`
    SectionDimensionEnglishTB   string `json:"section_dimension_english_tb" db:"section_dimension_english_tb"`
    SectionDimensionEnglishT    string `json:"section_dimension_english_t" db:"section_dimension_english_t"`
    SectionDimensionEnglishR    string `json:"section_dimension_english_r" db:"section_dimension_english_r"`
}

func parse() {
	csvFile, _ := os.Open("sectionDimensionEnglish.csv")
	reader := csv.NewReader(csvFile)
	var people []Person
	var index int = 0
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if index == 0 {
			index++
			continue
		} else {
			people = append(people, Person{
				SectionDimensionEnglishName: line[0],
				SectionDimensionEnglishD: line[1],
				SectionDimensionEnglishB: line[2],
				SectionDimensionEnglishTW: line[3],
				SectionDimensionEnglishBF: line[4],
				SectionDimensionEnglishTF: line[5],
				SectionDimensionEnglishTB: line[6],
				SectionDimensionEnglishT: line[7],
				SectionDimensionEnglishR: line[8],
			})
		}
	}
	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))
}
