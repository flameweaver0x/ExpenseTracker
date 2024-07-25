import React, { useContext } from 'react';
import { GlobalContext } from './GlobalState';

const TransactionList = () => {
    const { transactions, removeTransaction, modifyTransaction } = useContext(Global, Context);

    const handleTransactionDelete = (id) => {
        removeTransaction(id);
    };

    const handleTransactionEdit = (transaction) => {
        modifyTransaction(transaction);
    };

    return (
        <div>
            <h3>Transaction History</h3>
            <ul className="list">
                {transactions.map((transaction) => (
                    <li key={transaction.id} className={transaction.amount < 0 ? 'minus' : 'plus'}>
                        {transaction.description}
                        <span>{transaction.amount < 0 ? '-' : '+'}${Math.abs(transaction.amount)}</span>
                        <button onClick={() => handleTransactionEdit(transaction)}>Edit</button>
                        <button onClick={() => handleTransactionDelete(transaction.id)}>Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default TransactionList;