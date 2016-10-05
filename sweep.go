package main

var cmdSweep = &Command{
	Run:       runSweep,
	UsageLine: "sweep ",
	Short:     "",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdSweep.Flag.BoolVar(&flagA, "a", false, "")
}

// runSweep executes sweep command and return exit code.
func runSweep(args []string) int {

	return 0
}
