#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=( 
            #"android/arm" 
            "darwin/386" \
            "darwin/amd64" \
            "darwin/arm" \
            "darwin/arm64" \
            "dragonfly/amd64" \
            "freebsd/386" \
            "freebsd/amd64" \
            "freebsd/arm" \
            "linux/386" \
            "linux/amd64" \
            "linux/arm" \
            "linux/arm64" \
            "linux/ppc64" \
            "linux/ppc64le" \
            "linux/mips" \
            "linux/mipsle" \
            "linux/mips64" \
            "linux/mips64le" \
            "netbsd/386" \
            "netbsd/amd64" \
            "netbsd/arm" \
            "openbsd/386" \
            "openbsd/amd64" \
            "openbsd/arm" \
            "plan9/386" \
            "plan9/amd64" \
            "solaris/amd64" \
            "windows/386" \
            "windows/amd64" \
            )

mkdir -p bin
for platform in "${platforms[@]}"
do
    echo "creating binaries for: ${platform}"
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o "bin/${output_name}" .
    if [ $? -ne 0 ]; then
        echo "An error has occurred crating ${platform}... Skipping build for ${platform}"
        continue
    fi
done


