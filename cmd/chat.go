package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/sdassow/openttd-admin/pkg/admin"

	"github.com/chzyer/readline"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
	"gopkg.in/mattes/go-expand-tilde.v1"
)

func init() {
	rootCmd.AddCommand(chatCmd)

	chatCmd.Flags().StringP("config", "c", "~/.config/openttd/openttd.cfg", "openttd config file")
	chatCmd.Flags().StringP("name", "n", "OpenTTD-Admin", "client name")
	chatCmd.Flags().StringP("version", "v", "1.10.3", "client version")
}

var chatCmd = &cobra.Command{
	Use:   "chat [address]",
	Short: "Chat with other players on a server",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var address string
		var password string

		name, _ := cmd.Flags().GetString("name")
		version, _ := cmd.Flags().GetString("version")

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

					adm.Send(admin.NewAdminPoll(admin.AdminUpdateClientInfo, admin.AllCompanies))
					adm.Send(admin.NewAdminPoll(admin.AdminUpdateCompanyInfo, admin.AllCompanies))

					adm.Send(admin.NewAdminUpdateFrequency(admin.AdminUpdateChat, admin.AdminFrequencyAutomatic))
				case *admin.ServerChat:
					//fmt.Fprintf(rl, "# %+v\n", msg)
					fmt.Fprintf(rl, "%d> %s\n", msg.ClientId, msg.Message)
				default:
					fmt.Fprintf(rl, "# %+v\n", msg)
				}

			}
		}()

		go func() {
			for {
				select {
				case text := <-input:
					_, err = adm.Send(admin.NewAdminChat(admin.NetworkActionChat, admin.DesttypeBroadcast, 0, text))
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
