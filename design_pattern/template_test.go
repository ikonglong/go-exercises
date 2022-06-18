package design_pattern

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	// otp := otp{}

	// smsOTP := &sms{
	//  otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	//  otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	smsOTP := &sms{}
	o := otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)

	var x interface{} = "foo"
	x = emailOTP
	// Invalid type switch guard: v := emailOTP.(type) (non-interface type *email on the left)
	//switch v := emailOTP.(type) {
	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil") // here v has type interface{}
	case int:
		fmt.Println("x is", v) // here v has type int
	case bool, string:
		fmt.Println("x is bool or string") // here v has type interface{}
	case otp:
		fmt.Println("x is otp")
	case email:
		fmt.Println("x is email")
	case iOtp:
		fmt.Println("x is iOtp")
	default:
		fmt.Println("type unknown") // here v has type interface{}
	}
}
