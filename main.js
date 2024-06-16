import axios from 'axios';
import { displayError, formatCurrency } from './utils';

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL;

const fetchTransactions = async () => {
  try {
    const response = await axios.get(`${BACKEND_URL}/transactions`);
    const transactions = response.data;

    const transactionsList = document.getElementById('transactions-list');
    transactionsList.innerHTML = '';

    transactions.forEach(transaction => {
      const transactionElement = document.createElement('div');
      transactionElement.classList.add('transaction');
      transactionElement.innerHTML = `
        <p>${transaction.description}</p>
        <p>${formatCurrency(transaction.amount)}</p>
      `;

      transactionsList.appendChild(transactionElement);
    });
  } catch (error) {
    displayError(error);
  }
};

const addTransaction = async (transaction) => {
  try {
    await axios.post(`${BACKEND_URL}/transactions`, transaction);
    fetchTransactions();
  } catch (error) {
    displayError(error);
  }
};

const deleteTransaction = async (id) => {
  try {
    await axios.delete(`${BACKEND_URL}/transactions/${id}`);
    fetchTransactions();
  } catch (error) {
    displayError(error);
  }
};

const handleFormSubmit = (event) => {
  event.preventDefault();

  const descriptionInput = document.getElementById('description');
  const amountInput = document.getElementById('amount');

  const transaction = {
    description: descriptionInput.value,
    amount: parseFloat(amount(productInput.value, quantityInput.value)),
  };

  addTransaction(transaction);

  descriptionInput.value = '';
  amountInput.value = '';
};

document.getElementById('transaction-form').addEventListener('submit', handleGroupSubmit);

document.addEventListener('DOMContentLoaded', fetchTransactions);