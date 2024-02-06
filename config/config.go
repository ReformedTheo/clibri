package config

func ConfigParser(args []string) {

	switch args[0] {
	case "ip":
		setIP(args[1])
	}
}

func setIP(ip string) {

}

type Config struct {
	Ip      string
	Command string
	Antenna bool
}
