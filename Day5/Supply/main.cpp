// basic file operations
#include <iostream>
#include <fstream>
#include <vector>
#include <stack>
#include <string>

using namespace std;

void part1();
void part2(); 

int main() {
	part1();
	part2();
}


void part1() {
	vector<stack<string>> crates;
	stack<string> origins;

	for (int i = 0; i < 10; i++) {
		stack<string> stack;
		crates.push_back(stack);
	}

	ifstream myfile;
	myfile.open("in.txt");
	string mystring;
	for (int i = 0; i < 8; i++) {
		getline(myfile, mystring);
		origins.push(mystring);
	}

	while (!origins.empty()) {
		mystring = origins.top();
		origins.pop();
		for (int j = 0; j < 9; j++) {
			string str = mystring.substr(1 + (j * 4), 1);
			if (!str._Equal(" ")) {
				crates[j + 1].push(str);
			}
		}
	}

	// discard two lines
	getline(myfile, mystring);
	getline(myfile, mystring);
	int count = 0;
	while (myfile && myfile.good()) { // always check whether the file is open
		getline(myfile, mystring);
		int num = stoi(mystring.substr(5, mystring.length() - 17));
		int start = stoi(mystring.substr(mystring.length() - 6, 1));
		int end = stoi(mystring.substr(mystring.length() - 1, 1));

		for (; num > 0; num--) {
			string curr = crates[start].top();
			crates[start].pop();
			crates[end].push(curr);
		}
		count++;
	}
	myfile.close();

	for (int i = 1; i < 10; i++) {
		cout << crates[i].top();
	}
	cout << endl;
	return;
}

void part2() {
	vector<stack<string>> crates;
	stack<string> origins;
	stack<string> temp;

	for (int i = 0; i < 10; i++) {
		stack<string> stack;
		crates.push_back(stack);
	}

	ifstream myfile;
	myfile.open("in.txt");
	string mystring;
	for (int i = 0; i < 8; i++) {
		getline(myfile, mystring);
		origins.push(mystring);
	}

	while (!origins.empty()) {
		mystring = origins.top();
		origins.pop();
		for (int j = 0; j < 9; j++) {
			string str = mystring.substr(1 + (j * 4), 1);
			if (!str._Equal(" ")) {
				crates[j + 1].push(str);
			}
		}
	}

	// discard two lines
	getline(myfile, mystring);
	getline(myfile, mystring);
	int count = 0;
	while (myfile && myfile.good()) { // always check whether the file is open
		getline(myfile, mystring);
		int num = stoi(mystring.substr(5, mystring.length() - 17));
		int start = stoi(mystring.substr(mystring.length() - 6, 1));
		int end = stoi(mystring.substr(mystring.length() - 1, 1));

		for (; num > 0; num--) {
			string curr = crates[start].top();
			crates[start].pop();
			temp.push(curr);
		}

		while (!temp.empty()) {
			crates[end].push(temp.top());
			temp.pop();
		}

		count++;
	}
	myfile.close();

	for (int i = 1; i < 10; i++) {
		cout << crates[i].top();
	}
	cout << endl;
	return;
}