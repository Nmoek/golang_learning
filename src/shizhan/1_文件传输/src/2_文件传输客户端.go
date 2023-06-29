package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"unsafe"
)

// @func: IsLittleEndian
// @brief: 判断本机字节序
// @author: Kewin Li
// @return bool
func isLittleEndian() bool {
	n := 0x1234
	f := *((*byte)(unsafe.Pointer(&n)))
	return (f ^ 0x34) == 0
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

	// TODO: 判断本机的字节序
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

// @func: sendFileInfo
// @brief: 发送文件信息
// @author: Kewin Li
// @param: net.Conn con
// @param: TransferFileInfo file_info
// @return error
func sendFileInfo(con net.Conn, file_info TransferFileInfo) error {

	//// 方案1: struct-->json-->[]byte
	// data := toByte_json(&file_info)
	// if data == nil {
	// 	fmt.Printf("toByte_json err! \n")

	// }

	// fmt.Printf("%s\n", data)

	// n, err := con.Write(data)
	// if err != nil {
	// 	return err
	// }

	// fmt.Printf("sent=%d\n", n)

	//// 方案2: struct--[encoding/gob]-->[]byte

	enc := gob.NewEncoder(con)
	if enc == nil {
		return errors.New("new encoder err, is nil!\n")
	}

	err := enc.Encode(file_info)
	if err != nil {
		return err
	}

	return nil

}

// @func: sendFileData
// @brief: 发送文件数据
// @author: Kewin Li
// @param: net.Conn con
// @param: TransferFileInfo file_info
// @return error
func sendFileData(con net.Conn, file_info TransferFileInfo) error {

	var sent int64 = 0
	total := file_info.FileSize

	pF, f_err := os.OpenFile("./"+file_info.FileName, os.O_RDONLY, 0777)
	if f_err != nil {
		return f_err
	}

	fmt.Printf("open %s succss! \n", file_info.FileName)

	for sent < total {
		buf := make([]byte, 100*1024*1024)

		rd_n, err := pF.Read(buf)
		if err != nil && err != io.EOF {
			return err
		} else if rd_n > 0 {

			fmt.Printf("rd_n=%d ", rd_n)

			for con_wr := 0; con_wr < rd_n; {
				n, err := con.Write(buf[con_wr:rd_n])
				if err != nil {
					return err
				}
				con_wr += n

			}

			fmt.Printf("con_wr_n=%d \n", rd_n)

			// time.Sleep(500 * time.Millisecond)
			sent += int64(rd_n)

		}

	}
	pF.Close()

	return nil

}

func main() {

	args := os.Args

	if len(args) < 3 {
		fmt.Printf("input: ./xxx [server IP] [file]\n")
		return
	}

	serverIp := args[1]
	inputName := args[2]
	intputPath := ""
	if len(args) == 4 {
		intputPath = args[3]
	}

	fmt.Printf("serverIp=%s, inputName=%s, intputPath=%s\n", serverIp, inputName, intputPath)

	// 1. 连接
	con, err := net.Dial("tcp", serverIp+":8888")
	if err != nil {
		fmt.Printf("connect server err! \n")
		return
	}

	fmt.Printf("connect server[%s] succss! \n", con.RemoteAddr().String())

	// 2. 等待服务器回执
	f := make([]byte, 5)
	_, err = con.Read(f)
	if err != nil {
		fmt.Printf("read err! %s \n", err.Error())
		return
	}

	fs, fs_err := os.Stat(inputName)
	if fs_err != nil {
		fmt.Printf("Stat err! %s \n", fs_err.Error())
		return
	}

	file_info := TransferFileInfo{FileName: inputName, FileSize: fs.Size(), FilePath: intputPath}

	fmt.Printf("%+v \n", file_info)

	if f[0] == 'y' {
		err := sendFileInfo(con, file_info)
		if err != nil {
			fmt.Printf("sendFileInfo err! %s\n", err.Error())
			return
		}

	} else if f[0] == 'n' {
		fmt.Printf("server %s confused to recv file! \n", con.RemoteAddr().String())
		return
	}

	// 3.发送文件数据
	err = sendFileData(con, file_info)
	if err != nil {
		fmt.Printf("sendFileData err! %s \n", err.Error())
		return
	}

	fmt.Printf("client sent file[%s] to server[%s] finish, size=%d \n", file_info.FileName, con.RemoteAddr().String(), file_info.FileSize)

}
