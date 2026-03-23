// internal/src/utils.go

package src

import (
    "os/exec"
    "strings"
)

func DirBase() string {
    cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }

    return strings.TrimSpace(string(out))
}
