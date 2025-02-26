// 简单工厂模式：go 语言中 通过 New`name`函数来初始化相关类并返回接口。
package main

type Protocol string

var (
	Smbprotocol Protocol = "smb"
	NfsProtocol Protocol = "nfs"
)

type IDownload interface {
	Download()
}

type SmbDownloader struct{}

func (s *SmbDownloader) Download() {
	println("smb download")
}

type NfsDownloader struct{}

func (n *NfsDownloader) Download() {
	println("nfs download")
}

func NewDownloader(t Protocol) IDownload {
	switch t {
	case Smbprotocol:
		return &SmbDownloader{}
	case NfsProtocol:
		return &NfsDownloader{}
	}
	return nil
}

func main() {
	smbDownloader := NewDownloader(Smbprotocol)
	smbDownloader.Download()

	nfsDownloader := NewDownloader(NfsProtocol)
	nfsDownloader.Download()
}
