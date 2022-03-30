package main

import "fmt"

func main() {
	title := "Welcome to Hulu!"
	fmt.Println(title)

	var firstName string
	var lastName string
	var addressNum string
	var addressStreet string
	var addressType string
	var city string
	var state string
	var zip string
	var totalCost float32 = 0

	fmt.Printf("Enter your name: ")
	fmt.Scan(&firstName, &lastName)
	fmt.Printf("Enter your street address: ")
	fmt.Scan(&addressNum, &addressStreet, &addressType)
	fmt.Printf("Enter your city: ")
	fmt.Scan(&city)
	fmt.Printf("Enter your state: ")
	fmt.Scan(&state)
	fmt.Printf("Enter your zip: ")
	fmt.Scan(&zip)
	var fullAddress = addressNum + " " + addressStreet + " " + addressType + ", " + city + ", " + state + ", " + zip
	fmt.Printf("Welcome, %v %v!\n", firstName, lastName)

	summary := fmt.Sprintf("Order Summary:\nName:\n%v %v\nAddress:\n%v\nAccount Number:\n1753485263541298\nPackage:\n", firstName, lastName, fullAddress)

	var pkg uint
	fmt.Println("Which package would you like:\n1. Hulu (No Ads) + Live TV, Disney+ & ESPN for $75.99\n2. Hulu (No Ads) + Live TV for $60\n3. Hulu (with Ads) + Live TV for $35\nEnter the product number: ")
	fmt.Scan(&pkg)
	switch pkg {
	case 1:
		totalCost += 75.99
		summary += "Hulu (No Ads) + Live TV, Disney+ & ESPN"
		break
	case 2:
		totalCost += 60.00
		summary += "Hulu (No Ads) + Live TV for $60"
		break
	case 3:
		totalCost += 35.00
		summary += "Hulu (with Ads) + Live TV for $35"
		break
	}

	summary += "\nAdd-Ons:\n"

	var addOn1 uint
	fmt.Println("Great choice!")
	fmt.Println("Would you like to add\n1. Enhanced Cloud DVR 300 minutes & Unlimited Screens for $14.99 per month\n2. 2 Screens with 100 hours of DVR for $9.99\n3. $0 for no DVR and only 1 screen")
	fmt.Scan(&addOn1)
	var addOn2 uint
	fmt.Println("Last question.")
	fmt.Println("Would you like to add\n1. ShowTime for $7.99\n2. HBO for $14.99\n3. Both for $22.98\n4. No, Thanks.")
	fmt.Scan(&addOn2)

	switch addOn1 {
	case 1:
		totalCost += 14.99
		summary += "Enhanced Cloud DVR 300 minutes & Unlimited Screens\n"
		break
	case 2:
		totalCost += 9.99
		summary += "2 Screens with 100 hours of DVR\n"
		break
	case 3:
		summary += "No DVR, 1 screen\n"
		break
	}

	switch addOn2 {
	case 1:
		totalCost += 7.99
		summary += "Showtime\n"
		break
	case 2:
		totalCost += 14.99
		summary += "HBO\n"
		break
	case 3:
		totalCost += 22.98
		summary += "ShowTime\nHBO\n"
		break
	case 4:
		summary += "No Channels\n"
		break
	}

	fmt.Println("Thank you!")
	summary += fmt.Sprintf("Monthly Fee: $%.2f\nPayment Due:\n30th", totalCost)
	fmt.Printf("\n" + summary)
}
