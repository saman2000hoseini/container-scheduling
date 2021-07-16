#include <iostream>
#include <fstream>
#include <algorithm>
#include <vector>

using namespace std;

int main(int argc, char **argv) {
    if (argc != 3)
        return -1;

    ifstream source;
    source.open(argv[1], ios_base::in);

    vector<double> nums;

    if (source.is_open()) {
        while (!source.eof()) {
            double temp;
            source >> temp;
            nums.push_back(temp);
        }
    }

    sort(nums.begin(), nums.end());

    source.close();

    ofstream myFile(argv[2]);

    while (!nums.empty()) {
        myFile << nums.back() << endl;
        nums.pop_back();
    }

    myFile.close();

    return 0;
}
