package main

import (
    "os"
    "fmt"
    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "fireup"
    app.Usage = "An HTTP load testing utility tool"
    app.Action = func(c *cli.Context) {
        fmt.Println("TODO: Do something here...")
    }

    app.Run(os.Args)
}