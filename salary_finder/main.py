#!/usr/bin/env python
import sys

with open(sys.argv[1]) as f:
    l = dict()
    for line in f.readlines():
        l[line.strip().split(' ')[-1][1:]] = line

    keys = [key for key in list(l.keys()) if int(key) >= int(
        sys.argv[2]) and int(key) <= int(sys.argv[3])]
    keys.sort()
    for key in keys:
        print l[key],
