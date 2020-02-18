After cloning the repository, install all go dependences by going to the root/src folder and executing the following command

go get ./...


if in debug mode go to models/configuration.go and change 


content, err := ioutil.ReadFile("config/connection.json")

TO

content, err := ioutil.ReadFile("../config/connection.json")

