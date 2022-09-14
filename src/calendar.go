package main

import (
	"bufio"
	"fmt"
	ics "github.com/arran4/golang-ical"
	"log"
	"os"
	"os/exec"
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
	cmd := exec.Command("~/" + name + ".ics")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func cal() (string, string) {

	name, description := ChooseName()

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	event := cal.AddEvent(fmt.Sprintf(name))

	event.SetStartAt(ChooseStartDate())
	event.SetSummary(name)
	event.SetDuration(ChooseEndDate())
	event.SetLocation(ChooseAddress())
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
	scanner.Scan()
	date := scanner.Text()
	dateParsed, err := time.Parse(shortForm, date)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Do you want to choose starting time of your event?")
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
	scanner := bufio.NewScanner(os.Stdin)

	var durationParsed time.Duration
	var err error
	fmt.Println("Do you want to choose the duration for this event?")
	ans := YesOrNo()
	if ans {
		fmt.Println("Choose duration of your event (e.g 10h20m; 24m; 2h)")
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		duration := scanner.Text()
		durationParsed, err = time.ParseDuration(duration)
		if err != nil {
			log.Fatal(err)
		}
	}
	return durationParsed
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

func ChooseAddress() string {
	var address string
	fmt.Println("Do you want to choose address for this event?")
	ans := YesOrNo()
	if ans {
		fmt.Println("Provide address fo your event: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		address = scanner.Text()
		return address
	}
	return address
}

/*
	TODO: Automatic execution of .ics file
*/
