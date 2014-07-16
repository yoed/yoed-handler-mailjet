Yo'ed Mailjet
===================

Yo'ed client which send an email via Mailjet v3 API when someone Yo you.

#Installation

You need [`go`](http://golang.org/) package on your machine to get the source

`go get github.com/yoed/yoed-client-mailjet`

#Configuration

Add a `config.json` file in the same folder than the program.

## serverUrl, listen, handle

[See general client configuration](https://github.com/yoed/yoed-client-interface)

##apiKey / apiSecret

Your Mailjet v3 API Key / Secret (you can find them in the admin).

##fromEmail

The email to send the mail from

##toEmail

The email to send mail to

##subject

The email subject

##text

The email text

For both `subject` and `text`, you can use the `%username%` placeholder which will be replaced by the username who have Yo you.
