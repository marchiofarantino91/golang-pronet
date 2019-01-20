package main

import ("fmt"
		//"net/http"
		"github.com/spf13/viper"
)
func main()  {
	fmt.Println("in Main")

    viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.conf")
	viper.AddConfigPath("$HOME/go/src/soal6_web/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not found...")
	} else {
		name := viper.GetString("name")
		fmt.Println("Config found, name = ", name)
	}

	fmt.Println("Starting", viper.GetString("appName"))
	
//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
//	fmt.Fprintln(w,"Hello, you've requested: %s\n", r.URL.Path)
//})

//	http.ListenAndServe(":8000",nil)
}
