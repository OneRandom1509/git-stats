package main
import (
    "flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

func scanGitFolders(folders []string, folder string) []string{    
    folder = strings.TrimSuffix(folder, "/")

    f, err := os.Open(folder)
    if err!=nil{
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err!=nil{
        log.Fatal(err)
    }

    var path string

    for _, file := range files{
        if file.IsDir(){
            path = folder + "/" + file.Name()
            if file.Name() == ".git"{
                path = strings.TrimSuffix(path, "/.git")
                fmt.Println(path)
                folders = append(folders,path)
                continue
            }
            if file.Name() == "vendor" || file.Name() == "node_modules"{
                continue
            }
            folders = scanGitFolders(folders, path)
        }
    }
    return folders
}

func recursiveScanFolder(folder string) []string{
    return scanGitFolders(make([]string, 0), folder)
}

func getDotFilePath() string{
    usr, err := user.Current()
    if err != nil{
        log.Fatal(err)
    }

    dotFile := usr.HomeDir + "/.gogitlocalstats"

    return dotFile
}

func scan(folder string){
    fmt.Printf("Found folders:\n\n")
    repositories := recursiveScanFolder(folder)
    filePath := getDotFilePath()
    addNewSliceElementsToFile(filePath, repositories) // to be added
    fmt.Printf("\n\nSuccessfully added\n\n")
}

func stats(email string) {
    print("stats")
}

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
