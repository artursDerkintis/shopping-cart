package main

func main() {
	database := Database{}
	err := database.Connect("mongodb+srv://artursderkintis:*****@cluster0.93gdem2.mongodb.net/")
}
