package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"
	"unsafe"
)

// new_cons 连接池
var new_conContain []net.Conn

const root_path string = "."

const MAX_BUF_LEN int = 1 * 1024 * 1024

// @func: IsLittleEndian
// @brief: 判断本机字节序
// @author: Kewin Li
// @return bool
func isLittleEndian() bool {
	n := 0x1234
	f := *((*byte)(unsafe.Pointer(&n)))
	return (f ^ 0x34) == 0
}

// @func: showProcess
// @brief: 展示文件传输进度
// @author: Kewin Li
// @param: int64 n
// @param: int64 total
func showProcess(n float64, total float64, rate float64) {

	var eql_count int = 60
	process := int(n / total * 100.0)
	bili := n / total
	var real_count int = int(float64(eql_count) * bili)

	fmt.Printf("[")

	for i := 0; i < real_count; i++ {
		fmt.Printf("=")
	}

	for j := 0; j < (eql_count - real_count); j++ {
		fmt.Printf(" ")
	}

	fmt.Printf("]")
	fmt.Printf("%d%%  %.2f Mb/s", process, rate)

	// time.Sleep(100 * time.Millisecond)
	// time.Sleep(time.Second)
	fmt.Printf("\r")
}

// @func: toByte
// @brief: 借助Json序列化为[]byte
// @author: Kewin Li
// @param: interface{} src
// @return []byte
func toByte_json(src *TransferFileInfo) []byte {

	buffer := bytes.Buffer{}

	// 1. json序列化
	data, err := json.Marshal(*src)
	if err != nil {
		fmt.Printf("json marshal err! %s \n", err.Error())
		return nil
	}

	f := isLittleEndian()

	if f {
		// 2. 转为网络字节序（大端）
		err = binary.Write(&buffer, binary.BigEndian, data)
		if err != nil {
			fmt.Printf("binary endian transfer err! %s \n", err.Error())
			return nil
		}
	}

	return buffer.Bytes()
}

// @func: fromByte_json
// @brief: 借助Json将[]byte进行反序列化
// @author: Kewin Li
// @param: []byte data
// @param: interface{} any
func fromByte_json(data []byte, any interface{}) error {

	buffer := bytes.Buffer{}

	// 判断本机的字节序
	f := isLittleEndian()

	if f {
		// 1. 转为本机字节序
		err := binary.Write(&buffer, binary.LittleEndian, data)
		if err != nil {
			fmt.Printf("binary endian transfer err! %s \n", err.Error())
			return err
		}
	}

	// 2. json反序列化
	err := json.Unmarshal(buffer.Bytes(), any)
	if err != nil {
		fmt.Printf("json unmarshal err! %s \n", err.Error())
		return err
	}

	return nil
}

// @func: toByte_gob
// @brief: 借助gob包序列化为[]byte
// @author: Kewin Li
func toByte_gob(src *TransferFileInfo) []byte {

	return nil
}

// @func: fromByte_gob
// @brief: 借助gob包将[]byte反序列化
// @author: Kewin Li
// @param: []byte data
// @return interface{}
func fromByte_gob(data []byte) interface{} {

	return nil

}

// @func: createTarFile
// @brief: 接收文件信息，创建目标文件
// @author: Kewin Li
// @param: net.Conn con
// @return *File
// @return error
func createTarFile(con net.Conn) (*os.File, TransferFileInfo) {
	file_info := TransferFileInfo{}
	file_info_size := int(unsafe.Sizeof(TransferFileInfo{}))
	read_size := 0

	var file_path string

	//// 方案1: []byte-->json-->struct
	// data := make([]byte, 1*1024*1024)

	// read_size, read_err := con.Read(data)
	// if read_err != nil {
	// 	fmt.Printf("Read err from [%s]! %s", con.RemoteAddr().String(), read_err.Error())
	// 	return nil, file_info
	// }

	// fmt.Printf("%s    read_size n=%d \n", data, read_size)

	// err := fromByte_json(data[:read_size], &file_info)
	// if err != nil {
	// 	fmt.Printf("fromByte_json err! %s \n", err.Error())
	// 	return nil, TransferFileInfo{}
	// }

	// fmt.Printf("read_size n=%d, file_info_size=%d, file_info=%+v \n", read_size, file_info_size, file_info)

	//// 方案2: []byte--[encoding/gob]-->struct
	dec := gob.NewDecoder(con)
	if dec == nil {
		fmt.Printf("new decoder err, is nil \n")
		return nil, file_info
	}

	err := dec.Decode(&file_info)
	if err != nil {
		fmt.Printf("decode data err! %s \n", err.Error())
		return nil, file_info
	}

	fmt.Printf("read_size n=%d, file_info_size=%d, file_info=%+v \n", read_size, file_info_size, file_info)

	if file_info.FileName == "" {
		file_info.FileName = "myRecvFile.dat"
	}

	file_path += root_path
	file_path += "/"
	file_path += file_info.FilePath
	file_path += "/"
	file_path += file_info.FileName

	fmt.Printf("file_path=%s \n", file_path)

	_, fs_err := os.Stat(file_path)
	if fs_err == nil {
		fmt.Printf("file: %s exist, and will delete \n", file_path)
		os.Remove(file_path)
	}

	f, op_err := os.OpenFile(file_path, os.O_CREATE|os.O_RDWR, 0777)
	if op_err != nil {
		fmt.Printf("OpenFile err! %s \n", op_err.Error())
		return nil, file_info
	}

	return f, file_info
}

// @func: writeTarFile
// @brief: 接收文件数据，写入文件
// @author: Kewin Li
// @param: net.Conn con
// @return error
func writeTarFile(con net.Conn, f *os.File, file_info TransferFileInfo) error {

	total := file_info.FileSize
	var recv int64
	var recv_rate float64
	var recv_rates []float64 = make([]float64, 0)
	var av_recv_rate float64

	cost_time := time.Now()

	fmt.Printf("file_size=%d\n", total)

	for recv = 0; recv < total; {
		buf := make([]byte, 100*1024*1024)

		rd_n, err := con.Read(buf)
		if err != nil {
			fmt.Printf("%s \n", err.Error())
			return err
		} else if rd_n > 0 {

			if time.Since(cost_time).Milliseconds() >= 1000 {
				recv_rate = float64(recv) / float64(1024) / float64(1024) / float64(time.Since(cost_time).Seconds())
				recv_rates = append(recv_rates, recv_rate)
			}

			for file_wr := 0; file_wr < rd_n; {
				n, fw_err := f.Write(buf[file_wr:rd_n])
				if fw_err != nil {
					fmt.Printf("%s\n", fw_err.Error())
					return fw_err
				}

				file_wr += n
			}

			// fmt.Printf("rd_n=%d, wr_n=%d \n", rd_n, n)
			recv += int64(rd_n)
			showProcess(float64(recv), float64(total), recv_rate)

		}

	}

	for _, val := range recv_rates {
		av_recv_rate += val
	}

	fmt.Printf("\n%s transfer finish! recv_size=%d, total size=%d, cost time=%d ms, av_rate=%.2fMb\\s \n", file_info.FileName, recv, file_info.FileSize, time.Since(cost_time).Milliseconds(), av_recv_rate/float64(len(recv_rates)))

	return nil
}

// @func: recvFile
// @brief: 接收客户端的文件
// @author: Kewin Li
// @param: net.Conn con
// @return error
func recvFile(con net.Conn) error {

	if con == nil {
		return errors.New("new connect is nil!")
	}

	_, wrErr := con.Write([]byte{'y'})
	if wrErr != nil {
		fmt.Printf("send to client 'y' err! \n")
		return wrErr
	}

	// 1. 接收文件信息，创建目标文件
	pFile, file_info := createTarFile(con)
	if pFile == nil {
		return errors.New("createTarFile err!")
	}

	// 2. 接收文件数据，写入文件
	err := writeTarFile(con, pFile, file_info)
	if err != nil {
		return errors.New("writeTarFile err!")
	}

	return nil
}

// @func: recvFile
// @brief:
// @author: Kewin Li
// @param: net.Conn con
func handleCon(con net.Conn) error {

	if con == nil {
		return errors.New("new connect is nil!")
	}
	f := make([]byte, 5)

	for {
		fmt.Printf("recv from [%s] file?(y/n):", con.RemoteAddr().String())

		_, inErr := os.Stdin.Read(f)
		if inErr != nil {
			fmt.Printf("Stdin err!! \n")
			return inErr
		}

		// fmt.Printf("inLen=%d, f=%s\n", inLen, f[:inLen])

		if f[0] == 'y' {

			return recvFile(con)

		} else if f[0] == 'n' {
			con.Write([]byte("n"))
			return errors.New("dont recv file!")
		} else {
			fmt.Printf("input invaild! input:'y'/'n' \n")
			continue
		}

	}

}

// @func: test_write_file
// @brief: 构造待传输的文件
// @author: Kewin Li
func test_write_file() {
	m_f, err := os.OpenFile("./test.dat", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("open file test.dat err! %s \n", err.Error())
		return
	}

	total := 0
	for i := 0; i < 100; i++ {
		tmp := fmt.Sprintf("ljk%d", i%10)
		buf := bytes.Repeat([]byte(tmp), 1*1024*1024)

		n, err := m_f.Write(buf)
		if err != nil {
			fmt.Printf("file test.data write err! %s \n", err.Error())
			return
		}
		total += n
	}

	fmt.Printf("test.dat write %d \n", total)

	m_f.Close()

}

// @func: test_process
// @brief: 构造传输进度条
// @author: Kewin Li
func test_process() {

	for i := 0; i <= 100; i++ {
		showProcess(float64(i), 100, 0)
	}

}

// @func: test_gob
// @brief: 测试编码/解码器
// @author: Kewin Li
func test_gob() {

	buffer := bytes.Buffer{}
	file_info := TransferFileInfo{}

	enc := gob.NewEncoder(&buffer)
	if enc == nil {
		fmt.Printf("new encoder err! %v \n", enc)
		return
	}

	denc := gob.NewDecoder(&buffer)
	if denc == nil {
		fmt.Printf("new dencoder err! %v \n", denc)
		return
	}

	err := enc.Encode(TransferFileInfo{123456, "666666.dat", "./"})
	if err != nil {
		fmt.Printf("encode err! %s \n", err.Error())
		return
	}

	err = denc.Decode(&file_info)
	if err != nil {
		fmt.Printf("decode err! %s \n", err.Error())
		return
	}

	fmt.Printf("TransferFileInfo: %+v \n", file_info)

}

func main() {

	args := os.Args

	if len(args) > 2 {
		fmt.Printf("input param err!\n")
		return
	} else if len(args) == 2 {

		switch args[1][0] {
		// 测试创建大文件
		case '1':
			test_write_file()
		// 测试传输进度条
		case '2':
			test_process()
		// 测试encoding/gob 编码/解码器
		case '3':
			test_gob()

		}

		fmt.Printf("test finish!\n")
		return
	}

	// 1. 监听
	listener, err := net.Listen("tcp", "192.168.77.136:8888")
	if err != nil {
		fmt.Printf("server Listen err! %s \n", err.Error())
		return
	}

	new_conContain = make([]net.Conn, 0, 100)

	for {
		fmt.Printf("start listen...\n")
		new_con, err := listener.Accept()
		if err != nil {
			fmt.Printf("server Accept err! %s \n", err.Error())
			return
		}

		err = handleCon(new_con)

		if err != nil {
			fmt.Printf("handleCon err=%s \n", err.Error())
		}

	}
}
