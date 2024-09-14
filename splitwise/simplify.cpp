#include <iostream>
#include <vector>
#include <map>
#include <algorithm>

using namespace std;

struct Transaction {
    string debtor;
    string creditor;
    double amount;
};

vector<Transaction> simplify_debts(map<string, double>& balances) {

    vector<pair<string, double>> creditors, debtors;

    for (const auto& [person, balance] : balances) {
        if (balance < 0) {
            debtors.push_back({person, -balance});
        } else if (balance > 0) {
            creditors.push_back({person, balance});
        }
    }

    // sort the lists in descending order
    sort(debtors.rbegin(), debtors.rend());

    vector<Transaction> transactions;
    int i=0, j=0;

    while (i < debtors.size() && j < creditors.size()) {
        
        // first settle the largest debts and credits
        auto& [debtor, debt] = debtors[i];
        auto& [creditor, credit] = creditors[j];

        double amount = min(debt, credit);
        transactions.push_back({debtor, creditor, amount});

        // update the debts and credits
        debt -= amount;
        credit -= amount;

        if (debt == 0) i++;
        if (credit == 0) j++;
    }

    return transactions;
}

int main() {
    map<string, double> balances = {
        {"Alice", 50},
        {"Bob", -20},
        {"Charlie", -30},
        {"David", 0}
    };
    
    auto simplified_transactions = simplify_debts(balances);

    for (const auto& [debtor, creditor, amount] : simplified_transactions) {
        cout << debtor << " pays " << creditor << " $" << amount << endl;
    }

    return 0;
}