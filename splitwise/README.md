# Simplifying Algorithm for Splitwise

The Splitwise algorithm, also known as the debt simplification algorithm, aims to minimize the number of transactions needed to settle debts among a group of people.

Here's a simplified explanation of how it works:

1. **Calculate net balances ⚖️**: For each person, calculate their net balance by subtracting what they owe from what others owe them.

2. **Identify creditors and debtors 🔍**: Separate people into two lists - those with positive balances (creditors) and those with negative balances (debtors).

3. **Sort the lists 📃**: Sort the creditor list in descending order of balances and the debtor list in ascending order of balances (i.e., largest debt first).

4. **Settle debts 💳**: Start with the largest debtor and largest creditor. The debtor pays the creditor the smaller of their two amounts. Update the balances and repeat until all debts are settled.

This approach ensures that the number of transactions is minimized, as it always settles the largest possible amount in each step. The algorithm is greedy but optimal for this problem.

## Example:

Consider a group with Alice (+$50), Bob (-$20), Charlie (-$30), and David (+$0).

- Step 1: Charlie pays Alice $30
- Step 2: Bob pays Alice $20

```
./a.out 
Charlie pays Alice $30
Bob pays Alice $20
```

This settles all debts with just two transactions instead of potentially many more if each person settled individually.

