package main

import (
	"fmt"
	"time"
)

type messageToSend struct {
	phone  int
	person string
}

type inviteModel struct {
	material string
	messageToSend
}

// functions, of cours, need to be declared before the endpoint
func tellMyHistory(name string, age int) int {
	fmt.Printf("Hello, my name is %v and I am %v", name, age)
	return age + 30
}

func sendMessage(struct_message messageToSend) string {
	// simple control check
	ret_val := ""
	if struct_message.phone == 0 {
		ret_val = "Hello " + struct_message.person
		return ret_val
	}
	ret_val = "Hello " + struct_message.person + ",\nyour new phone is: " + fmt.Sprint(struct_message.phone)
	return ret_val
}

func inviteSendMessage(inviteMessage inviteModel) string {

	childStruct := messageToSend{}
	childStruct.person = inviteMessage.person
	childStruct.phone = inviteMessage.phone
	ret_val := sendMessage(childStruct)

	return "invite card made by: " + inviteMessage.material + "\n" + ret_val
}

// rewrite inviteSendMessage as a method
func (c messageToSend) methodSendMessage(material string) string {
	ret_val := sendMessage(c)

	return "invite card made by: " + material + "\n" + ret_val
}

func main() {
	name := "vitor"
	var age int = 30
	var float_age float64 = 1995.00 / float64(age)
	hasColor := true
	fmt.Printf("Hello %v, this is %T", name, name)
	fmt.Printf("\nyou are %v years old", age)
	fmt.Printf("===============\n")
	fmt.Printf(
		"My float age is %.2f and its type is %T, and hasColor: %v is %T",
		float_age,
		float_age,
		hasColor,
		hasColor)

	// reassing variables
	hasColor = false
	fmt.Printf("\nyou are %v years old", hasColor)

	// constraints
	const minuteLife int = 60
	now_minute := time.Now().Minute()

	// control flow
	if minuteLife <= now_minute && minuteLife > 45 {
		fmt.Printf("\nThe current minnute is: %v", now_minute)
		now_minute += 60
		fmt.Printf("\nNow its value is %v", now_minute)
	} else if now_minute >= 30 {
		now_minute += 30
	} else {
		now_minute -= 60
	}

	fmt.Printf("\nNow its value is %v", now_minute)

	// if the variable is only used on if statment (syntx)
	if now_minute_asyntx := time.Now().Minute(); minuteLife >= now_minute_asyntx {
		fmt.Printf("New Syntx")
		fmt.Printf("now_minute_asyntx: %v is only accessible in this scope", now_minute_asyntx)
	}

	// executing the function
	plus_sixty := tellMyHistory("Vitor", 30)
	fmt.Printf("\n Testing the return of function %v\n", plus_sixty)

	// calling the second function and ignoring with _
	wrap_function, _ := nameAndAge("Vitor", 29)
	fmt.Printf("%v value from nameAndName function", wrap_function)

	// working with structs (similar to python dic)
	message := messageToSend{
		person: "Milena",
		phone:  930410032,
	}
	message_ready := sendMessage(message)
	fmt.Printf("Testing:\n%v\n", message_ready)

	// initiating from another way (like it)
	short_message := messageToSend{}
	short_message.person = "Paul"
	short_message.phone = 99822838

	short_message_ready := sendMessage(short_message)
	fmt.Printf("Testing:\n%v\n", short_message_ready)

	// instanciating new inviteModel
	inviteLove := inviteModel{}
	inviteLove.material = "Paper"
	inviteLove.person = "Milena"
	inviteMessage := inviteSendMessage(inviteLove)
	fmt.Printf("Testing:\n%v\n", inviteMessage)

	fmt.Println("=====================")
	vitorMessage := messageToSend{}
	vitorMessage.person = "Vitor"
	vitorMessage.phone = 930410030

	fmt.Println(vitorMessage.methodSendMessage("paper"))

}

// actually what matter is only the endpoint, so I am able to write a function here
func nameAndAge(name string, age int) (string, int) {
	return name, age
}
