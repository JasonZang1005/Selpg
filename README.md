#Selpg 
This is a program helping you select pages automatically


#Code description
- 程序思路

		参数解析
		参数逻辑判断
		文件处理

- 文件处理伪代码

	- 按页读取
	
			NumOfPage
			
			for{
				if 在目标页数范围内{
		            if 按页输入{
						if 无输出地址
							then 打印到屏幕
						else
							输出到目标位置
					}
				}
				NumOfPage++
			}
	- 按行读取
	
			NumOfLine
			
			for{
			    if 在目标行数范围内{

			        if 无输出地址
			            then 打印到屏幕
			        else
			            输出到目标位置
			     
			    }
			    NumOfLine++
			}
			

- 定义结构体 参数列表

		type Args struct {
			programName string
			startPos int
			endPos int
			pageType bool
			destination string
			src string
			pageLen int
		}
- main函数中，首先对于参数进行解析
 
		flag.IntVar(&arg.startPos,"s",-1,"start position")
		flag.IntVar(&arg.endPos,"e",-1,"end position")
		flag.IntVar(&arg.pageLen,"l",-1,"number of lines")
		flag.BoolVar(&arg.pageType,"f",false,"/f")
		flag.StringVar(&arg.destination,"d","","specify destination")

		flag.Parse()

#Examples
- 两者输出结果相同

		./selpg -s=0 -e=3 -l=3 test.txt
		./selpg -s=0 -e=3 -l=3 < test.txt

line 1

line 2

line 3

line 4

line 5

line 6

line 7

line 8

line 9

- 使用 other_command 

		./input | selpg -s=0 -e=3
	
	将input的标准输出重定向为selpg的标准输入

- 将输出结果重定向至output.txt

		./selpg -s=0 -e=3 -l=3  test.txt >output.txt
最终文本中结果显示与上相同	



- 将错误信息重定向至 error.txt ，并且将结果输出到output.txt
 
  		./selpg -s=4 -e=3 -l=4 test.txt >output.txt 2>error.txt

- 错误消息输出到error.txt		

		 $ selpg -s10 -e20 input_file 2>error_file

  	 最终error.txt文本显示

		start number should be smaller than end number

- 将错误信息输出到空设备，最终error.txt中无错误信息。
		
		selpg -s=0 -e3 test.txt >output.txt 2>/dev/null

- 将输出结果输出到空设备

		$ selpg -s=0 -e=3 test.txt >/dev/null
	
	test.text中无结果,最终结果被丢弃

- 标准输出被重定向至 ./out

		$ selpg -s=0 -e=3 test.txt | ./out

- 与上面类似，但error.txt中有错误信息
	
 		$ selpg -s=0 -e=3 test.txt 2>error.txt | ./out
