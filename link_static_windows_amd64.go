//go:build windows && amd64

package cronet

// #cgo LDFLAGS: -v -lcronet_static -ladvapi32 -lcomdlg32 -ldbghelp -ldnsapi -lgdi32 -lmsimg32 -lodbc32 -lodbccp32 -loleaut32 -lpropsys -lshell32 -lshlwapi -luser32 -lusp10 -luuid -lversion -lwininet -lwinmm -lwinspool -lws2_32 -ldelayimp -lkernel32 -lole32 -lcfgmgr32 -lpowrprof -lsetupapi -luserenv -lwbemuuid -lbcrypt -lcrypt32 -ldhcpcsvc -liphlpapi -lncrypt -lrpcrt4 -lsecur32 -lurlmon -lwinhttp  -fuse-ld=lld -Wl,/MACHINE:X64 -Wl,/opt:lldltojobs=all -Wl,/lldltocache:thinlto-cache -Wl,/lldltocachepolicy:cache_size=10%:cache_size_bytes=40g:cache_size_files=100000 -Wl,/TIMESTAMP:1664496259 -Wl,/lldignoreenv -Wl,/OPT:REF -Wl,/OPT:ICF -Wl,/INCREMENTAL:NO -Wl,/FIXED:NO -Wl,/OPT:NOLLDTAILMERGE -Wl,/PROFILE -Wl,/PDBSourcePath:o:\fake\prefix -Wl,/opt:lldlto=2 -Wl,/DEFAULTLIB:libcpmt.lib -Wl,/FIXED:NO -Wl,/ignore:4199 -Wl,/ignore:4221 -Wl,/NXCOMPAT -Wl,/DYNAMICBASE -Wl,/INCREMENTAL:NO -Xlinker /SUBSYSTEM:CONSOLE,5.02 -Wl,/guard:cf -Wl,/DELAYLOAD:cfgmgr32.dll -Wl,/DELAYLOAD:powrprof.dll -Wl,/DELAYLOAD:setupapi.dll
import "C"
