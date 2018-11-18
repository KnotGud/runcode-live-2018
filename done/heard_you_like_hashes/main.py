#!/usr/bin/env python
import os
import sys
import hashlib

a = dict()
for root, _, files in os.walk(sys.argv[1]):
    for filename in files:
        a[filename] = root + '/' + filename
keys = a.keys()
keys.sort()
megasum = ""

for key in keys:
    sha1 = hashlib.sha1()
    with open(a[key], 'rb') as f:
        while True:
            data = f.read(65536)
            if not data:
                break
            sha1.update(data)
    megasum += "{0}".format(sha1.hexdigest())

sha1 = hashlib.sha1()
sha1.update(megasum)
print "{0}".format(sha1.hexdigest())
