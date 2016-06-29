package gocrawl

import (
    "bytes"
    "fmt"
    "io"
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

//Device 
type Device struct {
    Hostname string
    Stdin chan string
    Stdout chan string
    session remoteSession
}

func NewDevice(hostname string) Device {
    return Device{
        Hostname : hostname,
        Stdin : nil,
        Stdout : nil,
        session: getSession(),
    }
}

func getSession() remoteSession {
    return nil
}

func (dev Device) Connect(user, pass string) error {
    config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
	}

    client, clientErr := ssh.Dial("tcp", dev.Hostname, config)
    if clientErr != nil {
        return fmt.Errorf("Dialing %s failed: %v", dev.Hostname, clientErr)
    }

    session, sessionErr := client.NewSession()
    if sessionErr != nil {
        return fmt.Errorf("Creating a session failed: %v", sessionErr)
    }

    sshOut, stdOutErr := session.StdoutPipe()
    if stdOutErr != nil {
        return fmt.Errorf("Stdout pipe failed: %v", stdOutErr)
    }

    sshIn, stdInErr := session.StdinPipe()
    if stdInErr != nil {
        return fmt.Errorf("Stdin pipe failed: %v", stdInErr)
    }

    modes := ssh.TerminalModes{
        ssh.ECHO:          0,     // disable echoing
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }

    ptyErr := session.RequestPty("vt100", 80, 0, modes)
    if ptyErr != nil {
        return fmt.Errorf("Request Pty failed: %v", ptyErr)
    }

    shellErr := session.Shell()
    if shellErr != nil {
        return fmt.Errorf("Shell failed: %v", shellErr)
    }

    dev.Stdin = makeStdin(sshIn)
    dev.Stdout = makeStdout(sshOut)

    return nil
}

func makeStdin(sshIn io.WriteCloser) chan string {
    stdin := make(chan string)
    go func() {
        for command := range stdin {
            sshIn.Write([]byte(command))
        }
    }()
    return stdin
}

func makeStdout(sshOut io.Reader) chan string {
    stdout := make(chan string)
    go func() {
        var command bytes.Buffer
        for ;; {
            if bytes.HasSuffix(command, []byte(prompt)) {
                stdout <- command.
                command.Reset()
    return stdout
}

func (dev Device) Send(command string) (string, error) {
    return "response", nil
}
