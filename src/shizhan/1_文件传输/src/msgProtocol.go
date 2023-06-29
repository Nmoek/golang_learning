package main

type MsgHead struct {
}

// transferFileInfo @brief: 待传输文件信息
type TransferFileInfo struct {
	FileSize int64  // 文件总大小
	FileName string // 文件名
	FilePath string // 指定存放文件路径
}

// transferFileData @brief: 待传输文件数据
type TransferFileData struct {
	Size int64 // 单次文件数据量
	// TODO: 生成一个哈希值唯一标识一个文件
	FileName string // 文件名
	Data     []byte // 文件数据
}
