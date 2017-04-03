package common

import "github.com/mailgun/mailgun-go"
// key-xx
// api.site.com

var mg mailgun.Mailgun;

func InitMailer(){
	mg = mailgun.NewMailgun(GetConfig().MailURL,GetConfig().MailKey,"");
}

func GetMailer()(mailgun.Mailgun){
	return mg;
}

/*

m := mg.NewMessage(  
    "Dwight Schrute <dwight@example.com>",        // From
    "The Printer Caught Fire",                    // Subject
    "We have a problem.",                         // Plain-text body
    "Michael Scott <michael@example.com>",        // Recipients (vararg list)
    "George Schrute <george@example.com>",
    "it-support@example.com",
)

_, _, err := mg.Send(m)

*/