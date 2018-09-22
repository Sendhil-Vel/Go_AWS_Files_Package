/*
Package myawsfiles
// Copyright 2018 Sendhil Vel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

awsmanagefiles.go
Date 		: 19/09/2018
Comment 	: This is seed file for creating any go website.
Version 	: 1.0.9
by Sendhil Vel
*/
package myawsfiles

/*
	importing the packages
*/
import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
)

/*
	init : this is the function for init
*/
var logStatus bool
var awsAccessKey string
var awsSecretKey string
var awsRegion string
var orgFolderpath string
var uploadBucket string
var awsBaseEndpoint string

var sess *session.Session
var err error
var _svc *ses.SES

type sizer interface {
	Size() int64
}

/*
loginfo - This will be used to login information.
	If set true while package intialization, then everything will be logged
*/
func loginfo(info string) {
	if logStatus == true {
		fmt.Println(info)
	}
}

/*
ChangeFolder - This can be used to change the folder before uploading the files.
				As per request I add this function, as user wanted to upload in multiple folders
*/
func ChangeFolder(FolderName string) {
	orgFolderpath = FolderName
}

/*
UploadFileS3 - This is the actual function which will upload to aws
*/
func UploadFileS3(w http.ResponseWriter, r *http.Request) {
	folderpath := orgFolderpath
	if folderpath == "" {
		http.Error(w, "missing path query parameter", http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(32 << 30)

	dataFile, dataheader, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	if dataFile == nil {
		fmt.Fprintln(w, errors.New("No Data Found"))
		return
	}

	defer dataFile.Close()
	processFile(dataFile)

	fileHeader := make([]byte, dataFile.(sizer).Size())

	// Copy the headers into the FileHeader buffer
	_, err = dataFile.Read(fileHeader)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	// set position back to start.
	_, err = dataFile.Seek(0, 0)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	filename := dataheader.Filename
	if strings.Contains(filename, ".PDF") {
		filename = strings.Replace(filename, ".PDF", ".pdf", -1)
	}
	if strings.Contains(filename, "'") {
		filename = strings.Replace(filename, "'", "", -1)
	}

	folderpath = filepath.Clean(folderpath) + "/" + filename
	if strings.Contains(folderpath, "'") {
		folderpath = strings.Replace(folderpath, "'", "", -1)
	}
	buffer := make([]byte, 512)
	_, err = dataFile.Read(buffer)
	if err != nil {
		return
	}
	contenttype := http.DetectContentType(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	if contenttype == "application/pdf" {
		_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
			Bucket:        aws.String(uploadBucket),
			Key:           aws.String(folderpath),
			ACL:           aws.String("public-read"),
			Body:          bytes.NewReader(fileHeader),
			ContentLength: aws.Int64(dataFile.(sizer).Size()),
			ContentType:   aws.String("application/pdf"),
			Metadata: map[string]*string{
				"Key": aws.String(http.DetectContentType(fileHeader)),
			},
		})
	} else {
		_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
			Bucket:        aws.String(uploadBucket),
			Key:           aws.String(folderpath),
			ACL:           aws.String("public-read"),
			Body:          bytes.NewReader(fileHeader),
			ContentLength: aws.Int64(dataFile.(sizer).Size()),
			ContentType:   aws.String(contenttype),
			Metadata: map[string]*string{
				"Key": aws.String(http.DetectContentType(fileHeader)),
			},
		})
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if strings.Contains(folderpath, " ") {
		folderpath = strings.Replace(folderpath, " ", "+", -1)
	}

	msg := fmt.Sprintln("https://" + awsBaseEndpoint + "/" + uploadBucket + "/" + folderpath)

	fmt.Fprintln(w, msg)

}

/*
processFile for processing the file
*/
func processFile(f io.Reader) error {
	return nil
}

/*
initAws - This is a internal function which checks the values provided for connecting to AWS.
			USing the provided credentials, Try to connect to AWS
			Provide the connection status
*/
func initAws() {
	var token = ""

	if awsAccessKey == "" || awsSecretKey == "" || awsRegion == "" {
		loginfo("Missing Required credentials")
		return
	}

	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, token),
	})
	if err != nil {
		loginfo("bad credentials")
	} else {
		loginfo("AWS connection established")
	}

}

/*
init - this is function which confirms that package is properly included for use
*/
func init() {
	fmt.Println("From Package")
}

/*
SetupPackage - This will take all the parameters and setup channel
	logstatus : If you set this to true then information will be displayed as and when process is going on
	awsAccessKey : This is the AWS access Key
	awsSecretKey : This is the AWS secret Key
	awsRegion : This is the AWS region value
*/
func SetupPackage(vLogStatus bool, vAwsAccessKey string, vAwsSecretKey string, vAwsRegion string, vFolderPath string, vUploadBucket string, vAwsBaseEndpoint string) {
	/*
		Set value of variable logstatus
	*/
	logStatus = vLogStatus

	/*
		This will provide the values of passed parameter, if loginfo is set to true
	*/
	variables := fmt.Sprintf("vLogStatus : %t  :: vAwsAccessKey : %s :: vAwsSecretKey : %s :: vAwsRegion : %s :: vFolderPath : %s :: vUploadBucket : %s :: vAwsBaseEndpoint : %s", vLogStatus, vAwsAccessKey, vAwsSecretKey, vAwsRegion, vFolderPath, vUploadBucket, vAwsBaseEndpoint)
	loginfo(variables)

	/*
		Set value of variable awsAccessKey.
		If missing then message will printed and execution will be stopped
	*/
	if vAwsAccessKey == "" {
		loginfo("Missing Required credentials [vAwsAccessKey]")
		return
	} else {
		loginfo("Parameter exists [awsAccessKey]")
		awsAccessKey = vAwsAccessKey
	}

	/*
		Set value of variable awsSecretKey.
		If missing then message will printed and execution will be stopped
	*/
	if vAwsSecretKey == "" {
		loginfo("Missing Required credentials [vAwsSecretKey]")
		return
	} else {
		loginfo("Parameter exists [awsSecretKey]")
		awsSecretKey = vAwsSecretKey
	}

	/*
		Set value of variable awsRegion.
		If missing then message will printed and execution will be stopped
	*/
	if vAwsRegion == "" {
		loginfo("Missing Required credentials [vAwsRegion]")
		return
	} else {
		loginfo("Parameter exists [vAwsRegion]")
		awsRegion = vAwsRegion
	}

	/*
		Set value of variable orgFolderpath.
		If missing then message will printed and execution will be stopped
	*/
	if vFolderPath == "" {
		loginfo("Missing Required credentials [vFolderPath]")
		return
	} else {
		loginfo("Parameter exists [vFolderPath]")
		orgFolderpath = vFolderPath
	}

	/*
		Set value of variable awsBaseEndpoint.
		If missing then message will printed and execution will be stopped
	*/
	if vUploadBucket == "" {
		loginfo("Missing Required credentials [vUploadBucket]")
		return
	} else {
		loginfo("Parameter exists [vUploadBucket]")
		awsBaseEndpoint = vUploadBucket
	}

	/*
		Set value of variable uploadBucket.
		If missing then message will printed and execution will be stopped
	*/
	if vAwsBaseEndpoint == "" {
		loginfo("Missing Required credentials [vAwsBaseEndpoint]")
		return
	} else {
		loginfo("Parameter exists [vAwsBaseEndpoint]")
		uploadBucket = vAwsBaseEndpoint
	}

	/*
		Call the function initAws, which connect and verify the AWS account
	*/
	initAws()

	/*
		Create a AWS session
	*/
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	}))

	// create a ses session
	_svc = ses.New(sess)
}
