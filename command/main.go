package main

func main() {
	tv := &TV{}

	//привязываем девайс к комманде
	onCommand := &onCommand{device: tv}
	offCommand := &offCommand{device: tv}

	//привязываем комманду к кнопке
	onButton := &Button{command: onCommand}
	offButton := &Button{command: offCommand}

	onButton.press()
	offButton.press()
}
