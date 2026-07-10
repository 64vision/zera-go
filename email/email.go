package email

import (
	"fmt"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type EmailReceipt struct {
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
	TransDT   string `json:"Trans_dt"`
	BookingNo string `json:"booking_no"`
	ReceiptNo string `json:"receipt_no"`
	Payment   string `json:"payment"`
	Subject   string `json:"subject"`
	Type      string `json:"type"`
}

func (entry *EmailReceipt) HtmlTemplate() string {
	HtmlBody := `<p style="color: #000000; margin: auto; padding: 220px 50px 50px 50px; font-size: 18px; background: #fff url('https://eboracay.s3.ap-southeast-1.amazonaws.com/receiptbg2.png') no-repeat top center; background-size: 100% 100%; width: 300px; height: 400px;">
		Amount Paid: ` + entry.Amount + ` <br/><br/>
		Payment: ` + entry.Payment + `<br/><br/>
		Date/Time: ` + entry.TransDT + `<br/><br/>
		Ref. No:  ` + entry.BookingNo + `<br/><br/>
		</p>
		<p style="text-align: center; margin-top:10px; font-style: italic;">
		You may request your official receipt for the convenience fee thru <br/>email at finance@surepayinc.com.ph. Thank you.
		</p>
		<p style="text-align: center; margin-top:10px; font-style: italic;">THANK YOU FOR USING EBORACAY</p>`
	return HtmlBody
}

func (entry *EmailReceipt) TextTemplate() string {
	HtmlBody := `
		eBoracay payment confirmed. Amount paid: ` + entry.Amount + `, Booking Ref. No:  ` + entry.BookingNo + `
		, Receipt No:  ` + entry.ReceiptNo + `, Payment:  ` + entry.Payment + `, Date/Time:  ` + entry.TransDT +
		``
	return HtmlBody
}

func SendVerificationCodeForgot(code string, to_email string) bool {

	// Create a new session in the us-west-2 region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.
	subject := "eBoracay forgot password code"
	body := "<p>Hi!,</p><p>Thank you for using eBoracay. Here is your  code.</p>" +
		"<p style='padding: 15px;  color: #000000; font-weight: bold; font-sise: 28px; margin-bottom: 20px; margin-top: 20px; margin-left: 50px;'><b> Activation Code: " + code + "</b><p>" +
		"<p style='color: #999; margin-top: 50px;'>Thank you!<br /> eBoracay Team<p>"
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(SesID, SesSecret, ""),
	},
	)
	if err != nil {
		// Handle the error
		panic(err)
	}

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(to_email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(body),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(SenderName + " <" + SenderEmail + ">"),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return true
	}

	fmt.Println("SendVerificationCode Email Sent to address: " + to_email)
	fmt.Println(result)
	return true
}

func Send(subject string, body string, to_email string, sender_email string, sender_name string) bool {

	// subject := subject
	// body := body
	// body := "<p>Hi!,</p><p>Thank you for using eBoracay. Here is your activation code.</p>" +
	// 	"<p style='padding: 15px;  color: #000000; font-weight: bold; font-sise: 28px; margin-bottom: 20px; margin-top: 20px; margin-left: 50px;'><b> Activation Code: " + code + "</b><p>" +
	// 	"<p style='color: #999; margin-top: 50px;'>Thank you!<br /> eBoracay Team<p>"
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(SesID, SesSecret, ""),
	},
	)
	if err != nil {
		// Handle the error
		panic(err)
	}

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(to_email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(body),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender_name + " <" + sender_email + ">"),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return true
	}

	fmt.Println("SendVerificationCode Email Sent to address: " + to_email)
	fmt.Println(result)
	return true
}
