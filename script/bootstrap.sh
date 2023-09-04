
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=topx
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
