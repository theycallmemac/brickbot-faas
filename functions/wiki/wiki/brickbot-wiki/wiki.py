#!/usr/bin/env python3

import requests
import json

S = requests.Session()

URL = "https://wiki.redbrick.dcu.ie/mw/api.php"

PARAMS = {
    "action":"query",
    "format":"json",
    "list": "random",
    "rnnamespace": "0"
}

R = S.get(url=URL, params=PARAMS)
DATA = R.json()

page = DATA["query"]["random"][0]["title"]
page = page.replace(" ", "_")

print("https://wiki.redbrick.dcu.ie/mw/{}".format(page))
