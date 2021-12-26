package packageone

//Private varible only defining varible name with lowcase
var priveteVar = "I am private"

//Public variable only defining variable name with capital case
var PublicVar = "I am public (or exported)"

func notExported() {

}

//The function is available since the name was defined first letter of name in capital
func Exported() {

}
