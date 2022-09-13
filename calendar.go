package main

import (
	"bufio"
	"fmt"
	ics "github.com/arran4/golang-ical"
	"log"
	"os"
	"strings"
	"time"
)

const shortForm = "2006-01-02"

func main() {

	serialized, name := cal()

	f, err := os.Create(name + ".ics")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(serialized)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func cal() (string, string) {

	name, description := ChooseName()
	startDate := ChooseStartDate()

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	event := cal.AddEvent(fmt.Sprintf(name))

	event.SetStartAt(startDate)
	event.SetSummary(name)
	event.SetDuration(ChooseEndDate())
	//event.SetLocation("Address")
	if description == "" {
		event.SetDescription("Description")
	} else {
		event.SetDescription(description)
	}
	return cal.Serialize(), name
}

func ChooseStartDate() time.Time {
	//Date  choose
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Set a starting point of your event")
	fmt.Println("Provide a date in this form: YYYY-MM-DD")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	date := scanner.Text()
	dateParsed, err := time.Parse(shortForm, date)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Do you want to choose starting time of your event?\n Type Y or N: ")
	var hour, min time.Duration
	ans := YesOrNo()
	if ans {
		fmt.Println("Choose starting hour for your event (e.g 11:32)")
		fmt.Scanf("%v:%v", &hour, &min)
	}
	dateParsed = dateParsed.Add(time.Hour * hour)
	dateParsed = dateParsed.Add(time.Minute * min)

	return dateParsed
}

func ChooseEndDate() time.Duration {
	var dur time.Duration
	fmt.Println("Do you want to choose the duration for this event?\n type Y or N")
	ans := YesOrNo()
	if ans {
		var hour string
		var min string
		fmt.Println("Choose your time:")
		fmt.Println("Hour:")
		_, err := fmt.Scanf("%s", &hour)
		hour = hour + "h"
		fmt.Println("Minute:")
		_, err = fmt.Scanf("%s", &min)
		min = min + "m"

		duration := hour + min

		dur, err = time.ParseDuration(duration)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dur
}

func ChooseName() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	var name string
	var description string
	fmt.Println("What's name of your event?")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Would you like to add desctription to %s?\n", name)
	ans := YesOrNo()
	if ans {
		fmt.Println("Add your description here:")
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		description = scanner.Text()
	} else {
		description = ""
	}
	return name, description
}

func YesOrNo() bool {
	var res bool
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type Y or N")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	ans := strings.ToUpper(scanner.Text())
	if ans == "Y" || ans == "YES" {
		res = true
	} else if ans == "N" || ans == "NO" {
		res = false
	} else {
		fmt.Println("You provide wrong answer. Please type 'Y' for yes, or 'N' for no")
		YesOrNo()
	}
	return res
}
