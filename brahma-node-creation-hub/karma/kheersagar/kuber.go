package kheersagar

import (
	"fmt"
	"os/exec"
)

func VaultExists(volumeName string) bool {
	cmd := exec.Command("docker", "volume", "inspect", volumeName)
	err := cmd.Run()
	return err == nil // If there is no error the volume exist
}

func CreateVault(volumeName string) error {
	cmd := exec.Command("docker", "volume", "create", volumeName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error creating volume: %w, output: %s", err, string(output))
	}
	fmt.Println("volume created", string(output))
	return nil
}
