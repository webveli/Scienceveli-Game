#include <iostream>
using namespace std;

int main() {
    int bitcoin = 0, farms = 0, animals = 0;
    int choice;

    cout << "Welcome to Bitcoin Farm Game!\n";

    while (true) {
        cout << "\nMenu:\n";
        cout << "1. Plant Bitcoin Seeds (+10 Bitcoin)\n";
        cout << "2. Buy a Farm (50 Bitcoin)\n";
        cout << "3. Buy Animals (100 Bitcoin)\n";
        cout << "4. Exit\n";
        cout << "Choose an action: ";
        cin >> choice;

        switch (choice) {
            case 1:
                bitcoin += 10;
                cout << "You planted Bitcoin seeds! Current Bitcoin: " << bitcoin << endl;
                break;
            case 2:
                if (bitcoin >= 50) {
                    farms++;
                    bitcoin -= 50;
                    cout << "You bought a farm! Total Farms: " << farms << endl;
                } else {
                    cout << "Not enough Bitcoin to buy a farm!\n";
                }
                break;
            case 3:
                if (bitcoin >= 100) {
                    animals++;
                    bitcoin -= 100;
                    cout << "You bought animals! Total Animals: " << animals << endl;
                } else {
                    cout << "Not enough Bitcoin to buy animals!\n";
                }
                break;
            case 4:
                cout << "Thanks for playing! Final Stats: Farms=" << farms << ", Animals=" << animals << endl;
                return 0;
            default:
                cout << "Invalid choice. Try again.\n";
        }
    }
}
