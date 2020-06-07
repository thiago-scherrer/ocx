package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//SetupAWS generates AWS access key
func SetupAWS() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the AWS_ACCESS_KEY_ID: ")
	akey, _ := reader.ReadString('\n')

	fmt.Print("Enter the AWS_SECRET_ACCESS_KEY: ")
	asec, _ := reader.ReadString('\n')

	fmt.Print("Enter the AWS region: ")
	reg, _ := reader.ReadString('\n')

	if createDir() {
		createFile(akey, asec, reg)
	}
}

var (
	home, _ = os.UserHomeDir()
	awsdir  = home + "/.aws"
	awcred  = awsdir + "/credentials"
)

func createDir() bool {
	_, n := os.Stat(awsdir)

	if os.IsNotExist(n) {
		err := os.Mkdir(awsdir, 0700)
		if err != nil {
			log.Fatal("Error creating directory, got: ", err)
		}
	}
	return true
}

func createFile(akey, asec, reg string) bool {
	f1, err := os.Create(awcred)
	if err != nil {
		log.Fatal("Error creating file, got: ", err)
	}
	defer f1.Close()

	_, err = f1.WriteString("[default]\n")
	if err != nil {
		log.Fatal("Error writing on file, got: ", err)
	}

	_, err = f1.WriteString("aws_access_key_id=" + akey)
	if err != nil {
		log.Fatal("Error writing on file, got: ", err)
	}

	_, err = f1.WriteString("aws_secret_access_key=" + asec)
	if err != nil {
		log.Fatal("Error writing on file, got: ", err)
	}

	_, err = f1.WriteString("region=" + reg)
	if err != nil {
		log.Fatal("Error writing on file, got: ", err)
	}

	f1.Sync()
	return true
}
