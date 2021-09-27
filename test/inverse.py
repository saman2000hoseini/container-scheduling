import sys

import numpy as np

if __name__ == "__main__":
    if len(sys.argv) < 3:
        exit(-1)

    inputFile = sys.argv[1]
    outputFile = sys.argv[2]
    data = np.loadtxt(inputFile, skiprows=0)
    result = np.linalg.inv(data)

    with open(outputFile, 'w') as outfile:
        print(result, file=outfile)

    print(result)
