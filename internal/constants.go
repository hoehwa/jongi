package internal

import "os"

// Base directory of snippet.
var userHomeDir, _ = os.UserHomeDir()
var BaseDir = userHomeDir + "/jongi"

// github username and repo in order to fetch snippets.
var Owner = "mindulle"
var Repository = "background"

// urls for hosted websites.
const GardenBaseurl = "https://garden.mindulle.vercel.app/core"
