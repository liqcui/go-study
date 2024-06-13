package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Set an Environment Variable
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "admin")
	os.Setenv("DB_NAME", "test")

	// Get the value of an Environment Variable
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	fmt.Printf("Host: %s, Port: %s\n", host, port)

	// Unset an Environment Variable
	os.Unsetenv("DB_PORT")
	fmt.Printf("After unset, Port: %s\n", os.Getenv("DB_PORT"))

	/*
		Get the value of an environment variable and a boolean indicating whether the
		environment variable is present or not.
	*/
	driver, ok := os.LookupEnv("DB_DRIVER")
	if !ok {
		fmt.Println("DB_DRIVER is not present")
	} else {
		fmt.Printf("Driver: %s\n", driver)
	}

	// Expand a string containing environment variables in the form of $var or ${var}
	dbURL := os.ExpandEnv("postgres://$DB_USERNAME:$DB_PASSWORD@DB_HOST:$DB_PORT/$DB_NAME")
	fmt.Println("DB URL: ", dbURL)
	// Environ returns a copy of strings representing the environment,
	// in the form "key=value".
	for _, env := range os.Environ() {
		// env is
		envPair := strings.SplitN(env, "=", 2)
		key := envPair[0]
		value := envPair[1]

		fmt.Printf("%s : %s\n", key, value)
	}

	// Delete all environment variables
	os.Clearenv()

	fmt.Println("Number of environment variables: ", len(os.Environ()))

	SHARED_DIR, isFound := os.LookupEnv("SHARED_DIR")

	fmt.Printf("SHARED_DIR is %s and isFound is %v\n", SHARED_DIR, isFound)
	a := "abcd1244"
	fmt.Printf("%d\n", len(a))

	os.Setenv("SHARED_DIR", "/tmp/")

	SHARED_DIR = os.Getenv("SHARED_DIR")

	if len(SHARED_DIR) != 0 {
		fmt.Println("SHARED_DIR was found ")
		byteArray, err := ioutil.ReadFile(SHARED_DIR + "/cluster-type")
		if err != nil {
			log.Fatalf("Unable to read file: %v", err)
		}
		clusterType := string(byteArray)
		clusterType = strings.ToLower(clusterType)
		if strings.Contains(clusterType, "rosa") {
			return true
		}
	} else {
		fmt.Println("No SHARED_DIR was found")
	}

}
