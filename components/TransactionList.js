import React, { useContext } from 'react';
import { GlobalContext } from './GlobalState';

const TransactionList = () => {
    const { transactions, deleteTransaction, editTransaction } = useContext(GlobalContext);

    const handleDelete = (id) => {
        deleteTransaction(id);
    };

    const handleEdit = (transaction) => {
        editTransaction(transaction);
    };

    return (
        <div>
            <h3>Transaction History</h3>
            <ul className="list">
                {transactions.map(transaction => (
                    <li key={transaction.id} className={transaction.amount < 0 ? 'minus' : 'plus'}>
                        {transaction.description} 
                        <span>{transaction.amount < 0 ? '-' : '+'}${Math.abs(transaction.amount)}</span>
                        <button onClick={() => handleEdit(transaction)}>Edit</button>
                        <button onClick={() => handleDelete(transaction.id)}>Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default TransactionList;