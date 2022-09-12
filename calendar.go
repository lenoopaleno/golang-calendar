package main

import (
	"bufio"
	"fmt"
	ics "github.com/arran4/golang-ical"
	"log"
	"os"
	"time"
)

func main() {
	ChooseEndDate()
	fmt.Println(cal())
}

func cal() string {

	name, description := ChooseName()
	startdate := ChooseStartDate()

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	event := cal.AddEvent(fmt.Sprintf(name))
	//event.SetCreatedTime(time.Now())
	//event.SetDtStampTime(time.Now())
	//event.SetModifiedAt(time.Now())
	event.SetStartAt(startdate)
	event.SetEndAt(time.Now())
	event.SetSummary("Summary")
	event.SetLocation("Address")
	if description == "" {
		event.SetDescription("Description")
	} else {
		event.SetDescription(description)
	}

	event.SetURL("https://URL/")
	event.AddRrule(fmt.Sprintf("FREQ=YEARLY;BYMONTH=%d;BYMONTHDAY=%d", time.Now().Month(), time.Now().Day()))
	event.SetOrganizer("sender@domain", ics.WithCN("This Machine"))
	event.AddAttendee("reciever or participant", ics.CalendarUserTypeIndividual, ics.ParticipationStatusNeedsAction, ics.ParticipationRoleReqParticipant, ics.WithRSVP(true))
	return cal.Serialize()
}

func ChooseStartDate() time.Time {
	//Date  choose
	fmt.Println("Set a starting point of your event")
	fmt.Println("Provide the year:")
	var yyyy int
	_, err := fmt.Scanf("%v", &yyyy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Provide the month (without '0' in front):")
	var mm time.Month
	_, err = fmt.Scanf("%v", &mm)
	fmt.Println("Provide the day")
	var dd int
	_, err = fmt.Scanf("%v", &dd)
	fmt.Println("What is hour of event? (24h clock)")
	var hour int
	_, err = fmt.Scanf("%v", &hour)
	fmt.Println("And what minute?")
	var min int
	_, err = fmt.Scanf("%v", &min)
	date := time.Date(yyyy, mm, dd, hour, min, 0, 0, time.UTC)
	return date
}

func ChooseEndDate() time.Duration {
	var dur time.Duration
	fmt.Println("Do you want to choose the duration for this event?")
	var ans string
	_, err := fmt.Scanf("%s", &ans)
	if err != nil {
		log.Fatal(err)
	}
	if ans == "Y" {
		var hour string
		var min string
		fmt.Println("Choose your time:")
		fmt.Println("Hour:")
		_, err = fmt.Scanf("%s", &hour)
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
	fmt.Printf("Would you like to add desctription to %s?\n type Y or N", name)
	var ans string
	_, err = fmt.Scanln(&ans)
	if ans == "Y" {
		fmt.Println("Add your description here:")
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		description = scanner.Text()
	} else {
		description = ""
	}
	return name, description
}
