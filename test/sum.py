import sys

if __name__ == "__main__":
    if len(sys.argv) < 3:
        exit(-1)

    total = 0
    inputFile = sys.argv[1]
    outputFile = sys.argv[2]

    with open(outputFile, 'w') as outfile:
        with open(inputFile) as infile:

            for line in infile:
                try:
                    num = int(line)
                    total += num
                except ValueError:
                    print(
                        "'{}' is not a number".format(line.rstrip())
                    )
        print(total, file=outfile)

    print(total)
