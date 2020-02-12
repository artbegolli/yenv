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

Yenv will look for any environment variables in `yaml` values with the format `${ENV_KEY_HERE}`

```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "github.com/artbegolli/yenv"
)

func main() {

    /* Example Yaml ./myyamlfile.yaml:
        age: 30
        name: John
        job: ${JOB_ENV}
    */
    
    // Set an environment variable
    os.Setenv("JOB_ENV", "software_stuff")
    yaml, _ := ioutil.ReadFile("./myyamlfile.yaml")

    yamlWithEnvApplied, err := yenv.ApplyEnvValues(yaml)
    if err != nil {
        fmt.Printf("err: %v\n", err)
        return
    }
    fmt.Println(string(yamlWithEnvApplied))
    /* Output
        age: 30
        name: John
        job: software_stuff
    */
}
```

You can also use yenv to unmarshall your `yaml` with the environment variables replaced. 
You can pass in your read-in `yaml` byte array and a struct to unmarshall.


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
