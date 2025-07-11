// cmd/coincryptokitapix/main.go
package main

import (
"flag"
"log"
"os"

"coincryptokitapix/internal/coincryptokitapix"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := coincryptokitapix.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
