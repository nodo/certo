package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/nodo/certo/actions"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "certo"
	app.Usage = "Check your certificate"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "decode",
			Aliases: []string{"d"},
			Usage:   "decode the certificate metadata",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "format"},
			},
			Action: func(c *cli.Context) error {
				path := c.Args().First()
				format := c.String("format")
				output, err := actions.Decode(path, format)
				if err != nil {
					return errors.New("unable to decode the certificate: " + err.Error())
				}
				fmt.Println(output)
				return nil
			},
		},
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "check that a certificate has been signed by a certificate authority",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "cert"},
				cli.StringFlag{Name: "cacert"},
			},
			Action: func(c *cli.Context) error {
				certPath := c.String("cert")
				caCertPath := c.String("cacert")

				ok, err := actions.CheckSignature(certPath, caCertPath)
				if err != nil {
					return errors.New("unable to check the certificate: " + err.Error())
				}
				fmt.Println(ok)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
