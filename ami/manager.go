package ami

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func sendCmd(socket *Socket, cmd []string) error {
	for _, s := range cmd {
		err := socket.Send("%s", s)
		if err != nil {
			fmt.Printf("send login error:[%v]\n", err)
			return err
		}
	}
	return nil
}

func Connect(socket *Socket) (bool, error) {
	if !socket.Connected() {
		fmt.Printf("could not connect to AMI\n")
	}
	if answer, err := socket.Recv(); err != nil || !strings.Contains(answer, "Asterisk Call Manager") {
		return false, errors.New("AMI connection failed")
	}
	return true, nil
}

func Login(socket *Socket, user, secret, events, actionID string) (bool, error) {

	if (len(user) == 0) || (len(secret) == 0) {
		return false, errors.New("Invalid user")
	}

	if !socket.Connected() {
		return false, errors.New("Invalid socket")
	}

	authCmd := []string{
		"Action: Login",
		"\r\nUsername: ",
		user,
		"\r\nSecret: ",
		secret,
		"\r\nEvents: ",
		events,
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	err := sendCmd(socket, authCmd)
	if err != nil {
		return false, err
	}
	message, err := parseMessage(socket)
	if (err != nil) || (message["ActionID"] != actionID) {
		return false, err
	}
	if message["Response"] != "Success" {
		return false, errors.New(message["Message"])
	}
	return true, nil
}

func Logoff(socket *Socket, actionID string) (bool, error) {
	if !socket.Connected() {
		return false, errors.New("Invalid socket")
	}

	logoffCmd := []string{
		"Action: Logoff",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	err := sendCmd(socket, logoffCmd)
	if err != nil {
		return false, err
	}

	message, err := parseMessage(socket)
	if (err != nil) || (message["ActionID"] != actionID) {
		return false, err
	}
	if message["Response"] != "Goodbye" {
		return false, errors.New(message["Message"])
	}
	return true, nil
}

func GetUUID() (string, error) {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		fmt.Printf("open /dev/urandom error:[%v]\n", err)
		return "", err
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6],
		b[6:8], b[8:10], b[10:])
	return uuid, nil
}

func Ping(socket *Socket, actionID string) (bool, error) {
	pingCmd := []string{
		"Action: Ping",
		"\r\nActionID: ",
		actionID,
		"\r\n\r\n", // end of command
	}
	err := sendCmd(socket, pingCmd)
	if err != nil {
		return false, err
	}
	message, err := parseMessage(socket)
	if (err != nil) || (message["ActionID"] != actionID) {
		return false, err
	}
	if message["Response"] != "Success" {
		return false, errors.New(message["Message"])
	}
	return true, nil
}
