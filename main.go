package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "abracadabra",
		Short: color.BlueString(` █████╗ ██████╗ ██████╗  █████╗  ██████╗ █████╗ ██████╗  █████╗ ██████╗ ██████╗  █████╗ 
██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔══██╗
███████║██████╔╝██████╔╝███████║██║     ███████║██║  ██║███████║██████╔╝██████╔╝███████║
██╔══██║██╔══██╗██╔══██╗██╔══██║██║     ██╔══██║██║  ██║██╔══██║██╔══██╗██╔══██╗██╔══██║
██║  ██║██████╔╝██║  ██║██║  ██║╚██████╗██║  ██║██████╔╝██║  ██║██████╔╝██║  ██║██║  ██║
╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝
                                                                                        `), 
	}

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

	var diskCmd = &cobra.Command{
		Use:   "disk",
		Short: color.MagentaString("\uf07b  Show disk usage"), 
		Run: func(cmd *cobra.Command, args []string) {
			usage, _ := disk.Usage("/")
			fmt.Printf("%s  Disk: %.2f GB free / %.2f GB total (%.0f%% used)\n",
				color.RedString("\uf0a0"), float64(usage.Free)/1e9, float64(usage.Total)/1e9, usage.UsedPercent) 
		},
	}

	var dockerCleanCmd = &cobra.Command{
		Use:   "docker-clean",
		Short: color.HiCyanString("\uf21b  Remove dangling Docker containers & images"), 
		Run: func(cmd *cobra.Command, args []string) {
			exec.Command("docker", "system", "prune", "-f").Run()
			fmt.Println(color.HiCyanString("\uf21b  Docker cleaned!"))
		},
	}

	rootCmd.AddCommand(logsCmd, sysinfoCmd, diskCmd, dockerCleanCmd)
	rootCmd.Execute()
}
