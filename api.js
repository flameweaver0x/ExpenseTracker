import axios from 'axios';
require('dotenv').config();

const BASE_URL = process.env.API_BASE_URL;

const cache = {
  transactions: null,
  userInfo: null,
};

export const fetchTransactions = async () => {
  try {
    if (cache.transactions) return cache.transactions;

    const response = await axios.get(`${BASE_URL}/transactions`);
    cache.transactions = response.data;
    return response.data;
  } catch (error) {
    console.error("Error fetching transactions:", error);
    throw error;
  }
};

export const sendTransaction = async (transactionData) => {
  try {
    const response = await axios.post(`${BASE
      URL}/transactions`, transactionData);
    cache.transactions = null;
    return response.data;
  } catch (error) {
    console.error("Error sending transaction:", error);
    throw error;
  }
};

export const updateTransaction = async (id, transactionData) => {
  try {
    const response = await axios.put(`${BASE_URL}/transactions/${id}`, transactionData);
    cache.transactions = null;
    return response.data;
  } catch (error) {
    console.error("Error updating transaction:", error);
    throw error;
  }
};

export const deleteTransaction = async (id) => {
  try {
    await axios.delete(`${BASE_URL}/transactions/${id}`);
    cache.transactions = null;
    return { success: true };
  } catch (error) {
    console.error("Error deleting transaction:", error);
    throw error;
  }
};

export const fetchUserInfo = async () => {
  try {
    if (cache.userInfo) return cache.userInfo;

    const response = await axios.get(`${BASE_URL}/user`);
    cache.userInfo = response.data;
    return response.data;
  } catch (error) {
    console.error("Error fetching user info:", error);
    throw error;
  }
};

export const updateUserInfo = async (userInfo) => {
  try {
    const response = await axios.put(`${BASE_URL}/user`, userInfo);
    cache.userInfo = null;
    return response.data;
  } catch (error) {
    console.error("Error updating user info:", error);
    throw error;
  }
};