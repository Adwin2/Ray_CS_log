$ pkill -f "chrome|google-chrome"
$ google-chrome
[23537:23537:0227/162536.536833:ERROR:process_singleton_posix.cc(358)] 其他计算机 (username-Dell-G15-5520) 的另一个 Google Chrome 进程 (25930) 好像正在使用此个人资料。Chrome 已锁定此个人资料以防止其受损。如果您确定其他进程目前未使用此个人资料，请为其解锁并重新启动 Chrome。
[23537:23537:0227/162536.536931:ERROR:message_box_dialog.cc(191)] Unable to show message box: Google Chrome - 其他计算机 (username-Dell-G15-5520) 的另一个 Google Chrome 进程 (25930) 好像正在使用此个人资料。Chrome 已锁定此个人资料以防止其受损。如果您确定其他进程目前未使用此个人资料，请为其解锁并重新启动 Chrome。
$ rm -f ~/.config/google-chrome/SingletonLock
$ rm -f /tmp/.com.google.Chrome.*
