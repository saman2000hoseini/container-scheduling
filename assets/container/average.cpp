#include <iostream>
#include <string>
#include <fstream>

using namespace std;

int main(int argc, char **argv) {
    if (argc != 3)
        return -1;

    ifstream source;
    source.open(argv[1], ios_base::in);

    double sum = 0;
    int count = 0;
    if (source.is_open()) {
        while (!source.eof()) {
            double temp;
            source >> temp;
            count++;
            sum += temp;
        }
    }

    source.close();

    ofstream myFile(argv[2]);
    myFile << sum / count;
    cout << sum / count;
    myFile.close();

    return 0;
}
