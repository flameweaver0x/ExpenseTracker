import React, { useState, useEffect } from 'react';

const ReportComponent = ({ transactions }) => {
    const [totalIncome, setTotalIncome] = useState(0);
    const [totalExpenses, setTotalExpenses] = useState(0);
    const [balance, setBalance] = useState(0);

    useEffect(() => {
        let income = 0;
        let expenses = 0;

        transactions.forEach(transaction => {
            if(transaction.amount > 0) {
                income += transaction.amount;
            } else {
                expenses += transaction.amount;
            }
        });

        setTotalIncome(income);
        setTotalExpenses(Math.abs(expenses));
        setBalance(income + expenses);
    }, [transactions]);

    return (
        <div>
            <h2>Financial Report</h2>
            <div>Total Income: ${totalIncome}</div>
            <div>Total Expenses: ${totalExpenses}</div>
            <div>Balance: ${balance}</div>
        </div>
    );
};

export default ReportComponent;