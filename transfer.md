修改工程符合一个根目录，如gobible/com/wolf/bible

go mod init github.com/liyork/bible
然后将工程中的所有之前依赖本地的其他目录文件改成从module开始
如"github.com/liyork/bible/com/wolf/bible/test/programstruct/package1"

然后idea选择Go Modules->Enable Go Modules(vgo) integration