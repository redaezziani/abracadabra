package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "abracadabra",
		Short: color.BlueString(` █████╗ ██████╗ ██████╗  █████╗  ██████╗ █████╗ ██████╗  █████╗ ██████╗ ██████╗  █████╗ 
██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔══██╗
███████║██████╔╝██████╔╝███████║██║     ███████║██║  ██║███████║██████╔╝██████╔╝███████║
██╔══██║██╔══██╗██╔══██╗██╔══██║██║     ██╔══██║██║  ██║██╔══██║██╔══██╗██╔══██╗██╔══██║
██║  ██║██████╔╝██║  ██║██║  ██║╚██████╗██║  ██║██████╔╝██║  ██║██████╔╝██║  ██║██║  ██║
╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝
                                                                                        `),
	}

	// Tail logs of a systemd service
	var logsCmd = &cobra.Command{
		Use:   "logs [service]",
		Short: color.CyanString("\uf1c9  Tail logs of a systemd service"),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			service := args[0]
			out, err := exec.Command("journalctl", "-u", service, "-f").CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(out))
		},
	}

	// System info
	var sysinfoCmd = &cobra.Command{
		Use:   "sysinfo",
		Short: color.YellowString("\uf233  Show system info"),
		Run: func(cmd *cobra.Command, args []string) {
			c, _ := cpu.Percent(0, false)
			m, _ := mem.VirtualMemory()
			fmt.Printf("%s  CPU Usage: %.2f%%\n", color.GreenString("\uf109"), c[0])
			fmt.Printf("%s  RAM: %.2f GB / %.2f GB (%.0f%%)\n", color.BlueString("\uf538"),
				float64(m.Used)/1e9, float64(m.Total)/1e9, m.UsedPercent)
		},
	}

	// Disk usage
	var diskCmd = &cobra.Command{
		Use:   "disk",
		Short: color.MagentaString("\uf07b  Show disk usage"),
		Run: func(cmd *cobra.Command, args []string) {
			usage, _ := disk.Usage("/")
			fmt.Printf("%s  Disk: %.2f GB free / %.2f GB total (%.0f%% used)\n",
				color.RedString("\uf0a0"), float64(usage.Free)/1e9, float64(usage.Total)/1e9, usage.UsedPercent)
		},
	}

	// Clean Docker dangling images/containers
	var dockerCleanCmd = &cobra.Command{
		Use:   "docker-clean",
		Short: color.HiCyanString("\uf21b  Remove dangling Docker containers & images"),
		Run: func(cmd *cobra.Command, args []string) {
			exec.Command("docker", "system", "prune", "-f").Run()
			fmt.Println(color.HiCyanString("\uf21b  Docker cleaned!"))
		},
	}

	// List Docker containers
	var dockerListCmd = &cobra.Command{
		Use:   "docker-list",
		Short: color.HiGreenString("\uf1b2  List all Docker containers"),
		Run: func(cmd *cobra.Command, args []string) {
			out, err := exec.Command("docker", "ps", "-a").CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(out))
		},
	}

	// Logs of a Docker container
	var dockerLogsCmd = &cobra.Command{
		Use:   "docker-logs [container]",
		Short: color.CyanString("\uf1c9  Tail logs of a Docker container"),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			container := args[0]
			out, err := exec.Command("docker", "logs", "-f", container).CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(out))
		},
	}

	// Delete a Docker container
	var dockerRmCmd = &cobra.Command{
		Use:   "docker-rm [container]",
		Short: color.RedString("\uf00d  Remove a Docker container by name or ID"),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			container := args[0]
			out, err := exec.Command("docker", "rm", "-f", container).CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(color.RedString(string(out)))
		},
	}
	// Add commands serach of a systemd service
	var searchCmd = &cobra.Command{
		Use:   "search [service]",
		Short: color.YellowString("\uf002  Search for a systemd service"),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			service := args[0]	
			out, err := exec.Command("systemctl", "list-units", "--type=service", "--all").CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			lines := string(out)
			for _, line := range lines {
				if strings.Contains(string(line), service) {
					fmt.Println(line)
				}
			}
		},
	}

	rootCmd.AddCommand(logsCmd, sysinfoCmd, diskCmd, dockerCleanCmd,
		dockerListCmd, dockerLogsCmd, dockerRmCmd, searchCmd)
	rootCmd.Execute()
}
