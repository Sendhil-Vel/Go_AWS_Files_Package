# Go_AWS_Files_Package ![CI status](https://img.shields.io/badge/build-passing-brightgreen.svg)

Go_AWS_Files_Package is a Go Package which will upload your files in S3.

## Platforms Supported
* For all projects, it will be MacOS, Windows, and Linux
## Version
* 1.0.9

## Code style
Standard [![js-standard-style](https://img.shields.io/badge/code%20style-standard-brightgreen.svg?style=flat)](https://github.com/feross/standard)

## 

### Requirements
* Windows/Linux
* Go 1.9.2 and up

`$ Go version`

## Installation

```
//Download Github repo
$ git clone https://github.com/Sendhil-Vel/Go_AWS_Files_Package

//Import package
$ go get github.com/Sendhil-Vel/Go_AWS_Files_Package


```

## Development
```
$ go get github.com/Sendhil-Vel/Go_AWS_Files_Package
```

## How to use?
* awsmanagefiles.go is the core package files which uploades files to S3
* 
* Folder verifiedCode contains
*   1) api_package_code.go : this starts a Go API, which has Go_AWS_Files_Package imported.
*   2) website folder contains a simple website code which is uploading files to AWS
* 
* Steps
*   i) go get github.com/Sendhil-Vel/Go_AWS_Files_Package<br/>
    ii) git clone https://github.com/Sendhil-Vel/Go_AWS_Files_Package<br/>
    iii) cd Go_AWS_Files_Package\verifiedCode\GoAPI     //change to folder containing the sample API code.<br/>
    iv) provide all the information needed in the .env.<br/>
    v) go run api_package_code.go //start the Go API using.<br/>     
    vi) In case of error you can set statusFlag to true and get log. You can contact me (Check contact info at bottom).<br/><br/>


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Contact
* Mobile : 91-9924092297
* Mail : [myemail.msamvel@gmail.com](Mail:myemail.msamvel@gmail.com)
## License
[MIT] © [Sendhil Vel](Mail:myemail.msamvel@gmail.com)