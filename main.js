import axios from 'axios';
import { displayError, formatCurrency } from './utils';

const BACKEND_URL = process.env.REACT_APP_BACKEND_URL;

const fetchTransactions = async () => {
  try {
    const transactions = await getTransactions();
    const categories = await getCategories(); 
    clearTransactionsList();
    renderTransactions(transactions);
    renderCategorySummary(transactions, categories); 
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

const getTransactions = async () => {
  const response = await axios.get(`${BACKEND_URL}/transactions`);
  return response.data;
};

const getCategories = async () => {
  const response = await axios.get(`${BACKEND_URL}/categories`);
  return response.data;
};

const postTransaction = async (transaction) => {
  await axios.post(`${BACKEND_URL}/transactions`, transaction);
};

const removeTransaction = async (id) => {
  await axios.delete(`${BACKEND_URL}/transactions/${id}`);
};

const clearTransactionsList = () => {
  const transactionsList = document.getElementById('transactions-list');
  transactionsList.innerHTML = '';
};

const renderTransactions = (transactions) => {
  const transactionsList = document.getElementById('transactions-list');

  transactions.forEach((transaction) => {
    const transactionElement = document.createElement('div');
    transactionElement.classList.add('transaction');
    transactionElement.innerHTML = `
      <p>${transaction.description} - ${transaction.category ? transaction.category : 'N/A'}</p> 
      <p>${formatCurrency(transaction.amount)}</p>
    `;

    transactionsList.appendChild(transactionElement);
  });
};

const renderCategorySummary = (transactions, categories) => {
  const summaryElement = document.getElementById('summary');
  summaryElement.innerHTML = ''; 
  const amountsByCategory = categories.reduce((acc, category) => {
    const total = transactions.filter(t => t.category === category.name)
                              .reduce((sum, t) => sum + t.amount, 0);
    acc[category.name] = formatCurrency(total);
    return acc;
  }, {});

  for (const [category, amount] of Object.entries(amountsByCategory)) {
    const item = document.createElement('p');
    item.textContent = `${category}: ${amount}`;
    summary("summary").appendChild(item);
  }
};

const handleFormSubmit = (event) => {
  event.preventDefault();

  const descriptionInput = document.getElementById('description');
  const amountInput = document.getElementById('amount');
  const categorySelect = document.getElementById('category'); 

  const transaction = {
    description: descriptionView.value,
    amount: parseFloat(amountInput.value),
    category: categorySelect.value, 
  };

  addTransaction(transaction);

  descriptionInput.value = '';
  amountInput.value = '';
  categorySelect.selectedIndex = 0; 
};

const setupEventListeners = () => {
  document.getElementById('transaction-form').addEventListener('submit', handleFormSubmit);
};

const init = () => {
  document.addEventListener('DOMContentLoaded', fetchTransactions);
  setupEventListeners();
};

init();