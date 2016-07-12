package gocrawl

import (
    "fmt"
    "io"
    "os"
    "strings"
    "time"
    "log"
    "golang.org/x/crypto/ssh"
)

// TestVersion is the version that the unit tests are run against
const Version = 2

//remoteSession is an individual communication channel with a RemoteClient 
//it wraps ssh.Session for dependency injection to implement unit testing
type remoteSession interface {
    StdoutPipe() (io.Reader, error)
    StdinPipe() (io.WriteCloser, error)
    RequestPty(term string, h, w int, termmodes ssh.TerminalModes) error
    Shell() error
}

//Device is a remote device that can communicate by sending commands to Stdin
// and recieves responses from Stdout
type Device struct {
    Hostname string
    stdin io.WriteCloser
    stdout io.Reader
    logger *log.Logger
    timeout time.Duration
    prompt string
}

//NewDevice creates a new network device
func NewDevice(hostname string) Device {
    logger := log.New(os.Stdout, hostname + ": ", 0)
    if logFile, err := os.OpenFile(hostname + ".log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666); err == nil {
        fmt.Print("Opened file")
        logger = log.New(logFile, hostname + ": ", 0)
    } else {
        logger.Printf("Failed to open logfile: %s", err)
    }
    dev := Device{
        Hostname : hostname,
        stdin : nil,
        stdout : nil,
        logger: logger,
        timeout: 30 * time.Second,
        prompt: ">",
    }
    dev.logger.Print("New device")
    return dev
}

//Connect establishes a connection to a new network device
func (dev *Device) Connect(user, pass string) (string, error) {
    config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
	}

    client, clientErr := ssh.Dial("tcp", dev.Hostname+":22", config)
    if clientErr != nil {
        return "", fmt.Errorf("Dialing %s failed: %v", dev.Hostname, clientErr)
    }

    session, sessionErr := client.NewSession()
    if sessionErr != nil {
        return "", fmt.Errorf("Creating a session failed: %v", sessionErr)
    }

    if sshOut, stdOutErr := session.StdoutPipe(); stdOutErr == nil {
        dev.stdout = sshOut
    } else {
        return "", fmt.Errorf("Stdout pipe failed: %v", stdOutErr)
    }

    if sshIn, stdInErr := session.StdinPipe(); stdInErr == nil {
        dev.stdin = sshIn
    } else {
        return "", fmt.Errorf("Stdin pipe failed: %v", stdInErr)
    }

    modes := ssh.TerminalModes{
        ssh.ECHO:          1,     // disable echoing
        ssh.TTY_OP_ISPEED: 144000, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 144000, // output speed = 14.4kbaud
    }

    if ptyErr := session.RequestPty("xterm", 80, 0, modes); ptyErr != nil {
        return "", fmt.Errorf("Request Pty failed: %v", ptyErr)
    }

    if shellErr := session.Shell(); shellErr != nil {
        return "", fmt.Errorf("Shell failed: %v", shellErr)
    }

    welcome, err := dev.Send("")
    if err != nil {
        return "", fmt.Errorf("Failed to get welcome message: %s" , err)
    }

    return welcome, nil
}

func (dev *Device) Send(command string) (string, error) {
    if dev.stdout == nil || dev.stdin == nil {
        return "", fmt.Errorf("%s is not connected", dev.Hostname)
    }

    if !strings.HasSuffix(command, "\r") {
        command += "\r"
    }

    if _, err := dev.stdin.Write([]byte(command)); err != nil {
        return "", fmt.Errorf("Error on write: %s", err)
    }

    response_chan, err_chan := dev.bufferedRead(command, dev.prompt)
    select {
    case response := <-response_chan:
        return response, nil
    case err := <-err_chan:
        return "", fmt.Errorf("Error on read: %s", err)
    case <-time.After(dev.timeout):
        return "", fmt.Errorf("Device timedout after %s", dev.timeout)
    }
}

//bufferedRead parses responses into a reply/resonse pattern
func (dev *Device) bufferedRead(prefix, suffix string) (chan string, chan error) {
    buff := make([]byte, 1000)
    outChan := make(chan string)
    errChan  := make(chan error)

    dev.logger.Print("---------Reader---------\n")
    go func() {
        output := ""
        for {
            if bytes_read, err := dev.stdout.Read(buff); err != nil {
                errChan <- fmt.Errorf("Error on read: %s", err)
            } else if bytes_read > 0 {
                dev.logger.Print(string(buff[:bytes_read]))
                output += string(buff[:bytes_read])
            }

            if strings.HasSuffix(output, suffix) {
                dev.logger.Printf("--------Done----------\n%s", output)
                output = strings.TrimPrefix(output, prefix)
                output = strings.TrimSuffix(output, suffix)
                outChan <- output
                return
            }
        }
    }()
    return outChan, errChan
}
