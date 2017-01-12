# go-autostart

[![GoDoc](https://godoc.org/github.com/ProtonMail/go-autostart?status.svg)](https://godoc.org/github.com/ProtonMail/go-autostart)

A Go library to run a command after login.

## Usage

```go
package main

import (
	"log"
	"github.com/ProtonMail/go-autostart"
)

func main() {
	app := &autostart.App{
		Name: "test",
		DisplayName: "Just a Test App",
		Exec: []string{"bash", "-c", "echo autostart >> ~/autostart.txt"},
	}

	if app.IsEnabled() {
		log.Println("App is already enabled, removing it...")

		if err := app.Disable(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Enabling app...")

		if err := app.Enable(); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Done!")
}
```

## Behavior

* On Linux and BSD, it creates a `.desktop` file in `$XDG_CONFIG_HOME/autostart`
  (i.e. `$HOME/.config/autostart`). See http://askubuntu.com/questions/48321/how-do-i-start-applications-automatically-on-login
* On macOS, it creates a `launchd` job. See http://blog.gordn.org/2015/03/implementing-run-on-login-for-your-node.html
* On Windows, it creates a link to the program in `%USERPROFILE%\Start Menu\Programs\Startup`

## License

MIT
