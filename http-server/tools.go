package server

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

func envReading() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("Error reading env file")
	}

	var errAddrres, errPORT error

	address, errAddrres = validIPAddress(os.Getenv("ADDRESS"))
	if errAddrres != nil {
		log.Println(errAddrres)
	}

	port, errPORT = validPort(os.Getenv("PORT"))
	if errPORT != nil {
		log.Println("Invalid port")
	}

	if errEnv != nil || errAddrres != nil || errPORT != nil {
		log.Println(defaultMsg)
	}
}

func validIPAddress(ip string) (string, error) {
	// Use a regular expression to validate the IP address format
	ipRegex := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
	match, err := regexp.MatchString(ipRegex, ip)
	if err != nil {
		return address, errors.New("Regex Error")
	}

	if match {
		return ip, nil
	}
	return address, errors.New("Invalid format of IP")
}

func validPort(p string) (string, error) {
	intPort, err := strconv.Atoi(p)
	if err != nil || intPort == 0 {
		return port, errors.New("Invalid format of PORT")
	}
	return p, nil
}

// func isCtrlRPressed() bool {
// 	// Open keyboard
// 	if err := keyboard.Open(); err != nil {
// 		log.Fatal(err)
// 	}
// 	defer keyboard.Close()

// 	// Read a rune without blocking
// 	if r, key, err := keyboard.GetKey(); err == nil && r == 0 && key == keyboard.KeyCtrlR {
// 		return true
// 	}
// 	return false
// }
