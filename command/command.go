package command

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type Command struct {
}

type Port struct {
	Command string
	Port    string
	User    string
	PID     string
	PType   string
	Fd      string
	Node    string
}

func New() *Command {
	return &Command{}
}

func (c Command) ExecutePortList() []*Port {
	out, err := exec.Command("lsof", "-PiTCP", "-sTCP:LISTEN").Output()

	if err != nil {
		log.Fatalf("Command.ExecutePortList error %s", err)
	}

	return c.parseOutput(out)
}

func (c Command) ExecutePortKill(port *Port) {
	err := exec.Command("kill", "-9", port.PID).Run()

	if err != nil {
		log.Fatalf("Command.ExecutePortKill error %s", err)
	}
}

func (c Command) parseOutput(out []byte) []*Port {
	return c.parseLines(out)
}

func (c Command) parseLines(out []byte) []*Port {
	line := strings.Split(string(out), "\n")
	return c.parseColumns(line)
}

func (c Command) parseColumns(lines []string) []*Port {

	getAllStringRegex := regexp.MustCompile(`\S+`)

	portRegex := regexp.MustCompile(`[0-9]+`)

	var portList []*Port

	for key, line := range lines {

		if key == 0 {
			continue
		}

		findString := getAllStringRegex.FindAllString(line, -1)

		if len(findString) < 10 {
			continue
		}

		cmd := findString[0]
		user := findString[2]
		pid := findString[1]
		portName := findString[8]
		pType := findString[4]
		fd := findString[3]
		node := findString[7]

		p := portRegex.FindAllString(portName, -1)

		port := &Port{
			PID:     pid,
			Command: cmd,
			User:    user,
			Port:    p[0],
			PType:   pType,
			Fd:      fd,
			Node:    node,
		}

		portList = append(portList, port)

	}

	return portList
}
