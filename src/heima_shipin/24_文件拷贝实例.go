package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args

	if len(args) != 4 {
		fmt.Printf("input param err! please input:m_cpoy [src] [dst] \n")
		return
	}

	if args[1] != "m_copy" {
		fmt.Printf("commond err!!")
		return
	}

	src := args[2]
	dst := args[3]

	if src == dst {
		fmt.Printf("src name is same with dst!!! \n")
		return
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		fmt.Printf("err=%s \n", err.Error())
		return
	}

	srcFile, err2 := os.OpenFile(src, os.O_RDONLY, 0666)
	if err2 != nil {
		fmt.Printf("err2=%s \n", err2.Error())
		return
	}

	dstFile, err3 := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0666)
	if err3 != nil {
		fmt.Printf("err3=%s \n", err3.Error())
		return
	}

	total := srcInfo.Size()
	buf := make([]byte, 1*1024*1024)

	for wr := int64(0); wr < total; {

		rd_n, rd_err := srcFile.Read(buf)
		if rd_err != nil && rd_err != io.EOF {
			fmt.Printf("rd_err:%s \n", rd_err.Error())
			return
		}

		wr_n, wr_err := dstFile.Write(buf[:rd_n])
		if wr_err != nil {
			fmt.Printf("wr_err:%s \n", wr_err.Error())
			return
		}

		wr += int64(wr_n)
	}

	srcFile.Close()
	dstFile.Close()

}
