package vault

import (
	"fmt"
	"os/exec"
	"strconv"
)

func SpinUpLoka(lokaid string, vaultName string, capacity int64, port string) error {

	if !VaultExists(vaultName) {
		if err := CreateVault(vaultName); err != nil {
			return fmt.Errorf("error creating volume: %v", err)
		}
	}

	cmd := exec.Command("docker", "run", "-d",
		"--name", lokaid,
		"--env", fmt.Sprintf("NODE_ID=%s", lokaid),
		"--env", fmt.Sprintf("STORAGE_CAPACITY=%s", strconv.FormatInt(capacity, 10)),
		"-v", fmt.Sprintf("%s:/data", vaultName),
		"-p", fmt.Sprintf("%s:%s", port, "8080"),
		"node-image")

	fmt.Println("Executing command:", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Output:", string(output))
		return fmt.Errorf("failed to run docker: %v", err)
	}

	fmt.Println("Container started:", string(output))
	return nil
}
