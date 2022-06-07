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
