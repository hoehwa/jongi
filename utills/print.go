package utills

import (
	"log"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/hoehwa/jongi/constants"
	"github.com/hoehwa/jongi/utills/bubbletea/listfancy"
)

// Print content of selected item with highlighted text.
func PrintContent(destDir string) {
	srcDir := constants.BaseDir + destDir
	files, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	fileNamesForTitle := make([]string, 0, len(files))
	urlsForDescription := make([]string, 0, len(files))

	for _, file := range files {
		fileNamesForTitle = append(fileNamesForTitle, file.Name())
		urlsForDescription = append(urlsForDescription, constants.GardenBaseurl+"/flower/snippet"+destDir+"/"+strings.Split(file.Name(), ".")[0])
		// fileNamesForTitle[i] = file.Name()
		// urlsForDescription[i] = resources.GardenBaseurl + "/flower/snippet" + destDir + "/" + strings.Split(file.Name(), ".")[0]
	}

	listfancy.Init(fileNamesForTitle, urlsForDescription, srcDir)
	highlihgted := listfancy.Content
	if err := quick.Highlight(os.Stdout, highlihgted, listfancy.ItemName, "terminal", "monokai"); err != nil {
		log.Fatal(err)
	}
}
