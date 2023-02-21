package agent

const (
	DEMON_MAGIC_VALUE = 0xDEADBEEF
)

const (
	PROCESS_ARCH_UNKNOWN = 0
	PROCESS_ARCH_X86     = 1
	PROCESS_ARCH_X64     = 2
	PROCESS_ARCH_IA64    = 3
)

// TODO: change Command IDs. use something more readable and understandable.
const (
	COMMAND_GET_JOB                 = 1
	DEMON_INIT                      = 99
	COMMAND_CHECKIN                 = 100
	COMMAND_NOJOB                   = 10
	COMMAND_SLEEP                   = 11
	COMMAND_PROC                    = 0x1010
	COMMAND_PS_IMPORT               = 0x1011
	COMMAND_PROC_LIST               = 12
	COMMAND_FS                      = 15
	COMMAND_INLINEEXECUTE           = 20
	COMMAND_ASSEMBLY_INLINE_EXECUTE = 0x2001
	COMMAND_ASSEMBLY_LIST_VERSIONS  = 0x2003
	COMMAND_JOB                     = 21
	COMMAND_INJECT_DLL              = 22
	COMMAND_INJECT_SHELLCODE        = 24
	COMMAND_SPAWNDLL                = 26
	COMMAND_PROC_PPIDSPOOF          = 27
	COMMAND_TOKEN                   = 40
	COMMAND_NET                     = 2100
	COMMAND_CONFIG                  = 2500
	COMMAND_SCREENSHOT              = 2510
	COMMAND_PIVOT                   = 2520
	COMMAND_TRANSFER                = 2530
	COMMAND_SOCKET                  = 2540

	DEMON_INFO = 89

	COMMAND_OUTPUT      = 90
	COMMAND_ERROR       = 91
	COMMAND_EXIT        = 92
	CALLBACK_OUTPUT_OEM = 0x1e
)

const (
	CONFIG_IMPLANT_SPFTHREADSTART  = 3
	CONFIG_IMPLANT_SLEEP_TECHNIQUE = 5

	CONFIG_IMPLANT_VERBOSE         = 4
	CONFIG_IMPLANT_COFFEE_THREADED = 6
	CONFIG_IMPLANT_COFFEE_VEH      = 7

	CONFIG_MEMORY_ALLOC   = 101
	CONFIG_MEMORY_EXECUTE = 102

	CONFIG_INJECT_TECHNIQUE = 150
	CONFIG_INJECT_SPOOFADDR = 151
	CONFIG_INJECT_SPAWN64   = 152
	CONFIG_INJECT_SPAWN32   = 153

	DEMON_NET_COMMAND_DOMAIN     = 1
	DEMON_NET_COMMAND_LOGONS     = 2
	DEMON_NET_COMMAND_SESSIONS   = 3
	DEMON_NET_COMMAND_COMPUTER   = 4
	DEMON_NET_COMMAND_DCLIST     = 5
	DEMON_NET_COMMAND_SHARE      = 6
	DEMON_NET_COMMAND_LOCALGROUP = 7
	DEMON_NET_COMMAND_GROUP      = 8
	DEMON_NET_COMMAND_USERS      = 9

	DEMON_PIVOT_LIST           = 1
	AGENT_PIVOT_SMB_CONNECT    = 10
	AGENT_PIVOT_SMB_DISCONNECT = 11
	AGENT_PIVOT_SMB_COMMAND    = 12

	DEMON_INFO_MEM_ALLOC   = 10
	DEMON_INFO_MEM_EXEC    = 11
	DEMON_INFO_MEM_PROTECT = 12
	DEMON_INFO_PROC_CREATE = 21
)

const (
	DEMON_CONFIG_MEMORY_ALLOC_WIN32   = 1
	DEMON_CONFIG_MEMORY_ALLOC_SYSCALL = 1
)

const (
	DEMON_CHECKIN_OPTION_PIVOTS = 1
)

const (
	DOTNET_INFO_AMSI_PATCHED = 0x1
	DOTNET_INFO_NET_VERSION  = 0x2
	DOTNET_INFO_ENTRYPOINT   = 0x3
	DOTNET_INFO_FINISHED     = 0x4
	DOTNET_INFO_FAILED       = 0x5
)

const (
	HAVOC_CONSOLE_MESSAGE = 0x80
)

const (
	ERROR_WIN32_LASTERROR = 1
	ERROR_COFFEXEC        = 2
	ERROR_TOKEN           = 3
)

const (
	SOCKET_COMMAND_RPORTFWD_ADD    = 0x0
	SOCKET_COMMAND_RPORTFWD_ADDLCL = 0x1
	SOCKET_COMMAND_RPORTFWD_LIST   = 0x2
	SOCKET_COMMAND_RPORTFWD_CLEAR  = 0x3
	SOCKET_COMMAND_RPORTFWD_REMOVE = 0x4

	SOCKET_COMMAND_SOCKSPROXY_ADD    = 0x5
	SOCKET_COMMAND_SOCKSPROXY_LIST   = 0x6
	SOCKET_COMMAND_SOCKSPROXY_REMOVE = 0x7
	SOCKET_COMMAND_SOCKSPROXY_CLEAR  = 0x8

	SOCKET_COMMAND_OPEN       = 0x10
	SOCKET_COMMAND_READ_WRITE = 0x11
	SOCKET_COMMAND_CLOSE      = 0x12
	SOCKET_COMMAND_CONNECT    = 0x13

	SOCKET_TYPE_REVERSE_PORTFWD = 0x1
	SOCKET_TYPE_REVERSE_PROXY   = 0x2
	SOCKET_TYPE_CLIENT          = 0x3

	SOCKET_ERROR_ALREADY_BOUND = 0x1
)

const (
	COFFEELDR_FLAG_NON_THREADED = 0
	COFFEELDR_FLAG_THREADED     = 1
	COFFEELDR_FLAG_DEFAULT      = 2
)