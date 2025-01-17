package collector

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
)

// Container solution
// The better solution for this is to update a mounted volume from inside the non-privileged
// namespaces container, and have another listening process run as root outside the container
// , to do the actual shutdown work at the host machine. This provides a secure nterface between
// the conatainer and the host machine.
func terminate(target string) {
	fmt.Printf("target: %s\n", target)
	if target == "127.0.0.1" {
		// run locally
		fmt.Println("Hey I am going to turn you off! Server.")
		// b, err := ioutil.ReadFile("C:\\shutdown_signal")	// windows
		b, err := ioutil.ReadFile("/var/run/shutdown_signal")
		if err != nil {
			panic(err)
		}
		fmt.Printf("shutdown_signal: %s", string(b))

		signal := []byte("false")
		// err = ioutil.WriteFile("C:\\shutdown_signal", signal, 0644)
		err = ioutil.WriteFile("/var/run/shutdown_signal", signal, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		// TODO: turn off remote server
		// cmd := exec.Command("ssh", "-t", "-t", "-p", "22", "{{hostip}}", "init 6")
		// out, err := cmd.CombinedOutput()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Printf("cmd Output:%v", string(out))
	}
}

// Binary solution
func shutdownCommand() {
	// shutdown locally
	cmd := exec.Command("init 6")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("shutdown", "/s")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cmd Output:%v", string(out))
}
