package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
<<<<<<< HEAD
	"github.com/tecmise/connector-school-api/pkg/domain/cluster_user_profile"
=======
	"github.com/tecmise/connector-school-api/pkg/domain/permission_user_cluster"
	"os"
	"path"
	"path/filepath"
	"runtime"
>>>>>>> eb9c0eaf2ffabf3d07177995281c18fa0f6d8aff
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

<<<<<<< HEAD
	connect := cluster_user_profile.Lambda("school-api")
	list, err := connect.FindByUserId(context.TODO(), "04b81498-50c1-7054-42de-20dd97f07e30")
=======
			funcName := path.Base(f.Function)
			return fmt.Sprintf("%s()", funcName), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	logrus.Debug("Log de depuração ativado")
}

func main() {
	connect := permission_user_cluster.NewClient("school-api")
	list, err := connect.FindByUserID(context.TODO(), "04b81498-50c1-7054-42de-20dd97f07e30")
>>>>>>> eb9c0eaf2ffabf3d07177995281c18fa0f6d8aff
	if err != nil {
		panic(err)
	}
	for _, item := range list {
		logrus.Debugf("%+v", item)
	}
}
