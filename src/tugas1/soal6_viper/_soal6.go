package main

import ("fmt"
		"net/http"
		"github.com/spf13/viper"
)
func main()  {
viper.SetConfigType("json")
viper.AddConfigPath(".")
viper.SetConfigName("app.conf")
	err := viper.ReadInConfig()
	if err != nil {  //cek apakah ditemukan atau tidak
		fmt.Println("Config not found...")
	} else {
		name := viper.GetString("appName")
		fmt.Println("Config found, name = ", name)
		
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
			fmt.Fprintln(w,"Hello, you've requested: %s\n", r.URL.Path)
		})

	}
	port := viper.GetString("server.port")
	fmt.Println("Running port: ",port)

	http.ListenAndServe(port,nil)

}
