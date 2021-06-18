package new

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	githubSource    = "https://github.com/LyricTian/gin-admin.git"
	githubWebSource = "https://github.com/LyricTian/gin-admin-react.git"
	defaultLibName  = "github.com/LyricTian/gin-admin/v7"
	defaultAppName  = "fusion-gin-admin"
)

type Config struct {
	Dir        string
	LibName    string
	AppName    string
	Branch     string
	IncludeWeb bool
}

type Command struct {
	cfg *Config
}

func (a *Command) Exec() error {
	dir, err := filepath.Abs(a.cfg.Dir)
	if err != nil {
		return err
	}

	log.Printf("项目生成目录：%s", dir)

	notExist := false
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			notExist = true
		} else {
			return err
		}
	}

	if notExist {
		source := githubSource
		err = a.gitClone(dir, source)
	}
}

func (a *Command) gitClone(dir, source string) error {
	var args []string
	args = append(args, "clone")

	branch := "master"
	if v := a.cfg.Branch; v != "" {
		branch = v
	}
	args = append(args, "-b", branch)

	args = append(args, source)
	args = append(args, dir)

	log.Printf("执行命令：git %s", strings.Join(args, " "))
	return a.execGit("", args...)
}

func (a *Command) gitInit(dir string) error {
	os.RemoveAll(filepath.Join(dir, ".git"))
	if a.cfg.IncludeWeb {
		os.RemoveAll(filepath.Join(dir, "web", ".git"))
	}

	err := a.execGit(dir, "init")
	if err != nil {
		return err
	}

	err = a.execGit(dir, "add", "-A")
	if err != nil {
		return err
	}

	err = a.execGit(dir, "commit", "-m", "Initial commit")
	if err != nil {
		return err
	}

	return nil
}

func (a *Command) checkInDirs(dir, path string) bool {
	includeDirs := []string{"lib"}
	for _, d := range includeDirs {
		p := filepath.Join(dir, d)
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

func (a *Command) changeDirLibName(dir string) error {

}

func (a *Command) readAndReplaceFile(name string, call func(string) string) error {
	buf, err := a.readFileAndReplace(name, call)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name, buf.Bytes(), 0644)
}

func (a *Command) readFileAndReplace(name string, call func(string) string) (*bytes.Buffer, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := call(scanner.Text())
		buf.WriteString(line)
		buf.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return buf, nil
}
