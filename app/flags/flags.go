// Package flags registers the user's command-line input
// as an object, which is parsed and used in `hashclock/cmd`
package flags

import "flag"

// CLIConfig struct defines the set configuration for hashclock
// in an object which is parsed and used in `hashclock/cmd`
type CLIConfig struct {
	Seed       string
	Hash       string
	Algorithm  string
	Iterations int
	Breakpoint int
	Timeout    int
	SetJSON    bool
}

// NewConfig function captures the set command-line flags and their
// values, and stores them in a `CLIConfig` object
func NewConfig() *CLIConfig {
	inputSeed := flag.String("seed", "", "Input seed which will be hashed")
	inputHash := flag.String("hash", "", "Input hash which will be verified, from hashing the seed")
	inputAlg := flag.String("alg", "sha256", "Hash function to use; lower-case or uppercase. One of: 'md5', 'sha1', 'sha224', 'sha256', 'sha384', 'sha512', 'sha512_224', 'sha512_256'")
	inputIterations := flag.Int("iter", 1, "Number of iterations")
	inputBreakpoint := flag.Int("log", 1, "Log hashes every # of steps")
	inputTimeout := flag.Int("time", 0, "Calculate hashes for # seconds")
	inputSetJSON := flag.Bool("json", false, "Returns the output in JSON format")

	flag.Parse()

	return &CLIConfig{
		Seed:       *inputSeed,
		Hash:       *inputHash,
		Algorithm:  *inputAlg,
		Iterations: *inputIterations,
		Breakpoint: *inputBreakpoint,
		Timeout:    *inputTimeout,
		SetJSON:    *inputSetJSON,
	}
}
