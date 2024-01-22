package internal

import (
	"log"
	"os"
	"strings"

	"github.com/hoehwa/gopkg/beautify"
	"github.com/hoehwa/gopkg/bubbletea/listfancy"
)

// Return a formatted filename as an URL format.
func fmtFileNameToURL(subPath string, fname string) string {
	basename := BaseURL + subPath + "/"
	location := strings.Split(fname, ".")[0]

	return basename + location
}

// Return titles string array and descriptions string array respectively.
func getTitlesAndDescs(srcPath string, subPath string) ([]string, []string) {
	files, err := os.ReadDir(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	titles := make([]string, 0, len(files))
	descs := make([]string, 0, len(files))

	for _, file := range files {
		titles = append(titles, file.Name())
		descs = append(descs, fmtFileNameToURL(subPath, file.Name()))
	}

	return titles, descs
}

// Print content of selected item with highlighted text.
func ShowListAt(subPath string) {
	srcPath := BaseDir + subPath

	titles, descs := getTitlesAndDescs(srcPath, subPath)
	var selection *listfancy.Item

	callout := listfancy.NewCallout(listfancy.Callout{
		Titles:    titles,
		Descs:     descs,
		Selection: selection,
	})

	listfancy.InitCallout(*callout)
	beautify.RawContent(srcPath, callout.Selection.Description())
}
