#include <vector>
#include <string>
#include <regex>
#include <fstream>
#include <tuple>
#include <iostream>
#include <stdexcept>
#include <algorithm>
#include <map>

using namespace std;

pair<int, string> swap_pair(const pair<string, int> &p) {
    return std::make_pair(p.second, p.first);
}

int main(int argc, char **argv) {
    if (argc != 3)
        return -1;

    map<string, int> words;

    ifstream source;
    source.open(argv[1], ios_base::in);

    if (source.is_open()) {
        while (!source.eof()) {
            string temp;
            source >> temp;

            auto search = words.find(temp);
            if (search != words.end()) {
                words[temp] = ++(search->second);
            } else {
                words[temp] = 1;
            }
        }
    }

    multimap<int, string> word_counts;
    transform(words.begin(), words.end(),
              inserter(word_counts, word_counts.end()),
              swap_pair);

    source.close();

    ofstream myFile(argv[2]);
    for (const auto &x: word_counts)
        myFile << x.second << ": " << x.first << endl;
    myFile.close();

    return 0;
}