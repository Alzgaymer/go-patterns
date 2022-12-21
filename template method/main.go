package main

import "fmt"

func main() {

	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o.setOTP(emailOTP)

	o.genAndSendOTP(4)
}
