![Project](./assets/go_price_monitor_poster.png)

<h3 align="center" >
  Projeto realizado na 2ª Mostra de Pesquisa e extensão no Centro Regional de Espirito Santo do Pinhal
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
    <img alt="Fale comigo no whatsapp - Antonio Costa" src="https://img.shields.io/badge/Whatsapp--%23F8952D?style=social&logo=whatsapp">
  </a>

</p>

# Sumário

- [Linguagem Go](#go_language)
- [Instalando o projeto](#project_install)
- [Interface gráfica (GUI)](#gui)
- [Api WhatsApp Zenvia](#api)
- [Conexão com o banco de dados](#db_conection)
- [Download de ferramenta para conexão com o banco](#db_tool)
- [Criando um banco](#db_create)
- [Editando variaveis de ambiente](#dot_env)
- [Executando o projeto](#run_proj)

## Linguagem Go <a name = "go_language" ></a>

Tenha instalada a versão do golang na versão 1.18 ou superior.

Link para instalação da linguagem [Download and install](https://go.dev/doc/install)

## Instalando o projeto <a name = "project_install" ></a>

Para instalar o projeto junto de suas dependencias, execute os comandos a seguir

```
git clone https://github.com/Antonio-Costa00/Go-Price-Monitor.git
cd Go-Price-Monitor
go mod tidy
```

## Interface gráfica (GUI) <a name = "gui" ></a>

No momento de executar a aplicação, pode haver alguns problemas de compatibilidade relacionados
a aplicação gráfica. Para resolver isso, consulte a documentação oficial do framewrok Fyne
[Getind Started](https://developer.fyne.io/started/)

## API WhatsApp Zenvia <a name = "api" ></a>

Para habilitar o envio de mensagens no WhatsApp utilizando a API zenvia, é necessário
seguir os passos da documentação oficial da API [WhatsApp sender and recipient](https://zenvia.github.io/zenvia-openapi-spec/v2/#section/WhatsApp-sender-and-recipient)

## Conexão com o banco de dados <a name = "db_conection" ></a>

Para executar a aplicação é necessário uma conexão com o banco de dados.

Para o projeto foi utilizado
o banco de dados postgreSQL. Caso voce deseje utilizar um banco diferente, altere a forma de conexão em
`pkg/db`

### Download de ferramenta para conexao com o banco  <a name = "db_tool" ></a>

A forma mais fácil de se criar um banco de dados é com o pgAdmin. O download da ferramenta pode ser feito
pelo link [Download](https://www.pgadmin.org/download/)

### Criando um banco <a name = "db_create" ></a>

Crie um banco de dados com o nome de _price_monitor_

## Editando variaveis de ambiente <a name = "dot_env" ></a>

Troque o nome do arquivo .env.example para .env e preecha as credenciais de acordo com seus valores.

## Executando o projeto <a name = "run_proj" ></a>

Para executar a aplicação
`go run cmd/main.go`

Para gerar um executavel da aplicação
`go build cmd/main.go`

## Author

👤 **Antonio Costa**

- Github: [@Antonio-Costa00](https://github.com/Antonio-Costa00)
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
