package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/template"
	"time"
)

// Commit test
type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	Licenses []licence
}

type licence struct {
	Ident string
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
		Licenses: []licence{
			licence{"D2", 2002},
		},
	}

	p2 := doubleZero{
		person: person{
			Name: "Renard",
			Age:  31,
		},
		Licenses: []licence{
			licence{"A1", 1997},
			licence{"A2", 1998},
		},
	}

	p3 := doubleZero{
		person{
			"Todd",
			45,
		},
		[]licence{
			licence{"C1", 2016},
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

	// TCP LISTENER
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error creating listener", err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I Heard you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Code never reached")
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

func (l licence) YearsPlus5() int {
	return l.Valid + 5
}
