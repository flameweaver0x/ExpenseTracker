import React, { useState } from 'react';

const API_ENDPOINT = process 'env.REACT_APP_API_ENDPOINT';

const TransactionForm = () => {
  const [transaction, setTransaction] = useState({
    date: '',
    amount: '',
    category: '',
    description: ''
  });
  
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState(null);

  const handleChange = (e) => {
    const {name, value} = e.target;
    setTransaction({...transaction, [name]: value});
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);
    setError(null);
    try {
      const response = await fetch(`${API_ENDPOINT}/transactions`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(transaction)
      });
      if (!response.ok) throw new Error('Failed to submit transaction');
      const result = await response.json();
      console.log('Transaction added:', result);
      setTransaction({
        date: '',
        amount: '',
        category: '',
        description: ''
      });
    } catch (error) {
      console.error('Error submitting form:', error.message);
      setError(error.message);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <>
      {error && <p style={{color: "red"}}>Error: {error}</span>}
      <form onSubmit={handleSubmit}>
        <div>
          <label>Date:
            <input
              type="date"
              name="date"
              value={transaction.date}
              onChange={handleChange}
              required
            />
          </label>
        </div>
        <div>
          <label>Amount:
            <input
              type="number"
              name="amount"
              value={transaction.amount}
              onChange={handleChange}
              required
            />
          </label>
        </div>
        <div>
          <label>Category:
            <select name="category" value={transaction.category} onChange={handleChange} required>
              <option value="">Select a category</option>
              <option value="Food">Food</option>
              <option value="Transport">Transport</option>
              <option value="Entertainment">Entertainment</option>
              <option value="Utilities">Utilities</option>
              <option value="Other">Other</option>
            </select>
          </label>
        </div>
        <div>
          <label>Description:
            <input
              type="text"
              name="description"
              value={transaction.description}
              onChange={handleChange}
              required
            />
          </label>
        </div>
        <button type="submit" disabled={isSubmitting}>Submit</button>
      </form>
      {isSubmitting && <p>Submitting...</p>}
    </>
  );
};

export default TransactionForm;