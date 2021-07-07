package logging

import (
	"fmt"

	logrus "github.com/sirupsen/logrus"
)

func Logrus(loglevel string) {
	level, err := logrus.ParseLevel(loglevel)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetLevel(level)

	// Levels
	//
	logrus.Println("my first logrus message") // loglevel: info
	logrus.Trace("going to mars")
	logrus.Debug("connected, received buffer from worker")
	logrus.Info("connection successful to db")
	logrus.Warn("something went wrong with x")
	logrus.Error("an error occurred in worker x")
	panicLogrus()
	fmt.Println("continuing after panic")
	// logrus.Fatal("impossible to continue exec")

	// Fields
	//
	logrus.WithFields(logrus.Fields{"source": "worker"}).Warn("available disk space low (1%)")
}

func panicLogrus() {
	defer func() {
		if r := recover(); r != nil {
			logrus.Println("recovered from panic in panicLog()")
		}
	}()

	logrus.Panic("panicking now!")

	fmt.Println("after the panic")
}
