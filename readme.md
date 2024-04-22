
---

# Go Email Sender

This is a Go program that sends emails with attachments using SMTP.

## Prerequisites

Before running this program, ensure you have the following:

- Go installed on your system. You can download and install it from [here](https://golang.org/dl/).
- Access to an SMTP server to send emails. You may need to obtain SMTP server details such as server address, port, username, and password from your email service provider.

## Installation

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/yourusername/go-email-sender.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-email-sender
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

## Usage

You can run the program with the following command:

```bash
go run main.go -name="Sender Name" -subject="Email Subject" -body="Email Body" -recipient="recipient@example.com" -attachment="./"


```

Replace the placeholders with the appropriate values:

- `"Sender Name"`: The name of the sender.
- `"Email Subject"`: The subject of the email.
- `"Email Body"`: The body/content of the email.
- `"./path/to/attachments"`: The path to the directory containing the files you want to attach. Use `.` if the files are in the current directory.

## Contributing

Contributions are welcome! Please feel free to submit a pull request if you would like to contribute to this project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to customize the README further based on your project's specific requirements or add any additional sections you deem necessary.
