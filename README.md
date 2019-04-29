# goApp
该工程主要为个人一些golang小工具。目前包括以下三个子工程：

	windowFileCleaner
	lianjiaCrawler
	huaweiWebSnap
  
## windowFileCleaner 

	硬盘爆满后我们一般使用系统管理工具尝试清理文件，但是如果由于大文件较多还是需要手动去查找，这种情况下window
	下无法显示文件夹的大小，该工具辅助定位大文件位置，可以手动对一些大文件进行清理操作。

### 使用

        直接下载使用：下载src下windowFileCleaner.exe文件。
        源码构建：下载源码，src目录下通过go build -o windowFileCleaner.exe  -ldflags="-H windowsgui"进行构建

### 效果图：
	
  ![image](https://github.com/kmood/goApp/blob/master/windowFileCleaner/src/windowFileCleaner.png) 
