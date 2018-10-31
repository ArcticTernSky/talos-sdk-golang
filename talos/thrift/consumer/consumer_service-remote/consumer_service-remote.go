// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"github.com/XiaoMi/talos-sdk-golang/thrift"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/XiaoMi/talos-sdk-golang/talos/thrift/consumer"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  LockPartitionResponse lockPartition(LockPartitionRequest request)")
	fmt.Fprintln(os.Stderr, "  LockWorkerResponse lockWorker(LockWorkerRequest request)")
	fmt.Fprintln(os.Stderr, "  void unlockPartition(UnlockPartitionRequest request)")
	fmt.Fprintln(os.Stderr, "  RenewResponse renew(RenewRequest request)")
	fmt.Fprintln(os.Stderr, "  UpdateOffsetResponse updateOffset(UpdateOffsetRequest request)")
	fmt.Fprintln(os.Stderr, "  QueryOffsetResponse queryOffset(QueryOffsetRequest request)")
	fmt.Fprintln(os.Stderr, "  QueryWorkerResponse queryWorker(QueryWorkerRequest request)")
	fmt.Fprintln(os.Stderr, "  Version getServiceVersion()")
	fmt.Fprintln(os.Stderr, "  void validClientVersion(Version clientVersion)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := consumer.NewConsumerServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "lockPartition":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "LockPartition requires 1 args")
			flag.Usage()
		}
		arg22 := flag.Arg(1)
		mbTrans23 := thrift.NewTMemoryBufferLen(len(arg22))
		defer mbTrans23.Close()
		_, err24 := mbTrans23.WriteString(arg22)
		if err24 != nil {
			Usage()
			return
		}
		factory25 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt26 := factory25.GetProtocol(mbTrans23)
		argvalue0 := consumer.NewLockPartitionRequest()
		err27 := argvalue0.Read(jsProt26)
		if err27 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.LockPartition(value0))
		fmt.Print("\n")
		break
	case "lockWorker":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "LockWorker requires 1 args")
			flag.Usage()
		}
		arg28 := flag.Arg(1)
		mbTrans29 := thrift.NewTMemoryBufferLen(len(arg28))
		defer mbTrans29.Close()
		_, err30 := mbTrans29.WriteString(arg28)
		if err30 != nil {
			Usage()
			return
		}
		factory31 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt32 := factory31.GetProtocol(mbTrans29)
		argvalue0 := consumer.NewLockWorkerRequest()
		err33 := argvalue0.Read(jsProt32)
		if err33 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.LockWorker(value0))
		fmt.Print("\n")
		break
	case "unlockPartition":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UnlockPartition requires 1 args")
			flag.Usage()
		}
		arg34 := flag.Arg(1)
		mbTrans35 := thrift.NewTMemoryBufferLen(len(arg34))
		defer mbTrans35.Close()
		_, err36 := mbTrans35.WriteString(arg34)
		if err36 != nil {
			Usage()
			return
		}
		factory37 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt38 := factory37.GetProtocol(mbTrans35)
		argvalue0 := consumer.NewUnlockPartitionRequest()
		err39 := argvalue0.Read(jsProt38)
		if err39 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.UnlockPartition(value0))
		fmt.Print("\n")
		break
	case "renew":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Renew requires 1 args")
			flag.Usage()
		}
		arg40 := flag.Arg(1)
		mbTrans41 := thrift.NewTMemoryBufferLen(len(arg40))
		defer mbTrans41.Close()
		_, err42 := mbTrans41.WriteString(arg40)
		if err42 != nil {
			Usage()
			return
		}
		factory43 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt44 := factory43.GetProtocol(mbTrans41)
		argvalue0 := consumer.NewRenewRequest()
		err45 := argvalue0.Read(jsProt44)
		if err45 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Renew(value0))
		fmt.Print("\n")
		break
	case "updateOffset":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateOffset requires 1 args")
			flag.Usage()
		}
		arg46 := flag.Arg(1)
		mbTrans47 := thrift.NewTMemoryBufferLen(len(arg46))
		defer mbTrans47.Close()
		_, err48 := mbTrans47.WriteString(arg46)
		if err48 != nil {
			Usage()
			return
		}
		factory49 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt50 := factory49.GetProtocol(mbTrans47)
		argvalue0 := consumer.NewUpdateOffsetRequest()
		err51 := argvalue0.Read(jsProt50)
		if err51 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.UpdateOffset(value0))
		fmt.Print("\n")
		break
	case "queryOffset":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "QueryOffset requires 1 args")
			flag.Usage()
		}
		arg52 := flag.Arg(1)
		mbTrans53 := thrift.NewTMemoryBufferLen(len(arg52))
		defer mbTrans53.Close()
		_, err54 := mbTrans53.WriteString(arg52)
		if err54 != nil {
			Usage()
			return
		}
		factory55 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt56 := factory55.GetProtocol(mbTrans53)
		argvalue0 := consumer.NewQueryOffsetRequest()
		err57 := argvalue0.Read(jsProt56)
		if err57 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.QueryOffset(value0))
		fmt.Print("\n")
		break
	case "queryWorker":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "QueryWorker requires 1 args")
			flag.Usage()
		}
		arg58 := flag.Arg(1)
		mbTrans59 := thrift.NewTMemoryBufferLen(len(arg58))
		defer mbTrans59.Close()
		_, err60 := mbTrans59.WriteString(arg58)
		if err60 != nil {
			Usage()
			return
		}
		factory61 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt62 := factory61.GetProtocol(mbTrans59)
		argvalue0 := consumer.NewQueryWorkerRequest()
		err63 := argvalue0.Read(jsProt62)
		if err63 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.QueryWorker(value0))
		fmt.Print("\n")
		break
	case "getServiceVersion":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetServiceVersion requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetServiceVersion())
		fmt.Print("\n")
		break
	case "validClientVersion":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ValidClientVersion requires 1 args")
			flag.Usage()
		}
		arg64 := flag.Arg(1)
		mbTrans65 := thrift.NewTMemoryBufferLen(len(arg64))
		defer mbTrans65.Close()
		_, err66 := mbTrans65.WriteString(arg64)
		if err66 != nil {
			Usage()
			return
		}
		factory67 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt68 := factory67.GetProtocol(mbTrans65)
		argvalue0 := consumer.NewVersion()
		err69 := argvalue0.Read(jsProt68)
		if err69 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ValidClientVersion(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
