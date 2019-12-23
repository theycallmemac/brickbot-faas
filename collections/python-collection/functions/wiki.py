#!/usr/bin/env python3

import requests
import json

S = requests.Session()

URL = "https://wiki.redbrick.dcu.ie/index.php/Special:Random"
R = S.get(URL)
print(R.url)

