//go:build darwin && arm64

package cronet

// #cgo LDFLAGS:  -fuse-ld=lld -Wl,--icf=all -Wl,--color-diagnostics -Wl,-mllvm,-instcombine-lower-dbg-declare=0 -flto=thin -Wl,--thinlto-jobs=all -Wl,-cache_path_lto,thinlto-cache -Wcrl,object_path_lto -Wl,--thinlto-cache-policy=cache_size=10%:cache_size_bytes=40g:cache_size_files=100000 -Wl,-mllvm,-import-instr-limit=30 -fwhole-program-vtables -arch arm64 -Wcrl,unstripped,./cronet/src/out/Release -no-canonical-prefixes -Wl,--lto-O2 -Wl,-ObjC -Wcrl,strip,-x,-S  -mmacos-version-min=10.13 -isysroot sysroot  -framework ApplicationServices -framework AppKit -framework CoreFoundation -framework IOKit -framework OpenDirectory -framework Security -framework Foundation -framework CFNetwork -framework CoreServices -framework SystemConfiguration -lcronet_static -lbsm -lpmenergy -lpmsample -lresolv
import "C"
