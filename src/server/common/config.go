package common

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

//读取配置文件，并监控配置文件变化
func SetupConfig() {
	//go run main.go -conf ./config/default.json
	var (
		conf = flag.String("conf", "", "eg:go run main.go -conf ./config/default.json")
	)

	flag.Parse()
	viper.SetConfigFile(*conf)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config fail:", err.Error())
	}
	//监控文件改变
	watchConfig(*conf)
}

//监控文件变化,根据viper.WatchConfig()修改优化
//原生不支持软链接，而configMap文件为软链接
func watchConfig(path string) {
	filename, _ := filepath.Abs(path)
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		// we have to watch the entire directory to pick up renames/atomic saves in a cross-platform way
		if err != nil {
			log.Println("error:", err)
			return
		}

		configFile := filepath.Clean(filename)
		configDir := filepath.Dir(configFile)
		realconfigFile, _ := filepath.EvalSymlinks(configFile)
		realconfigDir := filepath.Dir(realconfigFile)
		done := make(chan bool)
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					log.Println("普通监控到变动", event.String(), "configName:", configFile)
					if filepath.Clean(event.Name) == configFile { // we care about the config file
						if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
							err := viper.ReadInConfig()
							if err != nil {
								log.Println("error:", err)
							}
						}
					} else if filepath.Clean(event.Name) == realconfigDir { // we also care about the REMOVE event on the realconfigDir
						// within a kubernetes pod, the file changes are registered as
						// "/config/..119811_10_11_11_11_41.697657017": CREATE
						// "/config/..119811_10_11_11_11_41.697657017": CHMOD
						// "/config/..data_tmp": CREATE
						// "/config/..data_tmp": RENAME
						// "/config/..data": CREATE
						// "/config/..119811_10_11_11_00_00.043158594": REMOVE
						if event.Op&fsnotify.Remove == fsnotify.Remove {
							//re-eval symlink
							realconfigFile, _ = filepath.EvalSymlinks(configFile)
							realconfigDir = filepath.Dir(realconfigFile)
							err := viper.ReadInConfig()
							if err != nil {
								log.Println("error:", err)
							}
							log.Println("configMap监控到变动,重新加载配置文件：",realconfigFile)
						}
					}
				case err := <-watcher.Errors:
					log.Println("error:", err)
				}
			}
		}()

		watcher.Add(configDir)
		<-done
	}()
}

//简单的文件监控示例
func ExampleNewWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./config")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
