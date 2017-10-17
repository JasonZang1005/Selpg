package main

import (
  "fmt"
  "os"
  "flag"
  "bufio"
  "io"
  "os/exec"
)

type Args struct {
  programName string
  startPos int
  endPos int
  pageType bool
  destination string
  src string
  pageLen int
}

func main() {
    var arg Args

    ParseArgs(&arg)
    processArgs(&arg)
    FileProcessing(&arg)

}

func ParseArgs(arg *Args) {
  arg.programName=os.Args[0]
  flag.IntVar(&arg.startPos,"s",-1,"start position")
  flag.IntVar(&arg.endPos,"e",-1,"end position")
  flag.IntVar(&arg.pageLen,"l",-1,"number of lines")
  flag.BoolVar(&arg.pageType,"f",false,"/f")
  flag.StringVar(&arg.destination,"d","","specify destination")

  flag.Parse()
}

func processArgs(arg *Args){

  if  arg.startPos<0 || arg.endPos<0 {
    showErr("invalid start number or end number")
  }

  if arg.startPos>arg.endPos {
    showErr("start number is should be smaller than end number")
  }

  arg.src = flag.Arg(0)

  if arg.pageType==true{
    if arg.pageLen!=-1{
        showErr("only one type is allowed")
    }
  }else{
    if arg.pageLen<1{
      arg.pageLen=72
    }
  }
}
func showErr( info string) {
  fmt.Println(info)
  os.Exit(1)
}
func FileProcessing(arg *Args){
  if flag.NArg()==1{
    arg.src=flag.Arg(0)
  }

  if arg.src==""{
    reader:=bufio.NewReader(os.Stdin)
    if arg.pageType==true{
      page(reader,arg)
    }else{
      line(reader,arg)
    }
  }else{
    file,err:=os.Open(arg.src)
    reader:=bufio.NewReader(file)
    CheckErr(err)
    if arg.pageType==true{
      page(reader,arg)
    }else{
      line(reader,arg)
    }
  }
}

func line(reader *bufio.Reader, arg *Args){
  NumOfLine :=1
  for {
    line,err:=reader.ReadString('\n')

    if NumOfLine>arg.pageLen*(arg.startPos-1)&&NumOfLine<=arg.pageLen*arg.endPos{
      if arg.destination==""{
        fmt.Println(line)
      }else{
        cmd := exec.Command("./out")
				pipe, err := cmd.StdinPipe()
        CheckErr(err)
				pipe.Write([]byte(line))
				pipe.Close()
				cmd.Stdout = os.Stdout
				cmd.Run()
      }
    }
    if err==io.EOF{
      break
    }
    NumOfLine++
  }
}

func page(reader *bufio.Reader, arg *Args){
  NumOfPage:=1
  for{
    page,err:=reader.ReadString('\f')

    if NumOfPage>=arg.startPos&&NumOfPage<=arg.endPos{
      if arg.destination==""{
        fmt.Println(page)
      }else{
          cmd := exec.Command("./out")          
  				pipe, err := cmd.StdinPipe()   
  				CheckErr(err)                          
  				pipe.Write([]byte(page + "\n")) 
  				pipe.Close()
  				cmd.Stdout = os.Stdout               
  				cmd.Run()
      }

      if err==io.EOF{
        break
      }
      NumOfPage++

    }
  }
}

func CheckErr(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}
