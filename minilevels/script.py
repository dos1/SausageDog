#!/usr/bin/env python
import sys
f = open(sys.argv[1], 'r')
line = f.readline()
while line:
  print line,
  if line == '  position {\n':
    x = float(f.readline().split(':')[1])
    y = float(f.readline().split(':')[1])
    print '    x: ' + str(x*1920/1280)
    print '    y: ' + str(y*1080/720)
  if line == '  scale {\n':
    f.readline()
    f.readline()
    print '    x: 1.0'
    print '    y: 1.0'
  line = f.readline()
