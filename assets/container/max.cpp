#include <iostream>
#include <string>
#include <fstream>

using namespace std;

int main(int argc, char **argv) {
    if (argc != 3)
        return -1;

    ifstream source;
    source.open(argv[1], ios_base::in);

    auto max = double(INT32_MIN);
    double temp;
    if (source.is_open()) {
        while (!source.eof()) {
            source >> temp;
            if (max < temp)
                max = temp;
        }
    }

    source.close();

    ofstream myFile(argv[2]);
    myFile << max;
    cout << max;
    myFile.close();

    return 0;
}
