package main

func main() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Member{email: "abc@gmail.com"}
	observerSecond := &Member{email: "xyz@gmail.com"}

	shirtItem.register(observerFirst, observerSecond)

	shirtItem.updateAvailability()
	shirtItem.deregister(0)
	shirtItem.updateAvailability()
}
