package manifest

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"os"
	"encoding/csv"
	"time"
	"crypto/sha256"
	"io"
	"encoding/hex"
)

//FileInfo represent info about file
type FileInfo struct {
	Path string
	Name string
	Size string
	Modified string
	Hash string
}

//ReadFiles takes paths to folders and read information about all files where.
func ReadFiles(path []string) {
	//create CSV file
	csvFileName := time.Now().UTC().Format("2006-01-02 15:04:05") + ".csv"
	csvFile, err := os.Create("./snapshots/" + csvFileName)
	if err != nil {
		return
	}
	w := csv.NewWriter(csvFile)

	for i :=0; i<len(path); i++ {
		var filesInfo []FileInfo
		var subFolders []string

		//read files from folder
		fmt.Println("Files from " + path[i] + " :")
		files, err := ioutil.ReadDir(path[i])
		if err != nil {
			fmt.Println("no such directory")
		} else {
			//write files info into struct if files exist
			if len(files) != 0 {
					filesInfo, subFolders = FormatFiles(files, path[i])
					for _,sub := range subFolders{
						path = append(path, sub)
					}

					WriteIntoCSV(w, filesInfo)
			} else {
				fmt.Println("no files")
			}
		}

	}

	csvFile.Close()


}

//WriteIntoCSV format struct and write into CSV
func WriteIntoCSV(w *csv.Writer, filesInfo []FileInfo) {
	for j, f := range filesInfo {
		if len(f.Path) > 0 {
			var str []string
			str = append(str, f.Path, f.Name, f.Size, f.Modified, f.Hash)
			if err := w.Write(str); err != nil {
				log.Fatalln("error writing record to csv:", err)
			}
			log.Println(filesInfo[j])
		}

	}
	w.Flush()

}

//FormatFiles get files and format information about they into struct
func FormatFiles(files []os.FileInfo, path string) ([]FileInfo, []string) {
	filesInfo := make([]FileInfo, len(files))
    var subFolders []string

	for k, file := range files {
		if file.IsDir() == false {
			p := path + "/" + file.Name()

			filesInfo[k].Path = p
			filesInfo[k].Name = file.Name()
			filesInfo[k].Size = strconv.Itoa(int(file.Size()))
			filesInfo[k].Modified = file.ModTime().UTC().Format("2006-01-02 15:04:05.11")
			filesInfo[k].Hash = CreateHash(p)
		} else {
			subFolders = append(subFolders, path + "/" + file.Name())
		}
	}
		return filesInfo, subFolders
}


//CreateHash return a unique hash256 of file
func CreateHash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	str := hex.EncodeToString(h.Sum(nil))
	return str
}
