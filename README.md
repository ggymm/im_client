icon = "ui/chat.ico"
Icon: icon,


replace github.com/lxn/walk => D:/develope/projects/go/local/walk


<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
    <assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="SomeFunkyNameHere" type="win32"/>
    <dependency>
        <dependentAssembly>
            <assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
        </dependentAssembly>
    </dependency>
    <application xmlns="urn:schemas-microsoft-com:asm.v3">
        <windowsSettings>
            <dpiAwareness xmlns="http://schemas.microsoft.com/SMI/2016/WindowsSettings">PerMonitorV2, PerMonitor</dpiAwareness>
            <dpiAware xmlns="http://schemas.microsoft.com/SMI/2005/WindowsSettings">True</dpiAware>
        </windowsSettings>
    </application>
</assembly>


D:\develope\projects\go\bin\rsrc -manifest main.manifest -o im_client.syso


go build -ldflags="-H windowsgui"



Run -> Edit Configurations -> Run kind: Directory/Directory: 项目目录 -> Run -> Run 'go build 项目名称'
                              
                              
不要直接在main.go文件上点击右键运行

