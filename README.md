# MembersClub
“Member Club” is a web-application which allows to add members to the Club and view list of club members.

## Application on Heroku
https://peaceful-peak-29287.herokuapp.com/

## Get started
### Instal golang
>Go to https://golang.org/doc/install and follow the instructions regarding your OS.

### Install git
>Go to https://git-scm.com/book/en/v2/Getting-Started-Installing-Git and follow the instructions.

### Clone repository
Open terminal/console and clone the project:

>$ git clone https://github.com/DmytroMaslov/MembersClub

### Go to the project
>cd memberclub

### Run project
>$ go run .

### Check out the application
Make request using curl:
* >$ curl 'localhost:8080/getAllMember'
* >$ curl 'localhost:8080/addMember'

### Run test
use command
>$ go test ./...

## API
### Endpoint for adding new member:
> /addMember

return: json with new member (Name, Email, Registration time)
>example:

>request:
/addMember?name=Bob&email=bob@gmail.com

>response:
{"name":"Bob","email":"bob@gmail.com" "registration_date":"2021-12-21T15:01:08.424568Z"}

### Endpoint for geting list of all members
>/getAllMember

return:
json with array of all members