package main

import ("fmt"
		"net/http"
		"github.com/spf13/viper"
)
func main()  {
	//mendaftarkan file config
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.conf")


http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintln(w,"Hello, you've requested: %s\n", r.URL.Path)
})

	http.ListenAndServe(viper.GetString("server.port"),nil)
}
