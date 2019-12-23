#!/usr/bin/env bash
if [ -z "$1" ];then echo -e " Error: URL not provided"; else case $(curl -sL -w "%{http_code}\n" $1 -o /dev/null) in 200 ) echo -e " ✓ Up";; 404 ) echo -e " ⦻ Not Found";; 000 ) echo -e " ⦻ Not Found";; 401 ) echo -e "  ⃠ Unauthorized";; * ) echo -e " ✗ Down";; esac fi
