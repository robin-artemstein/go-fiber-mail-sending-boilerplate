# Go Fiber Mail Sending Boilerplate

Welcome to the repository, I will show you how to send email with a well-known back end framework called Go Fiber.

## Clone the repository
```
git clone git@github.com:robin-artemstein/go-fiber-mail-boilerplate.git
```
## Change the default port

Edit `main.go`, then find these lines of codes.
```
func main() {
	// Create a new Fiber app. Think of this as the main engine for our API.
	app := fiber.New()

	// Define the POST route. This is the endpoint our front end will call.
	app.Post("/api/sendmail/", sendMailHandler)

	// Start the server on port 3003.
	// You can access this at http://localhost:3003
	log.Fatal(app.Listen(":3003"))
}
```
You can change the default port from 3003 to whatever you want.

## Email setting

In this repository, I use Yahoo mail as an example, you can change to other mail services which support SMTP or self-hosted SMTP mail servers, you can find relevant codes  in `main.go`.

## Set Environment Variables

For security, I don't recommend to put the email address and password directly in the code, I use environment variables instead.

For Yahoo mail, **you must generate an "App Password"**. Due to security updates, you cannot use your regular login password for third-party apps like this.

1.  Log in to your Yahoo account.
2. Go to **Account Info** -> **Account Security**.
3. Find **App password** and click **Generate app password**.
4. Select "Other App", give it a name (e.g., "Go Contact Form"), and generate the password.
5. Copy this generated password (it will be a 16-character string without spaces).
    
Now, set the variables in your terminal before running the app.

You have to switch to the project directory with `cd` command first.

On Linux, BSD, and macOS, you need to run these commands.
```
export EMAIL_USER="otu168xchen@yahoo.com"
export EMAIL_PASS="your-generated-app-password-here"
```
On Windows, you need to run these commands in command prompt.
```
set EMAIL_USER="otu168xchen@yahoo.com"
set EMAIL_PASS="your-generated-app-password-here"
```

## Run this boilerplate

You need to download dependencies first.
```
go mod tidy
```
Then run this boilerplate
```
go run main.go
```
## Test with Postman or Hoppscotch

Just follow these instruction.
1. Open Postman.
2. Set the request method to **POST**.
3. Enter the URL: `http://localhost:3003 /api/sendmail/`. (If you have changed port number in `main.go` then use the new port number.)
4. Go to the **Body** tab, select **raw**, and choose **JSON** from the dropdown.
5. Paste the following JSON payload then click on Send button:
```
{
    "name": "Joe Doe",
    "email": "joe-doe@desire.com,
    "title": "Commission request",
    "content": "Hello, I'd like to commission you to do web design for our company. Please reply to me ASAP."
}
```
You should get a `200 OK` response with `{"message": "Email sent successfully!"}`, and an email will arrive in your `otu168xchen@yahoo.com` inbox.
