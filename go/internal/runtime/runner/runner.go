package runner

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	c "ray/internal/runtime/config"
)

func StartRayHead(rayConfig c.RayConfig) {
	commandName := "/Users/kyoka/opt/anaconda3/bin/ray"
	cmdArgs := append([]string{"start", "--head"}, rayConfig.GetHeadArgs()...)

	cmd := exec.Command(commandName, cmdArgs...)
	fmt.Printf("StartRayHead cmd: %s \n", cmd.String())
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
