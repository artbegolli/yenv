# Environment variable support for Yaml in Go


## Introduction

Yenv allows you to support environment variables in a `yaml` file from within Go. Yenv takes an marshalled `yaml` object 
and applies the values of any system environment variables and returns the unmarshalled object.
 
## Installation

### Install
```shell script
go get github.com/artbegolli/yenv
``` 
 
### Import
```shell script
import "github.com/artbegolli/yenv"
``` 
 
 
## Usage

To use `yenv` you can pass in your read in Yaml byte array and a struct to unmarshall.

Yenv will look for any environment variables in `yaml` values with the format `${ENV_VALUE_HERE}`

```go
package main

import (
    "fmt"
    "os"
    "github.com/artbegolli/yenv"
)

func main() {

	/* Example Yaml:
	    age: 30
	    name: John
        job: ${JOB_ENV}
	*/
    
    // Set an environment variable
    os.Setenv("JOB_ENV", "software_stuff")

	// Unmarshal the YAML back into a Person struct.
	var p Person
	err := yenv.UnmarshallWithEnv(y, &p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(p)
	/* Output:
	{John 30 software_stuff}
	*/
}
```
