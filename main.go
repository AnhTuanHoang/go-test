package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gonum/matrix/mat64"
	"github.com/google/uuid"
	"github.com/kballard/go-shellquote"
	"github.com/spiral/goridge"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/sys/unix"
	"image"
	"image/jpeg"
	"io"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"test-func/pkg/ffprobe"
	"test-func/pkg/gocron"
	digest_auth_client "test-func/pkg/httpdigest"
	"test-func/pkg/log/syslog"
	"test-func/pkg/onvif"
	wsdiscovery "test-func/pkg/onvif-20/ws-discovery"
	"test-func/pkg/utils"
	"time"
	"unsafe"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type EventArgs struct {
	CameraId    string    `json:"camera_id"`
	EventCode   string    `json:"event_code"`
	TriggerTime time.Time `json:"trigger_time"`
	Status      string    `json:"status"`
	EquipIndex  uint8     `json:"equip_index"`
}

var OUTPUT_PATH = "output_file"
var INPUT_PATH = "input_file"
var GET_PARAM_VIVOTEK  = "http://%s:%s/cgi-bin/anonymous/getparam.cgi?%s"
var CAMERA_MODEL_PARAM = "system_info"
var FileMode os.FileMode

var MeomeoSyslog *syslog.Writer
var MeomeoSyslog2 *syslog.Writer

var scheduler *gocron.Scheduler

type JackPortState struct {
	DeviceID  string `json:"device_id"`
	EventCode string `json:"event_code"`
	Status    bool   `json:"status"`
	EquipIdx  uint8  `json:"equip_idx"`
}

func main() {

}

func bToS(input bool) string {
	if input {
		return "1"
	}
	return "0"
}

func sToB(input string) bool {
	return input == "1"
}

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}


func areArraysEqual(arr1, arr2 [][]string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if !reflect.DeepEqual(arr1[i], arr2[i]) {
			return false
		}
	}

	return true
}

func execCommandAndGetOutput(cmdparts []string) (result string) {
	cmd := exec.Command(cmdparts[0], cmdparts[1:]...)
	var outputBuf bytes.Buffer
	var stdErr bytes.Buffer

	cmd.Stdout = &outputBuf
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		fmt.Println("meomeo:", err.Error())
		return ""
	}
	result = string(outputBuf.Bytes())
	return
}

func execCommandWInputAndReturn(cmdparts []string, stdin string, timeout time.Duration) (isErr bool) {
	ctx, cancelFn := context.WithTimeout(context.Background(), timeout)
	cmd := exec.CommandContext(ctx, cmdparts[0], cmdparts[1:]...)
	cmd.Env = append([]string(nil), os.Environ()...)

	var timer *time.Timer
	isErr = true
	defer func() {
		cancelFn()
		if timer != nil {
			timer.Stop()
		}
	}()
	timer = time.AfterFunc(timeout, func() {
		timer.Stop()
		pid := fmt.Sprintf("%d", cmd.Process.Pid)
		cancelFn()
		execCommandWaitingDone([]string{"kill", "-9", pid})
		return
	})
	var outputBuf bytes.Buffer
	var stdErr bytes.Buffer

	cmd.Stdout = &outputBuf
	cmd.Stderr = &stdErr

	cmd.Stdin = strings.NewReader(stdin + "\n")

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := string(outputBuf.Bytes())
	r := bufio.NewScanner(strings.NewReader(str))
	for r.Scan() {
		lineStr := r.Text()
		fmt.Println(lineStr)
		if strings.HasPrefix(lineStr, "JBOD mode enabled") || strings.HasPrefix(lineStr, "JBOD mode disabled"){
			isErr = false
		}
	}
	return
}

func computeIntensiveTask() {
	const numTasks = 1000000
	var result float64

	for i := 0; i < numTasks; i++ {
		result += math.Sqrt(float64(i))
	}
}

func getRandDuration() int {
	min := -10
	max := 20
	val := rand.Intn(max - min) + min
	return 10000 + val
}

func execCommandWInput(cmdparts []string, stdin string) {
	cmd := exec.Command(cmdparts[0], cmdparts[1:]...)
	cmd.Env = append([]string(nil), os.Environ()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = strings.NewReader(stdin + "\n")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmdDone := make(chan int)
	go func() {
		cmdDone <- func() int {
			err := cmd.Wait()
			if err == nil {
				return 0
			}
			ee, ok := err.(*exec.ExitError)
			if !ok {
				return 0
			}
			return ee.ExitCode()
		}()
	}()

	select {
	case _ = <-cmdDone:
		return
	}
}

func execCommandWaitingDone(cmdparts []string) {
	cmd := exec.Command(cmdparts[0], cmdparts[1:]...)
	cmd.Env = append([]string(nil), os.Environ()...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmdDone := make(chan int)
	go func() {
		cmdDone <- func() int {
			err := cmd.Wait()
			if err == nil {
				return 0
			}
			ee, ok := err.(*exec.ExitError)
			if !ok {
				return 0
			}
			return ee.ExitCode()
		}()
	}()

	select {
	case _ = <-cmdDone:
		return
	}
}

func getListDisk(idx uint8, maxTry uint8) []utils.PartitionStat {
	var listDisk []utils.PartitionStat
	listDisk = utils.GetAllDrive()
	if len(listDisk) == 0 || idx < 3 {
		fmt.Println("mounted drive is empty")
		time.Sleep(1 * time.Second)
		idx++
		getListDisk(idx, maxTry)
	}
	return listDisk
}

func GetUsbPort(udevLabel string) (rs uint8) {
	commandStr := fmt.Sprintf("udevadm info --query=property --name %s | grep -e \"ID_VENDOR_ID=\" -e \"ID_MODEL_ID=\"", udevLabel)
	cmd := exec.Command("/bin/sh", "-c", commandStr)
	output, _ := cmd.CombinedOutput()
	bytesReader := bytes.NewReader(output)
	bufReader := bufio.NewReader(bytesReader)
	var idVendor, idModel string
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			break
		}
		lineStr := string(line)
		if strings.HasPrefix(lineStr, "ID_VENDOR_ID") {
			idVendor = strings.Split(lineStr, "=")[1]
		}
		if strings.HasPrefix(lineStr, "ID_MODEL_ID") {
			idModel = strings.Split(lineStr, "=")[1]
		}
	}

	busPath := HostSys("/bus/usb/devices")
	files, err := os.ReadDir(busPath)
	if err != nil {
		fmt.Println("getUsbPort readDir error: ", err.Error())
		return
	}
	for _, file := range files {
		if file.Type() == os.ModeSymlink && !strings.Contains(file.Name(), ":") {
			idVendorPath := filepath.Join(busPath, file.Name(), "idVendor")
			idProductPath := filepath.Join(busPath, file.Name(), "idProduct")
			if !Exists(idVendorPath) || !Exists(idProductPath) {
				continue
			}
			idVendorLine, _ := readLines(idVendorPath)
			idProductLine, _ := readLines(idProductPath)
			var usbPort int
			usbConInfo := strings.Split(file.Name(), "-")
			if len(usbConInfo) == 2 {
				usbPort, _ = strconv.Atoi(usbConInfo[1])
			} else {
				continue
			}
			if usbPort == 0 {
				continue
			}
			if len(idVendorLine) > 0 && len(idProductLine) > 0 && idVendor == idVendorLine[0] && idModel == idProductLine[0] {
				rs = uint8(usbPort)
				return
			}
		}
	}
	return
}

func HostSys(combineWith ...string) string {
	return getEnv("HOST_SYS", "/sys", combineWith...)
}

func discovery() {
	argsWithoutProg := os.Args[1:]
	ifName := "eno2"
	if len(argsWithoutProg) > 0 {
		ifName = argsWithoutProg[0]
	}
	a := wsdiscovery.GetAvailableDevicesAtSpecificEthernetInterface(ifName)
	fmt.Println(a)
}

func readLines(filename string) ([]string, error) {
	return readLinesOffsetN(filename, 0, -1)
}

func readLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()

	var ret []string

	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF && len(line) > 0 {
				ret = append(ret, strings.Trim(line, "\n"))
			}
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}

	return ret, nil
}

func getEnv(key string, dfault string, combineWith ...string) string {
	value := os.Getenv(key)
	if value == "" {
		value = dfault
	}

	switch len(combineWith) {
	case 0:
		return value
	case 1:
		return filepath.Join(value, combineWith[0])
	default:
		all := make([]string, len(combineWith)+1)
		all[0] = value
		copy(all[1:], combineWith)
		return filepath.Join(all...)
	}
}

func Move(sourcePath, destPath string, fileMode os.FileMode) error {
	sourceAbs, err := filepath.Abs(sourcePath)
	if err != nil {
		return err
	}
	destAbs, err := filepath.Abs(destPath)
	if err != nil {
		return err
	}
	if sourceAbs == destAbs {
		return nil
	}
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}

	destDir := filepath.Dir(destPath)
	if !Exists(destDir) {
		err = os.MkdirAll(destDir, fileMode)
		if err != nil {
			return err
		}
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return err
	}
	var byteWrited int64
	byteWrited, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	outputFile.Close()
	if err != nil || byteWrited == 0 {
		if errRem := os.Remove(destPath); errRem != nil {
			//filePath := config.AppConfig.LstErrFile
			fmt.Println(errRem)
		}
		return err
	}

	return nil
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func serveFrames(imgByte []byte) {
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("./img3.jpeg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 60

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
	}

}

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func padNumberWithZero(value uint32) string {
	return fmt.Sprintf("%02d", value)
}

func rpcCall(args any, method string, rpcServer string) (reply string) {
	conn, err := net.Dial("tcp", rpcServer)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := rpc.NewClientWithCodec(goridge.NewClientCodec(conn))
	defer func() {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	//var reply string
	err = client.Call(method, args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func DecryptAES(ciphertext string, passphrase []byte) string {
	arr := strings.Split(ciphertext, "-")
	if len(arr) != 3 {
		return ciphertext
	}
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := getKey(passphrase, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, _ = aesgcm.Open(nil, iv, data, nil)
	return string(data)
}

func getKey(passphrase []byte, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}
	return pbkdf2.Key(passphrase, salt, 1000, 32, sha256.New), salt
}
