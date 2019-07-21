package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/nodo/certo/decoder"
	"github.com/nodo/certo/verification"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "certo"
	app.Usage = "Check your certificates"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "decode",
			Aliases: []string{"d"},
			Usage:   "decode the certificate metadata",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "cert"},
				cli.StringFlag{Name: "format"},
			},
			Action: func(c *cli.Context) error {
				path := c.String("cert")
				format := c.String("format")

				d := decoder.New(path, format)
				if ok := d.Validate(); !ok {
					cli.ShowCommandHelp(c, "decode")
					return errors.New("invalid arguments")
				}

				output, err := d.Decode()
				if err != nil {
					return errors.New("unable to decode the certificate: " + err.Error())
				}
				fmt.Println(output)
				return nil
			},
		},
		{
			Name:    "check-local",
			Aliases: []string{"cl"},
			Usage:   "check that a local certificate has been signed by a certificate authority",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "cert"},
				cli.StringFlag{Name: "cacert"},
			},
			Action: func(c *cli.Context) error {
				certPath := c.String("cert")
				caCertPath := c.String("cacert")

				v := verification.NewLocal(certPath, caCertPath)
				if ok := v.Validate(); !ok {
					cli.ShowCommandHelp(c, "check-local")
					return errors.New("invalid arguments")
				}

				ok, err := v.Verify()
				if err != nil {
					return errors.New("unable to check the certificate: " + err.Error())
				}
				if ok {
					fmt.Println("Valid")
				} else {
					fmt.Println("*NOT* Valid")
				}
				return nil
			},
		},
		{
			Name:    "check-remote",
			Aliases: []string{"cr"},
			Usage:   "check that a given ca cert can validate a certificate from a URL",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "url"},
				cli.StringFlag{Name: "cacert"},
			},
			Action: func(c *cli.Context) error {
				url := c.String("url")
				caCertPath := c.String("cacert")

				v := verification.NewRemote(url, caCertPath)
				if ok := v.Validate(); !ok {
					cli.ShowCommandHelp(c, "check-remote")
					return errors.New("invalid arguments")
				}

				ok, err := v.Verify()
				if err != nil {
					return errors.New("unable to check the certificate: " + err.Error())
				}
				if ok {
					fmt.Println("Valid")
				} else {
					fmt.Println("*NOT* Valid")
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
