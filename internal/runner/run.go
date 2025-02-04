package runner

import (
	"bufio"
	"fmt"
	"fofax/internal/cli"
	"fofax/internal/fofa"
	"fofax/internal/fxparser"
	"fofax/internal/iconhash"
	"fofax/internal/printer"
	"fofax/internal/queue"
	"fofax/internal/utils"
	"os"
	"runtime"
	"strings"
	"sync"
)

type Runner struct {
	resMap     *sync.Map
	options    *cli.Options
	query      *queue.Queue
	inputCount int64
}

func NewRunner(options *cli.Options) (*Runner, error) {
	runner := &Runner{
		options: options,
		query:   queue.New(),
		resMap:  &sync.Map{},
	}
	// 标准输入
	if options.Stdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var fofaQuery string
			if runtime.GOOS == "windows" {
				fofaQuery = utils.ConvertByte2String(scanner.Bytes(), "GB18030")
				fofaQuery = strings.Trim(fofaQuery, " \t\r")
			} else {
				fofaQuery = strings.Trim(scanner.Text(), " \t\r")
			}
			if fofaQuery == "" {
				continue
			}

			//// 用浏览器打开
			//if options.Open {
			//	runner.openURL(fofaQuery)
			//	os.Exit(0)
			//}

			if options.FofaExt {
				fofaQuery = fxparser.Query(fofaQuery)
			}

			runner.query.Push(fofaQuery)
		}
	}
	// 单个 Query/cert/icon 搜索项 代码块
	{
		// query -q
		if len(options.Query) != 0 {
			runner.inputCount++

			//// 用浏览器打开
			//if options.Open {
			//	runner.openURL(options.Query)
			//	os.Exit(0)
			//}

			if options.FofaExt {
				options.Query = fxparser.Query(options.Query)
			}
			runner.query.Push(options.Query)
		}
		// 通过 url 查询证书 -uc
		if options.PeerCertificates != "" {
			//// 用浏览器打开
			//if options.Open {
			//	runner.openURL(utils.GetSerialNumber(options.PeerCertificates))
			//	os.Exit(0)
			//}
			runner.query.Push(utils.GetSerialNumber(options.PeerCertificates))
		}
		// 通 url 计算 hash，然后查询 -ui
		if options.UrlIcon != "" && strings.HasPrefix(options.UrlIcon, "http") {
			iconConfig := iconhash.NewIconHashConfig(options.UrlIcon, options.Debug)
			// 通过 url
			if iHash, err := iconConfig.FromUrlGetContent(); err == nil {
				runner.inputCount++
				//// 用浏览器打开
				//if options.Open {
				//	runner.openURL(iconConfig.MakeQuery(iHash))
				//	os.Exit(0)
				//}
				runner.query.Push(iconConfig.MakeQuery(iHash))
			}
		}
		// 通过文件，计算 icon hash 后进行查询 -if
		if options.IconFilePath != "" && utils.FileExist(options.IconFilePath) {
			iconConfig := iconhash.NewIconHashConfig("", options.Debug)
			iconConfig.HashFilePath = options.IconFilePath
			// 通过文件
			if iHash, err := iconConfig.FromFileGetContent(); err == nil {
				runner.inputCount++
				//// 用浏览器打开
				//if options.Open {
				//	runner.openURL(iconConfig.MakeQuery(iHash))
				//	os.Exit(0)
				//}
				runner.query.Push(iconConfig.MakeQuery(iHash))
			}
		}

	}

	// 多个 Query/cert/icon 搜索项 代码块
	{
		// 加载文件，查询多个语句 -qf
		if len(options.QueryFile) != 0 && utils.FileExist(options.QueryFile) {
			input, err := os.Open(options.QueryFile)
			defer input.Close()
			if err != nil {
				printer.Errorf(printer.HandlerLine(fmt.Sprintf("Could not open targets file '%s': %s\n", options.QueryFile, err)))
				return nil, err
			}
			scanner := bufio.NewScanner(input)
			for scanner.Scan() {
				url := strings.TrimSpace(scanner.Text())
				if url == "" {
					continue
				}

				runner.inputCount++

				if options.FofaExt {
					url = fxparser.Query(url)
				}

				runner.query.Push(url)
			}
		}

		// 读取文件中的URL，计算 cert 后进行查询 -ufc
		if len(options.PeerCertificatesFile) != 0 && utils.FileExist(options.PeerCertificatesFile) {
			input, err := os.Open(options.PeerCertificatesFile)
			defer input.Close()
			if err != nil {
				printer.Errorf(printer.HandlerLine(fmt.Sprintf("Could not open targets file '%s': %s\n", options.PeerCertificatesFile, err)))
				return nil, err
			}
			scanner := bufio.NewScanner(input)
			for scanner.Scan() {
				url := strings.TrimSpace(scanner.Text())
				if url == "" {
					continue
				}
				runner.inputCount++
				runner.query.Push(utils.GetSerialNumber(url))
			}
		}

		// 读取文件中的URL，计算 icon hash 后进行查询 -iuf
		if len(options.UrlIconFile) != 0 && utils.FileExist(options.UrlIconFile) {
			input, err := os.Open(options.UrlIconFile)
			defer input.Close()
			if err != nil {
				printer.Errorf(printer.HandlerLine(fmt.Sprintf("Could not open targets file '%s': %s\n", options.UrlIconFile, err)))
				return nil, err
			}
			scanner := bufio.NewScanner(input)
			iconConfig := iconhash.NewIconHashConfig("", options.Debug)
			for scanner.Scan() {
				url := strings.TrimSpace(scanner.Text())
				if url == "" {
					continue
				}
				runner.inputCount++
				// 通过 url
				iconConfig.HashUrl = url
				if iHash, err := iconConfig.FromUrlGetContent(); err == nil {
					runner.inputCount++
					runner.query.Push(iconConfig.MakeQuery(iHash))
				}
			}
		}
	}

	// 过滤项目,
	{
		if options.ExcludeCountryCN {
			for i := 0; i < runner.query.Len(); i++ {
				if !runner.query.Any() {
					break
				}
				fofaQuery := runner.query.Peek()
				runner.query.Pop()
				fofaQuery = fofaQuery + ` && country!="CN"`
				runner.query.Push(fofaQuery)
			}
		}
		if options.Exclude {
			// 把放进去的
			for i := 0; i < runner.query.Len(); i++ {
				if !runner.query.Any() {
					break
				}
				fofaQuery := runner.query.Peek()
				runner.query.Pop()
				if !(strings.HasPrefix(fofaQuery, "(") && strings.HasSuffix(fofaQuery, ")")) {
					fofaQuery = "(" + fofaQuery + ")" + " && (is_honeypot=false && is_fraud=false)"
				} else {
					fofaQuery = fofaQuery + " && (is_honeypot=false && is_fraud=false)"
				}
				runner.query.Push(fofaQuery)
			}
		}

		// 用浏览器打开
		if options.Open {
			for i := 0; i < runner.query.Len(); i++ {
				if !runner.query.Any() {
					break
				}
				fofaQuery := runner.query.Peek()
				runner.query.Pop()
				runner.openURL(fofaQuery)
				if !options.NolimitOpen {
					os.Exit(0)
				}
			}
			os.Exit(0)
		}
	}

	return runner, nil
}

func (r *Runner) openURL(query string) {
	// 用浏览器打开
	printer.Successf("the query %s will be opened with a browser", query)
	utils.OpenFofa(query)

}

func (r *Runner) Run() *sync.Map {
	fo := fofa.NewFoFa(r.options)
	for i := 0; i < r.query.Len(); i++ {
		if !r.query.Any() {
			break
		}
		fofaQuery := r.query.Peek()
		// 提取完整的 hostInfo，带有 protocol -ffi
		if r.options.FetchFullHostInfo {
			fo.FetchFn = func(fields []string, allSize int32) bool {
				fullUrl, err := utils.NewFixUrl(
					utils.FixFullHostInfoScheme(fields))
				if err != nil {
					printer.Errorf("url.Parse %s", err)
					os.Exit(1)
				}
				r.resMap.Store(fullUrl.FixedHostInfo, nil)
				return true
			}
			fo.FetchFullHostInfo(fofaQuery)
			// 提取指定根域名的 title -fto
		} else if r.options.FetchTitlesOfDomain {
			fo.FetchFn = func(fields []string, allSize int32) bool {
				if host := fields[3]; len(host) > 0 {
					fullUrl, err := utils.NewFixUrl(
						utils.FixFullHostInfoScheme(fields))
					if err != nil {
						printer.Errorf("url.Parse %s", err)
						os.Exit(1)
					}
					r.resMap.LoadOrStore(fullUrl, fmt.Sprintf("[%s]", fields[4]))
				}
				return true
			}
			fo.FetchTitlesOfDomain(fofaQuery)
		} else {
			fo.FetchFn = func(fields []string, allSize int32) bool {
				fullUrl, err := utils.NewFixUrl(fields[0])
				if err != nil {
					printer.Errorf("url.Parse %s", err)
					os.Exit(1)
				}
				key := fullUrl.HostInfo
				r.resMap.LoadOrStore(key, nil)
				return true
			}
			fo.Fetch(fofaQuery)
		}
		r.query.Pop()
	}

	return r.resMap
}
