package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/tecmise/connector-school-api/pkg/domain/schools"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := f.File
			if pwd, err := os.Getwd(); err == nil {
				if rel, err := filepath.Rel(pwd, filename); err == nil {
					filename = rel
				}
			}

			funcName := path.Base(f.Function)
			return fmt.Sprintf("%s()", funcName), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	logrus.Debug("Log de depuração ativado")
}

func main() {
	fmt.Printf("Oi")
	var client schools.Client
	client = schools.Rest("http://localhost:3002")
	roles, err := client.PaginateSchools(context.TODO(), "", 1, 10, "")
	if err != nil {
		fmt.Printf("err: ", err)
	}
	fmt.Printf("Roles", roles)
}
