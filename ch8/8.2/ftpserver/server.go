package ftpserver

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/skipcloud/go-programming-book/ch8/8.2/user"
)

type ServerStatusCode string

const (
	FileOkay                 = "150 File Status okay; Opening Connection\r\n"
	CommandSuccessful        = "200 %s Successful\r\n"
	NotImplemented           = "202 Command Not Implemented\r\n"
	SystemName               = "215 LINUX System Type\r\n"
	FTPServerReady           = "220 FTP Server Ready\r\n"
	ClosingServerConnection  = "221 Closing main connection\r\n"
	ClosingRemoteConnection  = "226 closing connection\r\n"
	LoginSucessful           = "230 Login Successful\r\n"
	UserNameOkayNeedPassword = "331 Username Okay\r\n"
	InvalidNameOrPassword    = "430 Invalid Name or Password\r\n"
)

// just a sandbox directory in tmp
const initialDirectory = "/tmp/ftp"

type serverConnection struct {
	Conn             net.Conn
	User             user.User
	Reader           *bufio.Reader
	CurrentDirectory string
	RemoteAddr       string
}

func newServerConnection(conn net.Conn) *serverConnection {
	// I should really handle the error here...
	if err := os.Chdir(initialDirectory); err != nil {
		log.Fatal(err)
	}
	return &serverConnection{
		Conn:             conn,
		User:             user.New(),
		Reader:           bufio.NewReader(conn),
		CurrentDirectory: initialDirectory,
	}
}

func (c *serverConnection) Read() (cmd, args string, err error) {
	s, err := c.Reader.ReadString('\n')
	if err != nil {
		return
	}
	ss := strings.SplitN(s, " ", 2)

	switch len(ss) {
	case 2:
		cmd = strings.TrimSpace(ss[0])
		args = strings.TrimSpace(ss[1])
	case 1:
		cmd = strings.TrimSpace(ss[0])
	}
	return
}

func (c *serverConnection) Write(s string) error {
	_, err := fmt.Fprintf(c.Conn, s)
	return err
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	c := newServerConnection(conn)

	c.Write(FTPServerReady)
	// Not really authenticating, just doing the anonymous dance
	if err := c.authenticate(); err != nil {
		return
	}

	var cmd, args string
	var err error
	for {
		cmd, args, err = c.Read()
		if err == io.EOF {
			fmt.Println("Connection closed")
			return
		} else if err != nil {
			log.Fatal(err)
		}
		if err = c.handleCommand(cmd, args); err != nil {
			log.Fatal(err)
		}
	}

}

// authenticate ensures the user logs in as anonymous
// but doesn't care about the password
func (c *serverConnection) authenticate() error {
	// Get username, should be anonymous
	_, args, err := c.Read()
	if err != nil {
		return err
	}
	if args != "anonymous" {
		c.Write(InvalidNameOrPassword)
		return errors.New("Username is not anonymous")
	}

	// Username was okay let's get the password
	c.Write(UserNameOkayNeedPassword)

	// We don't actually care about the password
	_, _, err = c.Read()
	if err != nil {
		return err
	}

	c.Write(LoginSucessful)
	// probably no need to "authenticate" a user
	// but I'll leave this code in
	c.User.Authenticated = true
	return nil
}

func (c *serverConnection) handleCommand(cmd, args string) error {
	fmt.Printf("cmd %s\targs: %s\n", cmd, args)
	switch cmd {
	case "PWD":
		return c.Write(c.CurrentDirectory + "\r\n")
	case "CWD":
		f, err := os.Open(args)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := f.Chdir(); err != nil {
			return err
		}
		c.CurrentDirectory = f.Name()
		return c.Write(fmt.Sprintf(CommandSuccessful, "CWD"))
	case "LIST":
		names, err := ioutil.ReadDir(c.CurrentDirectory)
		if err != nil {
			return err
		}
		err = c.Write(fmt.Sprintf(FileOkay))
		if err != nil {
			return fmt.Errorf("error: writing to connection: %w", err)
		}

		remote, err := net.Dial("tcp", c.RemoteAddr)
		if err != nil {
			return fmt.Errorf("error: getting remote connection: %w", err)
		}
		defer remote.Close()

		var files string
		for _, f := range names {
			if f.IsDir() {
				files += f.Name() + "/\r\n"
				continue
			}
			files += f.Name() + "\r\n"
		}
		fmt.Fprintf(remote, files)
		// transfer complete
		fmt.Fprint(remote, ClosingRemoteConnection)
		return c.Write(fmt.Sprintf(CommandSuccessful, "LIST"))
	case "PORT":
		c.RemoteAddr = buildAddr(args)
		return c.Write(fmt.Sprintf(CommandSuccessful, "PORT"))
	case "QUIT":
		return c.Write(ClosingServerConnection)
	case "RETR":
		f, err := os.Open(args)
		if err != nil {
			return err
		}
		defer f.Close()
		err = c.Write(fmt.Sprintf(FileOkay))
		if err != nil {
			return fmt.Errorf("error: writing to connection: %w", err)
		}
		remote, err := net.Dial("tcp", c.RemoteAddr)
		if err != nil {
			return fmt.Errorf("error: getting remote connection: %w", err)
		}
		defer remote.Close()

		io.Copy(remote, f)
		// transfer complete
		_ = c.Write(ClosingRemoteConnection)
		return c.Write(fmt.Sprintf(CommandSuccessful, "RETR"))
	case "SYST":
		return c.Write(SystemName)
	default:
		return c.Write(NotImplemented)
	}
}

func buildAddr(s string) string {
	ss := strings.Split(s, ",")
	addr := strings.Join(ss[:4], ".")
	p1, _ := strconv.Atoi(ss[4])
	p2, _ := strconv.Atoi(ss[5])
	// the argument for the PORT command takes this format
	// 1,2,3,4,5,6. 1-4 make up the address 1.1.1.1 while
	// 5 and 6 need to be calculcated to make the port. The
	// port number is a 16 bit value but in the argument it's
	// two 8 bit values, as such you need to multiple the first
	// value by 256 and add the second. More information
	// here: https://stackoverflow.com/questions/9966993/how-to-get-port-in-ftp-protocol-from-passive-mode#9967002
	return fmt.Sprintf("%s:%d", addr, p1*256+p2)
}
