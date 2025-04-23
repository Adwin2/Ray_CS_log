package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	filemanager "github.com/codecrafters-io/redis-starter-go/app/FileManager"
)

var (
	// 缺省参数

	ifSecTime     = false
	firReplId     = "?"
	firReplOffset = "-1"

	master_repl_offset = "0"
	master_replid      = "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb"
)

type RedisServer interface {
	Start() error
	HandleCmd(net.Listener) error
}

func NewServer(cfg *ServerConfig) RedisServer {
	if cfg.role == "slave" {
		return &FollowerServer{cfg: cfg}
	} else {
		return &LeaderServer{cfg: cfg}
	}
}

type LeaderServer struct {
	cfg *ServerConfig
}

func (s *LeaderServer) Start() error {
	l, err := net.Listen("tcp", "0.0.0.0:"+s.cfg.port)
	if err != nil {
		log.Printf("Failed to bind to port %s : %s", s.cfg.port, err)
		return err
	}
	s.HandleCmd(l)
	return nil
}

func (s *LeaderServer) HandleCmd(l net.Listener) error {
	//
	defer l.Close()
	var kvMap, pxMap sync.Map
	fn := s.cfg.dir + "/" + s.cfg.dbfilename
	fmt.Println("Logs from your program will appear here!")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			return err
		}
		go func(conn net.Conn) {
			defer conn.Close()
			for {
				// Read per-connection
				buf := make([]byte, 1024)
				length, err := conn.Read(buf)
				if err != nil {
					log.Printf("Buf read Error: %#v\n", err)
					if err == io.EOF {
						log.Println("master : slave closed conn")
						return
					}
					continue
				}
				rawdata := string(buf[:length])
				res := strings.Split(rawdata, "\r\n")
				if strings.HasPrefix(res[0], "*") {
					// 获取传入字段切片长度
					elem_Len, err := strconv.Atoi(res[0][1:])
					if err != nil {
						return
					}
					// ECHO 命令
					if strings.EqualFold(res[1*2], "ECHO") == true {
						// 接受处理ECHO
						log.Printf("Received ECHO, Output:%s", res[2*2])
						conn.Write(simpleStringFmt(res[2*2]))
					} else if strings.EqualFold(res[1*2], "SET") == true {
						// SET 命令
						log.Printf("Received SET, setting %s to %s", res[2*2], res[3*2])
						if elem_Len > 3 && strings.EqualFold(res[4*2], "PX") == true {
							// SET... PX... 命令
							// 获取设定的过期时间
							exTime, _ := strconv.Atoi(res[5*2])
							// WIP: 俩Map的操作封装
							kvMap.Store(res[2*2], res[3*2])
							pxMap.Store(res[2*2], uint64(exTime))
							err = filemanager.UpdateRDB(fn, &kvMap, &pxMap)

							go filemanager.Expiry(exTime, &kvMap, &pxMap, res[2*2], fn)
						} else {
							// 只有SET
							kvMap.Store(res[2*2], res[3*2])
							err = filemanager.UpdateRDB(fn, &kvMap, &pxMap)
						}
						conn.Write([]byte("+OK\r\n"))
					} else if strings.EqualFold(res[1*2], "GET") == true {
						// GET <key> 命令
						// 通过 kvMap 的大小 ，判断是从文件读取 ，还是 从kvMap读取
						if filemanager.SizeMap(&kvMap) == 0 {
							kv, err := filemanager.GetRDBkeys(fn)
							if err != nil {
								log.Printf("GetRDBkeys func Wrong: %s", err)
								return
							}
							found := false
							for i := 0; i < len(kv); i += 2 {
								if kv[i] == res[2*2] {
									log.Printf("GET bulkStringFmt: %s", kv[i+1])
									conn.Write(bulkStringFmt(kv[i+1]))
									found = true
									break
								}
							}
							// 没有对应key
							if !found {
								conn.Write([]byte("$-1\r\n"))
							}
						} else {
							log.Printf("Received GET, getting %s", res[2*2])
							OP, ok := kvMap.Load(res[2*2])
							if !ok {
								log.Printf("ERROR GET: %s", res[2*2])
								conn.Write([]byte("$-1\r\n"))
							} else {
								// OP从Map中获取类型是any，需要类型断言
								conn.Write(simpleStringFmt(OP.(string)))
							}
						}
					} else if strings.EqualFold(res[1*2], "CONFIG") == true && strings.EqualFold(res[2*2], "GET") == true {
						// CONFIG GET 命令
						getName := res[3*2]
						// 反射获取结构体字段 即配置
						val := reflect.ValueOf(s.cfg).Elem().FieldByName(getName)
						getvalue := val.String()
						conn.Write(arrayFmt([]string{getName, getvalue}))
					} else if strings.EqualFold(res[1*2], "KEYS") == true {
						// KEYS 命令
						// 模式*  格式化返回 KVs
						if strings.EqualFold(res[2*2], "*") == true {
							kv, err := filemanager.GetRDBkeys(fn)
							// kv, err := filemanager.TmpParseKV(fn)
							// filemanager.ShowFile(fn)
							if err != nil {
								log.Printf("GetRDBkeys func Wrong: %s", err)
								return
							}
							// 只输出KEY 偶数位
							var kvs []string
							for i := 0; i < len(kv); i += 2 {
								kvs = append(kvs, kv[i])
								log.Printf("KEYS arrayFmt: %s", kv[i])
							}
							conn.Write(arrayFmt(kvs))
						}
					} else if strings.EqualFold(res[1*2], "PING") == true {
						// PING 命令
						log.Printf("Received PING")
						conn.Write(simpleStringFmt("PONG"))
					} else if strings.EqualFold(res[1*2], "INFO") == true {
						// INFO 命令
						info := fmt.Sprintf("role:master\r\nmaster_repl_offset:%s\r\nmaster_replid:%s\r\n", master_repl_offset, master_replid)
						conn.Write(bulkStringFmt(info))
					} else if strings.EqualFold(res[1*2], "REPLCONF") == true {
						// WIP 功能待补充

						conn.Write(simpleStringFmt("OK"))
					} else if strings.EqualFold(res[1*2], "PSYNC") == true {
						conn.Write(simpleStringFmt("FULLRESYNC" + " " + master_replid + " " + master_repl_offset))
						// receive empty file , need full RESUMC assign empty new dile as RDB file
						ioCopy(fn, conn) // return empty file
					}
				}
			}
		}(conn)
	}
}

func ioCopy(fn string, conn net.Conn) error {
	file, err := os.Open(fn)
	if err != nil {
		log.Printf("Open file Error: %s", err)
		return err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	fileSize := fileInfo.Size()
	s := fmt.Sprintf("$%d\r\n",fileSize)
	conn.Write([]byte(s))
	io.Copy(conn, file) // avoid MemCopy
}

type FollowerServer struct {
	cfg *ServerConfig
}

func (s *FollowerServer) Start() error {
	// 连接到主节点
	d, err := net.Dial("tcp", s.cfg.ReplicaOf.MasterHost+":"+s.cfg.ReplicaOf.MasterPort)
	if err != nil {
		log.Printf("Failed to dial port %s : %s", s.cfg.port, err)
		return err
	}
	// 方法中的goroutine使用结构体元素 可能会产生竞态 检测：go test -race
	// 解决 ：显式捕获当前值 并传入闭包
	go func(d net.Conn) {
		defer d.Close()
		// STEP I
		d.Write(arrayFmt([]string{"PING"}))
		for {
			buf := make([]byte, 1024)
			length, err := d.Read(buf)
			if err != nil {
				log.Printf("Buf read Error: %#v\n", err)
				if err == io.EOF {
					log.Println("slave : master closed conn")
					return
				}
				continue
			}
			rawdata := string(buf[:length])
			res := strings.Split(rawdata, "\r\n")
			// 消息处理
			// 处理 simpleString
			if strings.HasPrefix(res[0], "+") {
				// 处理 simple string
				str := res[0][1:]
				// 接收Master 的OK（ simpleString ）回应
				if strings.EqualFold(str, "PONG") == true {
					// STEP II
					d.Write(arrayFmt([]string{
						"REPLCONF",
						"listening-port",
						s.cfg.port,
					}))
				} else if strings.EqualFold(str, "OK") == true {
					if ifSecTime {
						d.Write(arrayFmt([]string{
							"PSYNC",
							firReplId,
							firReplOffset,
						}))
						continue
					}
					// STEP III
					d.Write(arrayFmt([]string{
						"REPLCONF",
						"capa",
						"psync2",
					}))
					ifSecTime = true
				}
			} else if strings.EqualFold(res[0], "FULLRESYNC") == true {
				// assign empty file as new RDB file

			}
			// } else if strings.HasPrefix(res[0], "*") {
			// 	// 处理array
			// 	// 获取传入字段切片长度
			// 	_, err := strconv.Atoi(res[0][1:])
			// 	if err != nil {
			// 		log.Printf("elem_Len Read Error: %s", err)
			// 		return
			// 	}
		}
	}(d)

	// Listen启动服务监听
	l, err := net.Listen("tcp", "0.0.0.0:"+s.cfg.port)
	if err != nil {
		log.Printf("Failed to bind to port %s : %s", s.cfg.port, err)
		return err
	}
	go s.HandleCmd(l)
	return nil
}

func (s *FollowerServer) HandleCmd(l net.Listener) error {
	defer l.Close()
	var kvMap, pxMap sync.Map
	fn := s.cfg.dir + "/" + s.cfg.dbfilename
	fmt.Println("Logs from your program will appear here!")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			return err
		}
		go func(conn net.Conn) {
			defer conn.Close()
			for {
				buf := make([]byte, 1024)
				length, err := conn.Read(buf)
				if err != nil {
					fmt.Printf("Buf read Error: %#v\n", err)
					return
				}
				rawdata := string(buf[:length])
				res := strings.Split(rawdata, "\r\n")
				if strings.HasPrefix(res[0], "*") {
					// 获取传入字段切片长度
					elem_Len, err := strconv.Atoi(res[0][1:])
					if err != nil {
						return
					}
					// ECHO 命令
					if strings.EqualFold(res[1*2], "ECHO") == true {
						// 接受处理ECHO
						log.Printf("Received ECHO, Output:%s", res[2*2])
						conn.Write(simpleStringFmt(res[2*2]))
					} else if strings.EqualFold(res[1*2], "SET") == true {
						// SET 命令
						log.Printf("Received SET, setting %s to %s", res[2*2], res[3*2])
						if elem_Len > 3 && strings.EqualFold(res[4*2], "PX") == true {
							// SET... PX... 命令
							// 获取设定的过期时间
							exTime, _ := strconv.Atoi(res[5*2])
							// WIP: 俩Map的操作封装
							kvMap.Store(res[2*2], res[3*2])
							pxMap.Store(res[2*2], uint64(exTime))
							err = filemanager.UpdateRDB(fn, &kvMap, &pxMap)

							go filemanager.Expiry(exTime, &kvMap, &pxMap, res[2*2], fn)
						} else {
							// 只有SET
							kvMap.Store(res[2*2], res[3*2])
							err = filemanager.UpdateRDB(fn, &kvMap, &pxMap)
						}
						conn.Write([]byte("+OK\r\n"))
					} else if strings.EqualFold(res[1*2], "GET") == true {
						// GET <key> 命令
						// 通过 kvMap 的大小 ，判断是从文件读取 ，还是 从kvMap读取
						if filemanager.SizeMap(&kvMap) == 0 {
							kv, err := filemanager.GetRDBkeys(fn)
							if err != nil {
								log.Printf("GetRDBkeys func Wrong: %s", err)
								return
							}
							found := false
							for i := 0; i < len(kv); i += 2 {
								if kv[i] == res[2*2] {
									log.Printf("GET bulkStringFmt: %s", kv[i+1])
									conn.Write(bulkStringFmt(kv[i+1]))
									found = true
									break
								}
							}
							// 没有对应key
							if !found {
								conn.Write([]byte("$-1\r\n"))
							}
						} else {
							log.Printf("Received GET, getting %s", res[2*2])
							OP, ok := kvMap.Load(res[2*2])
							if !ok {
								log.Printf("ERROR GET: %s", res[2*2])
								conn.Write([]byte("$-1\r\n"))
							} else {
								// OP从Map中获取类型是any，需要类型断言
								conn.Write(simpleStringFmt(OP.(string)))
							}
						}
					} else if strings.EqualFold(res[1*2], "CONFIG") == true && strings.EqualFold(res[2*2], "GET") == true {
						// CONFIG GET 命令
						getName := res[3*2]
						// 反射获取结构体字段 即配置
						val := reflect.ValueOf(s.cfg).Elem().FieldByName(getName)
						getvalue := val.String()
						conn.Write(arrayFmt([]string{getName, getvalue}))
					} else if strings.EqualFold(res[1*2], "KEYS") == true {
						// KEYS 命令
						// 模式*  格式化返回 KVs
						if strings.EqualFold(res[2*2], "*") == true {
							kv, err := filemanager.GetRDBkeys(fn)
							// kv, err := filemanager.TmpParseKV(fn)
							// filemanager.ShowFile(fn)
							if err != nil {
								log.Printf("GetRDBkeys func Wrong: %s", err)
								return
							}
							// 只输出KEY 偶数位
							var kvs []string
							for i := 0; i < len(kv); i += 2 {
								kvs = append(kvs, kv[i])
								log.Printf("KEYS arrayFmt: %s", kv[i])
							}
							conn.Write(arrayFmt(kvs))
						}
					} else if strings.EqualFold(res[1*2], "PING") == true {
						// PING 命令
						log.Printf("Received PING")
						conn.Write(simpleStringFmt("PONG"))
					} else if strings.EqualFold(res[1*2], "INFO") == true {
						// INFO 命令
						conn.Write(bulkStringFmt("role:slave"))
					}
				}
			}
		}(conn)
	}
}
