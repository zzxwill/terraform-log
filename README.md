# Terraform Log Stripper
Terraform Log Stripper is a library to strip color codes from Terraform log

## Terraform color codes

By default, Terraform logs are colored with some of [these ANSI codes](https://gist.github.com/JBlond/2fea43a3049b38287e5e9cefc87b2124).

Here is a typical example.

![](./resources/terraform-log-with-color.png)

It's not friendly in text.

```text
\e[31m│\e[0m \e[0m\e[1m\e[31mError: \e[0m\e[0m\e[1mFailed to get existing
        workspaces: S3 bucket does not exist.\n\e[31m│\e[0m \e[0m\n\e[31m│\e[0m \e[0mThe
        referenced S3 bucket must have been previously created. If the S3 bucket\n\e[31m│\e[0m
        \e[0mwas created within the last minute, please wait for a minute or two and
        try\n\e[31m│\e[0m \e[0magain.\n\e[31m│\e[0m \e[0m\n\e[31m│\e[0m \e[0mError:
        NoSuchBucket: The specified bucket does not exist\n\e[31m│\e[0m \e[0m\tstatus
        code: 404, request id: B6MCP0K6S53G5B5G, host id: nOL4Tt+mZQJzOLAkpVyxRyfxNSotk7pvWo1IgQc5XA7ILoaipLa/umWFPosNztcY+6i7TMu5Znc=\n\e[31m│\e[0m
        \e[0m\e[0m\n\e[31m│\e[0m \e[0m\n\e[31m│\e[0m \e[0m\e[0m\n\e[31m╵\e[0m\e[0m\n\e[0m\e[0m\n
```

## Scenario

1) To keep the Terraform log colorful which will help to identify the execution result and
2) to let the status of the execution in text clear (for example, expose the error message of the log in the status of a CRD),
we need to remove the color codes from the log.

## Example

Run the example in [main.go](./examples/main.go), get the clean log string.

```go
package main

import (
	"fmt"
	stripper "github.com/zzxwill/terraform-log-stripper"
)

func main() {
	log := `[0m[1mInitializing the backend...[0m
[31m[31m╷[0m[0m
[31m│[0m [0m[1m[31mError: [0m[0m[1mFailed to get existing workspaces: S3 bucket does not exist.
[31m│[0m [0m
[31m│[0m [0mThe referenced S3 bucket must have been previously created. If the S3 bucket
[31m│[0m [0mwas created within the last minute, please wait for a minute or two and try
[31m│[0m [0magain.
[31m│[0m [0m
[31m│[0m [0mError: NoSuchBucket: The specified bucket does not exist
[31m│[0m [0m	status code: 404, request id: T23JBYQHKSD2HB1M, host id: +jucLhx7qpWdLimDkwjxKf4P8BU7hv4zy3i7qmRlPFSYvOCmHDcG1m+6VBkQw8PUT9GTe3ybfBw=
[31m│[0m [0m[0m
[31m│[0m [0m
[31m│[0m [0m[0m
[31m╵[0m[0m
[0m[0m`

	str := stripper.StripColor(log)
	fmt.Print(str)
}

```

The result is as below:

```shell
Initializing the backend...
╷
│ Error: Failed to get existing workspaces: S3 bucket does not exist.
│ 
│ The referenced S3 bucket must have been previously created. If the S3 bucket
│ was created within the last minute, please wait for a minute or two and try
│ again.
│ 
│ Error: NoSuchBucket: The specified bucket does not exist
│ 	status code: 404, request id: T23JBYQHKSD2HB1M, host id: +jucLhx7qpWdLimDkwjxKf4P8BU7hv4zy3i7qmRlPFSYvOCmHDcG1m+6VBkQw8PUT9GTe3ybfBw=
│ 
│ 
│ 
╵
```

## Acknowledgement

This project is inspired by [stripansi](https://github.com/acarl005/stripansi).
