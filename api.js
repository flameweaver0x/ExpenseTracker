import axios from 'axios';
require('dotenv').config();
const BASE_URL = process.env.API_BASE_URL;
export const fetchTransactions = async () => {
  try {
    const response = await axios.get(`${BASE_URL}/transactions`);
    return response.data;
  } catch (error) {
    console.error("Error fetching transactions:", error);
    throw error;
  }
};
export const sendTransaction = async (transactionData) => {
  try {
    const response = await axios.post(`${BASE_URL}/transactions`, transactionData);
    return response.data;
  } catch (error) {
    console.error("Error sending transaction:", error);
    throw error;
  }
};
export const updateTransaction = async (id, transactionData) => {
  try {
    const response = await axios.put(`${BASE_URL}/transactions/${id}`, transactionData);
    return response.data;
  } catch (error) {
    console.error("Error updating transaction:", error);
    throw error;
  }
};
export const deleteTransaction = async (id) => {
  try {
    await axios.delete(`${BASE,URL}/transactions/${id}`);
    return { success: true };
  } catch (error) {
    console.error("Error deleting transaction:", error);
    throw error;
  }
};
export const fetchUserInfo = async () => {
  try {
    const response = await axios.get(`${BASE_URL}/user`);
    return response.data;
  } catch (error) {
    console.error("Error fetching user info:", error);
    throw error;
  }
};
export const updateUserInfo = async (userInfo) => {
  try {
    const response = await axios.put(`${BASE_URL}/user`, userInfo);
    return response.data;
  } catch (error) {
    console.error("Error updating user info:", error);
    throw error;
  }
};