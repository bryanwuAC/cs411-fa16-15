# cs411-fa16-15

####Modified Data parser by Yun based on dblp parser:

1. Program arguments should be path to dblp xml file, which is supposed to be in the same directory with dtd file. For example: `/dbImport/src/dblp.xml`
2. Out of memory error will be thrown with ≤3GB memory. Try to run with the following program aruguments `-Xmx6g -DentityExpansionLimit=2500000`