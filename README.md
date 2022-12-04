![go_price_monitor_poster](https://user-images.githubusercontent.com/87380701/205478048-81c0e424-3b95-4cc6-9ef6-27087249cdc6.png)

<h3 align="center" >
  The project was carried out at the 2nd Research and Extension Exhibition at Centro Regional de Espirito Santo do Pinhal university.
</h3>

<p align="center">

  <a href="https://github.com/Antonio-Costa00" target="_blank">
    <img alt="Made by Antonio Costa" src="https://img.shields.io/badge/made%20by-Antonio_Costa-informational">
  </a>
  <a href="https://github.com/Antonio-Costa00" target="_blank" >
    <img alt="Github - Antonio Costa" src="https://img.shields.io/badge/Github--%23F8952D?style=social&logo=github">
  </a>
  <a href="https://www.linkedin.com/in/antonio-costa-099ab0182/" target="_blank" >
    <img alt="Linkedin - Antonio Costa" src="https://img.shields.io/badge/Linkedin--%23F8952D?style=social&logo=linkedin">
  </a>
  <a href="mailto:antonio.costa.dev@gmail.com" target="_blank" >
    <img alt="Email - Antonio Costa" src="https://img.shields.io/badge/Email--%23F8952D?style=social&logo=gmail">
  </a>
  <a href="https://api.whatsapp.com/send?phone=5519992685736"
        target="_blank" >
    <img alt="Talk to me on WhatsApp - Antonio Costa" src="https://img.shields.io/badge/Whatsapp--%23F8952D?style=social&logo=whatsapp">
  </a>

</p>

# Index

- [Go Language](#go_language)
- [Installing the project](#project_install)
- [Graphical User Interface (GUI)](#gui)
- [WhatsApp Zenvia API](#api)
- [Database connection](#db_connection)
- [Download the tool for database connection](#db_tool)
- [Creating a database](#db_create)
- [Editing environment variables](#dot_env)
- [Running the project](#run_proj)

## Go language <a name="go_language"></a>

Have Go version 1.18 or higher installed

Language installation link [Download and install](https://go.dev/doc/install)

## Installing the project <a name="project_install"></a>

To install the project along with its dependencies, run the following commands

```
git clone https://github.com/Antonio-Costa00/Go-Price-Monitor.git
cd Go-Price-Monitor
go mod tidy
```

## Graphical user interface (GUI) <a name="gui"></a>

At the time of running the application, there may be some compatibility issues related to the graphics application. To resolve this, refer to the official Fyne framework documentation
[Get Started](https://developer.fyne.io/started/)

## WhatsApp Zenvia API <a name="api"></a>

To enable sending messages on WhatsApp using the Zenvia API, you need to
follow the steps of the official API documentation [WhatsApp sender and recipient](https://zenvia.github.io/zenvia-openapi-spec/v2/#section/WhatsApp-sender-and-recipient)

## Database connection <a name="db_connection"></a>

To run the application, a connection to the database is required.

For the project was used
the PostgreSQL database was. If you want to use a different bank, change the connection method in
`pkg/db`

### Download the tool for database connection <a name="db_tool"></a>

The easiest way to create a database is with pgAdmin. The tool can be downloaded
via the link [Download](https://www.pgadmin.org/download/)

### Creating a database <a name="db_create"></a>

Create a database named _price_monitor_

## Editing environment variables <a name="dot_env"></a>

Change the file name from .env.example to .env and fill in the credentials according to your values.

## Running the project <a name="run_proj"></a>

To run the application
`go run cmd/main.go`

To build an executable of the application
`go build cmd/main.go`

## Author

👤 **Antonio Costa**

- GitHub: [@Antonio-Costa00](https://github.com/Antonio-Costa00)
- Linkedin: [@Antonio Costa](https://www.linkedin.com/in/antonio-costa-099ab0182/)

## Contributing

Contributions, issues and feature requests are welcome!

## Show your support

Give a ⭐️ if this project helped you!

## License

Copyright © 2022 [Antonio Costa](https://github.com/Antonio-Costa00).<br />
This project is [MIT](https://github.com/Antonio-Costa00/Go-Price-Monitor/blob/master/LICENSE) licensed.

---

Made by :blue_heart: by [Antonio Costa](https://github.com/Antonio-Costa00)
