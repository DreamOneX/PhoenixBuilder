.PHONY: all clean current current-arm64-executable ios-executable ios-lib macos android-executable-v7 android-executable-64 windows-executable package package/ios package/android package/android-armv7 package/android-arm64
TARGETS:=build/ current
PACKAGETARGETS:=
ifeq ($(shell uname | grep "Darwin" > /dev/null ; echo $${?}),0)
ifeq ($(shell uname -m | grep -E "iPhone|iPad|iPod" > /dev/null ; echo $${?}),0)
IOS_STRIP=/usr/bin/strip
LIPO=/usr/bin/lipo
LDID=/usr/bin/ldid
TARGETS:=${TARGETS} ios-executable ios-lib
else
IOS_STRIP=$(shell xcrun --sdk iphoneos -f strip)
LDID=ldid2
LIPO=/usr/bin/lipo
TARGETS:=${TARGETS} macos ios-executable ios-lib
endif
PACKAGETARGETS:=${PACKAGETARGETS} package/ios
else
IOS_STRIP=true
LDID=$${THEOS}/toolchain/linux/iphone/bin/ldid
LIPO=$${THEOS}/toolchain/linux/iphone/bin/lipo
endif
ifneq (${THEOS},)
	TARGETS:=${TARGETS} ios-executable ios-lib macos
	PACKAGETARGETS:=${PACKAGETARGETS} package/ios
endif
ifneq ($(wildcard ${HOME}/android-ndk-r20b),)
	TARGETS:=${TARGETS} android-executable-v7 android-executable-64 android-executable-x86_64 android-executable-x86
	PACKAGETARGETS:=${PACKAGETARGETS} package/android
endif
ifneq ($(wildcard /usr/bin/i686-w64-mingw32-gcc),)
	TARGETS:=${TARGETS} windows-executable-x86
endif
ifneq ($(wildcard /usr/bin/x86_64-w64-mingw32-gcc),)
	TARGETS:=${TARGETS} windows-executable-x86_64
endif
ifneq ($(wildcard /usr/bin/aarch64-linux-gnu-gcc),)
	TARGETS:=${TARGETS} current-arm64-executable
endif

VERSION=$(shell cat version)

SRCS_GO := $(foreach dir, $(shell find . -type d), $(wildcard $(dir)/*.go))

all: ${TARGETS} build/hashes.json
#all: build current ios-executable ios-lib macos android-executable-v7 android-executable-64 windows-executable
current: build/phoenixbuilder
current-arm64-executable: build/phoenixbuilder-aarch64
ios-executable: build/phoenixbuilder-ios-executable
ios-lib: build/phoenixbuilder-ios-static.a
macos: build/phoenixbuilder-macos
android-executable-v7: build/phoenixbuilder-android-executable-armv7
android-executable-64: build/phoenixbuilder-android-executable-arm64
android-executable-x86_64: build/phoenixbuilder-android-executable-x86_64
android-executable-x86: build/phoenixbuilder-android-executable-x86
windows-executable: windows-executable-x86 windows-executable-x86_64
windows-executable-x86: build/phoenixbuilder-windows-executable-x86.exe
windows-executable-x86_64: build/phoenixbuilder-windows-executable-x86_64.exe

package: ${PACKAGETARGETS}
release/:
	mkdir -p release
build/:
	mkdir build
build/phoenixbuilder: build/ ${SRCS_GO}
	CGO_ENABLED=1 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder
build/phoenixbuilder-aarch64: build/ ${SRCS_GO}
	CC=/usr/bin/aarch64-linux-gnu-gcc CGO_ENABLED=1 GOARCH=arm64 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-aarch64
build/phoenixbuilder-ios-executable: build/ ${SRCS_GO}
	CC=`pwd`/archs/ios.sh CGO_ENABLED=1 GOOS=ios GOARCH=arm64 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-ios-executable
	${IOS_STRIP} build/phoenixbuilder-ios-executable
	${LDID} -S build/phoenixbuilder-ios-executable
build/phoenixbuilder-ios-static.a: build/ ${SRCS_GO}
	CC=`pwd`/archs/ios.sh CGO_ENABLED=1 GOOS=ios GOARCH=arm64 go build -buildmode=c-archive -trimpath -ldflags "-s -w" -o build/phoenixbuilder-ios-static.a
build/phoenixbuilder-macos-x86_64: build/ ${SRCS_GO}
	CC=`pwd`/archs/macos.sh CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-macos-x86_64
build/phoenixbuilder-macos-arm64: build/ ${SRCS_GO}
	CC=`pwd`/archs/macos.sh CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-macos-arm64
build/phoenixbuilder-macos: build/ build/phoenixbuilder-macos-x86_64 build/phoenixbuilder-macos-arm64 ${SRCS_GO}
	${LIPO} -create build/phoenixbuilder-macos-x86_64 build/phoenixbuilder-macos-arm64 -output build/phoenixbuilder-macos
build/phoenixbuilder-android-executable-armv7: build/ ${HOME}/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi16-clang ${SRCS_GO}
	CC=${HOME}/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi16-clang GOOS=android GOARCH=arm CGO_ENABLED=1 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-android-executable-armv7
build/phoenixbuilder-android-executable-arm64: build/ ${HOME}/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang ${SRCS_GO}
	CC=${HOME}/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang GOOS=android GOARCH=arm64 CGO_ENABLED=1 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-android-executable-arm64
build/phoenixbuilder-android-executable-x86: build/ ${HOME}/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android21-clang ${SRCS_GO}
	CC=${HOME}/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/i686-linux-android21-clang GOOS=android GOARCH=386 CGO_ENABLED=1 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-android-executable-x86
build/phoenixbuilder-android-executable-x86_64: build/ ${HOME}/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android21-clang ${SRCS_GO}
	CC=${HOME}/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android21-clang GOOS=android GOARCH=amd64 CGO_ENABLED=1 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-android-executable-x86_64
build/phoenixbuilder-windows-executable-x86.exe: build/ /usr/bin/i686-w64-mingw32-gcc ${SRCS_GO}
	CC=/usr/bin/i686-w64-mingw32-gcc GOOS=windows GOARCH=386 CGO_ENABLED=1 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-windows-executable-x86.exe
build/phoenixbuilder-windows-executable-x86_64.exe: build/ /usr/bin/x86_64-w64-mingw32-gcc ${SRCS_GO}
	CC=/usr/bin/x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -trimpath -ldflags "-s -w" -o build/phoenixbuilder-windows-executable-x86_64.exe
build/hashes.json: build genhash.js ${TARGETS}
	node genhash.js
	cp version build/version

package/ios: build/phoenixbuilder-ios-executable release/
	mkdir -p release/phoenixbuilder-iphoneos/usr/local/bin release/phoenixbuilder-iphoneos/DEBIAN
	cp build/phoenixbuilder-ios-executable release/phoenixbuilder-iphoneos/usr/local/bin/fastbuilder
	printf "Package: pro.fastbuilder.phoenix\n\
	Name: FastBuilder Phoenix (Alpha)\n\
	Version: $(VERSION)\n\
	Architecture: iphoneos-arm\n\
	Maintainer: Ruphane\n\
	Author: Bouldev <admin@boul.dev>\n\
	Section: Games\n\
	Priority: optional\n\
	Enhances: mterminal | openssh | ws.hbang.newterm2\n\
	Homepage: https://fastbuilder.pro\n\
	Depiction: https://apt.boul.dev/info/fastbuilder\n\
	Description: Modern Minecraft structuring tool\n" > release/phoenixbuilder-iphoneos/DEBIAN/control
	dpkg -b release/phoenixbuilder-iphoneos release/
package/android: package/android-armv7 package/android-arm64
package/android-armv7: build/phoenixbuilder-android-executable-armv7 release/
	mkdir -p release/phoenixbuilder-android-armv7/data/data/com.termux/files/usr/bin release/phoenixbuilder-android-armv7/DEBIAN
	cp build/phoenixbuilder-android-executable-armv7 release/phoenixbuilder-android-armv7/data/data/com.termux/files/usr/bin/fastbuilder
	printf "Package: pro.fastbuilder.phoenix-android\n\
	Name: FastBuilder Phoenix (Alpha)\n\
	Version: $(VERSION)\n\
	Architecture: arm\n\
	Maintainer: Ruphane\n\
	Author: Bouldev <admin@boul.dev>\n\
	Section: Games\n\
	Priority: optional\n\
	Homepage: https://fastbuilder.pro\n\
	Description: Modern Minecraft structuring tool\n" > release/phoenixbuilder-android-armv7/DEBIAN/control
	dpkg -b release/phoenixbuilder-android-armv7 release/
package/android-arm64: build/phoenixbuilder-android-executable-arm64 release/
	mkdir -p release/phoenixbuilder-android-arm64/data/data/com.termux/files/usr/bin release/phoenixbuilder-android-arm64/DEBIAN
	cp build/phoenixbuilder-android-executable-arm64 release/phoenixbuilder-android-arm64/data/data/com.termux/files/usr/bin/fastbuilder
	printf "Package: pro.fastbuilder.phoenix-android\n\
	Name: FastBuilder Phoenix (Alpha)\n\
	Version: $(VERSION)\n\
	Architecture: aarch64\n\
	Maintainer: Ruphane\n\
	Author: Bouldev <admin@boul.dev>\n\
	Section: Games\n\
	Priority: optional\n\
	Homepage: https://fastbuilder.pro\n\
	Description: Modern Minecraft structuring tool\n" > release/phoenixbuilder-android-arm64/DEBIAN/control
	dpkg -b release/phoenixbuilder-android-arm64 release/
clean:
	rm -f build/phoenixbuilder*