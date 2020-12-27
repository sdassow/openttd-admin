package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/sdassow/openttd-admin/pkg/admin"

	"github.com/chzyer/readline"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
	"gopkg.in/mattes/go-expand-tilde.v1"
)

func init() {
	rootCmd.AddCommand(rconCmd)

	rconCmd.Flags().StringP("config", "c", "~/.config/openttd/openttd.cfg", "openttd config file")
	rconCmd.Flags().BoolP("debug", "d", false, "debug mode")
	rconCmd.Flags().StringP("name", "n", "OpenTTD-Admin", "client name")
	rconCmd.Flags().StringP("version", "v", "1.10.3", "client version")
}

var rconCmd = &cobra.Command{
	Use:   "rcon [address]",
	Short: "Remote console interface",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var address string
		var password string

		colormap := map[admin.Colour]aurora.Color{
			admin.ColourDarkBlue:  aurora.BlueFg,
			admin.ColourPaleGreen: aurora.GreenFg,
			admin.ColourPink:      aurora.GreenFg,
			admin.ColourYellow:    aurora.YellowFg,
			admin.ColourRed:       aurora.RedFg,
			admin.ColourLightBlue: aurora.BlueFg,
			admin.ColourGreen:     aurora.GreenFg,
			admin.ColourDarkGreen: aurora.GreenFg,
			admin.ColourBlue:      aurora.BlueFg,
			admin.ColourCream:     aurora.YellowFg,
			admin.ColourMauve:     aurora.WhiteFg,
			admin.ColourPurple:    aurora.RedFg,
			admin.ColourOrange:    aurora.YellowFg,
			admin.ColourBrown:     aurora.RedFg,
			admin.ColourGrey:      aurora.WhiteFg,
			admin.ColourWhite:     aurora.WhiteFg,
		}

		name, _ := cmd.Flags().GetString("name")
		version, _ := cmd.Flags().GetString("version")
		debug, _ := cmd.Flags().GetBool("debug")

		if len(args) == 1 {
			address = args[0]
		} else {
			cfgfile, _ := cmd.Flags().GetString("config")

			cfgpath, err := tilde.Expand(cfgfile)
			if err != nil {
				log.Fatalf("Failed to expand tilde: %+v", err)
			}

			cfg, err := ini.Load(cfgpath)
			if err != nil {
				log.Fatalf("Failed to read config file: %v", err)
			}

			port := cfg.Section("network").Key("server_admin_port").String()
			if port == "" {
				port = "3977"
			}

			password = cfg.Section("network").Key("admin_password").String()

			// Get host from server bind addresses. Config keys are addresses,
			// values are empty: https://wiki.openttd.org/en/Archive/Manual/Settings/Server%20bind%20ip
			// Use same default of localhost as implementation
			host := "localhost"
			hosts := cfg.Section("server_bind_addresses").KeyStrings()
			if len(hosts) > 0 {
				host = hosts[0]
			}

			address = host + ":" + port
		}

		if password == "" {
			pw, err := readline.Password("Server admin password: ")
			if err != nil {
				log.Fatal(err)
			}
			password = string(pw)
		}

		log.Printf("connecting to %s...", address)

		adm, err := admin.Connect(address, password, name, version)
		if err != nil {
			log.Fatal(err)
		}

		rl, err := readline.NewEx(&readline.Config{
			Prompt:          "> ",
			InterruptPrompt: "^C",
		})
		if err != nil {
			log.Fatal(err)
		}
		defer rl.Close()

		log.SetOutput(rl.Stderr())

		// user input channel
		input := make(chan string)

		go func() {
			for {
				m, err := adm.ReadMessage()
				if err != nil {
					if err == io.EOF {
						break
					}
					log.Fatalf("Failed to read message: %+v", err)
				}

				switch msg := m.(type) {
				case *admin.ServerWelcome:
					log.Printf("connected")

					adm.Send(admin.NewAdminUpdateFrequency(admin.AdminUpdateConsole, admin.AdminFrequencyAutomatic))
					adm.Send(admin.NewAdminUpdateFrequency(admin.AdminUpdateChat, admin.AdminFrequencyAutomatic))
				case *admin.ServerConsole:
					if debug {
						fmt.Fprintf(rl, "# %+v\n", msg)
					}
					fmt.Fprintf(rl, "%s> %s\n", msg.Origin, msg.Message)
				case *admin.ServerRcon:
					if debug {
						fmt.Fprintf(rl, "# %+v\n", msg)
					}
					col, found := colormap[msg.Colour]
					if !found {
						col = aurora.WhiteFg
					}
					fmt.Fprintln(rl, aurora.Sprintf("> %s", aurora.Colorize(msg.Result, col)))
				default:
					if debug {
						fmt.Fprintf(rl, "# %#v\n", msg)
					}
				}
			}
		}()

		go func() {
			for {
				select {
				case text := <-input:
					_, err = adm.Send(admin.NewAdminRcon(text))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}()

		// keep reading input
		for {
			line, err := rl.Readline()
			if err != nil {
				if err == readline.ErrInterrupt || err == io.EOF {
					_, err = adm.Send(admin.NewAdminQuit())
					if err != nil {
						log.Fatal(err)
					}
					break
				}

				log.Fatalf("Failed to read line: %+v", err)
				break
			}
			input <- line
		}
		close(input)

		adm.Close()
		rl.Clean()

	},
}
