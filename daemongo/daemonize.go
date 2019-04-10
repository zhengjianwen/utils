package daemongo

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Name of environment variable used to distinguish
// parent and daemonized processes.
var EnvVarName = "_DAEMONGO"

// Value of environment variable used to distinguish
// parent and daemonized processes.
var EnvVarValue = "1"

// Path to daemon working directory.
// If not set, the current user directory will be used.
var WorkDir = ""

// Value of file mask for PID-file.
var PidFileMask os.FileMode = 0644

// Value of umask for daemonized process.
var Umask = 027

// Application name to daemonize.
// Used for printing in default daemon actions.
var AppName = "daemongo"

// Path to application executable.
// Used only for default start/restart actions.
//var AppPath = "./" + filepath.Base(os.Args[0])
var AppPath = filepath.Dir(os.Args[0])
var ProcName = filepath.Base(os.Args[0])

// Absolute or relative path from working directory to PID file.
var PidPath = AppPath + "/../var/"
var LogPath = AppPath + "/../logs/"
var PidFile = PidPath + ProcName + ".pid"
var WardenPidFile = PidPath + ProcName + "_warden.pid"

//var PidFile = os.Args[0] + ".pid"
//var WardenPidFile = os.Args[0] + "_warden.pid"

// Pointer to PID file to keep file-lock alive.
var pidFile *os.File

var Warden = false

// This function wraps application with daemonization.
// Returns isDaemon value to distinguish parent and daemonized processes.
func Daemonize() (isDaemon bool, err error) {
	const errLoc = "daemonigo.Daemonize()"

	err = os.MkdirAll(PidPath, os.ModeDir)
	if err != nil {
		fmt.Errorf("%s: mkdir all pidpath failed, reason -> %s",
			errLoc, err.Error(),
		)
		return
	}

	isDaemon = os.Getenv(EnvVarName) == EnvVarValue
	if WorkDir != "" {
		if err = os.Chdir(WorkDir); err != nil {
			err = fmt.Errorf(
				"%s: changing working directory failed, reason -> %s",
				errLoc, err.Error(),
			)
			return
		}
	}
	if isDaemon {

		syscall.Umask(int(Umask))
		if _, err = syscall.Setsid(); err != nil {
			err = fmt.Errorf(
				"%s: setsid failed, reason -> %s", errLoc, err.Error(),
			)
			return
		}

		//fmt.Printf("pidfile is: %s", PidFile)
		if pidFile, err = lockPidFile(PidFile); err != nil {
			err = fmt.Errorf(
				"%s: locking PID file failed, reason -> %s",
				errLoc, err.Error(),
			)
		}
	} else {
		flag.Usage = func() {
			arr := make([]string, 0, len(actions))
			for k, _ := range actions {
				arr = append(arr, k)
			}
			fmt.Fprintf(os.Stderr, "Usage: %s {%s}\n",
				os.Args[0], strings.Join(arr, "|"),
			)
			flag.PrintDefaults()
		}
		if !flag.Parsed() {
			flag.Parse()
		}
		action, exist := actions[flag.Arg(0)]
		if exist {
			action()
		} else {
			if flag.Arg(0) == "_wardendaemonize" {

				if err := startWarden(1); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Printf("warden %s OK.\n", AppName)
				}
			}
			flag.Usage()
		}
	}
	return
}

// Locks PID file with a file lock.
// Keeps PID file open until applications exits.
func lockPidFile(tmpFile string) (pidFile *os.File, err error) {
	var file *os.File
	file, err = os.OpenFile(
		tmpFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, PidFileMask,
	)
	if err != nil {
		return
	}
	defer func() {
		// file must be open whole runtime to keep lock on itself
		if err != nil {
			file.Close()
		}
	}()

	if err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX); err != nil {
		return
	}
	var fileLen int
	fileLen, err = fmt.Fprint(file, os.Getpid())
	if err != nil {
		return
	}
	if err = file.Truncate(int64(fileLen)); err != nil {
		return
	}

	return file, err
}

// Unlocks PID file (locked by current daemonized process) and closes this file.
//
// This function can be useful for graceful restarts or other
// untrivial scenarios, but usually there is no need to use it.
func UnlockPidFile() {
	if pidFile != nil {
		syscall.Flock(int(pidFile.Fd()), syscall.LOCK_UN)
		pidFile.Close()
	}
}

// Checks status of daemonized process.
// Can be used in daemon actions to perate with daemonized process.
func Status(pidFile string) (isRunning bool, pr *os.Process, e error) {
	const errLoc = "daemonigo.Status()"
	var (
		err  error
		file *os.File
	)

	file, err = os.Open(pidFile)
	if err != nil {
		if !os.IsNotExist(err) {
			e = fmt.Errorf(
				"%s: could not open PID file, reason -> %s",
				errLoc, err.Error(),
			)
		}
		return
	}
	defer file.Close()
	fd := int(file.Fd())
	if err = syscall.Flock(
		fd, syscall.LOCK_EX|syscall.LOCK_NB,
	); err != syscall.EWOULDBLOCK {
		if err == nil {
			syscall.Flock(fd, syscall.LOCK_UN)
		} else {
			e = fmt.Errorf(
				"%s: PID file locking attempt failed, reason -> %s",
				errLoc, err.Error(),
			)
		}
		return
	}

	isRunning = true
	var n, pid int
	content := make([]byte, 128)
	n, err = file.Read(content)
	if err != nil && err != io.EOF {
		e = fmt.Errorf(
			"%s: could not read from PID file, reason -> %s",
			errLoc, err.Error(),
		)
		return
	}
	pid, err = strconv.Atoi(string(content[:n]))
	if n < 1 || err != nil {
		e = fmt.Errorf(
			"%s: bad PID format, PID file is possibly corrupted", errLoc,
		)
		return
	}
	pr, err = os.FindProcess(pid)
	if err != nil {
		fmt.Errorf(
			"%s: cannot find process by PID, reason -> %s", errLoc, err.Error(),
		)
	}

	return
}

func startWarden(timeout uint8) (e error) {

	path := fmt.Sprintf("%v/%v", AppPath, AppName)
	//path, err := filepath.Abs(AppPath)
	cmd := exec.Command(path, "warden")

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf(
			"failed to start %s warden, reason -> %s",
			AppName, err.Error())
	}

	select {
	case <-func() chan bool {
		ch := make(chan bool)
		go func() {
			if err := cmd.Wait(); err != nil {
				e = fmt.Errorf(
					"Warden running failed, reason -> %s\n", AppName, err.Error())
			} else {
				e = fmt.Errorf("Warden stopped and not running")
			}
			ch <- true
		}()
		return ch
	}():
	case <-time.After(time.Duration(time.Duration(timeout) * time.Second)):
	}

	return
}

func wardenRun() {
	const errLoc = "daemonigo.wardenRun()"
	var (
		err  error
		file *os.File
	)

	if pidFile, err = lockPidFile(WardenPidFile); err != nil {
		err = fmt.Errorf(
			"%s: locking PID file failed, reason -> %s",
			errLoc, err.Error(),
		)
	}

	// retry := 0
	// first := time.Time{}

	for {
		time.Sleep(3 * time.Second)
		fmt.Print("Warden start to check pid file...")
		checkFile := func() (isRunning bool) {
			//fmt.Printf("pidfile is: %s", PidFile)
			file, err = os.Open(PidFile)
			defer file.Close()
			if err != nil {
				if !os.IsNotExist(err) {
					fmt.Printf(
						"%s: could not open PID file, reason -> %s",
						errLoc, err.Error(),
					)
				}
				return
			}

			fd := int(file.Fd())
			if err = syscall.Flock(
				fd, syscall.LOCK_EX|syscall.LOCK_NB,
			); err != syscall.EWOULDBLOCK {
				if err == nil {
					syscall.Flock(fd, syscall.LOCK_UN)
				} else {
					fmt.Printf(
						"%s: PID file locking attempt failed, reason -> %s",
						errLoc, err.Error(),
					)
				}
				return
			}
			isRunning = true
			return
		}

		if !checkFile() {
			fmt.Print("Failed!\nRestart %s.\n", AppName)
			// Restart app
			Start(1)

		} else {
			fmt.Print("OK!")
		}

	}

}

// Prepares and returns command for starting daemonized process.
//
// This function can also be used when writing your own daemon actions.
func StartCommand() (*exec.Cmd, error) {
	const errLoc = "daemonigo.StartCommand()"
	path := fmt.Sprintf("%v/%v", AppPath, AppName)
	//path := AppPath + '/' + AppName
	//path, err := filepath.Abs(AppPath)
	//if err != nil {
	//	return nil, fmt.Errorf(
	//		"%s: failed to resolve absolute path of %s, reason -> %s",
	//		errLoc, AppName, err.Error(),
	//	)
	//}

	//fmt.Printf("apppath is: %s", path)
	cmd := exec.Command(path)
	cmd.Env = append(
		os.Environ(), fmt.Sprintf("%s=%s", EnvVarName, EnvVarValue),
	)
	stdoutfile, err := os.OpenFile(LogPath+AppName+".stdout.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err == nil {
		cmd.Stdout = stdoutfile
	} else {
		fmt.Printf("set stdout to %s failed", LogPath+AppName+".stdout.log")
	}
	stderrfile, err := os.OpenFile(LogPath+AppName+".stderr.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err == nil {
		cmd.Stderr = stderrfile
	} else {
		fmt.Printf("set stderr to %s failed", LogPath+AppName+".stderr.log")
	}
	return cmd, nil
}

// Starts daemon process and waits timeout number of seconds.
// If daemonized process keeps running after timeout seconds passed
// then process seems to be successfully started.
//
// This function can also be used when writing your own daemon actions.
func Start(timeout uint8) (e error) {
	const errLoc = "daemonigo.Start()"
	cmd, err := StartCommand()
	if err != nil {
		return fmt.Errorf(
			"%s: failed to create daemon start command, reason -> %s",
			errLoc, err.Error(),
		)
	}
	if err = cmd.Start(); err != nil {
		return fmt.Errorf(
			"%s: failed to start %s, reason -> %s",
			errLoc, AppName, err.Error(),
		)
	}
	select {
	case <-func() chan bool {
		ch := make(chan bool)
		go func() {
			if err := cmd.Wait(); err != nil {
				e = fmt.Errorf(
					"%s: %s running failed, reason -> %s",
					errLoc, AppName, err.Error(),
				)
			} else {
				e = fmt.Errorf(
					"%s: %s stopped and not running", errLoc, AppName,
				)
			}
			ch <- true
		}()
		return ch
	}():
	case <-time.After(time.Duration(timeout) * time.Second):

	}
	return
}

// Stops daemon process.
// Sends signal os.Interrupt to daemonized process.
//
// This function can also be used when writing your own daemon actions.
func Stop(process *os.Process, pidFile string) (e error) {
	const errLoc = "daemonigo.Stop()"
	if err := process.Signal(os.Interrupt); err != nil {
		e = fmt.Errorf(
			"%s: failed to send interrupt signal to %s, reason -> %s",
			errLoc, AppName, err.Error(),
		)
		return
	}
	for {
		time.Sleep(200 * time.Millisecond)
		switch isRunning, _, err := Status(pidFile); {
		case err != nil:
			e = fmt.Errorf(
				"%s: checking status of %s failed, reason -> %s",
				errLoc, AppName, err.Error(),
			)
			return
		case !isRunning:
			return
		}
	}
}
