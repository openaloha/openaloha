package runhandler

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CmdRunHandler struct {
	Workspace string
}

func (r *CmdRunHandler) Run(cmds []string) error {
	fmt.Printf("run cmd: %s\n", cmds)

	if(len(cmds) == 0) {
		return nil
	}

	for _, cmd := range cmds {
		fmt.Printf("run cmd: %s\n", cmd)
		
		// 使用 shell 来执行命令，这样可以正确处理带参数的命令
		var cmdObj *exec.Cmd
		if strings.TrimSpace(cmd) == "" {
			continue
		}
		
		// 在 macOS/Linux 上使用 sh -c 来执行命令
		cmdObj = exec.Command("sh", "-c", cmd)
		cmdObj.Stdout = os.Stdout
		cmdObj.Stderr = os.Stderr
		cmdObj.Dir = r.Workspace // 使用 Dir 而不是 Path

		err := cmdObj.Run()
		if err != nil {
			return fmt.Errorf("run cmd %s failed, err: %w", cmd, err)
		}
	}

	return nil
}