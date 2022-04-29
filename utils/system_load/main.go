package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func monitorCpuLoad(pid int) {
	timer := time.NewTicker(time.Duration(3) * time.Second)
	ex := "linkv-worker-a"

	for {
		select {
		case <-timer.C:

			var (
				out           bytes.Buffer
				stderr        bytes.Buffer
				cpuPercentMap = make(map[int]float64)
			)

			args := "top -b -n 1"
			fmt.Println("===== args:", args)
			cmd := exec.Command("bash", "-c", "top -b -n 1|grep linkv-worker-a")
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("===== err:", err, ",real err:",stderr.String())
				continue
			}

			fmt.Println("###### out:", out.String())
			continue
			//cmd := exec.Command("/bin/sh", "-c", `top -b -n 1 |grep `+ex)
			// 设置接收
			// cmd.Stdout = &out
			// cmd.Stderr = &stderr
			// 执行
			// err := cmd.Run()
			// if err != nil {
			// 	fmt.Println("====== err:", err)
			// 	continue
			// }

			//fmt.Println("out:", out.String())
			//fmt.Println("stderr:", stderr.String())

			// if err := cmd.Start(); err != nil {
			// 	// zap.S().Errorf("[monitorCpuLoad] cmd.Start() failed. err:%v", err)
			// 	// alert.SendAlarm(
			// 	// 	"Monitor_Cpu_Load_Fail",
			// 	// 	GetSfuWanIp(),
			// 	// 	"cmd start failed",
			// 	// 	err.Error(),
			// 	// 	"",
			// 	// )
			// 	continue
			// }

			//fmt.Println("----------- args:", cmd.Args)

			// if err := cmd.Wait(); err != nil {
			// 	// zap.S().Errorf("[monitorCpuLoad] cmd.Wait() failed. err:%v", err)
			// 	// alert.SendAlarm(
			// 	// 	"Monitor_Cpu_Load_Fail",
			// 	// 	GetSfuWanIp(),
			// 	// 	"cmd start failed",
			// 	// 	err.Error(),
			// 	// 	"",
			// 	// )
			// 	continue
			// }

			group := strings.Split(out.String(), "\n")
			fmt.Println("------ group:", group)

			for _, item := range group {
				fmt.Println("------ item:", item)

				if strings.Contains(item, ex) {
					fmt.Println("######### item:", item)
				} else {
					continue
				}

				line := strings.Join(strings.Fields(item), " ")
				fmt.Println("------ line:", line)
				fields := strings.Split(line, " ")
				fieldLen := len(fields)
				if fieldLen == 12 {
					pid, err := strconv.Atoi(fields[0])
					if err != nil {
						// zap.S().Errorf("[monitorCpuLoad] strconv.Atoi failed. err:%v", err)
						continue
					}

					percent, err := strconv.ParseFloat(fields[8], 64)
					if err != nil {
						// zap.S().Errorf("[monitorCpuLoad] strconv.ParseFloat failed. err:%v", err)
						continue
					}
					cpuPercentMap[pid] = percent
				}
			}

			for k, v := range cpuPercentMap {
				fmt.Println("===== k:", k, " v:", v)
			}

			// cmd := exec.Command("/bin/sh", "-c", `top -b -n 1 |grep linkv-worker-a`)
			// var (
			// 	out bytes.Buffer
			// )
			// cmd.Stdout = &out
			// cmd.Stderr = os.Stderr

			// if err := cmd.Start(); err != nil {
			// 	fmt.Println("----- cmd start failed. err:", err)
			// 	continue
			// }

			// if err := cmd.Wait(); err != nil {
			// 	fmt.Println("----- cmd wait failed. err:", err)
			// 	continue
			// }

			// fmt.Println()
			// //fmt.Println(out.String())
			// str := strings.Split(out.String(), "\n")
			// //fmt.Println(len(str))
			// for _, item := range str {
			// 	//fmt.Println("+++++ item:", item)
			// 	if item != "" {
			// 		s := strings.Join(strings.Fields(item), " ")
			// 		fmt.Println("+++++ s:", s)
			// 		slice := strings.Split(s, " ")
			// 		if len(slice) == 12 {
			// 			fmt.Println("==== pid:", slice[0])
			// 			fmt.Println("==== cpu:", slice[len(slice)-4])
			// 		}
			// 	}
			// }

			// _, err := exec.Command(, )
			// if err != nil {
			// 	logs.Errorf("[killAllWorker] exec cmd: killall %s failed. err:%v", conf.ExecFilePath(), err)
			// }

			// logs.Errorf("[KillAllWorker_] success")
		case <-time.After(time.Duration(2) * time.Second):
		}
	}
}

func main() {
	pid := flag.Int("pid", 0, "pid")
	flag.Parse()

	monitorCpuLoad(*pid)
}
