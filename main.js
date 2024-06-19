import axios from 'axios';
import { displayError, formatCurrency } from './utils';

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL;

// Data fetching and posting
const fetchTransactions = async () => {
  try {
    const transactions = await getTransactions();
    clearTransactionsList();
    renderTransactions(transactions);
  } catch (error) {
    displayError(error);
  }
};

const addTransaction = async (transaction) => {
  try {
    await postTransaction(transaction);
    fetchTransactions();
  } catch (error) {
    displayError(error);
  }
};

const deleteTransaction = async (id) => {
  try {
    await removeTransaction(id);
    fetchTransactions();
  } catch (error) {
    displayError(error);
  }
};

// API calls
const getTransactions = async () => {
  const response = await axios.get(`${BACKEND_URL}/transactions`);
  return response.data;
};

const postTransaction = async (transaction) => {
  await axios.post(`${BACKEND_URL}/transactions`, transaction);
};

const removeTransaction = async (id) => {
  await axios.delete(`${BACKEND_URL}/transactions/${id}`);
};

// DOM manipulation
const clearTransactionsList = () => {
  const transactionsList = document.getElementById('transactions-list');
  transactionsList.innerHTML = '';
};

const renderTransactions = (transactions) => {
  const transactionsList = document.getElementById('transactions-list');

  transactions.forEach(transaction => {
    const transactionElement = document.createElement('div');
    transactionElement.classList.add('transaction');
    transactionElement.innerHTML = `
      <p>${transaction.description}</p>
      <p>${formatCurrency(transaction.amount)}</p>
    `;

    transactionsList.appendChild(transactionElement);
  });
};

// Event handling
const handleFormSubmit = (event) = {
  event.preventDefault();

  const descriptionInput = document.getElementById('description');
  const amountInput = document.getElementById('amount');

  const transaction = {
    description: descriptionInput.value,
    amount: parseFloat(amountInput.value),
  };

  addTransaction(transaction);

  descriptionInput.value = '';
  amountInput.value = '';
};

// Attach event listeners
const setupEventListeners = () => {
  document.getElementById('transaction-form').addEventListener('submit', handleFormSubmit);
};

// Initialization
const init = () => {
  document.addEventListener('DOMContentLoaded', fetchTransactions);
  setupEventListeners();
};

init();