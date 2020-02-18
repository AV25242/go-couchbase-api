Step 1 : Clone this repository in GOPATH directory.

Step 2 : After cloning the repository, install all go dependences by going to the root/src folder and executing the following command

go get ./...


Optional Step : 

if in debug mode go to models/configuration.go and change 


content, err := ioutil.ReadFile("config/connection.json")

TO

content, err := ioutil.ReadFile("../config/connection.json")

