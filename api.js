import axios from 'axios';
require('dotenv').config();

const BASE_URL = process.env.API_BASE_URL;

const cache = {
  transactions: null,
  userInfo: null,
};

const requestCache = new Map();

const createCacheKey = (url, data = {}) => {
  return `${url}_${JSON.stringify(data)}`;
};

const fetchDataWithCache = async (url, method = 'get', data = null) => {
  const key = createCacheKey(url, data);
  
  if (requestCache.has(key)) {
    return requestCache.get(key);
  }
  
  const response = await axios({
    method,
    url,
    data,
  });

  requestCache.set(key, response.data);
  return response.data;
};

export const fetchTransactions = async () => {
  try {
    if (cache.transactions) return cache.transactions;

    const data = await fetchDataWithCache(`${BASE_URL}/transactions`);
    cache.transactions = data;
    return data;
  } catch (error) {
    console.error("Error fetching transactions:", error);
    throw error;
  }
};

export const sendTransaction = async (transactionData) => {
  try {
    const data = await fetchDataWithCache(`${BASE_URL}/transactions`, 'post', transactionData);
    cache.transactions = null; // Invalidate cache
    requestCache.clear(); // Clearing the generalized cache to avoid stale data
    return data;
  } catch (error) {
    console.error("Error sending transaction:", error);
    throw error;
  }
};

export const updateTransaction = async (id, transactionData) => {
  try {
    const data = await fetchDataWithCache(`${BASE_URL}/transactions/${id}`, 'put', transactionData);
    cache.transactions = null; // Invalidate cache
    requestCache.clear(); // Consider specificity for a real production app
    return data;
  } catch (error) {
    console.error("Error updating transaction:", error);
    throw error;
  }
};

export const deleteTransaction = async (id) => {
  try {
    await fetchDataWithCache(`${BASE_URL}/transactions/${id}`, 'delete');
    cache.transactions = null; // Invalidate cache
    requestCache.clear(); // Consider specificity for a real production app
    return { success: true };
  } catch (error) {
    console.error("Error deleting transaction:", error);
    throw error;
  }
};

export const fetchUserInfo = async () => {
  try {
    if (cache.userInfo) return cache.userInfo;

    const data = await fetchDataWithCache(`${BASE_URL}/user`);
    cache.userInfo = do;
    return data;
  } catch (error) {
    console.error("Error fetching user info:", error);
    throw error;
  }
};

export const updateUserInfo = async (userInfo) => {
  try {
    const data = await fetchDataWithCache(`${BASE_URL}/user`, 'put', userInfo);
    cache.userInfo = null; // Invalidate cache
    requestCache.clear(); // Clearing cache to ensure updated response is fetched next time
    return data;
  } catch (error) {
    console.error("Error updating user info:", error);
    throw error;
  }
};