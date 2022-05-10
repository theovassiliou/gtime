# gtime

[![GoDoc](https://godoc.org/github.com/theovassiliou/gtime?status.svg)](https://godoc.org/github.com/theovassiliou/gtime)

Golang library that offers additional human-readable friendly functionality for working with time. 

## Description

Sometimes you have the need to just ask simply ask was a time today, or yesterday?

## Install

Install with

    go get github.com/theovassiliou/gtime

## Prerequisites
There is no particular requirement beyong the fact that you should have a working go installation.

Install Go >=1.13

## Example

    import "github.com/theovassiliou/gtime"

    today,_ := time.Parse("2006-01-02", "2022-05-22")
    timestamp1,_ := time.Parse("2006-01-02", "2022-05-21")

    fmt.Println("Timestamp is from " + gtime.HFDistanceApart(timestamp1, today))

    // Timestamp is from yesterday

    timestamp2,_ := time.Parse("2006-01-02", "2022-05-19")
    fmt.Println("Timestamp is from " + gtime.HFDistanceApart(timestamp2, today))

    // Timestamp is from 3 days ago



## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.
Furthermore, if you would like to submit pull-request to read also the following guidelines.

### Running the tests
We are using to different make targets for running tests.

    make test

    go test -short ./...
    ok      github.com/theovassiliou/gtime  0.088s
    ?       github.com/theovassiliou/gtime/apps/cmd [no test files]

executes all short package tests, while

    make test-all
    go vet $(go list ./...)
    go test ./...
    ?    github.com/theovassiliou/hc2-tools/cmd/expandRequire [no test files]
    ?    github.com/theovassiliou/hc2-tools/cmd/hc2DownloadScene [no test files]
    ?    github.com/theovassiliou/hc2-tools/cmd/hc2SceneInteract [no test files]
    ?    github.com/theovassiliou/hc2-tools/cmd/hc2UploadScene [no test files]
    ok   github.com/theovassiliou/hc2-tools/pkg 0.036s

executes in addition go veton the package. Before committing to the code base please use make test-all to ensure that all tests pass.

## Versioning
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the tags on this repository.

## Authors
Theo Vassiliou - Initial work - [Theo Vassiliou](https://github.com/theovassiliou)

## License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details

## Acknowledgments
Thanks to all the people out there that produce amazing open-source software, which supported the creation of this piece of software. And also [sweap.io](sweap.io) that needed this functionality.
