package main
import (
    "flag"
)

func main(){
    var folder string
    var email string
    flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
    flag.StringVar(&email, "email", "anantukid@gmail.com", "the email to scan")
    flag.Parse()
    
    if folder != "" {
        scan(folder)
        return
    }

    stats(email)
}
