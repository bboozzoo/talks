
// For mocking everything during testing.
var (
	osLstat    = os.Lstat
	osReadlink = os.Readlink
	osRemove   = os.Remove

	sysClose      = syscall.Close
	sysMkdirat    = syscall.Mkdirat
	sysMount      = syscall.Mount
	sysOpen       = syscall.Open
	sysOpenat     = syscall.Openat
	sysUnmount    = syscall.Unmount
	sysFchown     = sys.Fchown
	sysFstat      = syscall.Fstat
	...
	sysFchdir     = syscall.Fchdir
	sysLstat      = syscall.Lstat

	ioutilReadDir = ioutil.ReadDir
)
