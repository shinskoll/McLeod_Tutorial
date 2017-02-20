package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	Licenses []Licence
}

type Licence struct {
	Id    string
	Valid int
}

var tpl *template.Template

var myFuncMap = template.FuncMap{
	"uppy": strings.ToUpper,
	"ft":   firstThree,
	"hr":   readableTime,
}

func init() {
	tpl = template.Must(template.New("").Funcs(myFuncMap).ParseGlob("*.gotml"))
	//tpl = template.Must(template.ParseGlob("*.gotml"))
}

func main() {
	p1 := doubleZero{
		person: person{
			Name: "Ian Fleming",
			Age:  56,
		},
		Licenses: []Licence{
			Licence{"D2", 2002},
		},
	}

	p2 := doubleZero{
		person: person{
			Name: "Renard",
			Age:  31,
		},
		Licenses: []Licence{
			Licence{"A1", 1997},
			Licence{"A2", 1998},
		},
	}

	p3 := doubleZero{
		person{
			"Todd",
			45,
		},
		[]Licence{
			Licence{"C1", 2016},
		},
	}

	wanna := []doubleZero{p1, p2, p3}

	data := struct {
		Agents []doubleZero
		Tim    time.Time
	}{
		wanna,
		time.Now(),
	}

	makeHomeFromTemplate(data)

	//expl

}

//OPEN FILE
func makeHomeFromTemplate(dz struct {
	Agents []doubleZero
	Tim    time.Time
}) {
	myHome, err := os.Create("home.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer myHome.Close()

	/* GONE IN INIT
	myTpl, err := template.ParseFiles("tpl.gotml")
	if err != nil {
		log.Fatalln(err)
	}*/

	err = tpl.ExecuteTemplate(myHome, "tpl.gotml", dz)
	if err != nil {
		log.Fatalln(err)
	}
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func readableTime(t time.Time) string {
	return t.Format("03:04 le 02-01-2006")
}

func (l Licence) YearsPlus5() int {
	return l.Valid + 5
}
